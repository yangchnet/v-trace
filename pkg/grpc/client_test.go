package grpc

import (
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	ctx := context.Background()

	r := registry.NewRegistry(ctx, &registry.Etcd{
		Addresses: []string{"127.0.0.1:2379"},
	})

	iamWatcher, err := r.Watch(ctx, "iam")
	if err != nil {
		t.Fatal(err)
	}

	g := NewConnGetter(ctx, iamWatcher, "iam")

	for {
		g.Get()

		time.Sleep(time.Second)
	}
}
