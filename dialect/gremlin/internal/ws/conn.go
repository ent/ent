// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ws

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/encoding"
	"entgo.io/ent/dialect/gremlin/encoding/graphson"

	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 5 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type (
	// A Dialer contains options for connecting to Gremlin server.
	Dialer struct {
		// Underlying websocket dialer.
		websocket.Dialer

		// Gremlin server basic auth credentials.
		user, pass string
	}

	// Conn performs operations on a gremlin server.
	Conn struct {
		// Underlying websocket connection.
		conn *websocket.Conn

		// Credentials for basic authentication.
		user, pass string

		// Goroutine tracking.
		ctx context.Context
		grp *errgroup.Group

		// Channel of outbound requests.
		send chan io.Reader

		// Map of in flight requests.
		inflight sync.Map
	}

	// inflight tracks request state.
	inflight struct {
		// partially received data
		frags []graphson.RawMessage

		// response channel
		result chan<- result
	}

	// represents an execution result.
	result struct {
		rsp *gremlin.Response
		err error
	}
)

var (
	// DefaultDialer is a dialer with all fields set to the default values.
	DefaultDialer = &Dialer{
		Dialer: websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 5 * time.Second,
			WriteBufferSize:  8192,
			ReadBufferSize:   8192,
		},
	}

	// ErrConnClosed is returned by the Conn's Execute method when
	// the underlying gremlin server connection is closed.
	ErrConnClosed = errors.New("gremlin: server connection closed")

	// ErrDuplicateRequest is returned by the Conns Execute method on
	// request identifier key collision.
	ErrDuplicateRequest = errors.New("gremlin: duplicate request")
)

// Dial creates a new connection by calling DialContext with a background context.
func (d *Dialer) Dial(uri string) (*Conn, error) {
	return d.DialContext(context.Background(), uri)
}

// DialContext creates a new Gremlin connection.
func (d *Dialer) DialContext(ctx context.Context, uri string) (*Conn, error) {
	c, rsp, err := d.Dialer.DialContext(ctx, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("gremlin: dialing uri %s: %w", uri, err)
	}
	defer rsp.Body.Close()

	conn := &Conn{
		conn: c,
		user: d.user,
		pass: d.pass,
		send: make(chan io.Reader),
	}
	conn.grp, conn.ctx = errgroup.WithContext(context.Background())

	conn.grp.Go(conn.sender)
	conn.grp.Go(conn.receiver)

	return conn, nil
}

// Execute executes a request against a Gremlin server.
func (c *Conn) Execute(ctx context.Context, req *gremlin.Request) (*gremlin.Response, error) {
	// buffered result channel prevents receiver block on context cancellation
	result := make(chan result, 1)

	// request id must be unique across inflight request
	if _, loaded := c.inflight.LoadOrStore(req.RequestID, &inflight{result: result}); loaded {
		return nil, ErrDuplicateRequest
	}

	pr, pw := io.Pipe()
	defer pr.Close()

	// stream graphson encoding into request
	c.grp.Go(func() error {
		err := graphson.NewEncoder(pw).Encode(req)
		if err != nil {
			err = fmt.Errorf("encoding request: %w", err)
		}
		pw.CloseWithError(err)
		return err
	})

	// local copy for single write
	send := c.send

	for {
		select {
		case <-c.ctx.Done():
			c.inflight.Delete(req.RequestID)
			return nil, ErrConnClosed
		case <-ctx.Done():
			c.inflight.Delete(req.RequestID)
			return nil, ctx.Err()
		case send <- pr:
			send = nil
		case result := <-result:
			return result.rsp, result.err
		}
	}
}

// Close connection with a Gremlin server.
func (c *Conn) Close() error {
	c.grp.Go(func() error { return ErrConnClosed })
	_ = c.grp.Wait()
	return nil
}

func (c *Conn) sender() error {
	pinger := time.NewTicker(pingPeriod)
	defer pinger.Stop()

	// closing connection terminates receiver
	defer c.conn.Close()

	for {
		select {
		case r := <-c.send:
			// ensure write completes within a window
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			// fetch next message writer
			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return fmt.Errorf("getting message writer: %w", err)
			}

			// write mime header
			if _, err := w.Write(encoding.GraphSON3Mime); err != nil {
				return fmt.Errorf("writing mime header: %w", err)
			}

			// write request body
			if _, err := io.Copy(w, r); err != nil {
				return fmt.Errorf("writing request: %w", err)
			}

			// finish message write
			if err := w.Close(); err != nil {
				return fmt.Errorf("closing message writer: %w", err)
			}
		case <-c.ctx.Done():
			// connection closing
			return c.conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
				time.Time{},
			)
		case <-pinger.C:
			// periodic connection keepalive
			if err := c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait)); err != nil {
				return fmt.Errorf("writing ping message: %w", err)
			}
		}
	}
}

func (c *Conn) receiver() error {
	// handle keepalive responses
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	// complete all in flight requests on termination
	defer c.inflight.Range(func(id, ifr any) bool {
		ifr.(*inflight).result <- result{err: ErrConnClosed}
		c.inflight.Delete(id)
		return true
	})

	for {
		// rely on sender connection close during termination
		_, r, err := c.conn.NextReader()
		if err != nil {
			return fmt.Errorf("writing ping message: %w", err)
		}

		// decode received response
		var rsp gremlin.Response
		if err := graphson.NewDecoder(r).Decode(&rsp); err != nil {
			return fmt.Errorf("reading response: %w", err)
		}

		ifr, ok := c.inflight.Load(rsp.RequestID)
		if !ok {
			// context cancellation aborts inflight requests
			continue
		}

		// handle incoming response
		if done := c.receive(ifr.(*inflight), &rsp); done {
			// stop tracking finished requests
			c.inflight.Delete(rsp.RequestID)
		}
	}
}

func (c *Conn) receive(ifr *inflight, rsp *gremlin.Response) bool {
	result := result{rsp: rsp}
	switch rsp.Status.Code {
	case gremlin.StatusSuccess:
		// quickly handle non fragmented responses
		if ifr.frags == nil {
			break
		}
		// handle fragment
		fallthrough
	case gremlin.StatusPartialContent:
		// append received fragment
		var frag []graphson.RawMessage
		if err := graphson.Unmarshal(rsp.Result.Data, &frag); err != nil {
			result.err = fmt.Errorf("decoding response fragment: %w", err)
			break
		}
		ifr.frags = append(ifr.frags, frag...)

		// partial response requires additional fragments
		if rsp.Status.Code == gremlin.StatusPartialContent {
			return false
		}

		// reassemble fragmented response
		if rsp.Result.Data, result.err = graphson.Marshal(ifr.frags); result.err != nil {
			result.err = fmt.Errorf("assembling fragmented response: %w", result.err)
		}
	case gremlin.StatusAuthenticate:
		// receiver should never block
		c.grp.Go(func() error {
			var buf bytes.Buffer
			if err := graphson.NewEncoder(&buf).Encode(
				gremlin.NewAuthRequest(rsp.RequestID, c.user, c.pass),
			); err != nil {
				return fmt.Errorf("encoding auth request: %w", err)
			}
			select {
			case c.send <- &buf:
			case <-c.ctx.Done():
			}
			return c.ctx.Err()
		})
		return false
	}

	ifr.result <- result
	return true
}
