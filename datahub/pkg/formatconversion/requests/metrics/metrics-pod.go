package metrics

import (
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/common"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/resources"
	FormatTypes "github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	"github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	ApiMetrics "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/metrics"
	"github.com/golang/protobuf/ptypes"
)

type CreatePodMetricsRequestExtended struct {
	ApiMetrics.CreatePodMetricsRequest
}

func (r *CreatePodMetricsRequestExtended) Validate() error {
	return nil
}

func (r *CreatePodMetricsRequestExtended) ProduceMetrics() DaoMetricTypes.PodMetricMap {
	podMetricMap := DaoMetricTypes.NewPodMetricMap()

	rateRange := int64(5)
	if r.GetRateRange() != 0 {
		rateRange = int64(r.GetRateRange())
	}

	for _, pod := range r.GetPodMetrics() {
		podMetric := DaoMetricTypes.NewPodMetric()
		podMetric.ObjectMeta = resources.NewObjectMeta(pod.GetObjectMeta())
		podMetric.RateRange = rateRange

		for _, container := range pod.GetContainerMetrics() {
			containerMetric := DaoMetricTypes.NewContainerMetric()
			containerMetric.ObjectMeta.ObjectMeta = resources.NewObjectMeta(pod.GetObjectMeta())
			containerMetric.ObjectMeta.PodName = podMetric.ObjectMeta.Name
			containerMetric.ObjectMeta.Name = container.GetName()
			containerMetric.RateRange = rateRange

			for _, data := range container.GetMetricData() {
				metricType := MetricTypeNameMap[data.GetMetricType()]
				for _, sample := range data.GetData() {
					timestamp, err := ptypes.Timestamp(sample.GetTime())
					if err != nil {
						scope.Error(" failed: " + err.Error())
					}
					sample := FormatTypes.Sample{
						Timestamp: timestamp,
						Value:     sample.GetNumValue(),
					}
					containerMetric.AddSample(metricType, sample)
				}
			}

			podMetric.ContainerMetricMap.AddContainerMetric(containerMetric)
		}

		podMetricMap.AddPodMetric(podMetric)
	}

	return podMetricMap
}

type ListPodMetricsRequestExtended struct {
	Request *ApiMetrics.ListPodMetricsRequest
}

func (r *ListPodMetricsRequestExtended) Validate() error {
	return nil
}

func (r *ListPodMetricsRequestExtended) SetDefaultWithMetricsDBType(dbType MetricsDBType) {
	q := normalizeListMetricsRequestQueryConditionWthMetricsDBType(*r.Request.QueryCondition, dbType)
	switch q.TimeRange.Step.Seconds {
	case 30:
		q.TimeRange.AggregateFunction = ApiCommon.TimeRange_MAX
	default:
		q.TimeRange.AggregateFunction = ApiCommon.TimeRange_AVG
	}
	r.Request.QueryCondition = &q
}

func (r *ListPodMetricsRequestExtended) ProduceRequest() DaoMetricTypes.ListPodMetricsRequest {
	request := DaoMetricTypes.NewListPodMetricsRequest()
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	request.RateRange = 5
	if r.Request.GetRateRange() != 0 {
		request.RateRange = int64(r.Request.GetRateRange())
	}
	objectMetas := make([]*metadata.ObjectMeta, len(r.Request.GetObjectMeta()))
	for i, objectMeta := range r.Request.GetObjectMeta() {
		copyObjectMeta := objectMeta
		o := resources.NewObjectMeta(copyObjectMeta)
		if o.IsEmpty() {
			objectMetas = nil
			break
		}
		objectMetas[i] = &o
	}
	request.ObjectMetas = objectMetas
	metricTypes := make([]FormatEnum.MetricType, 0)
	for _, metricType := range r.Request.GetMetricTypes() {
		metricTypes = append(metricTypes, MetricTypeNameMap[metricType])
	}
	if len(metricTypes) == 0 {
		metricTypes = append(metricTypes, MetricTypeNameMap[ApiCommon.MetricType_CPU_USAGE_SECONDS_PERCENTAGE])
		metricTypes = append(metricTypes, MetricTypeNameMap[ApiCommon.MetricType_MEMORY_USAGE_BYTES])
	}
	request.MetricTypes = metricTypes
	return request
}
