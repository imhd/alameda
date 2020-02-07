package influxdb

import (
	"errors"
	DaoDataTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	SchemaMgt "github.com/containers-ai/alameda/datahub/pkg/schema-mgt"
	//InternalCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	InfluxSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	Log "github.com/containers-ai/alameda/pkg/utils/log"
)

var (
	scope = Log.RegisterScope("dao-data-impl", "dao data implementation", 0)
)

type Data struct {
	InfluxDBConfig InternalInflux.Config
}

func NewDataWithConfig(config InternalInflux.Config) DaoDataTypes.DataDAO {
	return &Data{InfluxDBConfig: config}
}

func (p *Data) WriteData(request *DaoDataTypes.WriteDataRequest) error {
	if request.SchemaMeta == nil || request.WriteData == nil {
		return nil
	}

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]
	for _, w := range request.WriteData {
		m := schema.GetMeasurement(w.Measurement, w.MetricType, w.Boundary, w.Quota)
		if m == nil {
			scope.Error("measurement is not found when writing data")
			return errors.New("measurement is not found")
		}
		measurement := InternalInflux.NewMeasurement(SchemaMgt.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)
		err := measurement.Write(w.Columns, w.Rows)
		if err != nil {
			scope.Error(err.Error())
			return err
		}
	}

	return nil
}

func (p *Data) ReadData(request *DaoDataTypes.ReadDataRequest) (*DaoDataTypes.Data, error) {
	if request.SchemaMeta == nil || request.ReadData == nil {
		return nil, nil
	}

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]

	data := DaoDataTypes.Data{}
	data.SchemaMeta = InfluxSchema.NewSchemaMeta(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)
	for _, r := range request.ReadData {
		m := schema.GetMeasurement(r.Measurement, r.MetricType, r.Boundary, r.Quota)
		if m == nil {
			scope.Error("measurement is not found when reading data")
			return nil, errors.New("measurement is not found")
		}
		measurement := InternalInflux.NewMeasurement(SchemaMgt.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)
		groups, err := measurement.Read(InternalInflux.NewQuery(r.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return nil, err
		}
		rawdata := DaoDataTypes.Rawdata{
			Measurement: measurement.Name,
			MetricType:  measurement.MetricType,
			Boundary:    measurement.Boundary,
			Quota:       measurement.Quota,
			Groups:      groups,
		}
		data.Rawdata = append(data.Rawdata, &rawdata)
	}

	return &data, nil
}

func (p *Data) DeleteData(request *DaoDataTypes.DeleteDataRequest) error {
	if request.SchemaMeta == nil || request.DeleteData == nil {
		return nil
	}

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]

	data := DaoDataTypes.Data{}
	data.SchemaMeta = InfluxSchema.NewSchemaMeta(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)
	for _, d := range request.DeleteData {
		m := schema.GetMeasurement(d.Measurement, d.MetricType, d.Boundary, d.Quota)
		if m == nil {
			scope.Error("measurement is not found when reading data")
			return errors.New("measurement is not found")
		}
		measurement := InternalInflux.NewMeasurement(SchemaMgt.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)
		err := measurement.Drop(InternalInflux.NewQuery(d.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return err
		}
	}

	return nil
}
