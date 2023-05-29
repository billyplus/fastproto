package protohelper

import "google.golang.org/protobuf/encoding/protowire"

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
