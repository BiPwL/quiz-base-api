-- name: CreateCategory :one
INSERT INTO "categories" ("name",
                          "key")
VALUES ($1, $2)
RETURNING *;

-- name: GetCategory :one
SELECT *
FROM "categories"
WHERE "key" = $1
LIMIT 1;

-- name: ListCategories :many
SELECT *
FROM "categories"
ORDER BY "key"
LIMIT $1 OFFSET $2;

-- name: UpdateCategory :one
UPDATE "categories"
SET "name" = $2
WHERE "key" = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE
FROM "categories"
WHERE "key" = $1;