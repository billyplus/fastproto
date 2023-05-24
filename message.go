package fastproto

import "google.golang.org/protobuf/proto"

type Message interface {
	proto.Message
	Marshaler
	Unmarshaler
	Sizer
}
