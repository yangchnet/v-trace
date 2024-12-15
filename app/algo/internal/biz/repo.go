package biz

import "gitee.com/qciip-icp/v-trace/app/algo/internal/data"

//go:generate mockgen -destination=mock_repo.go -package=biz . AlgoRepo
type AlgoRepo interface {
	data.Store
}
