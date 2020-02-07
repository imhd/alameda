package scores

import (
	DaoScoreTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/scores/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/common"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	ApiScores "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/scores"
)

type ListSimulatedSchedulingScoresRequestExtended struct {
	Request *ApiScores.ListSimulatedSchedulingScoresRequest
}

func (r *ListSimulatedSchedulingScoresRequestExtended) ProduceRequest() DaoScoreTypes.ListRequest {
	var (
		queryCondition DBCommon.QueryCondition
	)

	queryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	listRequest := DaoScoreTypes.ListRequest{
		QueryCondition: queryCondition,
	}

	return listRequest
}
