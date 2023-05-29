package plugin

import (
	"github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal"
	"github.com/billyplus/fastproto/goimport"
	"github.com/billyplus/fastproto/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	protoMessage    = goimport.ProtoPackage.Ident("Message")
	equal           = goimport.FastProtoPackage.Ident("Equal")
	equalSlice      = goimport.FastProtoHelperPackage.Ident("EqualSlice")
	equalBytesSlice = goimport.FastProtoHelperPackage.Ident("EqualBytesSlice")
	equalProtoSlice = goimport.FastProtoHelperPackage.Ident("EqualProtoSlice")
	equalMap        = goimport.FastProtoHelperPackage.Ident("EqualMap")
	equalBytesMap   = goimport.FastProtoHelperPackage.Ident("EqualBytesMap")
	equalProtoMap   = goimport.FastProtoHelperPackage.Ident("EqualProtoMap")
	equalBytes      = goimport.BytesPackage.Ident("Equal")
)

func init() {
	internal.RegisterPlugin(newEqualer())
}

type equaler struct {
	*protogen.GeneratedFile
}

func newEqualer() internal.Plugin {
	return &equaler{}
}

func (p *equaler) Name() string {
	return "equaler"
}

func (p *equaler) Init() {
}

func (p *equaler) GenerateMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, idx int, m *protogen.Message) {
	if !options.IsEqualer(f.Desc, m.Desc) {
		return
	}

	p.GeneratedFile = g

	g.P(`func (x *`, m.GoIdent.GoName, `) Equal(v2 `, protoMessage, `) bool {`)
	g.P(`	vv, ok := v2.(*`, m.GoIdent.GoName, `)`)
	// g.P(`	if !ok {`)
	// g.P(`		return false`)
	// g.P(`	}`)
	if len(m.Fields) > 0 {
		p.fastCheckField(gen, g, f, m)
		p.slowCheckField(gen, g, f, m)
		// for _, field := range m.Fields {
		// 	p.generateField(f, field)
		// }
	}
	g.P(`	return true`)
	g.P(`}`)
	g.P()
}

func (p *equaler) fastCheckField(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	v := make([]interface{}, 0, 16)
	v = append(v, "		if !ok ")
	for _, field := range m.Fields {
		kind := field.Desc.Kind()
		fieldName := field.GoName

		if field.Desc.IsList() || field.Desc.IsMap() {
			v = append(v, "||\n len(x.", fieldName, ") != len(vv.", fieldName, ")")
			continue
		}

		switch kind {
		case protoreflect.MessageKind:
		case protoreflect.BytesKind:
			v = append(v, "||\n len(x.", fieldName, ") != len(vv.", fieldName, ")")
		default:
			v = append(v, "||\n x.", fieldName, " != vv.", fieldName)
		}
	}
	v = append(v, "{")
	p.P(v...)
	p.P(`        	return false`)
	p.P(`        }`)
}

func (p *equaler) slowCheckField(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	v := make([]interface{}, 0, 16)
	v = append(v, "		if true ")
	for _, field := range m.Fields {
		kind := field.Desc.Kind()
		fieldName := field.GoName

		// if field.Desc.IsList() || field.Desc.IsMap() {
		// 	continue
		// }
		if field.Desc.IsList() {
			switch kind {
			case protoreflect.MessageKind:
				// v = append(v, "|| \n!", equal, "(x.", fieldName, ", vv.", fieldName, ")")
				v = append(v, "|| \n!", equalProtoSlice, "(x.", fieldName, ", vv.", fieldName, ")")
			case protoreflect.BytesKind:
				v = append(v, "|| \n!", equalBytesSlice, "(x.", fieldName, ", vv.", fieldName, ")")
			default:
				v = append(v, "|| \n!", equalSlice, "(x.", fieldName, ", vv.", fieldName, ")")
			}
		} else if field.Desc.IsMap() {
			// key := field.Desc.MapKey()
			value := field.Desc.MapValue()
			switch value.Kind() {
			case protoreflect.MessageKind:
				v = append(v, "|| \n!", equalProtoMap, "(x.", fieldName, ", vv.", fieldName, ")")
			case protoreflect.BytesKind:
				v = append(v, "|| \n!", equalBytesMap, "(x.", fieldName, ", vv.", fieldName, ")")
			default:
				v = append(v, "|| \n!", equalMap, "(x.", fieldName, ", vv.", fieldName, ")")
			}
		} else {

			switch kind {
			case protoreflect.MessageKind:
				v = append(v, "|| \n!", equal, "(x.", fieldName, ", vv.", fieldName, ")")
			case protoreflect.BytesKind:
				v = append(v, "|| \n!", equalBytes, "(x.", fieldName, ", vv.", fieldName, ")")
			default:
			}
		}

	}
	v = append(v, "{")
	p.P(v...)
	p.P(`        	return false`)
	p.P(`        }`)
}

func (p *equaler) generateField(f *protogen.File, field *protogen.Field) {
	if field.Desc.IsMap() {

	} else if field.Desc.IsList() {

	} else {
		p.generateEntry(f, field.GoName, field.Desc)
	}
}

func (p *equaler) generateEntry(f *protogen.File, fieldName string, entryField protoreflect.FieldDescriptor) {
	kind := entryField.Kind()

	switch kind {
	case protoreflect.MessageKind:
		// p.P(`        if !`, fastprotoEqual, "(x.", fieldName, ", vv.", fieldName, "){")
		// p.P(`        	return false`)
		// p.P(`        }`)
	case protoreflect.BytesKind:
		// p.P(`        for i, b := range x.`, fieldName, " {")
		// p.P(`        	if vv.`, fieldName, "[i] != b {")
		// p.P(`        		return false`)
		// p.P(`        	}`)
		// p.P(`        }`)
	default:
		// p.P(`        if x.`, fieldName, " != vv.", fieldName, "{")
		// p.P(`        	return false`)
		// p.P(`        }`)
	}
}
