// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

// Gremlin server operations.
const (
	// OpsAuthentication used by the client to authenticate itself.
	OpsAuthentication = "authentication"

	// OpsBytecode used for a request that contains the Bytecode representation of a Traversal.
	OpsBytecode = "bytecode"

	// OpsEval used to evaluate a Gremlin script provided as a string.
	OpsEval = "eval"

	// OpsGather used to get a particular side-effect as produced by a previously executed Traversal.
	OpsGather = "gather"

	// OpsKeys used to get all the keys of all side-effects as produced by a previously executed Traversal.
	OpsKeys = "keys"

	// OpsClose used to get all the keys of all side-effects as produced by a previously executed Traversal.
	OpsClose = "close"
)

// Gremlin server operation processors.
const (
	// ProcessorTraversal is the default operation processor.
	ProcessorTraversal = "traversal"
)

const (
	// ArgsBatchSize allows to defines the number of iterations each ResponseMessage should contain
	ArgsBatchSize = "batchSize"

	// ArgsBindings allows to provide a map of key/value pairs to apply
	// as variables in the context of the Gremlin script.
	ArgsBindings = "bindings"

	// ArgsAliases allows to define aliases that represent globally bound Graph and TraversalSource objects.
	ArgsAliases = "aliases"

	// ArgsGremlin corresponds to the Traversal to evaluate.
	ArgsGremlin = "gremlin"

	// ArgsSideEffect allows to specify the unique identifier for the request.
	ArgsSideEffect = "sideEffect"

	// ArgsSideEffectKey allows to specify the key for a specific side-effect.
	ArgsSideEffectKey = "sideEffectKey"

	// ArgsAggregateTo describes how side-effect data should be treated.
	ArgsAggregateTo = "aggregateTo"

	// ArgsLanguage allows to change the flavor of Gremlin used (e.g. gremlin-groovy).
	ArgsLanguage = "language"

	// ArgsEvalTimeout allows to override the server setting that determines
	// the maximum time to wait for a script to execute on the server.
	ArgsEvalTimeout = "scriptEvaluationTimeout"

	// ArgsSasl defines the response to the server authentication challenge.
	ArgsSasl = "sasl"

	// ArgsSaslMechanism defines the SASL mechanism (e.g. PLAIN).
	ArgsSaslMechanism = "saslMechanism"
)
