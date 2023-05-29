package protohelper

import (
	"bytes"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/proto"
)

func EqualBytes(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	return bytes.Equal(b1, b2)
}

func EqualSlice[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true
}

func EqualBytesSlice(s1, s2 [][]byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if vv := s2[k]; 0 != bytes.Compare(v, vv) {
			return false
		}
	}

	return true
}

func EqualProtoSlice[V proto.Message](s1, s2 []V) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if vv := s2[k]; !fastproto.Equal(v, vv) {
			return false
		}
	}

	return true
}

func EqualMap[K comparable, V comparable](s1, s2 map[K]V) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if vv, ok := s2[k]; !ok || vv != v {
			return false
		}
	}

	return true
}

func EqualProtoMap[K comparable, V proto.Message](s1, s2 map[K]V) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if vv, ok := s2[k]; !ok || !fastproto.Equal(v, vv) {
			return false
		}
	}

	return true
}

func EqualBytesMap[K comparable](s1, s2 map[K][]byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if vv, ok := s2[k]; !ok || 0 != bytes.Compare(v, vv) {
			return false
		}
	}

	return true
}
