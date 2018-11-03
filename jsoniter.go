// Package jsoniter contains a codec to encode and decode entities in JSON
// format using jsoniter
package jsoniter

import (
	"github.com/json-iterator/go"
)

type jsoniterCodec struct {
	jsoniter.API
}

// Codec that encodes to and decodes from JSON using the fast jsoniter.
var Codec = jsoniterCodec{API: jsoniter.ConfigCompatibleWithStandardLibrary}

func (j jsoniterCodec) Name() string {
	return "jsoniter"
}
