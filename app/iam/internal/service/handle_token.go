package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// Token.
func (s *IamService) Token(ctx context.Context, req *v1.TokenRequest) (*v1.TokenResponse, error) {
	phone := req.GetPhone()
	passwd := req.GetPasswd()
	token, err := s.cas.CreateToken(ctx, phone, passwd)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.TokenResponse{
		Token: pbtools.ToProtoString(token),
	}, nil
}

// RefreshToken 更新令牌
func (s *IamService) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	username := ctxtools.GetSenderFromCtx(ctx)
	token, err := s.cas.RefreshToken(ctx, username)
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.RefreshTokenResponse{Token: pbtools.ToProtoString(token)}, nil
}
