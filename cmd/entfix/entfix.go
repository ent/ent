package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/gen"
	"github.com/alecthomas/kong"
)

type (
	// App configures the entfix CLI.
	App struct {
		GlobalID GlobalID `cmd:"" name:"globalid" help:"Migrate unique global id ent_types to ent global feature"`
	}
	// GlobalID represents the 'entfix globalid' command.
	GlobalID struct {
		Dialect string `name:"dialect" help:"Database dialect" required:"" enum:"mysql,postgres,sqlite3"`
		DSN     string `name:"dsn" help:"Data source name" required:""`
		Path    string `name:"path" help:"Path to the generated ent code" required:""`
	}
)

func main() {
	// Ensure to stop execution on Interrupt signal.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	app := kong.Parse(
		new(App),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.UsageOnError(),
	)
	app.FatalIfErrorf(func() error {
		err := app.Run()
		if err := context.Cause(ctx); err != nil {
			return err
		}
		if errors.Is(err, context.Canceled) {
			return nil
		}
		return err
	}())
}

func (cmd *GlobalID) Run(ctx context.Context) error {
	fmt.Print(`IMPORTANT INFORMATION

  'entfix globalid' will convert the allocated id ranges for your nodes from the 
  database stored 'ent_types' table to the new static configuration on the ent 
  schema itself.

  Please note, that the 'ent_types' table might differ between different environments 
  where your app is deployed. This is especially true if you are using 
  auto-migration instead of versioned migrations.

  Please check, that all 'ent_types' tables for all deployments are equal!

  Only 'yes' will be accepted to approve.

  Enter a value: `)
	switch c, err := bufio.NewReader(os.Stdin).ReadString('\n'); {
	case err != nil:
		return err
	case strings.TrimSpace(c) != "yes":
		fmt.Println("\nAborted.")
		return nil
	}
	db, err := sql.Open(cmd.Dialect, cmd.DSN)
	if err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := sql.Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := db.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	var ts []string
	if err := sql.ScanSlice(rows, &ts); err != nil {
		return err
	}
	is := make(gen.IncrementStarts, len(ts))
	for i, t := range ts {
		is[t] = int64(i << 32)
	}
	if err := is.WriteToDisk(cmd.Path); err != nil {
		return err
	}
	fmt.Println("\nSuccess! Please run code generation to complete the process.")
	return nil
}
