// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package querytest

import (
	"database/sql"
)

type Description struct {
	ID  string
	Txt sql.NullString
}
