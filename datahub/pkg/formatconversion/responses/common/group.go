package common

import (
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

func NewGroup(group *DBCommon.Group) *ApiCommon.Group {
	if group == nil {
		return nil
	}
	g := ApiCommon.Group{}
	g.Columns = group.Columns
	g.Rows = NewRows(group.Rows)
	return &g
}
