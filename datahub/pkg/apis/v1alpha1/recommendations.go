package v1alpha1

import (
	ApiRec "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) CreateRecommendations(ctx context.Context, in *ApiRec.CreateRecommendationsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListRecommendations(ctx context.Context, in *ApiRec.ListRecommendationsRequest) (*ApiRec.ListRecommendationsResponse, error) {
	response := ApiRec.ListRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Recommendations: &ApiRec.Recommendation{},
	}

	return &response, nil
}
