// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"log"

	"entgo.io/ent/cmd/internal/base"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "ent"}
	cmd.AddCommand(
		base.InitCmd(),
		base.DescribeCmd(),
		base.GenerateCmd(),
		base.DiffCmd(),
	)
	_ = cmd.Execute()
}
