-- name: GetRolesByUsername :many
-- GetRoles 获取用户的角色
SELECT
    `rolename`
FROM
    `user_role`
WHERE
    `username` = ?;

-- name: GrantRole :exec
-- GrantRole 为用户赋予角色
INSERT INTO
    `user_role` (`username`, `rolename`, created_at)
VALUES
    (?, ?, now());

-- name: DeleteRelation :exec
-- DeleteRelation 删除某用户的所有角色信息
DELETE FROM
    `user_role`
WHERE
    `username` = ?;

/*
-- name: RemoveRole :exec
-- RemoveRole 企业移除成员角色
DELETE FROM
    user_role
WHERE
        username = ? and rolename = ?;
 */

-- name: RemoveProducerRole :exec
DELETE FROM
    user_role
WHERE
    username = ? and rolename = producer;



