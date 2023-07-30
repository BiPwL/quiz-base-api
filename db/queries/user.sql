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

-- name: ListUserAnsweredQuestions :many
SELECT q.id, q.text, q.hint, q.category, q.created_at
FROM "answered_questions" AS aq
         JOIN "questions" AS q ON aq.question_id = q.id
WHERE aq.user_id = sqlc.arg(user_id)
  AND (sqlc.arg(category)::TEXT = '' OR q.category = sqlc.arg(category)::TEXT)
LIMIT $1 OFFSET $2;
