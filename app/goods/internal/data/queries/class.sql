-- name: CreateClass :execlastid
-- 创建一个产品类别记录
INSERT INTO
    class (name, des, status, creator, material_id, org_id, tm, created_at)
VALUES
    (?, ?, ?, ?, ?, ?, ?, now());

-- name: GetClassById :one
-- 根据id获取产品类别
SELECT
    *
FROM
    class
WHERE id = ?;

-- name: ListGoodClass :many
-- 列出产品类型
SELECT
    *
FROM
    `class`
WHERE
    org_id = ?
LIMIT  ?, ?;


-- name: UpdateGoodClass :exec
-- 更新产品类型
UPDATE
    class
SET name = ?, des = ?, status = ?, creator = ?, org_id = ?, tm = ?
WHERE
    material_id = ?;

-- name: GetOrgOfClass :one
-- 获取产品所属企业
SELECT
    org_id
FROM
    `class`
WHERE
    id = ?;