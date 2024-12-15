package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

func (s *GoodsService) CreateClass(ctx context.Context, req *v1.CreateClassRequest) (*v1.CreateClassResponse, error) {
	name := req.GetGoodsName()
	materialId := req.GetMaterial()
	des := req.GetGoodsDes().GetValue()
	creator := ctxtools.GetSenderFromCtx(ctx)
	orgId := req.GetOrgId()
	tm := req.GetTm().GetValue()

	class, err := s.cas.CreateClass(ctx, name, creator, des, materialId, orgId, tm)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.CreateClassResponse{
		GoodsId: pbtools.ToProtoInt32(int64(class.ID)),
	}, nil
}

// 获取种类
func (s *GoodsService) GetClass(ctx context.Context, req *v1.GetClassRequest) (*v1.GetClassResponse, error) {
	class, err := s.cas.GetClass(ctx, req.GetGoodsId())
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.GetClassResponse{Class: Class2Proto(class)}, nil
}

// 列出产品类型
func (s *GoodsService) ListGoodsClass(ctx context.Context, req *v1.ListGoodsClassRequest) (*v1.ListGoodsClassResponse, error) {
	classes, err := s.cas.ListGoodsClass(ctx, req.Offset, req.Limit, req.GetOrgId())
	if err != nil {
		return nil, verr.Error(s, err)
	}
	res := make([]*v1.Class, 0)
	for _, class := range classes {
		res = append(res, Class2Proto(class))
	}
	return &v1.ListGoodsClassResponse{GoodsClasses: res}, nil
}

// 更新产品类型信息
func (s *GoodsService) UpdateGoodsClass(ctx context.Context, req *v1.UpdateGoodsClassRequest) (*v1.UpdateGoodsClassResponse, error) {
	class := req.Class
	_, err := s.cas.UpdateGoodsClass(ctx, Proto2Class(class))
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.UpdateGoodsClassResponse{Ok: true}, nil
}

// GetOrgOfX 获取产品/类型/批次所属企业
func (s *GoodsService) GetOrgOfX(ctx context.Context, req *v1.GetOrgOfXRequest) (*v1.GetOrgOfXResponse, error) {
	orgId, err := s.cas.GetOrgOfX(ctx, req.GetX().Enum().String(), req.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.GetOrgOfXResponse{
		OrgId: pbtools.ToProtoInt32(orgId),
	}, nil
}
