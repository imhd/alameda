package v1alpha1

import (
	ApiPlannings "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/plannings"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) CreatePlannings(ctx context.Context, in *ApiPlannings.CreatePlanningsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListPlannings(ctx context.Context, in *ApiPlannings.ListPlanningsRequest) (*ApiPlannings.ListPlanningsResponse, error) {
	response := ApiPlannings.ListPlanningsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Plannings: &ApiPlannings.RawPlanning{},
	}

	return &response, nil
}
