package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"golang.org/x/sync/errgroup"
)

// App is a micro service which must implement registry.Registrar.
type App struct {
	// ctx
	ctx context.Context

	// options
	opts *options

	cancel func()

	instance *registry.ServiceInstance

	Err error
}

func New(opts ...Option) *App {
	o := options{
		sigs: []os.Signal{syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL},
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &App{
		ctx:  context.Background(),
		opts: &o,
		Err:  nil,
	}
}

func (a *App) Run() error {
	ctx := context.Background()

	eg, ctx := errgroup.WithContext(ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)
	eg.Go(func() error {
		wg.Done()

		return a.opts.server.Start(ctx)
	})

	wg.Wait()

	instance := a.buildInstance()
	if a.opts.registrar != nil {
		rctx, rcancel := context.WithTimeout(ctx, a.opts.registrarTimeout)
		defer rcancel()
		if err := a.opts.registrar.Register(rctx, instance); err != nil {
			return err
		}
		a.instance = instance
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Stop() {
	// 接收到信号后关闭服务器，若有请求正在服务，则等待请求服务完成或间隔一个时间间隔后再关闭
	a.opts.server.Stop(a.ctx)

	if a.opts.registrar != nil {
		if err := a.opts.registrar.Deregister(context.TODO(), a.instance); err != nil {
			fmt.Println(err)
		}
	}
}

func (a *App) buildInstance() *registry.ServiceInstance {
	s, ok := a.opts.server.(*grpc.Server)
	if !ok {
		return &registry.ServiceInstance{}
	}

	return &registry.ServiceInstance{
		ID:       a.opts.id,
		Name:     a.opts.name,
		Version:  a.opts.version,
		Metadata: a.opts.metadata,
		Endpoint: s.Endpoint().Host,
	}
}
