package influxdb

import (
	DaoSchemaTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/schemas/types"
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schema-mgt"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	Log "github.com/containers-ai/alameda/pkg/utils/log"
)

var (
	scope = Log.RegisterScope("dao-schema-impl", "dao schema implementation", 0)
)

type Schema struct {
	InfluxDBConfig InternalInflux.Config
}

func NewSchemaWithConfig(config InternalInflux.Config) DaoSchemaTypes.SchemaDAO {
	return &Schema{InfluxDBConfig: config}
}

func (p *Schema) CreateSchemas(request *DaoSchemaTypes.CreateSchemasRequest) error {
	if request.Schemas == nil {
		return nil
	}
	schemaMgt := SchemaMgt.NewSchemaManagement()
	schemaMgt.AddSchemas(request.Schemas)
	err := schemaMgt.Flush()
	if err != nil {
		scope.Error(err.Error())
		return err
	}
	return nil
}

func (p *Schema) ListSchemas(request *DaoSchemaTypes.ListSchemasRequest) ([]*schemas.Schema, error) {
	schemaMgt := SchemaMgt.NewSchemaManagement()
	if request.SchemaMeta == nil {
		return schemaMgt.GetSchemas(SchemaMgt.ScopeUndefined, "", ""), nil
	}
	return schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type), nil
}

func (p *Schema) DeleteSchemas(request *DaoSchemaTypes.DeleteSchemasRequest) error {
	if request.SchemaMeta == nil {
		return nil
	}
	schemaMgt := SchemaMgt.NewSchemaManagement()
	err := schemaMgt.DeleteSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)
	if err != nil {
		scope.Error(err.Error())
		return err
	}
	return nil
}
