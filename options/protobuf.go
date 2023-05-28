package options

import (
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func IsMarshaler(fileDesc protoreflect.FileDescriptor, msgDesc protoreflect.MessageDescriptor) bool {
	if GetExtension[bool](msgDesc.Options(), E_FastprotoMsgNoMarshaler) {
		// fmt.Println("msg no marshaler")
		return false
	}
	if GetExtension[bool](fileDesc.Options(), E_FastprotoNoMarshaler) {
		// fmt.Println("file no marshaler")
		return GetExtension[bool](msgDesc.Options(), E_FastprotoMsgMarshaler)
	}

	return true
}

func IsUnmarshaler(fileDesc protoreflect.FileDescriptor, msgDesc protoreflect.MessageDescriptor) bool {
	if GetExtension[bool](msgDesc.Options(), E_FastprotoMsgNoUnmarshaler) {
		return false
	}

	if GetExtension[bool](fileDesc.Options(), E_FastprotoNoUnmarshaler) {
		return GetExtension[bool](msgDesc.Options(), E_FastprotoMsgUnmarshaler)
	}

	return true
}

func IsSizer(fileDesc protoreflect.FileDescriptor, msgDesc protoreflect.MessageDescriptor) bool {
	if IsMarshaler(fileDesc, msgDesc) {
		return true
	}
	if GetExtension[bool](msgDesc.Options(), E_FastprotoMsgSizer) {
		return true
	}
	return GetExtension[bool](fileDesc.Options(), E_FastprotoSizer)
}

func GetExtension[T any](m proto.Message, xt protoreflect.ExtensionType) T {
	return proto.GetExtension(m, xt).(T)
}
