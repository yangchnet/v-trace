/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package services

import (
	"fmt"
	"time"

	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models/db"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

const TokenExpireSeconds int64 = 7200

type Claims struct {
	AppId   string
	AppRole db.AccessRole
	jwt.StandardClaims
}

type AppInfo struct {
	AppId   string
	AppKey  string
	AppRole db.AccessRole
}

func (c Claims) Valid() error {
	return c.StandardClaims.Valid()
}

// JWT generate token.
func GenToken(appId string, appRole db.AccessRole) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		appId,
		appRole,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(TokenExpireSeconds)).Unix(),
		},
	})
	return token.SignedString([]byte(utils.DefaultTokenSecretKey))
}

// JWT parse token.
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(utils.DefaultTokenSecretKey), nil
	})
	fmt.Println(token)
	fmt.Println(err)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token is not available")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token has expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("invalid token")
			} else {
				return nil, fmt.Errorf("token is not available")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("err invalid token")
}

// GetAppToken get token by appId and appKey.
func GetAppToken(appId, appKey string) (string, error) {
	appInfo, err := models.FindAppInfo(appId)
	if err != nil {
		logger.Error("get app token failed", zap.Error(err))
		return "", fmt.Errorf("get token failed: the Id inputted is wrong")
	}
	if appKey != appInfo.AppKey {
		logger.Error("get app token failed", zap.Error(err))
		return "", fmt.Errorf("get token failed: the key inputted is wrong")
	}
	token, err := GenToken(appId, appInfo.AppRole)
	if err != nil {
		logger.Error("get app token failed", zap.Error(err))
		return "", err
	}
	return token, nil
}

// InitAccessControl.
func InitAccessControl() (bool, error) {
	if !isUseAccessControlFromConfig() {
		logger.Info("the access control module is not enabled")
		return false, nil
	}
	logger.Info("init access control start")
	appInfos, err := checkAccessControlConf()
	if err != nil {
		err = fmt.Errorf("init access control failed: %s", err.Error())
		return false, err
	}

	for _, v := range appInfos {
		logger.Info("init access control", zap.Any("app info", v))
		_, err := models.FindAppInfo(v.AppId)
		if err != nil {
			err := models.InsertAppInfo(&db.AppInfo{
				AppId:   v.AppId,
				AppKey:  v.AppKey,
				AppRole: v.AppRole,
			})
			if err != nil {
				err = fmt.Errorf("init access control failed: %s", err.Error())
				return false, err
			}
		}
	}
	logger.Info("init access control end")
	return true, nil
}
