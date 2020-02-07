package enumconv

import (
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schema-mgt"
	Schemas "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var MetricTypeNameMap = map[Schemas.MetricType]ApiCommon.MetricType{
	SchemaMgt.MetricTypeUndefined:       ApiCommon.MetricType_METRICS_TYPE_UNDEFINED,
	SchemaMgt.CPUUsageSecondsPercentage: ApiCommon.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
	SchemaMgt.MemoryUsageBytes:          ApiCommon.MetricType_MEMORY_USAGE_BYTES,
	SchemaMgt.PowerUsageWatts:           ApiCommon.MetricType_POWER_USAGE_WATTS,
	SchemaMgt.TemperatureCelsius:        ApiCommon.MetricType_TEMPERATURE_CELSIUS,
	SchemaMgt.DutyCycle:                 ApiCommon.MetricType_DUTY_CYCLE,
	SchemaMgt.CurrentOffset:             ApiCommon.MetricType_CURRENT_OFFSET,
}

var ResourceBoundaryNameMap = map[Schemas.ResourceBoundary]ApiCommon.ResourceBoundary{
	SchemaMgt.ResourceBoundaryUndefined: ApiCommon.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED,
	SchemaMgt.ResourceRaw:               ApiCommon.ResourceBoundary_RESOURCE_RAW,
	SchemaMgt.ResourceUpperBound:        ApiCommon.ResourceBoundary_RESOURCE_UPPER_BOUND,
	SchemaMgt.ResourceLowerBound:        ApiCommon.ResourceBoundary_RESOURCE_LOWER_BOUND,
}

var ResourceQuotaNameMap = map[Schemas.ResourceQuota]ApiCommon.ResourceQuota{
	SchemaMgt.ResourceQuotaUndefined: ApiCommon.ResourceQuota_RESOURCE_QUOTA_UNDEFINED,
	SchemaMgt.ResourceLimit:          ApiCommon.ResourceQuota_RESOURCE_LIMIT,
	SchemaMgt.ResourceRequest:        ApiCommon.ResourceQuota_RESOURCE_REQUEST,
	SchemaMgt.ResourceInitialLimit:   ApiCommon.ResourceQuota_RESOURCE_INITIAL_LIMIT,
	SchemaMgt.ResourceInitialRequest: ApiCommon.ResourceQuota_RESOURCE_INITIAL_REQUEST,
}
