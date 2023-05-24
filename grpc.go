package fastproto

import (
	"fmt"

	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

// codec is a Codec implementation with protobuf, use fastproto to marshal and unmarshal messages.
type codec struct{}

func ProtoCodec() encoding.Codec {
	return codec{}
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	vv, ok := v.(Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want fastproto.Message", v)
	}
	return Marshal(vv)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	vv, ok := v.(Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want fastproto.Message", v)
	}
	return Unmarshal(data, vv)
}

func (codec) Name() string {
	return proto.Name
}
