package schemas

type SchemaMeta struct {
	Scope    Scope
	Category string
	Type     string
}

func NewSchemaMeta(scope Scope, category, schemaType string) *SchemaMeta {
	schemaMeta := SchemaMeta{}
	schemaMeta.Scope = scope
	schemaMeta.Category = category
	schemaMeta.Type = schemaType
	return &schemaMeta
}
