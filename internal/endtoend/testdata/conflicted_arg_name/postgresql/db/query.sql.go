// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const time2ByTime = `-- name: Time2ByTime :one
SELECT time2 FROM foo WHERE time=$1
`

func (q *Queries) Time2ByTime(ctx context.Context, argTime time.Time) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, time2ByTime, argTime)
	var time2 time.Time
	err := row.Scan(&time2)
	return time2, err
}

const uuid2ByUuid = `-- name: Uuid2ByUuid :one
SELECT uuid2 FROM foo WHERE uuid=$1
`

func (q *Queries) Uuid2ByUuid(ctx context.Context, argUuid uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, uuid2ByUuid, argUuid)
	var uuid2 uuid.UUID
	err := row.Scan(&uuid2)
	return uuid2, err
}
