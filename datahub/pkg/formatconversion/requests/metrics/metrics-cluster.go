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
	"github.com/pkg/errors"
)

type CreateClusterMetricsRequestExtended struct {
	ApiMetrics.CreateClusterMetricsRequest
}

func (r *CreateClusterMetricsRequestExtended) Validate() error {
	for _, m := range r.GetClusterMetrics() {
		if m == nil || m.ObjectMeta == nil || m.ObjectMeta.Name == "" {
			return errors.Errorf(`must provide "Name" in ObjectMeta`)
		}
	}
	return nil
}

func (r *CreateClusterMetricsRequestExtended) ProduceMetrics() DaoMetricTypes.ClusterMetricMap {
	metricMap := DaoMetricTypes.NewClusterMetricMap()

	for _, clusterMetric := range r.GetClusterMetrics() {
		if clusterMetric == nil {
			continue
		}
		m := DaoMetricTypes.NewClusterMetric()
		m.ObjectMeta = resources.NewObjectMeta(clusterMetric.GetObjectMeta())

		for _, data := range clusterMetric.GetMetricData() {
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
				m.AddSample(metricType, sample)
			}
		}

		metricMap.AddClusterMetric(m)
	}

	return metricMap
}

type ListClusterMetricsRequestExtended struct {
	Request *ApiMetrics.ListClusterMetricsRequest
}

func (r *ListClusterMetricsRequestExtended) Validate() error {
	return nil
}

func (r *ListClusterMetricsRequestExtended) SetDefaultWithMetricsDBType(dbType MetricsDBType) {
	q := normalizeListMetricsRequestQueryConditionWthMetricsDBType(*r.Request.QueryCondition, dbType)
	q.TimeRange.AggregateFunction = ApiCommon.TimeRange_AVG
	r.Request.QueryCondition = &q
}

func (r *ListClusterMetricsRequestExtended) ProduceRequest() DaoMetricTypes.ListClusterMetricsRequest {
	request := DaoMetricTypes.ListClusterMetricsRequest{}
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	objectMetas := make([]metadata.ObjectMeta, len(r.Request.GetObjectMeta()))
	for i, objectMeta := range r.Request.GetObjectMeta() {
		copyObjectMeta := objectMeta
		o := resources.NewObjectMeta(copyObjectMeta)
		if o.IsEmpty() {
			objectMetas = nil
			break
		}
		objectMetas[i] = o
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
