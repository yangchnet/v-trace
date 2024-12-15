package biz

import (
	"context"

	algoV1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
)

// 获取原料列表.
func (c *VTraceCase) GetMaterials(ctx context.Context) ([]*algoV1.Material, error) {
	return c.repo.GetMaterials(ctx)
}

// ListAlgoModels 列出所有算法模型
func (c *VTraceCase) ListAlgoModels(ctx context.Context) ([]*algoV1.Model, error) {
	return c.repo.ListAlgoModels(ctx)
}

// Predict 根据数据进行预测
func (c *VTraceCase) Predict(ctx context.Context, modelName string, data []byte) (*algoV1.Material, error) {
	return c.repo.Predict(ctx, modelName, data)
}
