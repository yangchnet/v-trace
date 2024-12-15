package biz

import (
	"context"
	"database/sql"
	"fmt"

	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
	"github.com/pkg/errors"
)

// GetTransRecords.
func (uc *TransCase) GetTransRecord(ctx context.Context, transId string) (*db.TransRecord, error) {
	record, err := uc.repo.GetTransByTransId(ctx, transId)
	if err != nil {
		logger.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrorRecordNotFound("不存在的transId:%s", transId)
		}

		return nil, err
	}

	return record, nil
}

// UpdateTrans.
func (uc *TransCase) UpdateTrans(ctx context.Context, transId, txHash string, success bool) error {
	var status string
	if success {
		status = v1.TransStatus_success.String()
	} else {
		status = v1.TransStatus_failed.String()
	}

	if _, err := uc.repo.UpdateTrans(ctx, &db.UpdateTransParams{
		Txhash: fmt.Sprintf("0x%s", txHash),
		Status: sql.NullString{
			String: status,
			Valid:  true,
		},
		Transid: transId,
	}); err != nil {
		logger.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return v1.ErrorRecordNotFound("不存在的transId:%s", transId)
		}

		if verr.IsDuplicate(err) {
			return v1.ErrorDuplicateErr("重复的记录:%v", err)
		}

		return err
	}

	return nil
}
