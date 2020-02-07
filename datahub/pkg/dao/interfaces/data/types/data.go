package types

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

type DataDAO interface {
	WriteData(*WriteDataRequest) error
	ReadData(*ReadDataRequest) (*Data, error)
	DeleteData(*DeleteDataRequest) error
}

type WriteDataRequest struct {
	SchemaMeta *schemas.SchemaMeta
	WriteData  []*WriteData
}

type ReadDataRequest struct {
	SchemaMeta *schemas.SchemaMeta
	ReadData   []*ReadData
}

type DeleteDataRequest struct {
	SchemaMeta *schemas.SchemaMeta
	DeleteData []*DeleteData
}

type WriteData struct {
	Measurement string
	MetricType  schemas.MetricType
	Boundary    schemas.ResourceBoundary
	Quota       schemas.ResourceQuota
	Columns     []string
	Rows        []*common.Row
}

type ReadData struct {
	Measurement    string
	MetricType     schemas.MetricType
	Boundary       schemas.ResourceBoundary
	Quota          schemas.ResourceQuota
	QueryCondition *common.QueryCondition
}

type DeleteData struct {
	Measurement    string
	MetricType     schemas.MetricType
	Boundary       schemas.ResourceBoundary
	Quota          schemas.ResourceQuota
	QueryCondition *common.QueryCondition
}

type Data struct {
	SchemaMeta *schemas.SchemaMeta
	Rawdata    []*Rawdata
}

type Rawdata struct {
	Measurement string
	MetricType  schemas.MetricType
	Boundary    schemas.ResourceBoundary
	Quota       schemas.ResourceQuota
	Groups      []*common.Group
}
