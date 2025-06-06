// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedColList = `-- name: SchemaScopedColList :many
SELECT foo.bar.id FROM foo.bar
`

func (q *Queries) SchemaScopedColList(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, schemaScopedColList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const schemaScopedList = `-- name: SchemaScopedList :many
SELECT id FROM foo.bar
`

func (q *Queries) SchemaScopedList(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, schemaScopedList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
