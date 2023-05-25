package fastproto

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

type Sizer interface {
	// Message
	Size() int
}

func Size(v proto.Message) int {
	if mm, ok := v.(Sizer); ok {
		return mm.Size()
	}

	return proto.Size(v)
}

func SizeVarintSlice[T int32 | uint32 | int64 | uint64](arr []T) int {
	sz := 0
	for _, v := range arr {
		sz += protowire.SizeVarint(uint64(v))
	}
	return sz
}

func SizeZigZagSlice[T int32 | uint32 | int64 | uint64](arr []T) int {
	sz := 0
	for _, v := range arr {
		sz += protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
	}
	return sz
}
