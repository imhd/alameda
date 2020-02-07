package schemas

import (
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/schemas/types"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

type CreateSchemasRequestExtended struct {
	*schemas.CreateSchemasRequest
}

type ListSchemasRequestExtended struct {
	*schemas.ListSchemasRequest
}

type DeleteSchemasRequestExtended struct {
	*schemas.DeleteSchemasRequest
}

func (p *CreateSchemasRequestExtended) Validate() error {
	return nil
}

func (p *CreateSchemasRequestExtended) ProduceRequest() *types.CreateSchemasRequest {
	request := types.CreateSchemasRequest{}
	for _, schema := range p.GetSchemas() {
		request.Schemas = append(request.Schemas, NewSchema(schema))
	}
	return &request
}

func (p *ListSchemasRequestExtended) Validate() error {
	return nil
}

func (p *ListSchemasRequestExtended) ProduceRequest() *types.ListSchemasRequest {
	request := types.ListSchemasRequest{}
	if p.GetSchemaMeta() != nil {
		request.SchemaMeta = NewSchemaMeta(p.GetSchemaMeta())
	}
	return &request
}

func (p *DeleteSchemasRequestExtended) Validate() error {
	return nil
}

func (p *DeleteSchemasRequestExtended) ProduceRequest() *types.DeleteSchemasRequest {
	request := types.DeleteSchemasRequest{}
	if p.GetSchemaMeta() != nil {
		request.SchemaMeta = NewSchemaMeta(p.GetSchemaMeta())
	}
	return &request
}
