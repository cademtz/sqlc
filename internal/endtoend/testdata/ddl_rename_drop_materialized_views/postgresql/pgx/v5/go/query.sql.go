// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const getAuthorMv = `-- name: GetAuthorMv :one
SELECT id, name, bio FROM authors_mv
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthorMv(ctx context.Context, id int64) (AuthorsMv, error) {
	row := q.db.QueryRow(ctx, getAuthorMv, id)
	var i AuthorsMv
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}
