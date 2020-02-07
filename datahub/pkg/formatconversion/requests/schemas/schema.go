package schemas

import (
	InternalSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewSchema(schema *ApiSchema.Schema) *InternalSchema.Schema {
	if schema != nil {
		s := InternalSchema.Schema{}
		s.SchemaMeta = NewSchemaMeta(schema.SchemaMeta)
		for _, measurement := range schema.Measurements {
			s.Measurements = append(s.Measurements, NewMeasurement(measurement))
		}
		return &s
	}
	return nil
}
