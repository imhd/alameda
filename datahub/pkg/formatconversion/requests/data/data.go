package data

import (
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

func NewWriteData(writeData *data.WriteData) *types.WriteData {
	if writeData == nil {
		return nil
	}
	data := types.WriteData{}
	data.Measurement = writeData.GetMeasurement()
	data.MetricType = enumconv.MetricTypeNameMap[writeData.GetMetricType()]
	data.Boundary = enumconv.ResourceBoundaryNameMap[writeData.GetResourceBoundary()]
	data.Quota = enumconv.ResourceQuotaNameMap[writeData.GetResourceQuota()]
	data.Columns = writeData.GetColumns()
	data.Rows = common.NewRows(writeData.GetRows())
	return &data
}

func NewReadData(readData *data.ReadData) *types.ReadData {
	if readData == nil {
		return nil
	}
	data := types.ReadData{}
	data.Measurement = readData.GetMeasurement()
	data.MetricType = enumconv.MetricTypeNameMap[readData.GetMetricType()]
	data.Boundary = enumconv.ResourceBoundaryNameMap[readData.GetResourceBoundary()]
	data.Quota = enumconv.ResourceQuotaNameMap[readData.GetResourceQuota()]
	data.QueryCondition = common.NewQueryCondition(readData.GetQueryCondition())
	return &data
}

func NewDeleteData(deleteData *data.DeleteData) *types.DeleteData {
	if deleteData == nil {
		return nil
	}
	data := types.DeleteData{}
	data.Measurement = deleteData.GetMeasurement()
	data.MetricType = enumconv.MetricTypeNameMap[deleteData.GetMetricType()]
	data.Boundary = enumconv.ResourceBoundaryNameMap[deleteData.GetResourceBoundary()]
	data.Quota = enumconv.ResourceQuotaNameMap[deleteData.GetResourceQuota()]
	data.QueryCondition = common.NewQueryCondition(deleteData.GetQueryCondition())
	return &data
}
