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

func MarshalTo(data []byte, m Marshaler) (int, error) {
	return MarshalOptions{}.MarshalTo(data, m)
}
