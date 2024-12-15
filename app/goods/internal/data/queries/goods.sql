-- name: CreateGoods :execlastid
-- 创建一个产品
INSERT INTO
    goods (status, creator, serial_id, created_at)
VALUES
    (?, ?, ?, now());

-- name: GetGoodsById :one
SELECT
    *
FROM
    goods
WHERE id = ?;


-- name: ListGoods :many
SELECT
    g.*, s.*, c.*
FROM `class` c
    INNER JOIN `serial` s ON c.id = s.`class_id`
    INNER JOIN `goods` g ON s.id = g.`serial_id`
WHERE
    c.org_id = ?
LIMIT ?, ?;


-- name: UpdateGood :exec
-- 更新产品信息
UPDATE
    goods
SET status = ?, creator = ?
WHERE
    serial_id = ?;

-- name: GetOrgOfGoods :one
-- 获取产品所属企业
SELECT
    org_id
FROM `class` c
    INNER JOIN `serial` s ON c.id = s.`class_id`
    INNER JOIN `goods` g ON s.id = g.`serial_id`
WHERE
    g.id = ?;