package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// Register.
func (s *VTraceService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	nickname := req.GetNickname()
	phone := req.GetPhone()
	passwd := req.GetPasswd()

	// 创建用户
	user, token, err := s.cas.Register(ctx, nickname, passwd, phone)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterResponse{
		User:  user,
		Token: pbtools.ToProtoString(token),
	}, nil
}

// Token.
func (s *VTraceService) Token(ctx context.Context, req *v1.TokenRequest) (*v1.TokenResponse, error) {
	phone := req.GetPhone()
	passwd := req.GetPassword()
	token, err := s.cas.GetToken(ctx, phone, passwd)
	if err != nil {
		return nil, err
	}

	return &v1.TokenResponse{
		Token: token,
	}, nil
}

// RefreshToken 更新令牌
func (s *VTraceService) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	token, err := s.cas.UpdateToken(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshTokenResponse{Token: token}, nil
}

// Profile 获取用户信息.
func (s *VTraceService) Profile(ctx context.Context, req *v1.ProfileRequest) (*v1.ProfileResponse, error) {
	pbUser, err := s.cas.Profile(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.ProfileResponse{
		User: pbUser,
	}, nil
}

// IdentityAuth 实名认证.
func (s *VTraceService) IdentityAuth(ctx context.Context, req *v1.IdentityAuthRequest) (*v1.IdentityAuthResponse, error) {
	pbUser, err := s.cas.IdentityAuth(ctx, req.Realname, req.Idcard)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.IdentityAuthResponse{
		User: pbUser,
	}, nil
}

// 企业认证.
func (s *VTraceService) OrgAuth(ctx context.Context, req *v1.OrgAuthRequest) (*v1.OrgAuthResponse, error) {
	pbOrg, err := s.cas.OrgAuth(ctx, req.GetLegalName(), req.GetOrgName(), req.GetCode(), req.GetLegalPhone(), req.GetOrgInfo().GetValue())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.OrgAuthResponse{
		Org: pbOrg,
	}, nil
}

// 企业增加成员
func (s *VTraceService) Member(ctx context.Context, req *v1.MemberRequest) (*v1.MemberResponse, error) {
	if err := s.cas.Member(ctx, req.GetOrgId(), req.GetUsername()); err != nil {
		return nil, err
	}

	return &v1.MemberResponse{}, nil
}

// DeleteUser 删除用户
func (s *VTraceService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	if err := s.cas.DeleteUser(ctx, req.GetUsername()); err != nil {
		return nil, err
	}

	return &v1.DeleteUserResponse{}, nil
}

// 不需实现.
func (s *VTraceService) mustEmbedUnimplementedVTraceInterfaceServer() {
	panic("not implemented")
}

// UpdateUser 用户信息更新
func (s *VTraceService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	user, err := s.cas.UpdateUser(ctx, req.User)
	if err != nil {
		return &v1.UpdateUserResponse{}, err
	}
	return &v1.UpdateUserResponse{User: user}, nil
}

// GetOrgUser 查询用户所属企业
func (s *VTraceService) GetOrgUser(ctx context.Context, req *v1.GetOrgOfUserRequest) (*v1.GetOrgOfUserResponse, error) {
	org, err := s.cas.GetOrgUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.GetOrgOfUserResponse{Org: org}, nil
}

// OrgRemoveMember 企业删除成员
func (s *VTraceService) OrgRemoveMember(ctx context.Context, req *v1.OrgRemoveMemberRequest) (*v1.OrgRemoveMemberResponse, error) {
	if err := s.cas.OrgRemoveMember(ctx, req.OrgId, req.Username); err != nil {
		return &v1.OrgRemoveMemberResponse{}, err
	}
	return &v1.OrgRemoveMemberResponse{Ok: true}, nil
}

// ListOrgMember 企业查询成员列表
func (s *VTraceService) ListOrgMember(ctx context.Context, req *v1.ListOrgMemberRequest) (*v1.ListOrgMemberResponse, error) {
	members, err := s.cas.ListOrgMember(ctx, req.OrgId, req.Offset, req.Limit)
	if err != nil {
		return &v1.ListOrgMemberResponse{}, err
	}
	return &v1.ListOrgMemberResponse{Users: members}, nil
}

// UpdateOrg 企业信息更新
func (s *VTraceService) UpdateOrg(ctx context.Context, req *v1.UpdateOrgRequest) (*v1.UpdateOrgResponse, error) {
	org, err := s.cas.UpdateOrg(ctx, req.Org)
	if err != nil {
		return &v1.UpdateOrgResponse{}, nil
	}
	return &v1.UpdateOrgResponse{Org: org}, nil
}
