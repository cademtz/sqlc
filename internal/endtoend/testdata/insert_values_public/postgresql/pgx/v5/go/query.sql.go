// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertValues = `-- name: InsertValues :exec
INSERT INTO public.foo (a, b) VALUES ($1, $2)
`

type InsertValuesParams struct {
	A pgtype.Text
	B pgtype.Int4
}

func (q *Queries) InsertValues(ctx context.Context, arg InsertValuesParams) error {
	_, err := q.db.Exec(ctx, insertValues, arg.A, arg.B)
	return err
}
