package fastproto

import "google.golang.org/protobuf/proto"

type Marshaler interface {
	MarshalTo(data []byte) (n int, err error)
	Marshal() ([]byte, error)
	AppendToSizedBuffer(data []byte) ([]byte, error)
}

// type MarshalOptions struct {
// }

// func (opt MarshalOptions) Marshal(m Marshaler) ([]byte, error) {
// 	return m.Marshal()
// }

// func (opt MarshalOptions) MarshalTo(data []byte, m Marshaler) (int, error) {
// 	return m.MarshalTo(data)
// }

func Marshal(m proto.Message) ([]byte, error) {
	if mm, ok := m.(Marshaler); ok {
		return mm.Marshal()
	}
	return proto.Marshal(m)
}

// data must have enough space for message which means cap(data) >= msg.Size(), or else it would return error
// the return int indicate how many bytes of data is used.
// data[:n] is encoded message.
func MarshalTo(data []byte, m proto.Message) (int, error) {
	if mm, ok := m.(Marshaler); ok {
		return mm.MarshalTo(data)
	}
	b, err := proto.Marshal(m)
	if err != nil {
		return 0, err
	}
	copy(data[:], b)
	return len(b), nil
}

func AppendToSizedBuffer(data []byte, m proto.Message) ([]byte, error) {
	if mm, ok := m.(Marshaler); ok {
		return mm.AppendToSizedBuffer(data)
	}

	b, err := proto.Marshal(m)
	if err != nil {
		return data, err
	}

	return append(data, b...), nil
}
