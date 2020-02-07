package common

import (
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

func NewWhereCondition(conditions []*ApiCommon.Condition) []*common.Condition {
	if conditions != nil {
		whereCondition := make([]*common.Condition, 0)
		for _, condition := range conditions {
			whereCondition = append(whereCondition, NewCondition(condition))
		}
		return whereCondition
	}
	return nil
}

func NewCondition(condition *ApiCommon.Condition) *common.Condition {
	if condition != nil {
		c := common.Condition{}
		c.Keys = condition.GetKeys()
		c.Values = condition.GetValues()
		c.Operators = condition.GetOperators()
		for _, dataType := range condition.GetTypes() {
			c.Types = append(c.Types, enumconv.DataTypeNameMap[dataType])
		}
		return &c
	}
	return nil
}
