package common

import (
	Common "github.com/containers-ai/api/common"
	"time"
)

// Order enumerator
type Order = int

// Aggregate function enumerator
type AggregateFunction = int

// Data type enumerator
type DataType int

// Sort order definition
const (
	NoneOrder Order = iota
	Asc             // Represent ascending order
	Desc            // Represent descending order
)

// Aggregation function definition
const (
	None        AggregateFunction = 0
	MaxOverTime AggregateFunction = 1
	AvgOverTime AggregateFunction = 2
)

// Data type definition
const (
	Invalid DataType = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
)

type Row struct {
    Time   *time.Time
    Values []string
}

type Group struct {
	Columns []string
	Rows    []*Row
}

var (
	AggregationOverTime = map[AggregateFunction]string{
		MaxOverTime: "max_over_time",
		AvgOverTime: "avg_over_time",
	}

	TimeRange2AggregationOverTime = map[Common.TimeRange_AggregateFunction]AggregateFunction{
		Common.TimeRange_NONE: None,
		Common.TimeRange_MAX:  MaxOverTime,
		Common.TimeRange_AVG:  AvgOverTime,
	}
)
