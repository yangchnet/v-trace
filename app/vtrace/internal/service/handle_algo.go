package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// 获取原材料列表.
func (s *VTraceService) ListMaterial(ctx context.Context, req *v1.ListMaterialRequest) (*v1.ListMaterialResponse, error) {
	pbMaterials, err := s.cas.GetMaterials(ctx)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.ListMaterialResponse{
		Materials: pbMaterials,
	}, nil
}

// ListModels 列出所有的算法模型
func (s *VTraceService) ListModels(ctx context.Context, req *v1.ListModelsRequest) (*v1.ListModelsResponse, error) {
	models, err := s.cas.ListAlgoModels(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.ListModelsResponse{
		Models: models,
	}, nil
}

// Predict 使用算法模型进行预测
func (s *VTraceService) Predict(ctx context.Context, req *v1.PredictRequest) (*v1.PredictResponse, error) {
	material, err := s.cas.Predict(ctx, req.GetModelName(), req.GetData())
	if err != nil {
		return nil, err
	}

	return &v1.PredictResponse{
		Material: material,
	}, nil
}
