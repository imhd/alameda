package schema_mgt

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func DefaultSchemasInit() {
	schemaMgt := NewSchemaManagement()
	schemaList := make([]*schemas.Schema, 0)

	// Application
	schemaList = append(schemaList, genSchemaApplicationKafkaTopic())
	schemaList = append(schemaList, genSchemaApplicationKafkaCG())

	// Metric
	schemaList = append(schemaList, genSchemaMetricKafkaTopic())
	schemaList = append(schemaList, genSchemaMetricKafkaCG())

	// Prediction
	schemaList = append(schemaList, genSchemaPredictionKafkaTopic())
	schemaList = append(schemaList, genSchemaPredictionKafkaCG())

	// Recommendation
	schemaList = append(schemaList, genSchemaRecommendationKafkaCG())

	schemaMgt.AddSchemas(schemaList)
	schemaMgt.Flush() // TODO: ONLY DO one time !!!
}

func genSchemaApplicationKafkaTopic() *schemas.Schema {
	// Application: kafka topic
	schema := schemas.NewSchema(Application, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka_topic", MetricTypeUndefined, ResourceBoundaryUndefined, ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func genSchemaApplicationKafkaCG() *schemas.Schema {
	// Application: kafka consumer group
	schema := schemas.NewSchema(Application, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka_consumer_group", MetricTypeUndefined, ResourceBoundaryUndefined, ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("topic_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_namespace", true, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_name", true, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_kind", true, schemas.Field, common.String)
	measurement.AddColumn("resource_custom_name", true, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_replicas", true, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_spec_replicas", true, schemas.Field, common.String)
	measurement.AddColumn("policy", true, schemas.Field, common.String)
	measurement.AddColumn("enable_execution", true, schemas.Field, common.Bool)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func genSchemaMetricKafkaTopic() *schemas.Schema {
	// Metric: kafka topic
	schema := schemas.NewSchema(Metric, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka_topic_partition_current_offset", CurrentOffset, ResourceBoundaryUndefined, ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func genSchemaMetricKafkaCG() *schemas.Schema {
	// Metric: kafka consumer group
	schema := schemas.NewSchema(Metric, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka_consumer_group_current_offset", CurrentOffset, ResourceBoundaryUndefined, ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("topic_name", true, schemas.Tag, common.String)
	measurement.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func genSchemaPredictionKafkaTopic() *schemas.Schema {
	// Prediction: kafka topic
	schema := schemas.NewSchema(Prediction, "kafka", "topic")

	// Raw
	raw := schemas.NewMeasurement("kafka_topic_partition_current_offset", CurrentOffset, ResourceRaw, ResourceQuotaUndefined)
	raw.AddColumn("name", true, schemas.Tag, common.String)
	raw.AddColumn("namespace", true, schemas.Tag, common.String)
	raw.AddColumn("cluster_name", true, schemas.Tag, common.String)
	raw.AddColumn("granularity", true, schemas.Tag, common.String)
	raw.AddColumn("model_id", true, schemas.Field, common.String)
	raw.AddColumn("prediction_id", true, schemas.Field, common.String)
	raw.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, raw)

	// Upper bound
	upperBound := schemas.NewMeasurement("kafka_topic_partition_current_offset_upper_bound", CurrentOffset, ResourceUpperBound, ResourceQuotaUndefined)
	upperBound.AddColumn("name", true, schemas.Tag, common.String)
	upperBound.AddColumn("namespace", true, schemas.Tag, common.String)
	upperBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperBound.AddColumn("granularity", true, schemas.Tag, common.String)
	upperBound.AddColumn("model_id", true, schemas.Field, common.String)
	upperBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperBound.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, upperBound)

	// Lower bound
	lowerBound := schemas.NewMeasurement("kafka_topic_partition_current_offset_lower_bound", CurrentOffset, ResourceLowerBound, ResourceQuotaUndefined)
	lowerBound.AddColumn("name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerBound.AddColumn("model_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, lowerBound)

	return schema
}

func genSchemaPredictionKafkaCG() *schemas.Schema {
	// Metric: kafka consumer group
	schema := schemas.NewSchema(Prediction, "kafka", "consumer_group")

	// Raw
	raw := schemas.NewMeasurement("kafka_consumer_group_current_offset", CurrentOffset, ResourceRaw, ResourceQuotaUndefined)
	raw.AddColumn("name", true, schemas.Tag, common.String)
	raw.AddColumn("namespace", true, schemas.Tag, common.String)
	raw.AddColumn("cluster_name", true, schemas.Tag, common.String)
	raw.AddColumn("topic_name", true, schemas.Tag, common.String)
	raw.AddColumn("granularity", true, schemas.Tag, common.String)
	raw.AddColumn("model_id", true, schemas.Field, common.String)
	raw.AddColumn("prediction_id", true, schemas.Field, common.String)
	raw.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, raw)

	// Upper bound
	upperBound := schemas.NewMeasurement("kafka_consumer_group_current_offset_upper_bound", CurrentOffset, ResourceUpperBound, ResourceQuotaUndefined)
	upperBound.AddColumn("name", true, schemas.Tag, common.String)
	upperBound.AddColumn("namespace", true, schemas.Tag, common.String)
	upperBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	upperBound.AddColumn("topic_name", true, schemas.Tag, common.String)
	upperBound.AddColumn("granularity", true, schemas.Tag, common.String)
	upperBound.AddColumn("model_id", true, schemas.Field, common.String)
	upperBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	upperBound.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, upperBound)

	// Lower bound
	lowerBound := schemas.NewMeasurement("kafka_consumer_group_current_offset_lower_bound", CurrentOffset, ResourceLowerBound, ResourceQuotaUndefined)
	lowerBound.AddColumn("name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("namespace", true, schemas.Tag, common.String)
	lowerBound.AddColumn("cluster_name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("topic_name", true, schemas.Tag, common.String)
	lowerBound.AddColumn("granularity", true, schemas.Tag, common.String)
	lowerBound.AddColumn("model_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("prediction_id", true, schemas.Field, common.String)
	lowerBound.AddColumn("value", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, lowerBound)

	return schema
}

func genSchemaRecommendationKafkaCG() *schemas.Schema {
	// Recommendation: kafka consumer group
	schema := schemas.NewSchema(Recommendation, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka_consumer_group", MetricTypeUndefined, ResourceBoundaryUndefined, ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("create_time", true, schemas.Field, common.String)
	measurement.AddColumn("current_replicas", true, schemas.Field, common.Int64)
	measurement.AddColumn("desired_replicas", true, schemas.Field, common.Int64)
	measurement.AddColumn("execution_time", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
