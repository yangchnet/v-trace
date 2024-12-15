package data

import (
	"context"

	iamV1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
)

func (r *Data) CreateUser(ctx context.Context, nickname, passwd, phone string) (*iamV1.User, error) {
	resp, err := r.Iam().CreateUser(ctx, &iamV1.CreateUserRequest{
		Nickname: nickname,
		Password: passwd,
		Phone:    phone,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return resp.User, nil
}

func (r *Data) GetToken(ctx context.Context, phone, passwd string) (string, error) {
	resp, err := r.Iam().Token(ctx, &iamV1.TokenRequest{
		Phone:  phone,
		Passwd: passwd,
	})
	if err != nil {
		logger.Errorf("get token failed: %v", err)
		return "", err
	}

	return resp.GetToken().GetValue(), nil
}

// RefreshToken 更新令牌
func (r *Data) UpdateToken(ctx context.Context, token string) (string, error) {
	username := ctxtools.GetSenderFromCtx(ctx)

	logger.Warn(username)

	newToken, err := r.Iam().RefreshToken(ctxtools.WithMetadata(ctx), &iamV1.RefreshTokenRequest{})
	if err != nil {
		logger.Errorf("Refresh Token failed: %v\n", err)
		return "", err
	}
	return newToken.GetToken().GetValue(), nil
}

// GetUserInfo 获取用户信息.
func (r *Data) GetUserInfo(ctx context.Context, username string) (*iamV1.User, error) {
	resp, err := r.Iam().GetUser(ctx, &iamV1.GetUserRequest{
		Username: username,
	})
	if err != nil {
		logger.Errorf("get user[%s] failed", username, err)
		return nil, err
	}

	return resp.User, nil
}

// UpdateUser 用户信息更新
func (r *Data) UpdateUser(ctx context.Context, user *iamV1.User) (*iamV1.User, error) {
	updateUser, err := r.Iam().UpdateUser(ctx, &iamV1.UpdateUserRequest{User: user})
	if err != nil {
		logger.Errorf("update user[%s] failed", user, err)
		return nil, err
	}
	return updateUser.User, nil
}

// CreateOrg 创建企业
func (r *Data) CreateOrg(ctx context.Context, owner, orgName, orgCode, legalName, legalPhone string, canProduce bool, orgInfo []byte) (*iamV1.Org, error) {
	pbOrg, err := r.Iam().CreateOrg(ctx, &iamV1.CreateOrgRequest{
		Owner:      owner,
		OrgName:    orgName,
		OrgCode:    orgCode,
		LegalName:  legalName,
		LegalPhone: legalPhone,
		CanProduce: true, // TODO: 当前默认可以生产err
		OrgInfo:    orgInfo,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return pbOrg.Org, nil
}

// GetOrg 获取企业信息
func (r *Data) GetOrg(ctx context.Context, orgId int32) (*iamV1.Org, error) {
	pbOrg, err := r.Iam().GetOrg(ctx, &iamV1.GetOrgRequest{
		OrgId: int64(orgId),
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return pbOrg.Org, nil
}

// AddMember 企业增加成员
func (r *Data) AddMember(ctx context.Context, orgId int32, username string) error {
	if _, err := r.Iam().OrgAddMember(ctx, &iamV1.OrgAddMemberRequest{
		OrgId:    int64(orgId),
		Username: username,
	}); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// 删除用户
func (r *Data) DeleteUser(ctx context.Context, username string) error {
	_, err := r.Iam().DeleteUser(ctx, &iamV1.DeleteUserRequest{
		Username: username,
	})
	if err != nil {
		logger.Error(err)

		return err
	}

	return nil
}

// CreateIdentity 记录用户实名
func (r *Data) CreateIdentity(ctx context.Context, username, realname, idcard string) (*iamV1.User, error) {
	user, err := r.Iam().CreateIdentity(ctx, &iamV1.CreateIdentityRequest{
		Username: username,
		RealName: realname,
		IdCard:   idcard,
	})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return user.User, nil
}

// GetOrgUser 查询用户所属企业
func (r *Data) GetOrgOfUser(ctx context.Context, username string) (*iamV1.Org, error) {
	org, err := r.Iam().GetOrgOfUser(ctx, &iamV1.GetOrgOfUserRequest{Username: username})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return org.Org, nil
}

// OrgRemoveMember 企业删除成员
func (r *Data) OrgRemoveMember(ctx context.Context, id int32, username string) error {
	if _, err := r.Iam().DeleteOrgMember(ctx, &iamV1.DeleteOrgMemberRequest{
		OrgId:    int64(id),
		Username: username,
	}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// ListOrgMember 企业查询成员列表
func (r *Data) ListOrgMember(ctx context.Context, org_id, offset, limit int32) ([]*iamV1.User, error) {
	members, err := r.Iam().ListOrgMember(ctx, &iamV1.ListOrgMemberRequest{
		OrgId:  int64(org_id),
		Offset: int64(limit * offset),
		Limit:  int64(limit * (offset + 1)),
	})
	if err != nil {
		return nil, err
	}
	return members.Users, nil
}

// UpdateOrg 企业信息更新
func (r *Data) UpdateOrg(ctx context.Context, org *iamV1.Org) (*iamV1.Org, error) {
	updateOrg, err := r.Iam().UpdateOrg(ctx, &iamV1.UpdateOrgRequest{Org: org})
	if err != nil {
		return nil, err
	}
	return updateOrg.Org, nil
}
