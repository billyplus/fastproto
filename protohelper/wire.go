package protohelper

import (
	"errors"
	"fmt"
	"io"
	"math"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
	ErrOverflow             = errors.New("variable length integer overflow")
)

type VarintType interface {
	int32 | uint32 | int64 | uint64 | bool
}

const (
	_ = -iota
	errCodeTruncated
	errCodeFieldNumber
	errCodeOverflow
	errCodeReserved
	errCodeEndGroup
)

func ConsumeVarint[T int32 | uint32 | int64 | uint64](data []byte) (T, int) {
	v, n := protowire.ConsumeVarint(data)
	if n < 0 {
		return 0, n
	}
	return T(v), n
}

func ConsumeString(data []byte) (string, int) {
	v, n := protowire.ConsumeBytes(data)
	if n < 0 {
		return "", n
	}
	return string(v), n
}

func ConsumeSint[T int32 | int64](data []byte) (T, int) {
	v, n := protowire.ConsumeVarint(data)
	if n < 0 {
		return 0, n
	}
	return T(protowire.DecodeZigZag(v)), n
}

func ConsumeFloat32(data []byte) (float32, int) {
	v, n := protowire.ConsumeFixed32(data)
	if n < 0 {
		return 0, n
	}
	return math.Float32frombits(v), n
}

func ConsumeFloat64(data []byte) (float64, int) {
	v, n := protowire.ConsumeFixed64(data)
	if n < 0 {
		return 0, n
	}
	return math.Float64frombits(v), n
}

func ConsumeBool(data []byte) (bool, int) {
	v, n := protowire.ConsumeVarint(data)
	if n < 0 {
		return false, n
	}
	return v != 0, n
}

func ConsumeEnum[T ~int32](data []byte) (T, int) {
	v, n := protowire.ConsumeVarint(data)
	if n < 0 {
		return 0, n
	}
	return T(v), n
}

func ConsumeFixed32[T int32 | uint32 | float32](data []byte) (T, int) {
	v, n := protowire.ConsumeFixed32(data)
	if n < 0 {
		return 0, n
	}
	return T(v), n
}

func ConsumeFixed64[T int64 | uint64 | float64](data []byte) (T, int) {
	v, n := protowire.ConsumeFixed64(data)
	if n < 0 {
		return 0, n
	}
	return T(v), n
}

func CalcListLength(data []byte) (int, int) {
	v, n := protowire.ConsumeVarint(data)
	if n < 0 {
		return 0, n
	}
	// msglen := int(v)
	if int(v) < 0 || int(v)+n > len(data) {
		return 0, errCodeOverflow
	}

	return int(v), n
}

func Skip(data []byte) (n int, err error) {
	l := len(data)
	idx := 0
	depth := 0
	for idx < l {
		wire, n := protowire.ConsumeVarint(data)
		if n < 0 {
			return 0, protowire.ParseError(n)
		}
		idx += n
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if idx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				idx++
				if data[idx-1] < 0x80 {
					break
				}
			}
		case 1:
			idx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if idx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[idx]
				idx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			idx += length
		case 3:
			// deprecated
			depth++
		case 4:
			// deprecated
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			idx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if idx < 0 {
			return 0, ErrInvalidLength
		}
		if idx > len(data) {
			return 0, io.ErrUnexpectedEOF
		}
		if depth == 0 {
			return idx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

func KindToType(kind protoreflect.Kind) protowire.Type {
	switch kind {
	case
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.DoubleKind:

		return protowire.Fixed64Type

	case protoreflect.Fixed32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.FloatKind:

		return protowire.Fixed32Type

	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind,
		protoreflect.BoolKind,
		protoreflect.EnumKind:

		return protowire.VarintType

	case protoreflect.StringKind,
		protoreflect.BytesKind,
		protoreflect.MessageKind:

		return protowire.BytesType

	case protoreflect.GroupKind:

		return protowire.BytesType
	}
	panic("unreachable")
}

func GoTypeOfField(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64"
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.BytesKind:
		return "[]byte"
	case protoreflect.MessageKind:
		return "*" + string(field.Message().Name())
	case protoreflect.EnumKind:
		return string(field.Enum().Name())
	default:
		return ""
	}
}
