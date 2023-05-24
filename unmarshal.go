package fastproto

import (
	"google.golang.org/protobuf/encoding/protowire"
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

func (opt UnmarshalOptions) Unmarshal(b []byte, m Unmarshaler) error {
	if !opt.Merge {
		if opt.IgnoreMessageInfo {
			m.XxxReset()
		} else {
			m.Reset()
		}
	} else {
		if !opt.IgnoreMessageInfo {
			m.FillMessageInfo()
		}
	}

	// if x != nil {
	// 	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	// 	if ms.LoadMessageInfo() == nil {
	// 		mi := &file_test_msg_proto_msgTypes[2]
	// 		ms.StoreMessageInfo(mi)
	// 	}
	// }
	return m.Unmarshal(b)
}

func ConsumeMessage(data []byte, msg Unmarshaler) (int, error) {
	msglen, n := CalcListLength(data)
	if n < 0 {
		return 0, protowire.ParseError(n)
	}
	data = data[n:]
	if err := msg.Unmarshal(data[:msglen]); err != nil {
		return 0, err
	}
	return n + msglen, nil
}

// // Unmarshal parses the wire-format message in b and places the result in m.
// // if m does not implement unmarshaler interface, it will fallback to proto.Unmarshal
func Unmarshal(b []byte, m Unmarshaler) error {
	// proto.Reset(m)
	return UnmarshalOptions{}.Unmarshal(b, m)
}
