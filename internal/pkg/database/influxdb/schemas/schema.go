package schemas

import (
	"strings"
)

type Schema struct {
    SchemaMeta   *SchemaMeta
    Measurements []*Measurement
}

func NewSchema(scope Scope, category, schemaType string) *Schema {
	schema := Schema{}
	schema.SchemaMeta = NewSchemaMeta(scope, category, schemaType)
	schema.Measurements = make([]*Measurement, 0)
	return &schema
}

func (p *Schema) AddMeasurement(name string, metricType MetricType, boundary ResourceBoundary, quota ResourceQuota, columns string) error {
	measurement := NewMeasurement(name, metricType, boundary, quota)
	if err := measurement.Initialize(columns); err != nil {
		return err
	}
	p.Measurements = append(p.Measurements, measurement)
	return nil
}

func (p *Schema) GetMeasurement(name string, metricType MetricType, boundary ResourceBoundary, quota ResourceQuota) *Measurement {
	if name != "" {
		for _, m := range p.Measurements {
			if m.Name == name {
				return m
			}
		}
		return nil
	}
	for _, m := range p.Measurements {
		if m.MetricType == metricType && m.Boundary == boundary && m.Quota == quota {
			return m
		}
	}
	return nil
}

func (p *Schema) Validate() error {
	return nil
}

func (p *Schema) String() string {
	values := make([]string, 0)
	for _, measurement := range p.Measurements {
		values = append(values, measurement.Name)
	}
	return strings.Join(values, ",")
}
