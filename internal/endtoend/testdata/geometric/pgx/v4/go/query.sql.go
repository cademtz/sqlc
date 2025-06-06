// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const selectTest = `-- name: SelectTest :many
SELECT v_box_null, v_circle_null, v_line_null, v_lseg_null, v_path_null, v_point_null, v_polygon_null, v_box, v_circle, v_line, v_lseg, v_path, v_point, v_polygon
from test_table
`

func (q *Queries) SelectTest(ctx context.Context) ([]TestTable, error) {
	rows, err := q.db.Query(ctx, selectTest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestTable
	for rows.Next() {
		var i TestTable
		if err := rows.Scan(
			&i.VBoxNull,
			&i.VCircleNull,
			&i.VLineNull,
			&i.VLsegNull,
			&i.VPathNull,
			&i.VPointNull,
			&i.VPolygonNull,
			&i.VBox,
			&i.VCircle,
			&i.VLine,
			&i.VLseg,
			&i.VPath,
			&i.VPoint,
			&i.VPolygon,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
