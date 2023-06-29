-- name: CreateAnsweredQuestion :one
INSERT INTO "answered_questions" ("user_id",
                                  "question_id")
VALUES ($1, $2)
RETURNING *;

-- name: GetAnsweredQuestion :one
SELECT *
FROM "answered_questions"
WHERE "id" = $1
LIMIT 1;

-- name: ListAnsweredQuestions :many
SELECT *
FROM "answered_questions"
ORDER BY "id"
LIMIT $1 OFFSET $2;

-- name: DeleteAnsweredQuestion :exec
DELETE
FROM "answered_questions"
WHERE "id" = $1;