package common

import (
	InternalCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/golang/protobuf/ptypes"
)

func NewRows(rows []*common.Row) []*InternalCommon.Row {
	if rows == nil {
		return nil
	}
	rs := make([]*InternalCommon.Row, 0)
	for _, row := range rows {
		rs = append(rs, NewRow(row))
	}
	return rs
}

func NewRow(row *common.Row) *InternalCommon.Row {
	if row == nil {
		return nil
	}
	r := InternalCommon.Row{}
	if row.GetTime() != nil {
		timestamp, _ := ptypes.Timestamp(row.GetTime())
		r.Time = &timestamp
	}
	r.Values = row.GetValues()
	return &r
}
