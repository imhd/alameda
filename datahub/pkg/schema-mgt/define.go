package schema_mgt

import (
	Schema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

// Table enumerator
const (
	ScopeUndefined Schema.Scope = iota
	Application
	Metric
	Planning
	Prediction
	Recommendation
	Resource
)

// Metric type enumerator
const (
	MetricTypeUndefined Schema.MetricType = iota
	CPUUsageSecondsPercentage
	MemoryUsageBytes
	PowerUsageWatts
	TemperatureCelsius
	DutyCycle
	CurrentOffset
)

// Resource boundary enumerator
const (
	ResourceBoundaryUndefined Schema.ResourceBoundary = iota
	ResourceRaw
	ResourceUpperBound
	ResourceLowerBound
)

// Resource quota enumerator
const (
	ResourceQuotaUndefined Schema.ResourceQuota = iota
	ResourceLimit
	ResourceRequest
	ResourceInitialLimit
	ResourceInitialRequest
)

var MeasurementNameMap = map[Schema.Scope]string{
	Application:    "application",
	Metric:         "metric",
	Planning:       "planning",
	Prediction:     "prediction",
	Recommendation: "recommendation",
	Resource:       "resource",
}

var MeasurementSchemaNameMap = map[Schema.Scope]string{
	Application:    "application_schema",
	Metric:         "metric_schema",
	Planning:       "planning_schema",
	Prediction:     "prediction_schema",
	Recommendation: "recommendation_schema",
	Resource:       "resource_schema",
}

var DatabaseNameMap = map[Schema.Scope]string{
	Application:    "alameda_application",
	Metric:         "alameda_metric",
	Planning:       "alameda_planning",
	Prediction:     "alameda_prediction",
	Recommendation: "alameda_recommendation",
	Resource:       "alameda_resource",
}
