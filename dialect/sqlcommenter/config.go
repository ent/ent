package sqlcommenter

import (
	"context"
)

type (
	CommentsHandler func(context.Context) SqlComments
	Option          func(*options)
	options         struct {
		commenters     []CommentsHandler
		globalComments SqlComments
	}
)

// WithCommenter overrides the default comments generator handler.
// default comments added via WithComments will still be applied.
func WithCommenter(commentsHandlers ...CommentsHandler) Option {
	return func(opts *options) {
		opts.commenters = commentsHandlers
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
