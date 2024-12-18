// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 不允许的操作
func IsPermissionDenied(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_PERMISSION_DENIED.String() && e.Code == 403
}

// 不允许的操作
func ErrorPermissionDenied(format string, args ...interface{}) *errors.Error {
	return errors.New(403, Error_PERMISSION_DENIED.String(), fmt.Sprintf(format, args...))
}

// 服务未上线
func IsServiceOffline(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error_SERVICE_OFFLINE.String() && e.Code == 500
}

// 服务未上线
func ErrorServiceOffline(format string, args ...interface{}) *errors.Error {
	return errors.New(500, Error_SERVICE_OFFLINE.String(), fmt.Sprintf(format, args...))
}
