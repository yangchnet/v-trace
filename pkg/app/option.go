package app

import (
	"context"
	"net/url"
	"os"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/registry"
)

type Option func(*options)

type options struct {
	id       string
	name     string
	version  string
	ctx      context.Context
	metadata map[string]string
	sigs     []os.Signal
	//server           *grpc.Server // now only support single server
	server           Server
	endpoint         *url.URL
	registrar        registry.Registrar
	registrarTimeout time.Duration
}

func WithID(id string) Option {
	return func(op *options) {
		op.id = id
	}
}

func WithName(name string) Option {
	return func(op *options) {
		op.name = name
	}
}

func WithVersion(version string) Option {
	return func(op *options) {
		op.version = version
	}
}

func WithContext(ctx context.Context) Option {
	return func(op *options) {
		op.ctx = ctx
	}
}

func WithMetadata(metadata map[string]string) Option {
	return func(op *options) {
		op.metadata = metadata
	}
}

func WithSigs(sigs ...os.Signal) Option {
	return func(op *options) {
		op.sigs = append([]os.Signal{}, sigs...)
	}
}

func WithServer(server Server) Option {
	return func(op *options) {
		op.server = server
	}
}

func WithConfig(conf *AppConfig) Option {
	return func(op *options) {
		op.id = conf.ID
		op.name = conf.Name
		op.version = conf.Version
	}
}

func WithEndpoint(endpoint *url.URL) Option {
	return func(o *options) { o.endpoint = endpoint }
}

func WithRegistrar(registrar registry.Registrar) Option {
	return func(o *options) {
		o.registrar = registrar
	}
}

func WithRegistrarTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.registrarTimeout = timeout
	}
}
