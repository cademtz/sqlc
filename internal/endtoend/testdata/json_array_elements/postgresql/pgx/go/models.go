// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package querytest

type SysAction struct {
	ID        int64
	Code      string
	Menu      string
	Title     string
	Resources []byte
}
