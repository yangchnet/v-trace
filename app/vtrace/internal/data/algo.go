package data

import (
	"context"

	algoV1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
)

// 获取原料列表.
func (d *Data) GetMaterials(ctx context.Context) ([]*algoV1.Material, error) {
	resp, err := d.Algo().ListMaterials(ctx, &algoV1.ListMaterialsRequest{})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return resp.Materials, nil
}

// ListAlgoModels 列出所有算法模型
func (d *Data) ListAlgoModels(ctx context.Context) ([]*algoV1.Model, error) {
	resp, err := d.Algo().ListModels(ctx, &algoV1.ListModelsRequest{})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return resp.Models, nil
}

// Predict 根据数据进行预测
func (d *Data) Predict(ctx context.Context, modelName string, data []byte) (*algoV1.Material, error) {
	resp, err := d.Algo().Predict(ctx, &v1.PredictRequest{
		ModelName: modelName,
		Data:      data,
	})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return resp.Material, nil
}
