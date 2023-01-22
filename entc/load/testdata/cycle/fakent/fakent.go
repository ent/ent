// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package fakent fis a fake generated Ent package.
package fakent

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/load/testdata/cycle"
)

type Hook = ent.Hook

var _ = &cycle.Used{}
