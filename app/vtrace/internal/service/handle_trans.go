package service

// import (
// 	"context"

// 	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
// 	"gitee.com/qciip-icp/v-trace/pkg/logger"
// 	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
// )

// func (s *VTraceService) ContractName(ctx context.Context, req *v1.ContractNameRequest) (*v1.ContractNameResponse, error) {
// 	name, err := s.cas.ContractName(ctx)
// 	if err != nil {
// 		logger.Error(err)

// 		return nil, err
// 	}

// 	return &v1.ContractNameResponse{
// 		ContractName: pbtools.ToProtoString(name),
// 	}, nil
// }
