package v1alpha1

import (
	DaoSchema "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/schemas"
	FormatRequest "github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/schemas"
	FormatResponse "github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/schemas"
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schema-mgt"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	ApiSchemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) CreateSchemas(ctx context.Context, in *ApiSchemas.CreateSchemasRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateSchemas grpc function: " + AlamedaUtils.InterfaceToString(in))

	if in.GetSchemas() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	requestExt := FormatRequest.CreateSchemasRequestExtended{CreateSchemasRequest: in}
	if requestExt.Validate() != nil {
		return &status.Status{
			Code: int32(code.Code_INVALID_ARGUMENT),
		}, nil
	}

	// Get schema write lock
	SchemaMgt.RWLock.Lock()
	defer SchemaMgt.RWLock.Unlock()

	schemaDAO := DaoSchema.NewSchemaDAO(*s.Config)
	if err := schemaDAO.CreateSchemas(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to create schemas: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListSchemas(ctx context.Context, in *ApiSchemas.ListSchemasRequest) (*ApiSchemas.ListSchemasResponse, error) {
	scope.Debug("Request received from ListSchemas grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.ListSchemasRequestExtended{ListSchemasRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiSchemas.ListSchemasResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	// Get schema read lock
	SchemaMgt.RWLock.RLock()
	defer SchemaMgt.RWLock.RUnlock()

	schemaDAO := DaoSchema.NewSchemaDAO(*s.Config)
	schemaList, err := schemaDAO.ListSchemas(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ListApplications failed: %+v", err)
		return &ApiSchemas.ListSchemasResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	schemas := make([]*ApiSchemas.Schema, 0)
	for _, schema := range schemaList {
		schemas = append(schemas, FormatResponse.NewSchema(schema))
	}

	response := ApiSchemas.ListSchemasResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Schemas: schemas,
	}

	return &response, nil
}

func (s *ServiceV1alpha1) DeleteSchemas(ctx context.Context, in *ApiSchemas.DeleteSchemasRequest) (*status.Status, error) {
	scope.Debug("Request received from ListSchemas grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.DeleteSchemasRequestExtended{DeleteSchemasRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	// Get schema write lock
	SchemaMgt.RWLock.Lock()
	defer SchemaMgt.RWLock.Unlock()

	schemaDAO := DaoSchema.NewSchemaDAO(*s.Config)
	if err := schemaDAO.DeleteSchemas(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete schemas: %+v", err)
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
