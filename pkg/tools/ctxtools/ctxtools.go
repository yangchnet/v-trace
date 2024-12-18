package ctxtools

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	messageIdKey = "x-message-id"
	requestIdKey = "x-request-id"
	forwardKey   = "x-forwarded-for"
)

type getMetadataFromContext func(ctx context.Context) (md metadata.MD, ok bool)

var getMetadataFromContextFunc = []getMetadataFromContext{
	metadata.FromOutgoingContext,
	metadata.FromIncomingContext,
}

func GetValueFromContext(ctx context.Context, key string) []string {
	if ctx == nil {
		return []string{}
	}
	for _, f := range getMetadataFromContextFunc {
		md, ok := f(ctx)
		if !ok {
			continue
		}
		m, ok := md[key]
		if ok && len(m) > 0 {
			return m
		}
	}
	m, ok := ctx.Value(key).([]string)
	if ok && len(m) > 0 {
		return m
	}
	s, ok := ctx.Value(key).(string)
	if ok && len(s) > 0 {
		return []string{s}
	}

	return []string{}
}

func GetSenderFromCtx(ctx context.Context) string {
	username := GetValueFromContext(ctx, "username")
	if len(username) > 0 {
		return username[0]
	}

	return ""
}

func WithMetadata(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	return metadata.NewOutgoingContext(ctx, md)
}
