package fastproto

import (
	"fmt"

	"google.golang.org/grpc/encoding"
	protoenc "google.golang.org/grpc/encoding/proto"
	"google.golang.org/protobuf/proto"
)

// codec is a Codec implementation with protobuf, use fastproto to marshal and unmarshal messages.
type codec struct {
	// marshal   MarshalOptions
	unmarshal UnmarshalOptions
}

func ProtoCodec(opt ...ProtoCodecOption) encoding.Codec {
	c := &codec{
		// marshal:   MarshalOptions{},
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

// func ProtoCodecMarshal(m MarshalOptions) ProtoCodecOption {
// 	return func(c *codec) {
// 		c.marshal = m
// 	}
// }

func (codec) Marshal(v interface{}) ([]byte, error) {
	vv, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
	}
	return Marshal(vv)
}

func (c codec) Unmarshal(data []byte, v interface{}) error {
	vv, ok := v.(proto.Message)
	if ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
	}
	return c.unmarshal.Unmarshal(data, vv)
}

// func (c *codec) Marshal(v interface{}) ([]byte, error) {
// 	if vv, ok := v.(Message); ok {
// 		return vv.Marshal()
// 	}
// 	// if message is not genereated by fastproto, try standard proto
// 	if vv, ok := v.(proto.Message); ok {
// 		return proto.Marshal(vv)
// 	}

// 	return nil, fmt.Errorf("failed to marshal, message is %T, want fastproto.Message", v)
// }

// func (c *codec) Unmarshal(data []byte, v interface{}) error {
// 	if vv, ok := v.(Message);ok{
// 		return c.unmarshal.Unmarshal(data, vv)
// 	}
// 	if
// 	if !ok {
// 		return fmt.Errorf("failed to unmarshal, message is %T, want fastproto.Message", v)
// 	}
// }

func (codec) Name() string {
	return protoenc.Name
}
