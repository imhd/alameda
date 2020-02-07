package schemas

import (
	InternalSchema "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewMeasurement(measurement *InternalSchema.Measurement) *ApiSchema.Measurement {
	if measurement != nil {
		m := ApiSchema.Measurement{}
		m.Name = measurement.Name
		m.MetricType = ApiCommon.MetricType(measurement.MetricType)
		for _, column := range measurement.Columns {
			m.Columns = append(m.Columns, NewColumn(column))
		}
		return &m
	}
	return nil
}
