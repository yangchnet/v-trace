package biz

import (
	"context"

	iamV1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
)

// Register 注册.
func (uc *VTraceCase) Register(ctx context.Context, nickname, passwd, phone string) (*iamV1.User, string, error) {
	u, err := uc.repo.CreateUser(ctx, nickname, passwd, phone)
	if err != nil {
		logger.Error(err)

		return nil, "", err
	}

	_, err = uc.repo.CreateCert(ctx, u.GetUsername().GetValue())
	if err != nil {
		return nil, "", err
	}

	tokenResp, err := uc.repo.GetToken(ctx, phone, passwd)
	if err != nil {
		return nil, "", err
	}

	return u, tokenResp, nil
}

// GetToken 获取token.
func (uc *VTraceCase) GetToken(ctx context.Context, phone, passwd string) (string, error) {
	return uc.repo.GetToken(ctx, phone, passwd)
}

// UpdateToken 更新令牌
func (uc *VTraceCase) UpdateToken(ctx context.Context) (string, error) {
	username := ctxtools.GetSenderFromCtx(ctx)

	logger.Warn(username)
	return uc.repo.UpdateToken(ctx, "")
}

// Profile 获取个人信息.
func (uc *VTraceCase) Profile(ctx context.Context) (*iamV1.User, error) {
	username := ctxtools.GetSenderFromCtx(ctx)

	user, err := uc.repo.GetUserInfo(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// IdentityAuth 个人实名认证.
func (uc *VTraceCase) IdentityAuth(ctx context.Context, realname, idcard string) (*iamV1.User, error) {
	username := ctxtools.GetSenderFromCtx(ctx)

	// TODO 1. 检查实名信息

	return uc.repo.CreateIdentity(ctx, username, realname, idcard)
}

// OrgAuth 企业认证.
func (uc *VTraceCase) OrgAuth(ctx context.Context, legalName, company, code, phone string, info []byte) (*iamV1.Org, error) {
	username := ctxtools.GetSenderFromCtx(ctx)

	// TODO 1. 检查认证信息

	// 2. 创建企业信息
	org, err := uc.repo.CreateOrg(ctx, username, company, code, legalName, phone, true, info)
	if err != nil {
		return nil, err
	}

	return org, nil
}

// 企业增加成员
func (uc *VTraceCase) Member(ctx context.Context, orgId int32, username string) error {
	org, err := uc.repo.GetOrg(ctx, orgId)
	if err != nil {
		return err
	}

	operator := ctxtools.GetSenderFromCtx(ctx)
	if org.GetOwner().GetValue() != operator {
		return v1.ErrorPermissionDenied("[%s]不是企业[%s]的所有者", operator, org.GetOrgName().GetValue())
	}

	return uc.repo.AddMember(ctx, orgId, username)
}

// 删除用户
func (uc *VTraceCase) DeleteUser(ctx context.Context, username string) error {
	return uc.repo.DeleteUser(ctx, username)
}

// UpdateUser 用户信息更新
func (uc *VTraceCase) UpdateUser(ctx context.Context, user *iamV1.User) (*iamV1.User, error) {
	if err := checkSelf(ctx, user.Username.GetValue()); err != nil {
		return nil, err
	}
	return uc.repo.UpdateUser(ctx, user)
}

// GetOrgUser 查询用户所属企业
func (uc *VTraceCase) GetOrgUser(ctx context.Context, username string) (*iamV1.Org, error) {
	return uc.repo.GetOrgOfUser(ctx, username)
}

// OrgRemoveMember 企业删除成员
func (uc *VTraceCase) OrgRemoveMember(ctx context.Context, id int32, username string) error {
	if err := uc.checkOrgOwner(ctx, id); err != nil {
		return err
	}
	return uc.repo.OrgRemoveMember(ctx, id, username)
}

// ListOrgMember 企业查询成员列表
func (uc *VTraceCase) ListOrgMember(ctx context.Context, org_id, offset, limit int32) ([]*iamV1.User, error) {
	if err := uc.checkOrgOwner(ctx, org_id); err != nil {
		return nil, err
	}
	return uc.repo.ListOrgMember(ctx, org_id, limit*offset, limit*(offset+1))
}

// UpdateOrg 企业信息更新
func (uc *VTraceCase) UpdateOrg(ctx context.Context, org *iamV1.Org) (*iamV1.Org, error) {
	if err := uc.checkOrgOwner(ctx, int32(org.Id.Value)); err != nil {
		return nil, err
	}
	return uc.repo.UpdateOrg(ctx, org)
}

// checkSelf 检查要操作的用户是否是自身.
func checkSelf(ctx context.Context, username string) error {
	sender := ctxtools.GetSenderFromCtx(ctx)
	if sender != username {
		logger.Errorf("permission error: [%s] required [%s]'s profile", sender, username)

		return ErrPermissionDenied
	}
	return nil
}

// checkOrgOwner 检查企业所属人与操作者是否为同一人
func (uc *VTraceCase) checkOrgOwner(ctx context.Context, org_id int32) error {
	org, err := uc.repo.GetOrg(ctx, org_id)
	if err != nil {
		return err
	}
	operator := ctxtools.GetSenderFromCtx(ctx)
	if org.Owner.GetValue() != operator {
		return v1.ErrorPermissionDenied("[%s]不是企业[%s]的所有者", operator, org.GetOrgName().GetValue())
	}
	return nil
}
