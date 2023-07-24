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
SET "text" = $2,
    "hint" = $3
WHERE "id" = $1
RETURNING *;


-- name: DeleteQuestion :exec
DELETE
FROM "questions"
WHERE "id" = $1;

-- name: ListQuestionAnswers :many
SELECT *
FROM "answers"
WHERE "question_id" = $1
ORDER BY "created_at"
LIMIT $2 OFFSET $3;
