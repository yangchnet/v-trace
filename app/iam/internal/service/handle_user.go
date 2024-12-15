package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// User Manage.

// CreateUser 创建一个新用户.
func (s *IamService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	nickname := req.GetNickname()
	phone := req.GetPhone()
	passwd := req.GetPassword()
	email := req.GetEmail().GetValue()
	avatar := req.GetAvatar().GetValue()
	user, err := s.cas.CreateUser(ctx, nickname, passwd, phone, email, avatar, v1.Status_name[int32(v1.Status_enable)])
	if err != nil {
		logger.Error(err)
		return nil, verr.Error(s, err)
	}

	return &v1.CreateUserResponse{
		User: User2Proto(user, constants.NormalRole),
	}, nil
}

// DeleteUser 删除用户.
func (s *IamService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	if err := s.cas.DeleteUser(ctx, req.GetUsername()); err != nil {
		return &v1.DeleteUserResponse{
			Result: v1.OperationResult_FAIL,
		}, err
	}

	return &v1.DeleteUserResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

// GetUser 根据用户名获取用户信息.
func (s *IamService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	user, err := s.cas.GetUser(ctx, req.GetUsername())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	role, err := s.cas.GetRolesByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.GetUserResponse{
		User: User2Proto(user, role),
	}, nil
}

// GetRole 根据用户名获取用户角色.
func (s *IamService) GetRole(ctx context.Context, req *v1.GetRoleRequest) (*v1.GetRoleResponse, error) {
	role, err := s.cas.GetRolesByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.GetRoleResponse{
		Role: pbtools.ToProtoString(role),
	}, nil
}

// InquireMemberList 企业查询成员列表
func (s *IamService) ListOrgMember(ctx context.Context, req *v1.ListOrgMemberRequest) (*v1.ListOrgMemberResponse, error) {
	users, err := s.cas.ListOrgMember(ctx, int(req.OrgId), int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, verr.Error(s, err)
	}

	pbUsers := make([]*v1.User, 0)
	for _, user := range users {
		pbUsers = append(pbUsers, User2Proto(user, constants.ProducerRole))
	}

	return &v1.ListOrgMemberResponse{Users: pbUsers}, nil
}

// UpdateUser 用户信息更新
func (s *IamService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	user, err := s.cas.UpdateUser(ctx, Proto2User(req.User))
	if err != nil {
		logger.Error(err)
		return nil, verr.Error(s, err)
	}

	role, err := s.cas.GetRolesByUsername(ctx, user.Username)
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.UpdateUserResponse{
		User: User2Proto(user, role),
	}, nil
}

// 记录实名信息
// 1. 记录实名信息
// 2. 授予transporter权限
func (s *IamService) CreateIdentity(ctx context.Context, req *v1.CreateIdentityRequest) (*v1.CreateIdentityResponse, error) {
	user, err := s.cas.CreateIdentity(ctx, req.GetUsername(), req.GetRealName(), req.GetIdCard())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	role, err := s.cas.GetRolesByUsername(ctx, user.Username)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.CreateIdentityResponse{
		User: User2Proto(user, role),
	}, nil
}
