// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

// this is the mood type
type FooMood string

const (
	FooMoodSad   FooMood = "sad"
	FooMoodOk    FooMood = "ok"
	FooMoodHappy FooMood = "happy"
)

func (e *FooMood) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooMood(s)
	case string:
		*e = FooMood(s)
	default:
		return fmt.Errorf("unsupported scan type for FooMood: %T", src)
	}
	return nil
}

type NullFooMood struct {
	FooMood FooMood
	Valid   bool // Valid is true if FooMood is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFooMood) Scan(value interface{}) error {
	if value == nil {
		ns.FooMood, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FooMood.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFooMood) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FooMood), nil
}

// this is the bar table
type FooBar struct {
	// this is the baz column
	Baz string
}

// this is the bat view
type FooBat struct {
	Baz string
}
