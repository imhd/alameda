package metric

import (
	"fmt"
	EntityInfluxMetricContainer "github.com/containers-ai/alameda/datahub/pkg/entity/influxdb/metric/container"
	Metric "github.com/containers-ai/alameda/datahub/pkg/metric"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var (
	scope = log.RegisterScope("cluster_status_db_measurement", "cluster_status DB measurement", 0)
)

// ContainerRepository is used to operate node measurement of cluster_status database
type ContainerRepository struct {
	influxDB *InternalInflux.InfluxClient
}

// NewContainerRepositoryWithConfig New container repository with influxDB configuration
func NewContainerRepositoryWithConfig(influxDBCfg InternalInflux.Config) *ContainerRepository {
	return &ContainerRepository{
		influxDB: &InternalInflux.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

// ListContainerPredictionsByRequest list containers' prediction from influxDB
func (r *ContainerRepository) ListContainerMetrics(in *datahub_v1alpha1.ListPodMetricsRequest) ([]*datahub_v1alpha1.PodMetric, error) {
	podMetricList := make([]*datahub_v1alpha1.PodMetric, 0)

	groupByTime := fmt.Sprintf("%s(%ds)", EntityInfluxMetricContainer.PodTime, in.GetQueryCondition().GetTimeRange().GetStep().GetSeconds())
	selectedField := fmt.Sprintf("sum(%s) as %s", EntityInfluxMetricContainer.Value, EntityInfluxMetricContainer.Value)

	rateRange := in.GetRateRange()
	measurementName := EntityInfluxMetricContainer.MetricMeasurementName
	if rateRange != 0 && rateRange != 5 {
		measurementName = fmt.Sprintf("%s_%dm", measurementName, rateRange)
	}

	influxdbStatement := InternalInflux.Statement{
		Measurement:    InternalInflux.Measurement(measurementName),
		SelectedFields: []string{selectedField},
		GroupByTags:    []string{EntityInfluxMetricContainer.PodNamespace, EntityInfluxMetricContainer.PodName, EntityInfluxMetricContainer.Name, EntityInfluxMetricContainer.MetricType, groupByTime},
	}

	influxdbStatement.AppendWhereClause(EntityInfluxMetricContainer.PodNamespace, "=", in.GetNamespacedName().GetNamespace())
	influxdbStatement.AppendWhereClause(EntityInfluxMetricContainer.PodName, "=", in.GetNamespacedName().GetName())
	influxdbStatement.AppendWhereClauseWithTime(">=", in.GetQueryCondition().GetTimeRange().GetStartTime().GetSeconds())
	influxdbStatement.AppendWhereClauseWithTime("<=", in.GetQueryCondition().GetTimeRange().GetEndTime().GetSeconds())
	influxdbStatement.SetLimitClauseFromQueryCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()

	cmd := influxdbStatement.BuildQueryCmd()

	results, err := r.influxDB.QueryDB(cmd, EntityInfluxMetricContainer.MetricDatabaseName)
	if err != nil {
		return podMetricList, errors.Wrap(err, "list container prediction failed")
	}

	rows := InternalInflux.PackMap(results)

	podMetricList = r.getPodMetricsFromInfluxRows(rows)
	return podMetricList, nil
}

func (r *ContainerRepository) getPodMetricsFromInfluxRows(rows []*InternalInflux.InfluxRow) []*datahub_v1alpha1.PodMetric {
	podMap := map[string]*datahub_v1alpha1.PodMetric{}

	podContainerMap := map[string]*datahub_v1alpha1.ContainerMetric{}
	podContainerMetricMap := map[string]*datahub_v1alpha1.MetricData{}
	podContainerMetricSampleMap := map[string][]*datahub_v1alpha1.Sample{}

	for _, row := range rows {
		podNamespace := row.Tags[EntityInfluxMetricContainer.PodNamespace]
		podName := row.Tags[EntityInfluxMetricContainer.PodName]
		containerName := row.Tags[EntityInfluxMetricContainer.Name]
		metricType := row.Tags[EntityInfluxMetricContainer.MetricType]

		metricValue := datahub_v1alpha1.MetricType(datahub_v1alpha1.MetricType_value[metricType])
		switch metricType {
		case Metric.TypeContainerCPUUsageSecondsPercentage:
			metricValue = datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE
		case Metric.TypeContainerMemoryUsageBytes:
			metricValue = datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES
		}

		podMap[podNamespace+"|"+podName] = &datahub_v1alpha1.PodMetric{}
		podMap[podNamespace+"|"+podName].NamespacedName = &datahub_v1alpha1.NamespacedName{
			Namespace: podNamespace,
			Name:      podName,
		}

		podContainerMap[podNamespace+"|"+podName+"|"+containerName] = &datahub_v1alpha1.ContainerMetric{}
		podContainerMap[podNamespace+"|"+podName+"|"+containerName].Name = containerName

		metricKey := podNamespace + "|" + podName + "|" + containerName + "|" + metricType
		podContainerMetricMap[metricKey] = &datahub_v1alpha1.MetricData{}
		podContainerMetricMap[metricKey].MetricType = metricValue

		for _, data := range row.Data {
			t, _ := time.Parse(time.RFC3339, data[EntityInfluxMetricContainer.PodTime])
			value := data[EntityInfluxMetricContainer.Value]

			googleTimestamp, _ := ptypes.TimestampProto(t)

			tempSample := &datahub_v1alpha1.Sample{
				Time:     googleTimestamp,
				NumValue: value,
			}
			podContainerMetricSampleMap[metricKey] = append(podContainerMetricSampleMap[metricKey], tempSample)
		}
	}

	for k := range podContainerMetricMap {
		podNamespace := strings.Split(k, "|")[0]
		podName := strings.Split(k, "|")[1]
		containerName := strings.Split(k, "|")[2]
		metricType := strings.Split(k, "|")[3]

		containerKey := podNamespace + "|" + podName + "|" + containerName
		metricKey := podNamespace + "|" + podName + "|" + containerName + "|" + metricType

		podContainerMetricMap[metricKey].Data = podContainerMetricSampleMap[metricKey]
		podContainerMap[containerKey].MetricData = append(podContainerMap[containerKey].MetricData, podContainerMetricMap[metricKey])
	}

	for k := range podContainerMap {
		podNamespace := strings.Split(k, "|")[0]
		podName := strings.Split(k, "|")[1]
		containerName := strings.Split(k, "|")[2]

		podKey := podNamespace + "|" + podName
		containerKey := podNamespace + "|" + podName + "|" + containerName

		podMap[podKey].ContainerMetrics = append(podMap[podKey].ContainerMetrics, podContainerMap[containerKey])
	}

	podList := make([]*datahub_v1alpha1.PodMetric, 0)
	for k := range podMap {
		podList = append(podList, podMap[k])
	}

	return podList
}
