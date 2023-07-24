-- name: CreateUser :one
INSERT INTO "users" ("email",
                     "password")
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM "users"
WHERE "id" = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM "users"
ORDER BY "id"
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE "users"
SET "password" = $2
WHERE "id" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM "users"
WHERE "id" = $1;

-- name: GetUsersCount :one
SELECT COUNT(*)
FROM "users";
