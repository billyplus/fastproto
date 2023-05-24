package fastproto

import (
	"fmt"

	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

// codec is a Codec implementation with protobuf, use fastproto to marshal and unmarshal messages.
type codec struct {
	marshal   MarshalOptions
	unmarshal UnmarshalOptions
}

func ProtoCodec(opt ...ProtoCodecOption) encoding.Codec {
	c := &codec{
		marshal:   MarshalOptions{},
		unmarshal: UnmarshalOptions{},
	}
	for _, o := range opt {
		o(c)
	}

	return c
}

type ProtoCodecOption func(c *codec)

func ProtoCodecUnmarshal(u UnmarshalOptions) ProtoCodecOption {
	return func(c *codec) {
		c.unmarshal = u
	}
}

func ProtoCodecMarshal(m MarshalOptions) ProtoCodecOption {
	return func(c *codec) {
		c.marshal = m
	}
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	vv, ok := v.(Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want fastproto.Message", v)
	}
	return c.marshal.Marshal(vv)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	vv, ok := v.(Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want fastproto.Message", v)
	}
	return c.unmarshal.Unmarshal(data, vv)
}

func (codec) Name() string {
	return proto.Name
}
