// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sid

import (
	"crypto/rand"
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) Value() (driver.Value, error) {
	r, err := strconv.ParseInt(string(i), 10, 64)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (i *ID) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		return nil
	case int64:
		*i = ID(fmt.Sprint(v))
		return nil
	}
	return errors.New("not a valid ID")
}

func New() ID {
	return NewLength(10)
}

func NewLength(l int) ID {
	var out string
	for len(out) < l {
		result, _ := rand.Int(rand.Reader, big.NewInt(100))
		out += fmt.Sprint(result.Uint64() + 1)
	}
	return ID(out[:l])
}
