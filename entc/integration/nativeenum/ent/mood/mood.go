// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package mood

type Mood string

const (
	Happy Mood = "happy"
	Ok         = "ok"
	Sad        = "sad"
)

func (Mood) Values() (roles []string) {
	for _, r := range []Mood{Happy, Ok, Sad} {
		roles = append(roles, string(r))
	}
	return
}
