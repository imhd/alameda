package metrics

import (
	FormatEnum "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	Log "github.com/containers-ai/alameda/pkg/utils/log"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type MetricsDBType = string

const (
	MetricsDBTypePromethues MetricsDBType = "prometheus"
	MetricsDBTypeInfluxdb   MetricsDBType = "influxdb"
)

var (
	scope = Log.RegisterScope("request-extend-applications", "datahub(request-extend-applications) log", 0)
)

var MetricTypeNameMap = map[ApiCommon.MetricType]FormatEnum.MetricType{
	ApiCommon.MetricType_CPU_USAGE_SECONDS_PERCENTAGE: FormatEnum.MetricTypeCPUUsageSecondsPercentage,
	ApiCommon.MetricType_MEMORY_USAGE_BYTES:           FormatEnum.MetricTypeMemoryUsageBytes,
	ApiCommon.MetricType_POWER_USAGE_WATTS:            FormatEnum.MetricTypePowerUsageWatts,
	ApiCommon.MetricType_TEMPERATURE_CELSIUS:          FormatEnum.MetricTypeTemperatureCelsius,
	ApiCommon.MetricType_DUTY_CYCLE:                   FormatEnum.MetricTypeDutyCycle,
}

func normalizeListMetricsRequestQueryConditionWthMetricsDBType(q ApiCommon.QueryCondition, dbType MetricsDBType) ApiCommon.QueryCondition {

	t := q.TimeRange
	if t == nil {
		t = &ApiCommon.TimeRange{}
	}
	normalizeT := normalizeListMetricsRequestTimeRangeByMetricsDBType(*t, dbType)
	q.TimeRange = &normalizeT

	return q
}

func normalizeListMetricsRequestTimeRange(t ApiCommon.TimeRange) ApiCommon.TimeRange {

	defaultStartTime := timestamp.Timestamp{}
	defaultEndTime := *ptypes.TimestampNow()
	defaultStep := duration.Duration{
		Seconds: 30,
	}

	if t.StartTime == nil {
		t.StartTime = &defaultStartTime
	}
	if t.EndTime == nil {
		t.EndTime = &defaultEndTime
	}
	if t.Step == nil {
		t.Step = &defaultStep
	}

	return t
}

func normalizeListMetricsRequestTimeRangeByMetricsDBType(t ApiCommon.TimeRange, metricsDBType MetricsDBType) ApiCommon.TimeRange {

	t = normalizeListMetricsRequestTimeRange(t)

	t.StartTime = &timestamp.Timestamp{
		Seconds: t.StartTime.Seconds - t.StartTime.Seconds%t.Step.Seconds,
	}

	switch metricsDBType {
	case MetricsDBTypePromethues:
		t.EndTime = &timestamp.Timestamp{
			Seconds: t.EndTime.Seconds - t.EndTime.Seconds%t.Step.Seconds,
		}
	case MetricsDBTypeInfluxdb:
		t.EndTime = &timestamp.Timestamp{
			Seconds: t.EndTime.Seconds - t.EndTime.Seconds%t.Step.Seconds - 1,
		}
	}
	return t
}
