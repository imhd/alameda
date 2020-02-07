package v1alpha1

import (
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	ApiApps "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/applications"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) CreateApps(ctx context.Context, in *ApiApps.CreateApplicationsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateApps grpc function: " + AlamedaUtils.InterfaceToString(in))

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListApps(ctx context.Context, in *ApiApps.ListApplicationsRequest) (*ApiApps.ListApplicationsResponse, error) {
	scope.Debug("Request received from ListApps grpc function: " + AlamedaUtils.InterfaceToString(in))

	response := ApiApps.ListApplicationsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Applications: &ApiApps.Application{},
	}

	return &response, nil
}

func (s *ServiceV1alpha1) DeleteApps(ctx context.Context, in *ApiApps.DeleteApplicationsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
