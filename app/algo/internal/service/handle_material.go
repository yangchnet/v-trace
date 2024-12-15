package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// 获取算法模型支持的材料.
func (s *AlgoService) ListMaterials(ctx context.Context, req *v1.ListMaterialsRequest) (*v1.ListMaterialsResponse, error) {
	materials, models, err := s.cas.ListMaterials(ctx)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	pbMaterials := make([]*v1.Material, 0)
	for i, m := range materials {
		pbMaterials = append(pbMaterials, Material2Proto(m, models[i]))
	}

	return &v1.ListMaterialsResponse{
		Materials: pbMaterials,
	}, nil
}
