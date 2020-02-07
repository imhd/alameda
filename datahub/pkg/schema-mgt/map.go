package schema_mgt

import (
	"encoding/json"
	"fmt"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

type SchemaMap struct {
	Schemas map[schemas.Scope][]*schemas.Schema
}

func NewSchemaMap() *SchemaMap {
	schemaMap := SchemaMap{}
	schemaMap.Schemas = make(map[schemas.Scope][]*schemas.Schema)
	return &schemaMap
}

func (p *SchemaMap) AddSchema(schema *schemas.Schema) {
	if p.Schemas == nil {
		p.Schemas = make(map[schemas.Scope][]*schemas.Schema)
	}

	scope := schema.SchemaMeta.Scope
	if _, ok := p.Schemas[scope]; !ok {
		p.Schemas[scope] = make([]*schemas.Schema, 0)
	}

	// Do a deep copy of schema
	s := schemas.Schema{}
	DeepCopy(&s, schema)
	p.Schemas[scope] = append(p.Schemas[scope], &s)
}

func (p *SchemaMap) Empty() {
	if p.Schemas != nil {
		for k := range p.Schemas {
			delete(p.Schemas, k)
		}
		p.Schemas = nil
	}
}

func DeepCopy(dst interface{}, src interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("unable to unmarshal into dst: %s", err)
	}
	return nil
}
