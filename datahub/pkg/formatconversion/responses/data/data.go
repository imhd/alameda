package data

import (
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/common"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/schemas"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

func NewData(readData *types.Data) *data.Data {
	if readData == nil {
		return nil
	}
	rawdata := data.Data{}
	rawdata.SchemaMeta = schemas.NewSchemaMeta(readData.SchemaMeta)
	rawdata.Rawdata = NewRawdata(readData.Rawdata)
	return &rawdata
}

func NewRawdata(rawdata []*types.Rawdata) []*data.Rawdata {
	if rawdata == nil {
		return nil
	}
	rawdataList := make([]*data.Rawdata, 0)
	for _, d := range rawdata {
		rd := data.Rawdata{}
		rd.Measurement = d.Measurement
		rd.MetricType = enumconv.MetricTypeNameMap[d.MetricType]
		rd.ResourceBoundary = enumconv.ResourceBoundaryNameMap[d.Boundary]
		rd.ResourceQuota = enumconv.ResourceQuotaNameMap[d.Quota]
		for _, group := range d.Groups {
			rd.Groups = append(rd.Groups, common.NewGroup(group))
		}
		rawdataList = append(rawdataList, &rd)
	}
	return rawdataList
}
