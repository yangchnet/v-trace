-- name: CreateUser :execlastid
-- CreateUser 创建一个用户
INSERT INTO
    `user` (`username`, `nickname`, `passwd`, `phone`, `email`, created_at, `avatar`, `status`)
VALUES
    (?, ?, ?, ?, ?, now(), ?, ?);

-- name: UpdateIdentity :exec
-- UpdateIdentity 更新用户实名信息
UPDATE
    `user`
SET
    realname = ?, idcard = ?
WHERE
    username = ?;

-- name: UpdatePasswdByUsername :execlastid
-- UpdatePasswdByUsername 根据用户名更新密码
UPDATE
    `user`
SET
    `passwd` = ?
WHERE
    `username` = ?;

-- name: GetUserByUsername :one
-- GetUserByUsername 根据用户名获取用户信息
SELECT
    *
FROM
    `user`
WHERE `username` = ?;

-- name: GetUserByPhone :one
-- GetUserByPhone 根据手机号获取用户
SELECT
    *
FROM
    `user`
WHERE `phone` = ?;

-- name: GetUserByID :one
-- GetUserByPhone 根据ID获取用户
SELECT
    *
FROM
    `user`
WHERE `id` = ?;


-- name: DeleteUserByID :exec
-- DeleteUserByUsername 根据用户id删除用户
UPDATE
    user
SET
    status = ?
WHERE
    username = ?;

-- name: UpdateUser :exec
-- UpdateUser 用户信息更新
UPDATE
    user
SET
    `nickname` = ?, `passwd` = ?, avatar = ?
where
    `username` = ?;
