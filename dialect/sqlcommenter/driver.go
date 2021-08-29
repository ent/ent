package sqlcommenter

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
)

type (
	commenter struct {
		options
	}
	// CommentDriver is a driver that adds sql comments (see https://google.github.io/sqlcommenter/).
	CommentDriver struct {
		dialect.Driver // underlying driver.
		commenter
	}
	// CommentTx is a transaction implementation that adds sql comments.
	CommentTx struct {
		dialect.Tx                 // underlying transaction.
		ctx        context.Context // underlying transaction context.
		commenter
	}
)

func NewCommentDriver(drv dialect.Driver, options ...Option) dialect.Driver {
	defaultCommenters := []Commenter{NewDriverVersionCommenter()}
	opts := buildOptions(append(options, WithCommenter(defaultCommenters...)))
	return &CommentDriver{drv, commenter{opts}}
}

func (c commenter) getComments(ctx context.Context) SqlComments {
	cmts := make(SqlComments)
	cmts.Append(c.globalComments)
	for _, h := range c.commenters {
		cmts.Append(h.GetComments(ctx))
	}
	return cmts
}

func (c commenter) withComments(ctx context.Context, query string) string {
	return fmt.Sprintf("%s /*%s*/", query, c.getComments(ctx).Marshal())
}

// Exec adds sql comments to the original query and calls the underlying driver Exec method.
func (d *CommentDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return d.Driver.Exec(ctx, d.withComments(ctx, query), args, v)
}

// Query adds sql comments to the original query and calls the underlying driver Query method.
func (d *CommentDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return d.Driver.Query(ctx, d.withComments(ctx, query), args, v)
}

// Tx wraps the underlying Tx command with a commenter.
func (d *CommentDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &CommentTx{tx, ctx, d.commenter}, nil
}

// BeginTx wraps the underlying transaction with commenter and calls the underlying driver BeginTx command if it's supported.
func (d *CommentDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	drv, ok := d.Driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.BeginTx is not supported")
	}
	tx, err := drv.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &CommentTx{tx, ctx, d.commenter}, nil
}

// Exec adds sql comments and calls the underlying transaction Exec method.
func (d *CommentTx) Exec(ctx context.Context, query string, args, v interface{}) error {
	return d.Tx.Exec(ctx, d.withComments(ctx, query), args, v)
}

// Query adds sql comments and calls the underlying transaction Query method.
func (d *CommentTx) Query(ctx context.Context, query string, args, v interface{}) error {
	return d.Tx.Query(ctx, d.withComments(ctx, query), args, v)
}

// Commit calls the underlying Tx to commit.
func (d *CommentTx) Commit() error {
	return d.Tx.Commit()
}

// Rollback calls the underlying Tx to rollback.
func (d *CommentTx) Rollback() error {
	return d.Tx.Rollback()
}
