package common

import (
	InternalCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/golang/protobuf/ptypes"
)

func NewRows(rows []*InternalCommon.Row) []*common.Row {
	if rows == nil {
		return nil
	}
	rs := make([]*common.Row, 0)
	for _, row := range rows {
		rs = append(rs, NewRow(row))
	}
	return rs
}

func NewRow(row *InternalCommon.Row) *common.Row {
	if row == nil {
		return nil
	}
	r := common.Row{}
	if row.Time != nil {
		timestamp, _ := ptypes.TimestampProto(*row.Time)
		r.Time = timestamp
	}
	r.Values = row.Values
	return &r
}
