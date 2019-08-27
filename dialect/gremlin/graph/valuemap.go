// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// ValueMap models a .valueMap() gremlin response.
type ValueMap []map[string]interface{}

// Decode decodes a value map into v.
func (m ValueMap) Decode(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return errors.New("cannot unmarshal into a non pointer")
	}
	if rv.IsNil() {
		return errors.New("cannot unmarshal into a nil pointer")
	}

	if rv.Elem().Kind() != reflect.Slice {
		v = &[]interface{}{v}
	}
	return m.decode(v)
}

func (m ValueMap) decode(v interface{}) error {
	cfg := mapstructure.DecoderConfig{
		DecodeHook: func(f, t reflect.Kind, data interface{}) (interface{}, error) {
			if f == reflect.Slice && t != reflect.Slice {
				rv := reflect.ValueOf(data)
				if rv.Len() == 1 {
					data = rv.Index(0).Interface()
				}
			}
			return data, nil
		},
		Result:  v,
		TagName: "json",
	}

	dec, err := mapstructure.NewDecoder(&cfg)
	if err != nil {
		return errors.Wrap(err, "creating structure decoder")
	}
	if err := dec.Decode(m); err != nil {
		return errors.Wrap(err, "decoding value map")
	}
	return nil
}
