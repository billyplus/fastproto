package protohelper

import (
	"fmt"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

func ConsumeMessage(data []byte, msg proto.Message) (int, error) {
	msglen, n := CalcListLength(data)
	if n < 0 {
		return 0, protowire.ParseError(n)
	}
	data = data[n:]
	if err := fastproto.Unmarshal(data[:msglen], msg); err != nil {
		return 0, err
	}
	return n + msglen, nil
}

func ConsumeSlice[T int32 | int64 | uint32 | uint64](arr *[]T, data []byte, wireType protowire.Type) (int, error) {
	m := 0
	if wireType == 0 {
		v, n := protowire.ConsumeVarint(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		m += n
		// data = data[n:]
		*arr = append(*arr, T(v))
	} else if wireType == 2 {
		msglen, n := CalcListLength(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		m += n
		data = data[n:]
		elementCount := 0
		for _, i := range data[:msglen] {
			if i < 128 {
				elementCount++
			}
		}
		if elementCount > 0 {
			if len(*arr) == 0 {
				(*arr) = make([]T, 0, elementCount)
			} else {
				ss := make([]T, 0, elementCount+len((*arr)))
				ss = append(ss, (*arr)...)
				(*arr) = ss
			}

			for elementCount > 0 {
				v, n := protowire.ConsumeVarint(data)
				if n < 0 {
					return n, protowire.ParseError(n)
				}
				data = data[n:]
				m += n
				elementCount--
				(*arr) = append((*arr), T(v))
			}
		}
	} else {
		return m, fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
	}
	return m, nil
}

func ConsumeSignedSlice[T int32 | int64](arr *[]T, data []byte, wireType protowire.Type) (int, error) {
	m := 0
	if wireType == 0 {
		v, n := protowire.ConsumeVarint(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		m += n
		(*arr) = append((*arr), T(protowire.DecodeZigZag(v)))
	} else if wireType == 2 {
		msglen, n := CalcListLength(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		data = data[n:]
		m += n
		elementCount := 0
		for _, i := range data[:msglen] {
			if i < 128 {
				elementCount++
			}
		}
		if elementCount > 0 {
			if len((*arr)) == 0 {
				(*arr) = make([]T, 0, elementCount)
			} else {
				ss := make([]T, 0, elementCount+len((*arr)))
				ss = append(ss, (*arr)...)
				(*arr) = ss
			}
			for elementCount > 0 {
				v, n := protowire.ConsumeVarint(data)
				if n < 0 {
					return m, protowire.ParseError(n)
				}
				data = data[n:]
				m += n
				elementCount--
				(*arr) = append((*arr), T(protowire.DecodeZigZag(v)))
			}
		}
	} else {
		return m, fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
	}
	return m, nil
}

func ConsumeFixedSlice[T int32 | int64 | uint32 | uint64](arr *[]T, data []byte, wireType protowire.Type, consumefn func([]byte) (T, int), valSize int) (int, error) {
	m := 0
	if wireType == 1 {
		v, n := consumefn(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		m += n
		(*arr) = append((*arr), T(v))
	} else if wireType == 2 {
		msglen, n := CalcListLength(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		data = data[n:]
		m += n
		elementCount := msglen / valSize
		if elementCount > 0 {
			if len((*arr)) == 0 {
				(*arr) = make([]T, 0, elementCount)
			} else {
				ss := make([]T, 0, elementCount+len((*arr)))
				ss = append(ss, (*arr)...)
				(*arr) = ss
			}
			for elementCount > 0 {
				v, n := consumefn(data)
				if n < 0 {
					return m, protowire.ParseError(n)
				}
				data = data[n:]
				m += n
				elementCount--
				(*arr) = append((*arr), T(v))
			}
		}
	} else {
		return m, fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
	}
	return m, nil
}
