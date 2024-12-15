-- name: ListMaterials :many
SELECT * from material;

-- name: GetMaterialByID :one
SELECT
    *
FROM
    material
WHERE
    id = ?;