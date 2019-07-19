package graphson

import (
	"github.com/json-iterator/go"
)

var config = jsoniter.Config{}.Froze()

func init() {
	config.RegisterExtension(&encodeExtension{})
	config.RegisterExtension(&decodeExtension{})
}
