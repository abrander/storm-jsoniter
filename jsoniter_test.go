package jsoniter

import (
	"testing"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec"
)

func TestStormCodec(t *testing.T) {
	options := &storm.Options{}

	err := storm.Codec(Codec)(options)
	if err != nil {
		t.Errorf("storm.Codec() returned an error: %s", err.Error())
	}
}

func TestCodecName(t *testing.T) {
	name := Codec.Name()
	if name == "" {
		t.Errorf("Name() returned an empty string")
	}
}

// Make sure we implement the needed interface.
var _ codec.MarshalUnmarshaler = (*jsoniterCodec)(nil)
