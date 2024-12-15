package client

//
//import (
//"context"
//"errors"
//"sync"
//
//v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
//localGrpc "gitee.com/qciip-icp/v-trace/pkg/grpc"
//"gitee.com/qciip-icp/v-trace/pkg/logger"
//"gitee.com/qciip-icp/v-trace/pkg/registry"
//)
//
//// var ProviderSet = wire.NewSet(NewIamServiceClient)
//
//var (
//	iamClient v1.IamServiceClient
//	m         = sync.RWMutex{}
//)
//
//func IamServiceClient(ctx context.Context, watcher registry.Watcher) (v1.IamServiceClient, error) {
//	if iamClient != nil {
//		m.RLock()
//		defer m.RUnlock()
//		return iamClient, nil
//	}
//
//	instances, err := watcher.Next()
//	if err != nil {
//		logger.Errorf("registry watcher error: %v", err)
//		return nil, err
//	}
//	if len(instances) <= 0 {
//		logger.Errorf("iam service is note registered\n")
//		return nil, errors.New("iam service is note registered\n")
//	}
//	m.Lock()
//	defer m.Unlock()
//
//	logger.Warnf("iam instance: %v", instances[0])
//
//	conn, err := localGrpc.NewConn(ctx, instances[0].Endpoint)
//	if err != nil {
//		logger.Error("error dial rpc server: %s", instances[0].Endpoint)
//
//		return nil, err
//	}
//	iamClient = v1.NewIamServiceClient(conn)
//
//	go func(ctx context.Context) {
//		for {
//			instances, err := watcher.Next()
//			if err != nil || len(instances) <= 0 {
//				continue
//			}
//			update(ctx, instances[0])
//		}
//	}(ctx)
//
//	return iamClient, nil
//}
//
//func update(ctx context.Context, instance *registry.ServiceInstance) {
//	m.Lock()
//	defer m.Unlock()
//
//	conn, err := localGrpc.NewConn(ctx, instance.Endpoint)
//	if err != nil {
//		logger.Error("error dial rpc server: %s", instance.Endpoint)
//		return
//	}
//	iamClient = v1.NewIamServiceClient(conn)
//}
