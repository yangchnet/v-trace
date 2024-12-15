package biz

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	passwdtools "gitee.com/qciip-icp/v-trace/pkg/tools/passwd"
)

func (uc *IamCase) CreateToken(ctx context.Context, phone, passwd string) (string, error) {
	user, err := uc.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", v1.ErrorUserNotFound(fmt.Sprintf("不存在的用户: %s", phone))
		}
		return "", err
	}

	if err := passwdtools.CheckPassword(passwd, user.Passwd); err != nil {
		logger.Error("check password failed\n", err)

		return "", v1.ErrorPasswdErr("密码错误\n")
	}

	roleStrings, err := uc.repo.GetRolesByUsername(ctx, sql.NullString{
		String: user.Username,
		Valid:  true,
	})
	if err != nil {
		logger.Error(err)

		return "", err
	}

	roles := make([]string, 0)
	for i := range roleStrings {
		roles = append(roles, roleStrings[i].String)
	}

	token, err := uc.tokenMaker.CreateToken(user.Username, constants.ShouldRole(roles))
	if err != nil {
		logger.Errorf("create token failed: %v\n", err)

		return "", err
	}

	return token, nil
}

// RefreshToken 更新令牌
func (uc *IamCase) RefreshToken(ctx context.Context, username string) (string, error) {
	user, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Errorf("find user by token failed: %v\n", err)
		return "", err
	}

	roleStrings, err := uc.repo.GetRolesByUsername(ctx, sql.NullString{
		String: user.Username,
		Valid:  true,
	})

	roles := make([]string, 0)
	for i := range roleStrings {
		roles = append(roles, roleStrings[i].String)
	}

	token, err := uc.tokenMaker.CreateToken(user.Username, constants.ShouldRole(roles))
	if err != nil {
		logger.Errorf("refresh token failed: %v\n", err)
		return "", err
	}
	return token, nil
}
