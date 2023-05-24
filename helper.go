package fastproto

import (
	"strings"
	"unicode/utf8"

	"google.golang.org/protobuf/compiler/protogen"
)

func Bool2Int(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func FieldMessageName(gen *protogen.Plugin, file *protogen.File, field *protogen.Field) string {
	impFile, ok := gen.FilesByPath[field.Desc.ParentFile().Path()]
	if !ok {
		return field.GoName
	}

	if impFile.GoImportPath == file.GoImportPath {
		// types in the same Go package.
		return field.GoName
	}

	return string(field.Message.Desc.FullName())

}

func MessageTypesVarName(f *protogen.File) string {
	return fileVarName(f, "msgTypes")
}

func fileVarName(f *protogen.File, suffix string) string {
	prefix := f.GoDescriptorIdent.GoName
	_, n := utf8.DecodeRuneInString(prefix)
	prefix = strings.ToLower(prefix[:n]) + prefix[n:]
	return prefix + "_" + suffix
}
