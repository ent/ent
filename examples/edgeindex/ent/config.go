// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"github.com/facebookincubator/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver is the driver used for execute database requests.
	driver dialect.Driver
	// verbose enable a verbosity logging.
	verbose bool
	// log used for logging on verbose mode.
	log func(...interface{})
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.verbose {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Verbose sets the client logging to verbose.
func Verbose() Option {
	return func(c *config) {
		c.verbose = true
	}
}

// Log sets the client logging to verbose.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
