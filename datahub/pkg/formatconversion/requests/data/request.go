package data

import (
	"errors"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/schemas"
	"github.com/containers-ai/alameda/datahub/pkg/schema-mgt"
	InternalSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

type WriteDataRequestRequestExtended struct {
	*data.WriteDataRequest
}

type ReadDataRequestRequestExtended struct {
	*data.ReadDataRequest
}

type DeleteDataRequestRequestExtended struct {
	*data.DeleteDataRequest
}

func (p *WriteDataRequestRequestExtended) Validate() error {
	if p.GetSchemaMeta() == nil {
		return errors.New("schema meta is not given")
	}

	schemaMgt := schema_mgt.NewSchemaManagement()
	schemaMeta := schemas.NewSchemaMeta(p.GetSchemaMeta())

	if err := isSchemaMetaComplete(schemaMeta); err != nil {
		return err
	}

	schema := schemaMgt.GetSchemas(schemaMeta.Scope, schemaMeta.Category, schemaMeta.Type)
	if len(schema) == 0 {
		return errors.New("schema is not found")
	}

	for _, w := range p.GetWriteData() {
		metricType := enumconv.MetricTypeNameMap[w.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[w.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[w.GetResourceQuota()]
		measurement := schema[0].GetMeasurement(w.GetMeasurement(), metricType, boundary, quota)
		if measurement == nil {
			return errors.New("measurement is not supported in this schema meta")
		}
		err := measurement.Validate(w.GetColumns())
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *WriteDataRequestRequestExtended) ProduceRequest() *types.WriteDataRequest {
	request := types.WriteDataRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, w := range p.GetWriteData() {
		request.WriteData = append(request.WriteData, NewWriteData(w))
	}
	return &request
}

func (p *ReadDataRequestRequestExtended) Validate() error {
	if p.GetSchemaMeta() == nil {
		return errors.New("schema meta is not given")
	}

	schemaMgt := schema_mgt.NewSchemaManagement()
	schemaMeta := schemas.NewSchemaMeta(p.GetSchemaMeta())

	if err := isSchemaMetaComplete(schemaMeta); err != nil {
		return err
	}

	schema := schemaMgt.GetSchemas(schemaMeta.Scope, schemaMeta.Category, schemaMeta.Type)
	if len(schema) == 0 {
		return errors.New("schema is not found")
	}

	for _, r := range p.GetReadData() {
		metricType := enumconv.MetricTypeNameMap[r.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[r.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[r.GetResourceQuota()]
		if err := isMeasurementFound(schema[0], r.Measurement, metricType, boundary, quota); err != nil {
			return err
		}
	}

	return nil
}

func (p *ReadDataRequestRequestExtended) ProduceRequest() *types.ReadDataRequest {
	request := types.ReadDataRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, r := range p.GetReadData() {
		request.ReadData = append(request.ReadData, NewReadData(r))
	}
	return &request
}

func (p *DeleteDataRequestRequestExtended) Validate() error {
	if p.GetSchemaMeta() == nil {
		return errors.New("schema meta is not given")
	}

	schemaMgt := schema_mgt.NewSchemaManagement()
	schemaMeta := schemas.NewSchemaMeta(p.GetSchemaMeta())

	if err := isSchemaMetaComplete(schemaMeta); err != nil {
		return err
	}

	schema := schemaMgt.GetSchemas(schemaMeta.Scope, schemaMeta.Category, schemaMeta.Type)
	if len(schema) == 0 {
		return errors.New("schema is not found")
	}

	for _, d := range p.GetDeleteData() {
		metricType := enumconv.MetricTypeNameMap[d.GetMetricType()]
		boundary := enumconv.ResourceBoundaryNameMap[d.GetResourceBoundary()]
		quota := enumconv.ResourceQuotaNameMap[d.GetResourceQuota()]
		if err := isMeasurementFound(schema[0], d.Measurement, metricType, boundary, quota); err != nil {
			return err
		}
	}

	return nil
}

func (p *DeleteDataRequestRequestExtended) ProduceRequest() *types.DeleteDataRequest {
	request := types.DeleteDataRequest{}
	request.SchemaMeta = schemas.NewSchemaMeta(p.GetSchemaMeta())
	for _, d := range p.GetDeleteData() {
		request.DeleteData = append(request.DeleteData, NewDeleteData(d))
	}
	return &request
}

func isSchemaMetaComplete(schemaMeta *InternalSchema.SchemaMeta) error {
	if int(schemaMeta.Scope) == 0 {
		return errors.New("schema meta(scope) is not given")
	}
	if schemaMeta.Category == "" {
		return errors.New("schema meta(category) is not given")
	}
	if schemaMeta.Type == "" {
		return errors.New("schema meta(type) is not given")
	}
	return nil
}

func isMeasurementFound(schema *InternalSchema.Schema, measurement string, metricType InternalSchema.MetricType, boundary InternalSchema.ResourceBoundary, quota InternalSchema.ResourceQuota) error {
	m := schema.GetMeasurement(measurement, metricType, boundary, quota)
	if m == nil {
		return errors.New("measurement is not found in schema meta")
	}
	return nil
}
