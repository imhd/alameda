package v1alpha1

import (
	DaoData "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data"
	FormatRequest "github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/data"
	FormatResponse "github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/data"
	//InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	//InternalPromth "github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	ApiData "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	//Common "github.com/containers-ai/api/common"
	//"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) ReadData(ctx context.Context, in *ApiData.ReadDataRequest) (*ApiData.ReadDataResponse, error) {
	scope.Debug("Request received from ReadData grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.ReadDataRequestRequestExtended{ReadDataRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &ApiData.ReadDataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}

	if in.GetReadData() == nil {
		return &ApiData.ReadDataResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Data: nil,
		}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	readData, err := dataDAO.ReadData(requestExt.ProduceRequest())
	if err != nil {
		scope.Errorf("ReadData failed: %+v", err)
		return &ApiData.ReadDataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	response := ApiData.ReadDataResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Data: FormatResponse.NewData(readData),
	}

	return &response, nil
}

func (s *ServiceV1alpha1) WriteData(ctx context.Context, in *ApiData.WriteDataRequest) (*status.Status, error) {
	scope.Debug("Request received from WriteData grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.WriteDataRequestRequestExtended{WriteDataRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	if in.GetWriteData() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	if err := dataDAO.WriteData(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to write data: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}

func (s *ServiceV1alpha1) DeleteData(ctx context.Context, in *ApiData.DeleteDataRequest) (*status.Status, error) {
	scope.Debug("Request received from WriteData grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExt := FormatRequest.DeleteDataRequestRequestExtended{DeleteDataRequest: in}
	if err := requestExt.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	if in.GetDeleteData() == nil {
		return &status.Status{Code: int32(code.Code_OK)}, nil
	}

	dataDAO := DaoData.NewDataDAO(*s.Config)
	if err := dataDAO.DeleteData(requestExt.ProduceRequest()); err != nil {
		scope.Errorf("failed to delete data: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}
