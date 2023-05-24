package internal

import (
	"fmt"

	"github.com/billyplus/fastproto"
	"google.golang.org/protobuf/compiler/protogen"
)

// GenerateFile generates a _fast.pb.go file implement fast marshalling and unmarshalling.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_fast.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-fast. DO NOT EDIT.")
	g.P("// versions:")
	g.P(fmt.Sprintf("//  protoc-gen-go-fast %s", Version))
	g.P("//  protoc             ", fastproto.ProtocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates the kratos errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the packages it is being compiled against.")
	g.P("var _ = ", fastproto.FmtPackage.Ident("Errorf"))
	g.P("var _ = ", fastproto.MathPackage.Ident("MaxFloat32"))
	g.P("var _ = ", fastproto.ProtoWirePackage.Ident("MinValidNumber"))
	g.P("var _ = ", fastproto.FastProtoPackage.Ident("Skip"))
	g.P("var _ = ", fastproto.ProtoImplPackage.Ident("MinVersion"))
	g.P()

	// for i, imps := 0, file.Desc.Imports(); i < imps.Len(); i++ {
	// 	genImport(gen, file, g, imps.Get(i))
	// }

	for idx, message := range file.Messages {
		genMessage(gen, file, g, idx, message)
	}
}

func genMessage(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, idx int, m *protogen.Message) {
	if m.Desc.IsMapEntry() {
		return
	}

	for _, plugin := range plugins {
		plugin.GenerateMessage(gen, g, file, idx, m)
	}
}
