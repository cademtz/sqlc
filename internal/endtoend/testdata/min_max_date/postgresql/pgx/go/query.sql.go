// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const activityStats = `-- name: ActivityStats :one
SELECT COUNT(*) as NumOfActivities,
        MIN(event_time) as MinDate, 
        MAX(event_time) as MaxDate 
FROM activities 
WHERE account_id = $1
`

type ActivityStatsRow struct {
	Numofactivities int64
	Mindate         pgtype.Timestamptz
	Maxdate         pgtype.Timestamptz
}

func (q *Queries) ActivityStats(ctx context.Context, accountID int64) (ActivityStatsRow, error) {
	row := q.db.QueryRow(ctx, activityStats, accountID)
	var i ActivityStatsRow
	err := row.Scan(&i.Numofactivities, &i.Mindate, &i.Maxdate)
	return i, err
}
