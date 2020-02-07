package enumconv

import (
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schema-mgt"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	Schemas "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var AggregateFunctionNameMap = map[ApiCommon.TimeRange_AggregateFunction]DBCommon.AggregateFunction{
	ApiCommon.TimeRange_NONE: DBCommon.None,
	ApiCommon.TimeRange_MAX:  DBCommon.MaxOverTime,
	ApiCommon.TimeRange_AVG:  DBCommon.AvgOverTime,
}

var QueryConditionOrderNameMap = map[ApiCommon.QueryCondition_Order]DBCommon.Order{
	ApiCommon.QueryCondition_NONE: DBCommon.NoneOrder,
	ApiCommon.QueryCondition_ASC:  DBCommon.Asc,
	ApiCommon.QueryCondition_DESC: DBCommon.Desc,
}

var DataTypeNameMap = map[ApiCommon.DataType]DBCommon.DataType{
	ApiCommon.DataType_DATATYPE_UNDEFINED: DBCommon.Invalid,
	ApiCommon.DataType_DATATYPE_BOOL:      DBCommon.Bool,
	ApiCommon.DataType_DATATYPE_INT:       DBCommon.Int,
	ApiCommon.DataType_DATATYPE_INT8:      DBCommon.Int8,
	ApiCommon.DataType_DATATYPE_INT16:     DBCommon.Int16,
	ApiCommon.DataType_DATATYPE_INT32:     DBCommon.Int32,
	ApiCommon.DataType_DATATYPE_INT64:     DBCommon.Int64,
	ApiCommon.DataType_DATATYPE_UINT:      DBCommon.Uint,
	ApiCommon.DataType_DATATYPE_UINT8:     DBCommon.Uint8,
	ApiCommon.DataType_DATATYPE_UINT16:    DBCommon.Uint16,
	ApiCommon.DataType_DATATYPE_UINT32:    DBCommon.Uint32,
	ApiCommon.DataType_DATATYPE_UTIN64:    DBCommon.Uint64,
	ApiCommon.DataType_DATATYPE_FLOAT32:   DBCommon.Float32,
	ApiCommon.DataType_DATATYPE_FLOAT64:   DBCommon.Float64,
	ApiCommon.DataType_DATATYPE_STRING:    DBCommon.String,
}

var MetricTypeNameMap = map[ApiCommon.MetricType]Schemas.MetricType{
	ApiCommon.MetricType_METRICS_TYPE_UNDEFINED:       SchemaMgt.MetricTypeUndefined,
	ApiCommon.MetricType_CPU_USAGE_SECONDS_PERCENTAGE: SchemaMgt.CPUUsageSecondsPercentage,
	ApiCommon.MetricType_MEMORY_USAGE_BYTES:           SchemaMgt.MemoryUsageBytes,
	ApiCommon.MetricType_POWER_USAGE_WATTS:            SchemaMgt.PowerUsageWatts,
	ApiCommon.MetricType_TEMPERATURE_CELSIUS:          SchemaMgt.TemperatureCelsius,
	ApiCommon.MetricType_DUTY_CYCLE:                   SchemaMgt.DutyCycle,
	ApiCommon.MetricType_CURRENT_OFFSET:               SchemaMgt.CurrentOffset,
}

var ResourceBoundaryNameMap = map[ApiCommon.ResourceBoundary]Schemas.ResourceBoundary{
	ApiCommon.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED: SchemaMgt.ResourceBoundaryUndefined,
	ApiCommon.ResourceBoundary_RESOURCE_RAW:                SchemaMgt.ResourceRaw,
	ApiCommon.ResourceBoundary_RESOURCE_UPPER_BOUND:        SchemaMgt.ResourceUpperBound,
	ApiCommon.ResourceBoundary_RESOURCE_LOWER_BOUND:        SchemaMgt.ResourceLowerBound,
}

var ResourceQuotaNameMap = map[ApiCommon.ResourceQuota]Schemas.ResourceQuota{
	ApiCommon.ResourceQuota_RESOURCE_QUOTA_UNDEFINED: SchemaMgt.ResourceQuotaUndefined,
	ApiCommon.ResourceQuota_RESOURCE_LIMIT:           SchemaMgt.ResourceLimit,
	ApiCommon.ResourceQuota_RESOURCE_REQUEST:         SchemaMgt.ResourceRequest,
	ApiCommon.ResourceQuota_RESOURCE_INITIAL_LIMIT:   SchemaMgt.ResourceInitialLimit,
	ApiCommon.ResourceQuota_RESOURCE_INITIAL_REQUEST: SchemaMgt.ResourceInitialRequest,
}
