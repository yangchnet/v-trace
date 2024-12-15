package biz

import (
	"context"
	"database/sql"
	"errors"

	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/verr"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/data/db"
)

// CreateOrg 创建企业
func (c *IamCase) CreateOrg(ctx context.Context, owner, orgName, orgCode, legalName, legalPhone string, canProduce bool, orgInfo []byte) (*db.Org, error) {
	var orgId int
	if err := c.repo.ExecTx(ctx, func(q *db.Queries) error {
		_orgId, err := q.CreateOrg(ctx, &db.CreateOrgParams{
			OrgName:          orgName,
			OrgCode:          orgCode,
			LegalPersonName:  legalName,
			LegalPersonPhone: legalPhone,
			Owner:            owner,
			Info:             orgInfo,
		})
		if err != nil {
			return err
		}
		if err := q.GrantRole(ctx, &db.GrantRoleParams{
			Username: sql.NullString{
				String: owner,
				Valid:  true,
			},
			Rolename: sql.NullString{
				String: constants.BossRole,
				Valid:  true,
			},
		}); err != nil {
			if !verr.IsDuplicate(err) {
				return err
			}
		}
		_, err = q.AddMember(ctx, &db.AddMemberParams{
			Username: sql.NullString{
				String: owner,
				Valid:  true,
			},
			OrgID: sql.NullInt32{
				Int32: int32(_orgId),
				Valid: true,
			}})
		if err != nil {
			return err
		}

		orgId = int(_orgId)
		return err
	}); err != nil {
		return nil, err
	}

	key := cache.GenKey("iam", "org", "id", int(orgId))
	orgI, err := c.repo.CacheGet(ctx, key, func(q db.Querier) (any, error) {
		org, err := c.repo.GetOrgByID(ctx, int32(orgId))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorOrgNotExist(err.Error())
			}
			return nil, err
		}
		return org, nil
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	externalOrg := orgI.(*db.Org)

	return externalOrg, nil
}

// AddMember 企业增加成员
func (c *IamCase) AddMember(ctx context.Context, orgId int, username string) error {
	if err := c.repo.ExecTx(ctx, func(q *db.Queries) error {
		_, err := q.AddMember(ctx, &db.AddMemberParams{
			Username: sql.NullString{
				String: username,
				Valid:  true,
			},
			OrgID: sql.NullInt32{
				Int32: int32(orgId),
				Valid: true,
			},
		})
		if err != nil {
			return err
		}
		if err := q.GrantRole(ctx, &db.GrantRoleParams{
			Username: sql.NullString{
				String: username,
				Valid:  true,
			},
			Rolename: sql.NullString{
				String: constants.ProducerRole,
				Valid:  true,
			},
		}); err != nil {
			if !verr.IsDuplicate(err) {
				return err
			}
		}
		return nil
	}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// GetOrg 获取企业信息
func (c *IamCase) GetOrg(ctx context.Context, orgId int) (*db.Org, error) {
	key := cache.GenKey("iam", "org", "id", orgId)
	orgI, err := c.repo.CacheGet(ctx, key, func(q db.Querier) (any, error) {
		org, err := c.repo.GetOrgByID(ctx, int32(orgId))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorOrgNotExist(err.Error())
			}

			return nil, err
		}

		return org, nil
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return orgI.(*db.Org), nil
}

// DeleteOrgMember 企业删除成员
func (c *IamCase) DeleteOrgMember(ctx context.Context, org_id int, username string) error {
	if err := c.repo.ExecTx(ctx, func(q *db.Queries) error {
		if err := c.repo.DeleteOrgMember(ctx, &db.DeleteOrgMemberParams{
			OrgID: sql.NullInt32{
				Int32: int32(org_id),
				Valid: true,
			},
			Username: sql.NullString{
				String: username,
				Valid:  true,
			},
		}); err != nil {
			return err
		}

		if err := c.repo.RemoveRole(ctx, sql.NullString{
			String: username,
			Valid:  true,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// ListOrgMember 企业查询成员列表
func (c *IamCase) ListOrgMember(ctx context.Context, org_id, offset, limit int) ([]*db.User, error) {
	rows, err := c.repo.ListOrgMember(ctx, &db.ListOrgMemberParams{
		Orgid: sql.NullInt32{
			Int32: int32(org_id),
			Valid: true,
		},
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	users := make([]*db.User, 0)
	for _, row := range rows {
		users = append(users, &db.User{
			ID:        row.ID,
			Username:  row.Username,
			Nickname:  row.Nickname,
			Email:     row.Email,
			CreatedAt: row.CreatedAt,
			Realname:  row.Realname,
			Idcard:    row.Idcard,
			Phone:     row.Phone,
		})
	}
	return users, nil
}

// UpdateOrg 企业信息更新
func (c *IamCase) UpdateOrg(ctx context.Context, org *db.Org) (*db.Org, error) {
	key := cache.GenKey("iam", "org", "id", org.ID)
	if err := c.repo.CacheUpdate(ctx, key, func(q db.Querier) error {
		if err := c.repo.UpdateOrg(ctx, &db.UpdateOrgParams{
			OrgName:          org.OrgName,
			OrgCode:          org.OrgCode,
			LegalPersonName:  org.LegalPersonName,
			LegalPersonPhone: org.LegalPersonPhone,
			Owner:            org.Owner,
			Info:             org.Info,
			ID:               org.ID,
		}); err != nil {
			logger.Errorf("企业不存在： %s\n", org)
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return org, nil
}

// GetOrgOfUser 查询用户所属企业
func (c *IamCase) GetOrgOfUser(ctx context.Context, username string) (*db.Org, error) {
	org, err := c.repo.GetOrgOfUser(ctx, sql.NullString{
		String: username,
		Valid:  true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrorUserNotFound("用户：%s 不存在\n", username)
		}
		return nil, err
	}
	return org, nil
}
