// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package role

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
