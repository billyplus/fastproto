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

func ConsumeSlice[T ~int32 | int64 | uint32 | uint64 | float32 | float64 | bool](arr *[]T, data []byte, wireType protowire.Type, valueWireType protowire.Type, consumeValFn func([]byte) (T, int)) (int, error) {
	m := 0
	if wireType == valueWireType {
		v, n := consumeValFn(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		m += n
		// data = data[n:]
		*arr = append(*arr, v)
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
				v, n := consumeValFn(data)
				if n < 0 {
					return n, protowire.ParseError(n)
				}
				data = data[n:]
				m += n
				elementCount--
				(*arr) = append((*arr), v)
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

func ConsumeFixedSlice[T int32 | int64 | uint32 | uint64 | float32 | float64](arr *[]T, data []byte, wireType protowire.Type, consumeValFn func([]byte) (T, int), valSize int) (int, error) {
	m := 0
	if wireType == 1 {
		v, n := consumeValFn(data)
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
				v, n := consumeValFn(data)
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

func ConsumeMap[K int32 | int64 | uint32 | uint64 | string | bool, V any](dst *map[K]V, data []byte, wireType protowire.Type, keyWireType protowire.Type, valueWireType protowire.Type, consumeKey func([]byte) (K, int), consumeVal func([]byte) (V, int)) (int, error) {
	m := 0
	if wireType != 2 {
		return m, fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
	}
	msglen, n := CalcListLength(data)
	if n < 0 {
		return m, protowire.ParseError(n)
	}
	data = data[n:]
	m += n
	if *dst == nil {
		*dst = make(map[K]V)
	}
	var mapkey K
	var mapvalue V
	for msglen > 0 {
		subNum, subWireType, n := protowire.ConsumeTag(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		data, msglen = data[n:], msglen-n
		m += n
		if subNum == 1 {
			if subWireType != keyWireType {
				return m, fmt.Errorf("proto: wrong wireType = %d for field key", subWireType)
			}
			v, n := consumeKey(data)
			if n < 0 {
				return m, protowire.ParseError(n)
			}
			data, msglen = data[n:], msglen-n
			m += n
			mapkey = v
		} else if subNum == 2 {
			if subWireType != valueWireType {
				return m, fmt.Errorf("proto: wrong wireType = %d for field value", subWireType)
			}
			v, n := consumeVal(data)
			if n < 0 {
				return m, protowire.ParseError(n)
			}
			data, msglen = data[n:], msglen-n
			m += n
			mapvalue = v
		} else {
			if skippy, err := Skip(data); err != nil {
				return m, err
			} else {
				data = data[skippy:]
				m += skippy
				msglen -= skippy
			}
		}
	}
	(*dst)[mapkey] = mapvalue
	return m, nil
}

type IMessagePTR[T any] interface {
	proto.Message
	*T
}

func ConsumeMapMessage[K int32 | int64 | uint32 | uint64 | string | bool, IV IMessagePTR[V], V any](dst *map[K]IV, data []byte, wireType protowire.Type, keyWireType protowire.Type, consumeKey func([]byte) (K, int)) (int, error) {
	m := 0
	if wireType != 2 {
		return m, fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
	}
	msglen, n := CalcListLength(data)
	if n < 0 {
		return m, protowire.ParseError(n)
	}
	data = data[n:]
	m += n
	if *dst == nil {
		*dst = make(map[K]IV)
	}
	var mapkey K
	var mapvalue IV
	for msglen > 0 {
		subNum, subWireType, n := protowire.ConsumeTag(data)
		if n < 0 {
			return m, protowire.ParseError(n)
		}
		data, msglen = data[n:], msglen-n
		m += n
		if subNum == 1 {
			if subWireType != keyWireType {
				return m, fmt.Errorf("proto: wrong wireType = %d for field key", subWireType)
			}
			v, n := consumeKey(data)
			if n < 0 {
				return m, protowire.ParseError(n)
			}
			data, msglen = data[n:], msglen-n
			m += n
			mapkey = v
		} else if subNum == 2 {
			if subWireType != 2 {
				return m, fmt.Errorf("proto: wrong wireType = %d for field value", subWireType)
			}
			mapvalue = new(V)
			n, err := ConsumeMessage(data, mapvalue)
			if err != nil {
				return m, err
			}
			data, msglen = data[n:], msglen-n
			m += n
		} else {
			if skippy, err := Skip(data); err != nil {
				return m, err
			} else {
				data = data[skippy:]
				m += skippy
				msglen -= skippy
			}
		}
	}
	(*dst)[mapkey] = mapvalue
	return m, nil
}
