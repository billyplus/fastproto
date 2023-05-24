package fastproto

type Marshaler interface {
	MarshalTo(data []byte) (n int, err error)
	Marshal() ([]byte, error)
}

type MarshalOptions struct {
}

func (opt MarshalOptions) Marshal(m Marshaler) ([]byte, error) {
	return m.Marshal()
}

func (opt MarshalOptions) MarshalTo(data []byte, m Marshaler) (int, error) {
	return m.MarshalTo(data)
}

func Marshal(m Marshaler) ([]byte, error) {
	return MarshalOptions{}.Marshal(m)
}

// data must have enough space for message which means cap(data) >= msg.Size(), or else it would return error
// the return int indicate how many bytes of data is used.
// data[:n] is encoded message.
func MarshalTo(data []byte, m Marshaler) (int, error) {
	return MarshalOptions{}.MarshalTo(data, m)
}
