package schemas

import (
	InternalSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewSchemaMeta(schemaMeta *ApiSchema.SchemaMeta) *InternalSchema.SchemaMeta {
	if schemaMeta != nil {
		meta := InternalSchema.SchemaMeta{}
		meta.Scope = InternalSchema.Scope(schemaMeta.GetScope())
		meta.Category = schemaMeta.GetCategory()
		meta.Type = schemaMeta.GetType()
		return &meta
	}
	return nil
}
