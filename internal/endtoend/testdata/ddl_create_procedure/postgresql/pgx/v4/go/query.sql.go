// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const callInsertData = `-- name: CallInsertData :exec
CALL insert_data($1, $2)
`

type CallInsertDataParams struct {
	A int32
	B int32
}

func (q *Queries) CallInsertData(ctx context.Context, arg CallInsertDataParams) error {
	_, err := q.db.Exec(ctx, callInsertData, arg.A, arg.B)
	return err
}

const callInsertDataNamed = `-- name: CallInsertDataNamed :exec
CALL insert_data(b => $1, a => $2)
`

type CallInsertDataNamedParams struct {
	B int32
	A int32
}

func (q *Queries) CallInsertDataNamed(ctx context.Context, arg CallInsertDataNamedParams) error {
	_, err := q.db.Exec(ctx, callInsertDataNamed, arg.B, arg.A)
	return err
}

const callInsertDataNoArgs = `-- name: CallInsertDataNoArgs :exec
CALL insert_data(1, 2)
`

func (q *Queries) CallInsertDataNoArgs(ctx context.Context) error {
	_, err := q.db.Exec(ctx, callInsertDataNoArgs)
	return err
}

const callInsertDataSqlcArgs = `-- name: CallInsertDataSqlcArgs :exec
CALL insert_data($1, $2)
`

type CallInsertDataSqlcArgsParams struct {
	Foo int32
	Bar int32
}

func (q *Queries) CallInsertDataSqlcArgs(ctx context.Context, arg CallInsertDataSqlcArgsParams) error {
	_, err := q.db.Exec(ctx, callInsertDataSqlcArgs, arg.Foo, arg.Bar)
	return err
}
