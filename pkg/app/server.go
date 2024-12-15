package app

import "golang.org/x/net/context"

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
