package fastproto

import "google.golang.org/protobuf/compiler/protogen"

const (
	ProtoWirePackage = protogen.GoImportPath("google.golang.org/protobuf/encoding/protowire")
	ProtoImplPackage = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoimpl")
	FastProtoPackage = protogen.GoImportPath("github.com/billyplus/fastproto")
	ContextPackage   = protogen.GoImportPath("context")
	FmtPackage       = protogen.GoImportPath("fmt")
	MathPackage      = protogen.GoImportPath("math")
)
