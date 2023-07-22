package db

import (
	"context"
	"database/sql"
	"fmt"
)

const tableExists = `-- name: CleanTable :one
SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1);
`

func (q *Queries) TableExists(ctx context.Context, tableName string) (bool, error) {
	row := q.db.QueryRowContext(ctx, tableExists, sql.Named("table_name", tableName))
	var isExists bool
	err := row.Scan(&isExists)
	return isExists, err
}

const cleanTable = `-- name: CleanTable :exec
DELETE FROM %s;
ALTER SEQUENCE %s_id_seq RESTART WITH 1;
`

func (q *Queries) CleanTable(ctx context.Context, tableName string) error {
	_, err := q.db.ExecContext(ctx, fmt.Sprintf(cleanTable, tableName, tableName))
	return err
}
