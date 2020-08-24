// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package template

// Extension is a template extension.
type Extension struct {
	Type string
}

func (*Extension) Name() string {
	return "Extension"
}
