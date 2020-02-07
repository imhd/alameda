package schemas

import (
	"github.com/containers-ai/alameda/pkg/utils/log"
)

type Scope int
type MetricType int
type ResourceBoundary int
type ResourceQuota int
type ColumnType int

// Influxdb column type enumerator
const (
	ColumnTypeUndefined ColumnType = iota
	Tag
	Field
)

var (
	scope = log.RegisterScope("Database", "influxdb-schemas", 0)
)
