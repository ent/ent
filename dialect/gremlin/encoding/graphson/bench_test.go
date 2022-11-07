// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type book struct {
	ID       string   `json:"id" graphson:"g:UUID"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Pages    int      `json:"num_pages"`
	Chapters []string `json:"chapters"`
}

func generateObject() *book {
	return &book{
		ID:       "21d5dcbf-1fd4-493e-9b74-d6c429f9e4a5",
		Title:    "The Art of Computer Programming, Vol. 2",
		Author:   "Donald E. Knuth",
		Pages:    784,
		Chapters: []string{"Random numbers", "Arithmetic"},
	}
}

func BenchmarkMarshalObject(b *testing.B) {
	obj := generateObject()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalObject(b *testing.B) {
	b.ReportAllocs()
	out, err := Marshal(generateObject())
	if err != nil {
		b.Fatal(err)
	}

	obj := &book{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = Unmarshal(out, obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalInterface(b *testing.B) {
	b.ReportAllocs()
	data, err := jsoniter.Marshal(generateObject())
	if err != nil {
		b.Fatal(err)
	}

	var obj any
	if err = jsoniter.Unmarshal(data, &obj); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err = Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalInterface(b *testing.B) {
	b.ReportAllocs()
	data, err := Marshal(generateObject())
	if err != nil {
		b.Fatal(err)
	}

	var obj any

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = Unmarshal(data, &obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}
