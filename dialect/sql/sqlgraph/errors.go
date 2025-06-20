// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"errors"
	"strings"
)

// IsConstraintError returns true if the error resulted from a database constraint violation.
func IsConstraintError(err error) bool {
	var e *ConstraintError
	return errors.As(err, &e) ||
		IsUniqueConstraintError(err) ||
		IsForeignKeyConstraintError(err) ||
		IsCheckConstraintError(err)
}

// IsUniqueConstraintError reports if the error resulted from a DB uniqueness constraint violation.
// e.g. duplicate value in unique index.
func IsUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}
	for _, s := range []string{
		"Error 1062",                 // MySQL
		"violates unique constraint", // Postgres
		"UNIQUE constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}

// IsForeignKeyConstraintError reports if the error resulted from a database foreign-key constraint violation.
// e.g. parent row does not exist.
func IsForeignKeyConstraintError(err error) bool {
	if err == nil {
		return false
	}
	for _, s := range []string{
		"Error 1451",                      // MySQL (Cannot delete or update a parent row).
		"Error 1452",                      // MySQL (Cannot add or update a child row).
		"violates foreign key constraint", // Postgres
		"FOREIGN KEY constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}

// IsCheckConstraintError reports if the error resulted from a database check constraint violation.
// e.g. a value does not satisfy a check condition.
func IsCheckConstraintError(err error) bool {
	if err == nil {
		return false
	}
	for _, s := range []string{
		"Error 3819",                // MySQL
		"violates check constraint", // Postgres
		"CHECK constraint failed",   // SQLite
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}
