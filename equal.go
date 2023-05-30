package fastproto

import (
	"google.golang.org/protobuf/proto"
)

type Equaler interface {
	Equal(v2 proto.Message) bool
}

func Equal(m1, m2 proto.Message) bool {
	if m1 == m2 {
		return true
	}
	if mm1, ok := m1.(Equaler); ok {
		return mm1.Equal(m2)
	}

	return proto.Equal(m1, m2)
}
