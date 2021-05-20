// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package role

import (
	"database/sql/driver"
	"errors"
)

type Role string

const (
	Admin Role = "ADMIN"
	Owner Role = "OWNER"
	User  Role = "USER"
	Read  Role = "READ"
	Write Role = "WRITE"
)

func (Role) Values() (roles []string) {
	for _, r := range []Role{Admin, Owner, User, Read, Write} {
		roles = append(roles, string(r))
	}
	return
}

type Priority int

const (
	Unknown Priority = iota
	Low
	High
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "LOW"
	case High:
		return "HIGH"
	default:
		return "UNKNOWN"
	}
}

func (Priority) Values() []string {
	return []string{Unknown.String(), Low.String(), High.String()}
}

func (p Priority) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p *Priority) Scan(val interface{}) error {
	if val == nil {
		return nil
	}

	var s, ok = val.(string)
	if !ok {
		return errors.New("val must be of type string")
	}

	switch s {
	case "LOW":
		*p = Low
	case "HIGH":
		*p = High
	default:
		*p = Unknown
	}

	return nil
}
