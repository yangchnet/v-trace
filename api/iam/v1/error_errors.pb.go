// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 密码hash错误
func IsPasswdHashFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_PASSWD_HASH_FAILED.String() && e.Code == 500
}

// 密码hash错误
func ErrorPasswdHashFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, Error_PASSWD_HASH_FAILED.String(), fmt.Sprintf(format, args...))
}

// 创建用户失败
func IsCreateUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_CREATE_USER_FAILED.String() && e.Code == 500
}

// 创建用户失败
func ErrorCreateUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, Error_CREATE_USER_FAILED.String(), fmt.Sprintf(format, args...))
}

// 用户未找到
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_USER_NOT_FOUND.String() && e.Code == 404
}

// 用户未找到
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, Error_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 登录密码错误
func IsPasswdErr(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_PASSWD_ERR.String() && e.Code == 400
}

// 登录密码错误
func ErrorPasswdErr(format string, args ...interface{}) *errors.Error {
	return errors.New(400, Error_PASSWD_ERR.String(), fmt.Sprintf(format, args...))
}

// 参数错误
func IsParamsErr(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_PARAMS_ERR.String() && e.Code == 400
}

// 参数错误
func ErrorParamsErr(format string, args ...interface{}) *errors.Error {
	return errors.New(400, Error_PARAMS_ERR.String(), fmt.Sprintf(format, args...))
}

// 不存在的角色
func IsRoleNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_ROLE_NOT_EXIST.String() && e.Code == 400
}

// 不存在的角色
func ErrorRoleNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(400, Error_ROLE_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 企业不存在
func IsOrgNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_ORG_NOT_EXIST.String() && e.Code == 404
}

// 企业不存在
func ErrorOrgNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(404, Error_ORG_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 数据已存在
func IsEntryDuplicate(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_ENTRY_DUPLICATE.String() && e.Code == 400
}

// 数据已存在
func ErrorEntryDuplicate(format string, args ...interface{}) *errors.Error {
	return errors.New(400, Error_ENTRY_DUPLICATE.String(), fmt.Sprintf(format, args...))
}

// token过期
func IsTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_TOKEN_EXPIRED.String() && e.Code == 401
}

// token过期
func ErrorTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(401, Error_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

// 更新用户信息失败
func IsRefreshUserFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_REFRESH_USER_FAILED.String() && e.Code == 500
}

// 更新用户信息失败
func ErrorRefreshUserFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, Error_REFRESH_USER_FAILED.String(), fmt.Sprintf(format, args...))
}
