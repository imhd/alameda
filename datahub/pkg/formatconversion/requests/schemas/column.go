package schemas

import (
	InternalCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewColumn(column *ApiSchema.Column) *InternalSchema.Column {
	if column != nil {
		c := InternalSchema.Column{}
		c.Name = column.Name
		c.Required = column.Required
		c.ColumnType = InternalSchema.ColumnType(column.ColumnType)
		c.DataType = InternalCommon.DataType(column.DataType)
		return &c
	}
	return nil
}
