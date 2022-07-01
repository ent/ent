// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package schemadir is used as an example to test cases where
// package-name is not equal to the package-path (schemadir <> dir).
package schemadir

import (
	"net/http"
)

type Dir http.Dir
