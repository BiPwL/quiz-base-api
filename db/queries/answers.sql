-- name: CreateAnswer :one
INSERT INTO "answers" ("question_id",
                       "text",
                       "is_correct")
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnswer :one
SELECT *
FROM "answers"
WHERE "id" = $1
LIMIT 1;

-- name: ListAnswers :many
SELECT *
FROM "answers"
ORDER BY "id"
LIMIT $1 OFFSET $2;

-- name: UpdateAnswer :one
UPDATE "answers"
SET "text"       = CASE WHEN $3 = 'text' THEN $2 ELSE "text" END,
    "is_correct" = CASE WHEN $3 = 'is_correct' THEN $2 ELSE "is_correct" END
WHERE "id" = $1
RETURNING *;

-- name: DeleteAnswer :exec
DELETE
FROM "answers"
WHERE "id" = $1;