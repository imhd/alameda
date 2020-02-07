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

type CreateNamespaceMetricsRequestExtended struct {
	ApiMetrics.CreateNamespaceMetricsRequest
}

func (r *CreateNamespaceMetricsRequestExtended) Validate() error {
	for _, m := range r.GetNamespaceMetrics() {
		if m == nil || m.ObjectMeta == nil || m.ObjectMeta.Name == "" || m.ObjectMeta.ClusterName == "" {
			return errors.Errorf(`must provide "Name" and "ClusterName" in ObjectMeta`)
		}
	}
	return nil
}

func (r *CreateNamespaceMetricsRequestExtended) ProduceMetrics() DaoMetricTypes.NamespaceMetricMap {
	metricMap := DaoMetricTypes.NewNamespaceMetricMap()

	for _, namespaceMetric := range r.GetNamespaceMetrics() {
		if namespaceMetric == nil {
			continue
		}
		m := DaoMetricTypes.NewNamespaceMetric()
		m.ObjectMeta = resources.NewObjectMeta(namespaceMetric.GetObjectMeta())

		for _, data := range namespaceMetric.GetMetricData() {
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

		metricMap.AddNamespaceMetric(m)
	}

	return metricMap
}

type ListNamespaceMetricsRequestExtended struct {
	Request *ApiMetrics.ListNamespaceMetricsRequest
}

func (r *ListNamespaceMetricsRequestExtended) Validate() error {
	return nil
}

func (r *ListNamespaceMetricsRequestExtended) SetDefaultWithMetricsDBType(dbType MetricsDBType) {
	q := normalizeListMetricsRequestQueryConditionWthMetricsDBType(*r.Request.QueryCondition, dbType)
	q.TimeRange.AggregateFunction = ApiCommon.TimeRange_AVG
	r.Request.QueryCondition = &q
}

func (r *ListNamespaceMetricsRequestExtended) ProduceRequest() DaoMetricTypes.ListNamespaceMetricsRequest {
	request := DaoMetricTypes.ListNamespaceMetricsRequest{}
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
