-- name: CreateSerial :execlastid
-- 创建一个产品批次
INSERT INTO
    serial (product_time, creator, class_id, created_at)
VALUES
    (?, ?, ?, now());

-- name: GetSerialById :one
SELECT
    *
FROM
    serial
WHERE id = ?;

-- name: ListGoodsSerial :many
-- 列出产品批次
SELECT
    s.*, c.*
FROM
    `class` c JOIN `serial` s
ON
    c.`id` = s.`class_id`
WHERE
    c.`org_id` = ?
LIMIT ?, ?;

-- name: UpdateGoodSerial :exec
-- 更新产品批次
UPDATE
    serial
SET product_time = ?, creator = ?
WHERE
    class_id = ?;

-- name: GetOrgOfSerial :one
SELECT
    org_id
FROM
    `class` JOIN `serial`
ON
    `class`.`id` = `serial`.`class_id`
WHERE
    `serial`.`id` = ?;