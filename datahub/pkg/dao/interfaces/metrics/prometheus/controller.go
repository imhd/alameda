package prometheus

import (
	"context"
	"strings"

	EntityPromthMetric "github.com/containers-ai/alameda/datahub/pkg/dao/entities/prometheus/metrics"
	DaoClusterStatusTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	RepoInfluxClusterStatus "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	RepoPromthMetric "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/prometheus/metrics"
	"github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InternalPromth "github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type ControllerMetrics struct {
	PrometheusConfig InternalPromth.Config
	InfluxDBConfig   InternalInflux.Config

	influxControllerRepo *RepoInfluxClusterStatus.ControllerRepository
	influxPodRepo        *RepoInfluxClusterStatus.PodRepository

	clusterUID string
}

// NewControllerMetricsWithConfig Constructor of prometheus controller metric dao
func NewControllerMetricsWithConfig(promCfg InternalPromth.Config, influxCfg InternalInflux.Config, clusterUID string) DaoMetricTypes.ControllerMetricsDAO {
	return &ControllerMetrics{
		PrometheusConfig: promCfg,
		InfluxDBConfig:   influxCfg,

		influxControllerRepo: RepoInfluxClusterStatus.NewControllerRepository(&influxCfg),
		influxPodRepo:        RepoInfluxClusterStatus.NewPodRepository(&influxCfg),

		clusterUID: clusterUID,
	}
}

func (p ControllerMetrics) CreateMetrics(ctx context.Context, m DaoMetricTypes.ControllerMetricMap) error {
	return errors.New("not implemented")
}

// ListMetrics returns controller metrics, if length of req.ControllerObjectMetas equals 0, fetch all kinds of controller metrics in the cluster
// otherwise listing controller metadatas by function listControllerMetasFromRequest and fetching the returned controllers' metrics.
func (p ControllerMetrics) ListMetrics(ctx context.Context, req DaoMetricTypes.ListControllerMetricsRequest) (DaoMetricTypes.ControllerMetricMap, error) {

	options := []DBCommon.Option{
		DBCommon.StartTime(req.StartTime),
		DBCommon.EndTime(req.EndTime),
		DBCommon.StepTime(req.StepTime),
		DBCommon.AggregateOverTimeFunc(req.AggregateOverTimeFunction),
	}

	var metricMap DaoMetricTypes.ControllerMetricMap
	var err error
	controllerMetas, err := p.listControllerMetasFromRequest(ctx, req)
	if err != nil {
		return DaoMetricTypes.ControllerMetricMap{}, errors.Wrap(err, "list controller metadatas from request failed")
	}
	controllerMetas = p.filterObjectMetaByClusterUID(p.clusterUID, controllerMetas)
	if len(controllerMetas) == 0 {
		return DaoMetricTypes.ControllerMetricMap{}, nil
	}
	metricMap, err = p.getControllerMetricMapByObjectMetas(ctx, controllerMetas, options...)
	if err != nil {
		return DaoMetricTypes.ControllerMetricMap{}, errors.Wrap(err, "list controller metrics failed")
	}
	metricMap.SortByTimestamp(req.QueryCondition.TimestampOrder)
	metricMap.Limit(req.QueryCondition.Limit)
	return metricMap, nil
}

func (p *ControllerMetrics) listControllerMetasFromRequest(ctx context.Context, req DaoMetricTypes.ListControllerMetricsRequest) ([]DaoMetricTypes.ControllerObjectMeta, error) {

	controllers, err := p.influxControllerRepo.ListControllers(DaoClusterStatusTypes.ListControllersRequest{
		ObjectMeta: req.ObjectMetas,
		Kind:       strings.ToUpper(req.Kind),
	})
	if err != nil {
		return nil, errors.Wrap(err, "list controller metadatas failed")
	}
	metas := make([]DaoMetricTypes.ControllerObjectMeta, len(controllers))
	for i, controller := range controllers {
		metas[i] = DaoMetricTypes.ControllerObjectMeta{
			ObjectMeta: controller.ObjectMeta,
			Kind:       controller.Kind,
		}
	}
	return metas, nil
}

func (p *ControllerMetrics) getControllerMetricMapByObjectMetas(ctx context.Context, controllerMetas []DaoMetricTypes.ControllerObjectMeta, options ...DBCommon.Option) (DaoMetricTypes.ControllerMetricMap, error) {
	scope.Debugf("getControllerMetricMapByObjectMetas: controllerMetas: %+v", controllerMetas)

	metricMap := DaoMetricTypes.NewControllerMetricMap()
	metricChan := make(chan DaoMetricTypes.ControllerMetric)
	producerWG := errgroup.Group{}
	for _, controllerMeta := range controllerMetas {
		copyControllerMeta := controllerMeta
		producerWG.Go(func() error {
			m, err := p.getControllerMetric(ctx, copyControllerMeta, options...)
			if err != nil {
				return errors.Wrap(err, "get controller metric failed")
			}
			metricChan <- m
			return nil
		})
	}
	consumerWG := errgroup.Group{}
	consumerWG.Go(func() error {
		for m := range metricChan {
			copyM := m
			metricMap.AddControllerMetric(&copyM)
		}
		return nil
	})

	err := producerWG.Wait()
	close(metricChan)
	if err != nil {
		return metricMap, err
	}

	consumerWG.Wait()

	return metricMap, nil
}

func (p *ControllerMetrics) getControllerMetric(ctx context.Context, controllerMeta DaoMetricTypes.ControllerObjectMeta, options ...DBCommon.Option) (DaoMetricTypes.ControllerMetric, error) {

	emptyControllerMetric := DaoMetricTypes.ControllerMetric{
		ObjectMeta: controllerMeta,
	}

	pods, err := p.listPodMetasByControllerObjectMeta(ctx, controllerMeta)
	if err != nil {
		return emptyControllerMetric, errors.Wrap(err, "list monitored pods failed")
	} else if len(pods) == 0 {
		return emptyControllerMetric, nil
	}

	namespace := pods[0].Namespace
	podNames := make([]string, len(pods))
	for i, pod := range pods {
		podNames[i] = pod.Name
	}
	metricMap := DaoMetricTypes.NewControllerMetricMap()
	metricChan := make(chan DaoMetricTypes.ControllerMetric)
	producerWG := errgroup.Group{}
	producerWG.Go(func() error {
		podCPUUsageRepo := RepoPromthMetric.NewPodCPUUsageRepositoryWithConfig(p.PrometheusConfig)
		podCPUMetricEntities, err := podCPUUsageRepo.ListPodCPUUsageMillicoresEntitiesBySummingPodMetrics(ctx, namespace, podNames, options...)
		if err != nil {
			return errors.Wrap(err, "list sum of pod cpu usage metrics failed")
		}
		for _, e := range podCPUMetricEntities {
			controllerEntity := EntityPromthMetric.ControllerCPUUsageMillicoresEntity{
				NamespaceName:  controllerMeta.Namespace,
				ControllerName: controllerMeta.Name,
				ControllerKind: controllerMeta.Kind,
				Samples:        e.Samples,
			}
			m := controllerEntity.ControllerMetric()
			m.ObjectMeta = controllerMeta
			metricChan <- m
		}
		return nil
	})
	producerWG.Go(func() error {
		podMemoryUsageRepo := RepoPromthMetric.NewPodMemoryUsageRepositoryWithConfig(p.PrometheusConfig)
		podMemoryMetricEntities, err := podMemoryUsageRepo.ListPodMemoryUsageBytesEntityBySummingPodMetrics(ctx, namespace, podNames, options...)
		if err != nil {
			return errors.Wrap(err, "list sum of pod memory usage metrics failed")
		}
		for _, e := range podMemoryMetricEntities {
			controllerEntity := EntityPromthMetric.ControllerMemoryUsageBytesEntity{
				NamespaceName:  controllerMeta.Namespace,
				ControllerName: controllerMeta.Name,
				ControllerKind: controllerMeta.Kind,
				Samples:        e.Samples,
			}
			m := controllerEntity.ControllerMetric()
			m.ObjectMeta = controllerMeta
			metricChan <- m
		}
		return nil
	})

	consumerWG := errgroup.Group{}
	consumerWG.Go(func() error {
		for m := range metricChan {
			copyM := m
			metricMap.AddControllerMetric(&copyM)
		}
		return nil
	})

	err = producerWG.Wait()
	close(metricChan)
	if err != nil {
		return DaoMetricTypes.ControllerMetric{}, err
	}

	consumerWG.Wait()
	metric, exist := metricMap.MetricMap[controllerMeta]
	if !exist || metric == nil {
		return emptyControllerMetric, nil
	}
	return *metric, nil
}

func (p *ControllerMetrics) listPodMetasByControllerObjectMeta(ctx context.Context, controllerObjectMeta DaoMetricTypes.ControllerObjectMeta) ([]metadata.ObjectMeta, error) {

	pods, err := p.influxPodRepo.ListPods(DaoClusterStatusTypes.ListPodsRequest{
		ObjectMeta: []metadata.ObjectMeta{controllerObjectMeta.ObjectMeta},
		Kind:       strings.ToUpper(controllerObjectMeta.Kind),
	})
	if err != nil {
		return nil, errors.Wrap(err, "list pod metadatas by application failed")
	}
	podMetas := make([]metadata.ObjectMeta, len(pods))
	for i, pod := range pods {
		podMetas[i] = *pod.ObjectMeta
	}

	return podMetas, nil
}

func (p *ControllerMetrics) filterObjectMetaByClusterUID(clusterUID string, objectMetas []DaoMetricTypes.ControllerObjectMeta) []DaoMetricTypes.ControllerObjectMeta {
	newObjectMetas := make([]DaoMetricTypes.ControllerObjectMeta, 0, len(objectMetas))
	for _, objectMeta := range objectMetas {
		if objectMeta.ClusterName == clusterUID {
			newObjectMetas = append(newObjectMetas, objectMeta)
		}
	}
	return newObjectMetas
}