// Code generated by protoc-gen-go-fast. DO NOT EDIT.
// versions:
//  protoc-gen-go-fast v0.0.1
//  protoc             v4.25.3
// source: test/outer.proto

package pb

import (
	fmt "fmt"
	fastproto "github.com/billyplus/fastproto"
	protohelper "github.com/billyplus/fastproto/protohelper"
	protowire "google.golang.org/protobuf/encoding/protowire"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	math "math"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the packages it is being compiled against.
var _ = fmt.Errorf
var _ = math.MaxFloat32
var _ = protowire.MinValidNumber
var _ = protohelper.Skip
var _ = protoimpl.MinVersion

func (x *OuterMsg) XxxReset() {
	*x = OuterMsg{}
}

func (x *OuterMsg) Unmarshal(data []byte) error {
	for len(data) > 0 {
		num, wireType, n := protowire.ConsumeTag(data)
		if n < 0 {
			return protowire.ParseError(n)
		}
		if num <= 0 {
			return fmt.Errorf("proto: OuterMsg: illegal tag %d (wire type %d)", num, wireType)
		}
		prev := data
		data = data[n:]
		switch num {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Eid", wireType)
			}
			v, n := protohelper.ConsumeVarint[int64](data)
			if n < 0 {
				return protowire.ParseError(n)
			}
			data = data[n:]
			x.Eid = v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenId", wireType)
			}
			v, n := protohelper.ConsumeString(data)
			if n < 0 {
				return protowire.ParseError(n)
			}
			data = data[n:]
			x.OpenId = string(v)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			v, n := protohelper.ConsumeString(data)
			if n < 0 {
				return protowire.ParseError(n)
			}
			data = data[n:]
			x.Name = string(v)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Job", wireType)
			}
			v, n := protohelper.ConsumeVarint[int32](data)
			if n < 0 {
				return protowire.ParseError(n)
			}
			data = data[n:]
			x.Job = v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sex", wireType)
			}
			v, n := protohelper.ConsumeVarint[int32](data)
			if n < 0 {
				return protowire.ParseError(n)
			}
			data = data[n:]
			x.Sex = v
		default:
			n, err := protohelper.Skip(prev[:])
			if err != nil {
				return err
			}
			x.unknownFields = append(x.unknownFields, prev[:n]...)
			data = prev[n:]
		}
	}
	return nil
}

func (x *OuterMsg) MarshalTo(data []byte) (n int, err error) {
	maxN := cap(data)
	data, err = x.AppendToSizedBuffer(data[:0])
	if maxN < len(data) {
		return 0, fmt.Errorf("Not enough space for message(OuterMsg)")
	}
	return len(data), nil
}

func (x *OuterMsg) Marshal() ([]byte, error) {
	data := make([]byte, 0, fastproto.Size(x))
	data, err := x.AppendToSizedBuffer(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (x *OuterMsg) AppendToSizedBuffer(data []byte) (ret []byte, err error) {
	if x.Eid != 0 {
		data = protowire.AppendVarint(data, 8)
		data = protowire.AppendVarint(data, uint64(x.Eid))
	}
	if len(x.OpenId) > 0 {
		data = protowire.AppendVarint(data, 18)
		data = protowire.AppendVarint(data, uint64(len(x.OpenId)))
		data = append(data, x.OpenId...)
	}
	if len(x.Name) > 0 {
		data = protowire.AppendVarint(data, 26)
		data = protowire.AppendVarint(data, uint64(len(x.Name)))
		data = append(data, x.Name...)
	}
	if x.Job != 0 {
		data = protowire.AppendVarint(data, 32)
		data = protowire.AppendVarint(data, uint64(x.Job))
	}
	if x.Sex != 0 {
		data = protowire.AppendVarint(data, 40)
		data = protowire.AppendVarint(data, uint64(x.Sex))
	}
	return data, nil
}

func (x *OuterMsg) Size() (n int) {
	if x == nil {
		return 0
	}
	var l int
	_ = l
	if x.Eid != 0 {
		n += 1 + protowire.SizeVarint(uint64(x.Eid))
	}
	l = len(x.OpenId)
	if l > 0 {
		n += 1 + protowire.SizeBytes(l)
	}
	l = len(x.Name)
	if l > 0 {
		n += 1 + protowire.SizeBytes(l)
	}
	if x.Job != 0 {
		n += 1 + protowire.SizeVarint(uint64(x.Job))
	}
	if x.Sex != 0 {
		n += 1 + protowire.SizeVarint(uint64(x.Sex))
	}
	if x.unknownFields != nil {
		n += len(x.unknownFields)
	}
	x.sizeCache = int32(n)
	return
}
