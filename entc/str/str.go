// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package str provides function related to string manipulation
package str

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Title is wrapper around cases.Title with language.Und
// replacement for deprecated func strings.Title
func Title(t string) string {
	return cases.Title(language.Und).String(t)
}
