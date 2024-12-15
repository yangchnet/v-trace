-- name: CreateRecord :execlastid
INSERT INTO
    circ_record (`transId`, `objectId`, `circType`, `operator`, `from`, `to`, `formValue`, `times`, `status`, `created_at`)
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, now());

-- name: GetRecordById :one
-- 根据id查找流转记录
SELECT
    *
FROM
    circ_record
WHERE
    id = ?;

-- name: GetRecordByTransId :one
-- 根据transId查找流转记录
SELECT
    *
FROM
    circ_record
WHERE transId = ?;

-- name: GetRecordByObjIdDesc :many
-- 根据流转实体id查找流转记录, 按流转顺序降序排列，最新的在第一个
SELECT
    *
FROM
    circ_record
WHERE
    objectId = ? AND status="success"
ORDER BY
    times
DESC;

-- name: GetObjOwner :one
-- 获取商品当前owner
SELECT
    `to`
FROM
    circ_record
WHERE
    objectId = ? AND status="success"
ORDER BY
    times
DESC LIMIT 1;

-- name: UpdateStatus :exec
-- 更新流转记录状态
UPDATE
    circ_record
SET
    status=sqlc.arg(status)
WHERE
    transId = sqlc.arg(transId);
