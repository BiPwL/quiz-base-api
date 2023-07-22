package db

import (
	"context"
	"database/sql"
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
