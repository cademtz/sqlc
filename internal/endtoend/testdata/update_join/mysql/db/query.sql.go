// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package db

import (
	"context"
)

const updateJoin = `-- name: UpdateJoin :exec
UPDATE  join_table as jt
        JOIN primary_table as pt
            ON jt.primary_table_id = pt.id
SET     jt.is_active = ?
WHERE   jt.id = ?
        AND pt.user_id = ?
`

type UpdateJoinParams struct {
	IsActive bool
	ID       uint64
	UserID   uint64
}

func (q *Queries) UpdateJoin(ctx context.Context, arg UpdateJoinParams) error {
	_, err := q.db.ExecContext(ctx, updateJoin, arg.IsActive, arg.ID, arg.UserID)
	return err
}

const updateLeftJoin = `-- name: UpdateLeftJoin :exec
UPDATE  join_table as jt
        LEFT JOIN primary_table as pt
            ON jt.primary_table_id = pt.id
SET     jt.is_active = ?
WHERE   jt.id = ?
        AND pt.user_id = ?
`

type UpdateLeftJoinParams struct {
	IsActive bool
	ID       uint64
	UserID   uint64
}

func (q *Queries) UpdateLeftJoin(ctx context.Context, arg UpdateLeftJoinParams) error {
	_, err := q.db.ExecContext(ctx, updateLeftJoin, arg.IsActive, arg.ID, arg.UserID)
	return err
}

const updateRightJoin = `-- name: UpdateRightJoin :exec
UPDATE  join_table as jt
        RIGHT JOIN primary_table as pt
            ON jt.primary_table_id = pt.id
SET     jt.is_active = ?
WHERE   jt.id = ?
        AND pt.user_id = ?
`

type UpdateRightJoinParams struct {
	IsActive bool
	ID       uint64
	UserID   uint64
}

func (q *Queries) UpdateRightJoin(ctx context.Context, arg UpdateRightJoinParams) error {
	_, err := q.db.ExecContext(ctx, updateRightJoin, arg.IsActive, arg.ID, arg.UserID)
	return err
}
