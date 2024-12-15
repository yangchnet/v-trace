package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// CreateOrg 创建一个企业组织
func (s *IamService) CreateOrg(ctx context.Context, req *v1.CreateOrgRequest) (*v1.CreateOrgResponse, error) {
	org, err := s.cas.CreateOrg(
		ctx,
		req.GetOwner(),
		req.GetOrgName(),
		req.GetOrgCode(),
		req.GetLegalName(),
		req.GetLegalPhone(),
		req.GetCanProduce(),
		req.GetOrgInfo(),
	)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.CreateOrgResponse{
		Org: Org2Proto(org),
	}, nil
}

// OrgAddMember 企业增加成员
func (s *IamService) OrgAddMember(ctx context.Context, req *v1.OrgAddMemberRequest) (*v1.OrgAddMemberResponse, error) {
	if err := s.cas.AddMember(ctx, int(req.GetOrgId()), req.GetUsername()); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.OrgAddMemberResponse{}, nil
}

// GetOrg 获取企业信息
func (s *IamService) GetOrg(ctx context.Context, in *v1.GetOrgRequest) (*v1.GetOrgResponse, error) {
	org, err := s.cas.GetOrg(ctx, int(in.GetOrgId()))
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.GetOrgResponse{
		Org: Org2Proto(org),
	}, nil
}

// OrgRemoveMember 企业删除成员
func (s *IamService) OrgDeleteMember(ctx context.Context, req *v1.DeleteOrgMemberRequest) (*v1.DeleteOrgMemberResponse, error) {
	if err := s.cas.DeleteOrgMember(ctx, int(req.OrgId), req.Username); err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.DeleteOrgMemberResponse{Ok: true}, nil
}

// UpdateOrg 企业信息更新
func (s *IamService) UpdateOrg(ctx context.Context, req *v1.UpdateOrgRequest) (*v1.UpdateOrgResponse, error) {
	org, err := s.cas.UpdateOrg(ctx, Proto2Org(req.Org))
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.UpdateOrgResponse{Org: Org2Proto(org)}, nil
}

// GetOrgOfUser 查询用户所属企业
func (s *IamService) GetOrgOfUser(ctx context.Context, req *v1.GetOrgOfUserRequest) (*v1.GetOrgOfUserResponse, error) {
	org, err := s.cas.GetOrgOfUser(ctx, req.Username)
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.GetOrgOfUserResponse{Org: Org2Proto(org)}, nil
}
