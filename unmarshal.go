package fastproto

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

type Unmarshaler interface {
	Unmarshal(data []byte) error
	Reset()
	XxxReset()
	FillMessageInfo()
}

type UnmarshalOptions struct {
	// Merge merges the input into the destination message.
	// The default behavior is to always reset the message before unmarshaling,
	// unless Merge is specified.
	Merge bool

	// sometimes we don't need message info. so we can ignore message info when we unmarshal message
	IgnoreMessageInfo bool
}

func (opt UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
	mm, ok := m.(Unmarshaler)
	if !ok {
		return proto.UnmarshalOptions{Merge: opt.Merge}.Unmarshal(b, m)
	}

	if !opt.Merge {
		if opt.IgnoreMessageInfo {
			mm.XxxReset()
		} else {
			mm.Reset()
		}
	} else {
		if !opt.IgnoreMessageInfo {
			mm.FillMessageInfo()
		}
	}

	return mm.Unmarshal(b)
}

func ConsumeMessage(data []byte, msg proto.Message) (int, error) {
	msglen, n := CalcListLength(data)
	if n < 0 {
		return 0, protowire.ParseError(n)
	}
	data = data[n:]
	if err := Unmarshal(data[:msglen], msg); err != nil {
		return 0, err
	}
	return n + msglen, nil
}

// // Unmarshal parses the wire-format message in b and places the result in m.
// // if m does not implement unmarshaler interface, it will fallback to proto.Unmarshal
func Unmarshal(b []byte, m proto.Message) error {
	return UnmarshalOptions{IgnoreMessageInfo: true}.Unmarshal(b, m)
}
