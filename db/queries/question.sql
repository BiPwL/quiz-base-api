-- name: CreateQuestion :one
INSERT INTO "questions" ("text",
                         "hint",
                         "category")
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetQuestion :one
SELECT *
FROM "questions"
WHERE "id" = $1
LIMIT 1;

-- name: ListQuestions :many
SELECT *
FROM "questions"
ORDER BY "id"
LIMIT $1 OFFSET $2;

-- name: UpdateQuestion :one
UPDATE "questions"
SET "text" = CASE WHEN $3 = 'text' THEN $2 ELSE "text" END,
    "hint" = CASE WHEN $3 = 'hint' THEN $2 ELSE "hint" END
WHERE "id" = $1
RETURNING *;


-- name: DeleteQuestion :exec
DELETE
FROM "questions"
WHERE "id" = $1;