// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package datatype

import (
	"net/netip"
)

type Foo struct {
	Bar *netip.Addr
	Baz *netip.Prefix
}
