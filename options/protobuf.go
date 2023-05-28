package options

import (
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func IsMarshaler(fileDesc protoreflect.FileDescriptor, msgDesc protoreflect.MessageDescriptor) bool {
	if proto.HasExtension(msgDesc.Options(), E_FastprotoMsgMarshaler) {
		return true
	}
	return proto.HasExtension(fileDesc.Options(), E_FastprotoMarshaler)
}

func IsUnmarshaler(fileDesc protoreflect.FileDescriptor, msgDesc protoreflect.MessageDescriptor) bool {
	if proto.HasExtension(msgDesc.Options(), E_FastprotoMsgUnmarshaler) {
		return true
	}
	return proto.HasExtension(fileDesc.Options(), E_FastprotoUnmarshaler)
}

func IsSizer(fileDesc protoreflect.FileDescriptor, msgDesc protoreflect.MessageDescriptor) bool {
	if IsMarshaler(fileDesc, msgDesc) {
		return true
	}
	if proto.HasExtension(msgDesc.Options(), E_FastprotoMsgSizer) {
		return true
	}
	return proto.HasExtension(fileDesc.Options(), E_FastprotoSizer)
}
