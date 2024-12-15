// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: model.sql

package db

import (
	"context"
	"database/sql"
)

const getModelByID = `-- name: GetModelByID :one
SELECT
    id, name, version, status, des, metadata
FROM
    model
WHERE
    id = ?
`

// 根据模型id获取模型数据
func (q *Queries) GetModelByID(ctx context.Context, id int32) (*Model, error) {
	row := q.db.QueryRowContext(ctx, getModelByID, id)
	var i Model
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.Status,
		&i.Des,
		&i.Metadata,
	)
	return &i, err
}

const getModelByName = `-- name: GetModelByName :one
SELECT
    id, name, version, status, des, metadata
FROM
    model
WHERE
    name = ?
`

// 根据模型名获取模型数据
func (q *Queries) GetModelByName(ctx context.Context, name sql.NullString) (*Model, error) {
	row := q.db.QueryRowContext(ctx, getModelByName, name)
	var i Model
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.Status,
		&i.Des,
		&i.Metadata,
	)
	return &i, err
}

const listModels = `-- name: ListModels :many
SELECT id, name, version, status, des, metadata from model
`

func (q *Queries) ListModels(ctx context.Context) ([]*Model, error) {
	rows, err := q.db.QueryContext(ctx, listModels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Model
	for rows.Next() {
		var i Model
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.Status,
			&i.Des,
			&i.Metadata,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateModelStatus = `-- name: UpdateModelStatus :exec
UPDATE model
SET
    status = ?
WHERE
    name = ?
`

type UpdateModelStatusParams struct {
	Status sql.NullString `json:"status"`
	Name   sql.NullString `json:"name"`
}

// UpdateModelStatus 根据模型名更新模型状态
func (q *Queries) UpdateModelStatus(ctx context.Context, arg *UpdateModelStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateModelStatus, arg.Status, arg.Name)
	return err
}
