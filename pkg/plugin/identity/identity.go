package identity

import "context"

type IdentityAuther interface {
	// 用于个人实名认证
	RealNameAuth(context.Context, *RealNameRequest) (*RealNameResponse, error)

	// 用于企业认证
	EnterpriseAuth(context.Context, *EnterpriseAuthRequest) (*EnterpriseAuthResponse, error)
}

type RealNameRequest struct{}

type RealNameResponse struct{}

type EnterpriseAuthRequest struct{}

type EnterpriseAuthResponse struct{}
