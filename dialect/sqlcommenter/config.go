package sqlcommenter

import (
	"context"
)

type Commenter interface {
	GetComments(context.Context) SqlComments
}

type (
	CommentsHandler func(context.Context) SqlComments
	Option          func(*options)
	options         struct {
		commenters     []Commenter
		globalComments SqlComments
	}
)

// WithCommenter overrides the default comments generator handler.
// default comments added via WithComments will still be applied.
func WithCommenter(commenters ...Commenter) Option {
	return func(opts *options) {
		opts.commenters = append(opts.commenters, commenters...)
	}
}

// WithComments appends the given comments to every sql query.
func WithComments(comments SqlComments) Option {
	return func(opts *options) {
		opts.globalComments = comments
	}
}

func buildOptions(opts []Option) options {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
