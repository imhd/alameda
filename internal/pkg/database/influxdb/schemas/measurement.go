package schemas

import (
	"errors"
	"fmt"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"strings"
)

type Measurement struct {
	Name       string
	MetricType MetricType
	Boundary   ResourceBoundary
	Quota      ResourceQuota
	Columns    []*Column
}

func NewMeasurement(name string, metricType MetricType, boundary ResourceBoundary, quota ResourceQuota) *Measurement {
	measurement := Measurement{}
	measurement.Name = name
	measurement.MetricType = metricType
	measurement.Boundary = boundary
	measurement.Quota = quota
	measurement.Columns = make([]*Column, 0)
	return &measurement
}

func (p *Measurement) Initialize(columns string) error {
	if columns != "" {
		columnList := strings.Split(columns, ",")
		for _, columnStr := range columnList {
			column := NewColumn()
			if err := column.Parse(columnStr); err != nil {
				return err
			}
			p.Columns = append(p.Columns, column)
		}
	}
	return nil
}

func (p *Measurement) AddColumn(name string, required bool, columnType ColumnType, dataType common.DataType) {
	column := NewColumn()
	column.Name = name
	column.Required = required
	column.ColumnType = columnType
	column.DataType = dataType
	p.Columns = append(p.Columns, column)
}

func (p *Measurement) Validate(columns []string) error {
	// Check required columns
	for _, c := range p.Columns {
		found := false
		for _, name := range columns {
			if c.Name == name {
				found = true
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("column(%s) is not given in measurement(%s)", c.Name, p.Name))
		}
	}

	// Check not supported columns
	for _, name := range columns {
		found := false
		for _, column := range p.Columns {
			if name == column.Name {
				found = true
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("column(%s) is not supported in measurement(%s)", name, p.Name))
		}
	}

	return nil
}

func (p *Measurement) String() string {
	values := make([]string, 0)
	for _, column := range p.Columns {
		values = append(values, column.String())
	}
	return strings.Join(values, ",")
}
