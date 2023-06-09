// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO "categories" ("name",
                          "key")
VALUES ($1, $2)
RETURNING key, name
`

type CreateCategoryParams struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.Name, arg.Key)
	var i Category
	err := row.Scan(&i.Key, &i.Name)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE
FROM "categories"
WHERE "key" = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, key string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, key)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT key, name
FROM "categories"
WHERE "key" = $1
LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, key string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, key)
	var i Category
	err := row.Scan(&i.Key, &i.Name)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT key, name
FROM "categories"
ORDER BY "key"
LIMIT $1 OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.Key, &i.Name); err != nil {
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

const updateCategory = `-- name: UpdateCategory :one
UPDATE "categories"
SET "name" = $2
WHERE "key" = $1
RETURNING key, name
`

type UpdateCategoryParams struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.Key, arg.Name)
	var i Category
	err := row.Scan(&i.Key, &i.Name)
	return i, err
}
