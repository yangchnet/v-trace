package biz

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/idtools"
	"gitee.com/qciip-icp/v-trace/pkg/tools/passwd"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
	"github.com/google/uuid"
)

// CreateUser 创建一个新用户.
func (u *IamCase) CreateUser(ctx context.Context, nickname, realPasswd, phone, email, avatar, status string) (*db.User, error) {
	// 加密密码
	hashedPasswd, err := passwd.HashPassword(realPasswd)
	if err != nil {
		logger.Error(err)

		return nil, v1.ErrorPasswdHashFailed("hash密码失败: %v", err)
	}

	username := fmt.Sprintf("user-%s", idtools.NewId())

	var externalUser *db.User
	key := cache.GenKey("iam", "user", "username", username)

	if err := u.repo.ExecTx(ctx, func(q *db.Queries) error {
		userI, err := u.repo.CacheCreate(ctx, key, func(q db.Querier) (any, error) {
			id, err := u.repo.CreateUser(ctx, &db.CreateUserParams{
				Username: username,
				Nickname: sql.NullString{
					String: nickname,
					Valid:  true,
				},
				Passwd: hashedPasswd,
				Phone:  phone,
				Email: sql.NullString{
					String: uuid.New().String(),
					Valid:  true,
				},
				Avatar: sql.NullString{
					String: avatar,
					Valid:  true,
				},
				Status: sql.NullString{
					String: status,
					Valid:  true,
				},
			})
			if err != nil {
				return nil, v1.ErrorCreateUserFailed("创建用户失败: %v", err)
			}

			if err := u.repo.GrantRole(ctx, &db.GrantRoleParams{
				Username: sql.NullString{
					String: username,
					Valid:  true,
				},
				Rolename: sql.NullString{
					String: constants.NormalRole,
					Valid:  true,
				},
			}); err != nil {
				if !verr.IsDuplicate(err) {
					return nil, err
				}
			}

			user, err := u.repo.GetUserByID(ctx, int32(id))
			if err != nil {
				logger.Error(err)
				if errors.Is(err, sql.ErrNoRows) {
					return nil, v1.ErrorUserNotFound("不存在的用户id:%d", id)
				}

				return nil, err
			}

			return user, nil
		})
		if err != nil {
			return err
		}

		externalUser = userI.(*db.User)

		return nil
	}); err != nil {
		return nil, err
	}

	return externalUser, nil
}

// UpdateUser 用户信息更新
func (u *IamCase) UpdateUser(ctx context.Context, user *db.User) (*db.User, error) {
	// 加密密码
	hashedPasswd, err := passwd.HashPassword(user.Passwd)
	if err != nil {
		logger.Error(err)
		return nil, v1.ErrorPasswdHashFailed("hash密码失败: %v", err)
	}

	key := cache.GenKey("iam", "user", "username", user.Username)
	if err := u.repo.ExecTx(ctx, func(queries *db.Queries) error {
		return u.repo.CacheUpdate(ctx, key, func(q db.Querier) error {
			if err := queries.UpdateUser(ctx, &db.UpdateUserParams{
				Nickname: sql.NullString{
					String: user.Nickname.String,
					Valid:  user.Nickname.Valid,
				},
				Passwd:   hashedPasswd,
				Username: user.Username,
				Avatar: sql.NullString{
					String: user.Avatar.String,
					Valid:  true,
				},
			}); err != nil {
				return v1.ErrorRefreshUserFailed("用户信息更新失败: %v", err)
			}
			return nil
		})
	}); err != nil {
		logger.Error(err)
		return nil, err
	}
	return user, nil
}

// GetUser 根据唯一用户名查询用户.
func (u *IamCase) GetUser(ctx context.Context, username string) (*db.User, error) {
	key := cache.GenKey("iam", "user", "username", username)

	userI, err := u.repo.CacheGet(ctx, key, func(q db.Querier) (any, error) {
		user, err := q.GetUserByUsername(ctx, username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorUserNotFound("不存在的用户名:%s", username)
			}
			return nil, err
		}

		return user, nil
	})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return userI.(*db.User), nil
}

// DeleteUser 删除用户.
func (u *IamCase) DeleteUser(ctx context.Context, username string) error {
	key := cache.GenKey("iam", "user", "username", username)
	return u.repo.CacheDelete(ctx, key, func(q db.Querier) error {
		if err := u.repo.DeleteUserByID(ctx, &db.DeleteUserByIDParams{
			Status: sql.NullString{
				String: "deleted",
				Valid:  true,
			},
			Username: username,
		}); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return v1.ErrorUserNotFound("不存在的用户名:%s", username)
			}

			return err
		}

		return nil
	})
}

// GetRolesByUsername 获取用户角色.
func (u *IamCase) GetRolesByUsername(ctx context.Context, username string) (string, error) {
	rets, err := u.repo.GetRolesByUsername(ctx, sql.NullString{
		String: username,
		Valid:  true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", v1.ErrorUserNotFound("不存在的用户名:%s", username)
		}

		return "", err
	}

	roles := make([]string, 0)
	for _, role := range rets {
		roles = append(roles, role.String)
	}

	return constants.ShouldRole(roles), nil
}

// 用户实名认证
func (u *IamCase) CreateIdentity(ctx context.Context, username, realname, id_card string) (*db.User, error) {
	key := cache.GenKey("iam", "user", "username", username)

	var externalUser *db.User

	if err := u.repo.ExecTx(ctx, func(q *db.Queries) error {
		return u.repo.CacheUpdate(ctx, key, func(q db.Querier) error {
			if err := q.UpdateIdentity(ctx, &db.UpdateIdentityParams{
				Realname: sql.NullString{
					String: realname,
					Valid:  true,
				},
				Idcard: sql.NullString{
					String: id_card,
					Valid:  true,
				},
				Username: username,
			}); err != nil {
				return err
			}

			if err := q.GrantRole(ctx, &db.GrantRoleParams{
				Username: sql.NullString{
					String: username,
					Valid:  true,
				},
				Rolename: sql.NullString{
					String: constants.TransporterRole,
					Valid:  true,
				},
			}); err != nil {
				if !verr.IsDuplicate(err) {
					return err
				}
			}

			user, err := u.repo.GetUserByUsername(ctx, username)
			if err != nil {
				return err
			}

			externalUser = user

			return nil
		})

	}); err != nil {
		logger.Error(err)

		return nil, err
	}

	return externalUser, nil
}
