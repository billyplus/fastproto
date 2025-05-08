package fastproto

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Unmarshaler interface {
	Unmarshal(data []byte) error
	XxxReset()
}

type UnmarshalOptions struct {
	// Merge merges the input into the destination message.
	// The default behavior is to always reset the message before unmarshaling,
	// unless Merge is specified.
	Merge bool

	// Resolver is used for looking up types when unmarshaling extension fields.
	// If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)
		FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
	}
}

func (opt UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
	mm, ok := m.(Unmarshaler)
	if !ok {
		return proto.UnmarshalOptions{Merge: opt.Merge}.Unmarshal(b, m)
	}

	if !opt.Merge {
		mm.XxxReset()
	}

	return mm.Unmarshal(b)
}

// Unmarshal parses the wire-format message in b and places the result in m.
// if m does not implement unmarshaler interface, it will fallback to proto.Unmarshal
// merge = true is the default behavior
func Unmarshal(b []byte, m proto.Message) error {
	return UnmarshalOptions{}.Unmarshal(b, m)
}
