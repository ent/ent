// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "entgo.io/ent"

// Api represents an example schema with wrong usage of acronym (e.g., API).
// nolint:stylecheck
type Api struct {
	ent.Schema
}
