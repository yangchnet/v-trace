-- name: GetMaterialID :one
-- 根据model id和index获取条目
SELECT
    material_id
FROM
    relation
WHERE
    `model_id` = ? AND `index` = ?;

-- name: GetRelationByMaterial :many
-- 根据material获取条目
SELECT
    *
FROM
    relation
WHERE
   material_id = ?;

-- name: GetAvaiableModel :many
-- GetAvaiableModel 根据material id获取其可用的模型
SELECT
    a.*
FROM
    model a JOIN relation b
ON
    a.id = b.model_id
WHERE
    b.material_id = ?;