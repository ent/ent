// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	jsoniter "github.com/json-iterator/go"
)

var config = jsoniter.Config{}.Froze()

func init() {
	config.RegisterExtension(&encodeExtension{})
	config.RegisterExtension(&decodeExtension{})
}
