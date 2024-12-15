package service

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/iam/internal/data/db"
)

// IamCaseInterface 完成有关用户的业务操作.
type IamCaseInterface interface {
	// CreateUser 创建一个新用户.
	CreateUser(ctx context.Context, nickname, realPasswd, phone, email, avatar, status string) (*db.User, error)

	// GetUser 根据唯一用户名查询用户.
	GetUser(ctx context.Context, username string) (*db.User, error)

	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, username string) error

	// CreateToken 令牌签发
	CreateToken(ctx context.Context, phone, passwd string) (string, error)

	// RefreshToken 更新令牌
	RefreshToken(ctx context.Context, username string) (string, error)

	// GetRolesByUsername
	GetRolesByUsername(ctx context.Context, username string) (string, error)

	// CreateOrg 创建组织
	CreateOrg(ctx context.Context, owner, orgName, orgCode, legalName, legalPhone string, canProduce bool, orgInfo []byte) (*db.Org, error)

	// AddMember 企业增加成员
	AddMember(ctx context.Context, memberId int, username string) error

	// GetOrg 获取企业信息
	GetOrg(ctx context.Context, orgId int) (*db.Org, error)

	// DeleteOrgMember 	企业删除成员
	DeleteOrgMember(ctx context.Context, org_id int, username string) error

	// ListOrgMember 企业查询成员列表
	ListOrgMember(ctx context.Context, org_id, offset, limit int) ([]*db.User, error)

	// UpdateOrg 企业信息更新
	UpdateOrg(ctx context.Context, org *db.Org) (*db.Org, error)

	// UpdateUser 用户信息更新
	UpdateUser(ctx context.Context, user *db.User) (*db.User, error)

	// CreateIdentity 用户实名认证
	CreateIdentity(ctx context.Context, username, realname, id_card string) (*db.User, error)

	// GetOrgOfUser 查询用户所属企业
	GetOrgOfUser(ctx context.Context, username string) (*db.Org, error)
}
