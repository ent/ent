// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// +build go1.13

package sql

import "database/sql"

// NullTime represents a time.Time that may be null.
type NullTime = sql.NullTime
