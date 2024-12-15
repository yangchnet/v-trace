package grpc

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/tools/gotools"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var ClientOptions = []grpc.DialOption{
	grpc.WithTransportCredentials(insecure.NewCredentials()),
	grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}),
}

func NewClient(ctx context.Context, host string, port int) (*grpc.ClientConn, error) {
	endpoint := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.DialContext(ctx, endpoint,
		ClientOptions...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewConn(ctx context.Context, endpoint string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, endpoint,
		ClientOptions...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type ConnGetInterface interface {
	Get() grpc.ClientConnInterface
}

type offlineConn struct {
	name string
}

func (o offlineConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	return errors.New(fmt.Sprintf("%s 服务未上线！", o.name))
}

func (o offlineConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, errors.New(fmt.Sprintf("%s 服务未上线！", o.name))
}

var _ ConnGetInterface = (*ConnGetter)(nil)

type connEntry struct {
	id   string
	conn grpc.ClientConnInterface
}

type ConnGetter struct {
	sync.RWMutex

	conns       []*connEntry
	id2conn     sync.Map
	offlineConn grpc.ClientConnInterface

	watcher registry.Watcher

	ctx context.Context

	name string
}

func NewConnGetter(ctx context.Context, watcher registry.Watcher, name string) *ConnGetter {
	getter := &ConnGetter{
		conns:       make([]*connEntry, 0),
		id2conn:     sync.Map{},
		offlineConn: offlineConn{name: name},
		watcher:     watcher,
		RWMutex:     sync.RWMutex{},

		ctx: ctx,

		name: name,
	}

	gotools.Go(func() {
		getter.listen(ctx)
	})

	return getter
}

func (g *ConnGetter) Get() grpc.ClientConnInterface {
	g.RLock()
	defer g.RUnlock()

	if len(g.conns) <= 0 {
		return g.offlineConn
	}

	target := rand.Intn(len(g.conns))

	if target < len(g.conns) {
		return g.conns[target].conn
	}

	return g.offlineConn
}

func (g *ConnGetter) listen(ctx context.Context) {
	for {
		add, deleted, err := g.watcher.Next()
		if err != nil {
			continue
		}

		if len(add) > 0 {
			g.update(ctx, add)
		}

		if len(deleted) > 0 {
			g.drop(ctx, deleted)
		}
	}
}

func (g *ConnGetter) drop(ctx context.Context, keys []string) {
	g.Lock()
	defer g.Unlock()

	for _, key := range keys {
		instanceId := strings.Split(key, "/")[2]
		if _, ok := g.id2conn.LoadAndDelete(instanceId); ok {
			logger.Warnf("将删除[%s]服务的[%s]实例", g.name, instanceId)
			for i, e := range g.conns {
				if e.id == instanceId {
					g.conns = append(g.conns[:i], g.conns[i+1:]...)
					break
				}
			}
		}
	}

}

func (g *ConnGetter) update(ctx context.Context, instance []*registry.ServiceInstance) {
	g.Lock()
	defer g.Unlock()

	for _, ins := range instance {
		conn, err := NewConn(ctx, ins.Endpoint)
		if err != nil {
			logger.Error("error dial rpc server: %s", ins.Endpoint)
			continue
		}

		if _, ok := g.id2conn.LoadOrStore(ins.ID, true); !ok {
			logger.Warnf("将添加[%s]服务的[%s]实例", g.name, ins.ID)
			g.conns = append(g.conns, &connEntry{
				id:   ins.ID,
				conn: conn,
			})
		}
	}
}
