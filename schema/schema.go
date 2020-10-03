// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

// Annotation is used to attach arbitrary metadata to the schema objects in codegen.
// The object must be serializable to JSON raw value (e.g. struct, map or slice).
//
// Template extensions can retrieve this metadata and use it inside their templates.
// Read more about it in ent website: https://entgo.io/docs/templates/#annotations.
type Annotation interface {
	// Name defines the name of the annotation to be retrieved by the codegen.
	Name() string
}
