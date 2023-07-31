// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: answered_question.sql

package db

import (
	"context"
	"database/sql"
)

const createAnsweredQuestion = `-- name: CreateAnsweredQuestion :one
INSERT INTO "answered_questions" ("user_id",
                                  "question_id")
VALUES ($1, $2)
RETURNING id, user_id, question_id, answered_at
`

type CreateAnsweredQuestionParams struct {
	UserID     int64 `json:"user_id"`
	QuestionID int64 `json:"question_id"`
}

func (q *Queries) CreateAnsweredQuestion(ctx context.Context, arg CreateAnsweredQuestionParams) (AnsweredQuestion, error) {
	row := q.db.QueryRowContext(ctx, createAnsweredQuestion, arg.UserID, arg.QuestionID)
	var i AnsweredQuestion
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.AnsweredAt,
	)
	return i, err
}

const deleteAnsweredQuestion = `-- name: DeleteAnsweredQuestion :exec
DELETE
FROM "answered_questions"
WHERE "id" = $1
`

func (q *Queries) DeleteAnsweredQuestion(ctx context.Context, id int64) error {
	result, err := q.db.ExecContext(ctx, deleteAnsweredQuestion, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

const getAnsweredQuestion = `-- name: GetAnsweredQuestion :one
SELECT id, user_id, question_id, answered_at
FROM "answered_questions"
WHERE "id" = $1
LIMIT 1
`

func (q *Queries) GetAnsweredQuestion(ctx context.Context, id int64) (AnsweredQuestion, error) {
	row := q.db.QueryRowContext(ctx, getAnsweredQuestion, id)
	var i AnsweredQuestion
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.AnsweredAt,
	)
	return i, err
}

const listAnsweredQuestions = `-- name: ListAnsweredQuestions :many
SELECT id, user_id, question_id, answered_at
FROM "answered_questions"
ORDER BY "id"
LIMIT $1 OFFSET $2
`

type ListAnsweredQuestionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAnsweredQuestions(ctx context.Context, arg ListAnsweredQuestionsParams) ([]AnsweredQuestion, error) {
	rows, err := q.db.QueryContext(ctx, listAnsweredQuestions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnsweredQuestion{}
	for rows.Next() {
		var i AnsweredQuestion
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.QuestionID,
			&i.AnsweredAt,
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
