// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: question.sql

package db

import (
	"context"
)

const createQuestion = `-- name: CreateQuestion :one
INSERT INTO "questions" ("text",
                         "hint",
                         "category")
VALUES ($1, $2, $3)
RETURNING id, text, hint, category, created_at
`

type CreateQuestionParams struct {
	Text     string `json:"text"`
	Hint     string `json:"hint"`
	Category string `json:"category"`
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion, arg.Text, arg.Hint, arg.Category)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Hint,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}

const deleteQuestion = `-- name: DeleteQuestion :exec
DELETE
FROM "questions"
WHERE "id" = $1
`

func (q *Queries) DeleteQuestion(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteQuestion, id)
	return err
}

const getQuestion = `-- name: GetQuestion :one
SELECT id, text, hint, category, created_at
FROM "questions"
WHERE "id" = $1
LIMIT 1
`

func (q *Queries) GetQuestion(ctx context.Context, id int64) (Question, error) {
	row := q.db.QueryRowContext(ctx, getQuestion, id)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Hint,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}

const listQuestions = `-- name: ListQuestions :many
SELECT id, text, hint, category, created_at
FROM "questions"
ORDER BY "id"
LIMIT $1 OFFSET $2
`

type ListQuestionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListQuestions(ctx context.Context, arg ListQuestionsParams) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, listQuestions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Question{}
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.Text,
			&i.Hint,
			&i.Category,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuestion = `-- name: UpdateQuestion :one
UPDATE "questions"
SET "text" = $2,
    "hint" = $3
WHERE "id" = $1
RETURNING id, text, hint, category, created_at
`

type UpdateQuestionParams struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
	Hint string `json:"hint"`
}

func (q *Queries) UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, updateQuestion, arg.ID, arg.Text, arg.Hint)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Hint,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}
