-- name: CreateOrg :execlastid
-- CreateOrg 创建一个用户
INSERT INTO
    `org` (`org_name`, `org_code`, `legal_person_name`, `legal_person_phone`, `created_at`, `owner`, `info`)
VALUES
    (?, ?, ?, ?, now(), ?, ?);


-- name: AddMember :execlastid
-- AddMember 企业增加成员
INSERT INTO
    `member` (`username`, `org_id`, `created_at`)
VALUES
    (?, ?, now());

-- name: GetOrgByID :one
-- GetOrgByID 根据组织id获取组织信息
SELECT
    *
FROM
    `org`
WHERE `id` = ?;

-- name: DeleteMember :exec
-- DeleteMember 删除成员
DELETE FROM
    `member`
WHERE
    `username` = ?;

-- name: DeleteOrgMember :exec
-- DeleteOrgMember 企业移除成员
DELETE FROM
    member
WHERE
    org_id = ? and username = ?;

-- name: ListOrgMember :many
-- ListOrgMember 企业查询成员列表
SELECT
	`user`.`id`, `user`.`username`, `user`.`nickname`, `user`.`phone`, `user`.`email`, `user`.`created_at`, `user`.`realname`, `user`.`idcard`
FROM
	`member` JOIN `user`
ON
	`member`.`username` = `user`.`username`
WHERE
	`member`.`org_id` = sqlc.arg(orgId)
ORDER BY
	`user`.`id`
LIMIT
	?, ?;

-- name: UpdateOrg :exec
-- UpdateOrg 企业信息更新
UPDATE
    org
SET org_name = ?,
    org_code = ?,
    legal_person_name = ?,
    legal_person_phone = ?,
    created_at = now(),
    owner = ?,
    info = ?
WHERE
id = ?;

-- name: GetOrgOfUser :one
-- 查询用户所属企业
SELECT
    org.*
FROM
    org join member
WHERE
    member.username = ? and member.org_id = org.id
