package influxdb

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
)

type Functions = string

// Influxdb function definition
const (
	Last Functions = "LAST"
	Mean Functions = "MEAN"
	Max  Functions = "MAX"
)

var AggregateFuncMap = map[common.AggregateFunction]Functions{
	common.None:        Last,
	common.MaxOverTime: Max,
	common.AvgOverTime: Mean,
}
