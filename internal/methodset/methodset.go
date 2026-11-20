// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package methodset lets entc/load install the interface check that
// (*field.RType).Implements uses for RTypes decoded from JSON.
//
// The check iterates interface methods with reflect.Type.Method. When that is
// reachable, the linker keeps every exported method of every reachable type.
// (*field.RType).Implements is reachable in runtime binaries, because its
// signature matches reflect.Type.Implements. The check itself lives in
// entc/load, which runtime binaries don't import.
package methodset

import "reflect"

// Implements reports whether a *field.RType decoded from JSON
// implements typ, using its Methods. entc/load sets it.
var Implements func(rtype any, typ reflect.Type) bool
