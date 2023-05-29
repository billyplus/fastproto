package fastproto

import (
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
