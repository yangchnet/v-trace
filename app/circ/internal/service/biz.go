package service

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/circ/internal/data/db"
)

type CircCaseInterface interface {
	RecordInterface
}

type RecordInterface interface {
	CreateRecord(ctx context.Context, transId string, circType string, operator string, from string, to string, formValue []byte) (*db.CircRecord, error)
	BatchRecord(ctx context.Context, transIds []string, operator, from, to string, circType string, formValue []byte) ([]int, error)
	GetRecordByTransId(ctx context.Context, transId string) (*db.CircRecord, error)
	GetRecordByID(ctx context.Context, circId int) (*db.CircRecord, error)
	GetRecordByGoodsId(ctx context.Context, goodsId int) ([]*db.CircRecord, error)
	TransId(ctx context.Context, goodsId int32) (string, error)
	BatchTransId(ctx context.Context, goodsIds []int32) (map[int32]string, error)
	UpdateCircStatus(ctx context.Context, transId, status string) error
}
