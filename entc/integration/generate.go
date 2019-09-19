// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

//go:generate go run ../cmd/entc/entc.go generate --storage=sql,gremlin --idtype string --header "Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.\n// This source code is licensed under the Apache 2.0 license found\n// in the LICENSE file in the root directory of this source tree.\n\n// Code generated (@generated) by entc, DO NOT EDIT." ./ent/schema
//go:generate go run ../cmd/entc/entc.go generate --header "Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.\n// This source code is licensed under the Apache 2.0 license found\n// in the LICENSE file in the root directory of this source tree.\n\n// Code generated (@generated) by entc, DO NOT EDIT." ./migrate/entv1/schema
//go:generate go run ../cmd/entc/entc.go generate --header "Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.\n// This source code is licensed under the Apache 2.0 license found\n// in the LICENSE file in the root directory of this source tree.\n\n// Code generated (@generated) by entc, DO NOT EDIT." ./migrate/entv2/schema
//go:generate go run ../cmd/entc/entc.go generate --header "Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.\n// This source code is licensed under the Apache 2.0 license found\n// in the LICENSE file in the root directory of this source tree.\n\n// Code generated (@generated) by entc, DO NOT EDIT." ./template/ent/schema --template=template/ent/template
//go:generate go run ../cmd/entc/entc.go generate --header "Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.\n// This source code is licensed under the Apache 2.0 license found\n// in the LICENSE file in the root directory of this source tree.\n\n// Code generated (@generated) by entc, DO NOT EDIT." ./json/ent/schema
