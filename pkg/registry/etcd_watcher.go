package registry

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var _ Watcher = &watcher{}

type watcher struct {
	key         string
	ctx         context.Context
	cancel      context.CancelFunc
	watchChan   clientv3.WatchChan
	watcher     clientv3.Watcher
	kv          clientv3.KV
	first       bool
	serviceName string
}

func newWatcher(ctx context.Context, key, name string, client *clientv3.Client) (*watcher, error) {
	w := &watcher{
		key:         key,
		watcher:     clientv3.NewWatcher(client),
		kv:          clientv3.NewKV(client),
		first:       true,
		serviceName: name,
	}
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.watchChan = w.watcher.Watch(w.ctx, key, clientv3.WithPrefix())
	err := w.watcher.RequestProgress(context.Background())
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *watcher) Next() (add []*ServiceInstance, deleted []string, err error) {
	if w.first {
		item, err := w.getInstance()
		w.first = false
		return item, nil, err
	}

	select {
	case <-w.ctx.Done():
		return nil, nil, w.ctx.Err()
	case event := <-w.watchChan:
		if len(event.Events) > 0 {
			for i := 0; i < len(event.Events); i++ {
				switch event.Events[i].Type {
				case mvccpb.PUT:
					si, err := unmarshal(event.Events[i].Kv.Value)
					if err != nil {
						return nil, nil, err
					}
					add = append(add, si)
				case mvccpb.DELETE:
					deleted = append(deleted, string(event.Events[i].Kv.Key))
				}
			}

		}

		//add, err = w.getInstance()

		return
	}
}

func (w *watcher) Stop() error {
	w.cancel()
	return w.watcher.Close()
}

func (w *watcher) getInstance() ([]*ServiceInstance, error) {
	resp, err := w.kv.Get(w.ctx, w.key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	items := make([]*ServiceInstance, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		si, err := unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		if si.Name != w.serviceName {
			continue
		}
		items = append(items, si)
	}
	return items, nil
}
