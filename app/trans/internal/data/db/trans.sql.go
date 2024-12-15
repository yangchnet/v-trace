// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: trans.sql

package db

import (
	"context"
	"database/sql"
)

const createTrans = `-- name: CreateTrans :execlastid
INSERT INTO
    ` + "`" + `trans_record` + "`" + ` (` + "`" + `transId` + "`" + `, ` + "`" + `sender` + "`" + `, ` + "`" + `contract` + "`" + `, ` + "`" + `method` + "`" + `, ` + "`" + `params` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `txHash` + "`" + `, ` + "`" + `tx_params_hash` + "`" + `, created_at)
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, now())
`

type CreateTransParams struct {
	Transid      string
	Sender       string
	Contract     string
	Method       string
	Params       sql.NullString
	Status       sql.NullString
	Txhash       string
	TxParamsHash string
}

// CreateTrans 创建交易记录
func (q *Queries) CreateTrans(ctx context.Context, arg *CreateTransParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createTrans,
		arg.Transid,
		arg.Sender,
		arg.Contract,
		arg.Method,
		arg.Params,
		arg.Status,
		arg.Txhash,
		arg.TxParamsHash,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const getTransByTransId = `-- name: GetTransByTransId :one
SELECT
    id, transid, sender, contract, method, params, status, created_at, updated_at, txhash, tx_params_hash
FROM
    ` + "`" + `trans_record` + "`" + `
WHERE
    ` + "`" + `transId` + "`" + ` = ?
`

// 根据transId获取记录
func (q *Queries) GetTransByTransId(ctx context.Context, transid string) (*TransRecord, error) {
	row := q.db.QueryRowContext(ctx, getTransByTransId, transid)
	var i TransRecord
	err := row.Scan(
		&i.ID,
		&i.Transid,
		&i.Sender,
		&i.Contract,
		&i.Method,
		&i.Params,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Txhash,
		&i.TxParamsHash,
	)
	return &i, err
}

const updateTrans = `-- name: UpdateTrans :execlastid
UPDATE
    ` + "`" + `trans_record` + "`" + `
SET
    ` + "`" + `txHash` + "`" + ` = ?, ` + "`" + `status` + "`" + ` = ?, ` + "`" + `updated_at` + "`" + ` = now()
WHERE
    ` + "`" + `transId` + "`" + ` = ?
`

type UpdateTransParams struct {
	Txhash  string
	Status  sql.NullString
	Transid string
}

// 根据transId更新交易状态
func (q *Queries) UpdateTrans(ctx context.Context, arg *UpdateTransParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updateTrans, arg.Txhash, arg.Status, arg.Transid)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}