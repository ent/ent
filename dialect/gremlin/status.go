// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

const (
	// StatusSuccess is returned on success.
	StatusSuccess = 200

	// StatusNoContent means the server processed the request but there is no result to return.
	StatusNoContent = 204

	// StatusPartialContent indicates the server successfully returned some content, but there
	// is more in the stream to arrive wait for a success code to signify the end.
	StatusPartialContent = 206

	// StatusUnauthorized means the request attempted to access resources that
	// the requesting user did not have access to.
	StatusUnauthorized = 401

	// StatusAuthenticate denotes a challenge from the server for the client to authenticate its request.
	StatusAuthenticate = 407

	// StatusMalformedRequest means the request message was not properly formatted which means it could not be parsed at
	// all or the "op" code  was not recognized such that Gremlin Server could properly route it for processing.
	// Check the message format and retry the request.
	StatusMalformedRequest = 498

	// StatusInvalidRequestArguments means the request message was parsable, but the arguments supplied in the message
	// were in conflict or incomplete. Check the message format and retry the request.
	StatusInvalidRequestArguments = 499

	// StatusServerError indicates a general server error occurred that prevented the request from being processed.
	StatusServerError = 500

	// StatusScriptEvaluationError is returned when the script submitted for processing evaluated in the ScriptEngine
	// with errors and could not be processed. Check the script submitted for syntax errors or other problems
	// and then resubmit.
	StatusScriptEvaluationError = 597

	// StatusServerTimeout means the server exceeded one of the timeout settings for the request and could therefore
	// only partially responded or did not respond at all.
	StatusServerTimeout = 598

	// StatusServerSerializationError means the server was not capable of serializing an object that was returned from the
	// script supplied on the request. Either transform the object into something Gremlin Server can process within
	// the script or install mapper serialization classes to Gremlin Server.
	StatusServerSerializationError = 599
)

var statusText = map[int]string{
	StatusSuccess:                  "Success",
	StatusNoContent:                "No Content",
	StatusPartialContent:           "Partial Content",
	StatusUnauthorized:             "Unauthorized",
	StatusAuthenticate:             "Authenticate",
	StatusMalformedRequest:         "Malformed Request",
	StatusInvalidRequestArguments:  "Invalid Request Arguments",
	StatusServerError:              "Server Error",
	StatusScriptEvaluationError:    "Script Evaluation Error",
	StatusServerTimeout:            "Server Timeout",
	StatusServerSerializationError: "Server Serialization Error",
}

// StatusText returns status text of code.
func StatusText(code int) string {
	return statusText[code]
}
