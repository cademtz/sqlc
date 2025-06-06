// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const cTECount = `-- name: CTECount :many
WITH all_count AS (
	SELECT count(*) FROM bar
), ready_count AS (
	SELECT count(*) FROM bar WHERE ready = true
)
SELECT all_count.count, ready_count.count
FROM all_count, ready_count
`

type CTECountRow struct {
	Count   int64
	Count_2 int64
}

func (q *Queries) CTECount(ctx context.Context) ([]CTECountRow, error) {
	rows, err := q.db.Query(ctx, cTECount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CTECountRow
	for rows.Next() {
		var i CTECountRow
		if err := rows.Scan(&i.Count, &i.Count_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
