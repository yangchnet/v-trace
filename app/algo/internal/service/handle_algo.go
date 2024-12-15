package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

func (s *AlgoService) ListModels(ctx context.Context, req *v1.ListModelsRequest) (*v1.ListModelsResponse, error) {
	models, err := s.cas.ListModels(ctx)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	pbModels := make([]*v1.Model, 0)
	for _, model := range models {
		pbModels = append(pbModels, Model2Proto(model))
	}

	return &v1.ListModelsResponse{
		Models: pbModels,
	}, nil
}

func (s *AlgoService) Predict(ctx context.Context, req *v1.PredictRequest) (*v1.PredictResponse, error) {
	material, err := s.cas.Predict(ctx, req.GetModelName(), req.GetData())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.PredictResponse{
		Material: Material2Proto(material, nil),
	}, nil
}
