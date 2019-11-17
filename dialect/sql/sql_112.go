// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// +build !go1.13

package sql

import (
	"database/sql/driver"
	"time"
)

// NullTime represents a time.Time that may be null.
//
// NullTime implements the Scanner interface so it can
// be used as a scan destination, similar to NullString.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (n *NullTime) Scan(v interface{}) error {
	if v, ok := v.(time.Time); ok {
		n.Time = v
		n.Valid = true
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
