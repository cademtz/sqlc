// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedUpdate = `-- name: SchemaScopedUpdate :exec
UPDATE foo.bar SET name = $2 WHERE id = $1
`

type SchemaScopedUpdateParams struct {
	ID   int32
	Name string
}

func (q *Queries) SchemaScopedUpdate(ctx context.Context, arg SchemaScopedUpdateParams) error {
	_, err := q.db.ExecContext(ctx, schemaScopedUpdate, arg.ID, arg.Name)
	return err
}
