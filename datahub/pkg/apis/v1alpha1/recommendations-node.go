package v1alpha1

import (
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	ApiRecommendations "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) CreateNodeRecommendations(ctx context.Context, in *ApiRecommendations.CreateNodeRecommendationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateNodeRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListNodeRecommendations(ctx context.Context, in *ApiRecommendations.ListNodeRecommendationsRequest) (*ApiRecommendations.ListNodeRecommendationsResponse, error) {
	scope.Debug("Request received from ListNodeRecommendations grpc function: " + AlamedaUtils.InterfaceToString(in))

	return &ApiRecommendations.ListNodeRecommendationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
	}, nil
}
