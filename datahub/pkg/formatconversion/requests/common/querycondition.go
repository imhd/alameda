package common

import (
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/golang/protobuf/ptypes"
)

func NewQueryCondition(queryCondition *ApiCommon.QueryCondition) *common.QueryCondition {
	if queryCondition != nil {
		qc := common.QueryCondition{}
		if queryCondition.GetTimeRange() != nil {
			timeRange := queryCondition.GetTimeRange()
			if timeRange.GetStartTime() != nil {
				ts, _ := ptypes.Timestamp(timeRange.GetStartTime())
				qc.StartTime = &ts
			}
			if timeRange.GetEndTime() != nil {
				ts, _ := ptypes.Timestamp(timeRange.GetEndTime())
				qc.EndTime = &ts
			}
			if timeRange.GetStep() != nil {
				ts, _ := ptypes.Duration(timeRange.GetStep())
				qc.StepTime = &ts
			}
			qc.AggregateOverTimeFunction = enumconv.AggregateFunctionNameMap[timeRange.GetAggregateFunction()]
		}
		qc.TimestampOrder = enumconv.QueryConditionOrderNameMap[queryCondition.GetOrder()]
		qc.WhereClause = queryCondition.GetWhereClause()
		qc.WhereCondition = NewWhereCondition(queryCondition.GetWhereCondition())
		qc.Selects = queryCondition.GetSelects()
		qc.Groups = queryCondition.GetGroups()
		qc.Limit = int(queryCondition.GetLimit())
		return &qc
	}
	return nil
}
