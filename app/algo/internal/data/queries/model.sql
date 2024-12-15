-- name: ListModels :many
SELECT * from model;

-- name: UpdateModelStatus :exec
-- UpdateModelStatus 根据模型名更新模型状态
UPDATE model
SET
    status = ?
WHERE
    name = ?;

-- name: GetModelByName :one
-- 根据模型名获取模型数据
SELECT
    *
FROM
    model
WHERE
    name = ?;

-- name: GetModelByID :one
-- 根据模型id获取模型数据
SELECT
    *
FROM
    model
WHERE
    id = ?;