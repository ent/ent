// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"

	"entgo.io/ent/cmd/internal/base"
	"entgo.io/ent/entc/gen"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "entc"}
	cmd.AddCommand(
		base.InitCmd(),
		base.DescribeCmd(),
		base.GenerateCmd(migrate),
	)
	_ = cmd.Execute()
}

func migrate(c *gen.Config) {
	var (
		target = filepath.Join(c.Target, "generate.go")
		oldCmd = []byte("entgo.io/ent/cmd/entc")
	)
	buf, err := ioutil.ReadFile(target)
	if err != nil || !bytes.Contains(buf, oldCmd) {
		return
	}
	_ = ioutil.WriteFile(target, bytes.ReplaceAll(buf, oldCmd, []byte("entgo.io/ent/cmd/ent")), 0644)
}
