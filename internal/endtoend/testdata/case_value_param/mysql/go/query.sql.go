// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const update = `-- name: Update :exec
UPDATE testing
SET value = CASE ? WHEN true THEN 'Hello' WHEN false THEN 'Goodbye' ELSE value END
`

func (q *Queries) Update(ctx context.Context, value sql.NullString) error {
	_, err := q.db.ExecContext(ctx, update, value)
	return err
}
