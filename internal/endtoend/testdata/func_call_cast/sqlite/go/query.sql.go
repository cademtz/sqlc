// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const demo = `-- name: Demo :one
SELECT CAST(CHAR(1,2,3,4,5) AS BLOB) AS col1
`

func (q *Queries) Demo(ctx context.Context) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, demo)
	var col1 []byte
	err := row.Scan(&col1)
	return col1, err
}
