// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package db

import (
	"context"
)

const getRestrictedId = `-- name: GetRestrictedId :one
SELECT
  NULLIF(id, $1) restricted_id
FROM
  author
`

func (q *Queries) GetRestrictedId(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, getRestrictedId, id)
	var restricted_id bool
	err := row.Scan(&restricted_id)
	return restricted_id, err
}
