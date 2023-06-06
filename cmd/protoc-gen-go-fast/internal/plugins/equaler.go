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
	g.P(`	if !ok {`)
	g.P(`		return false`)
	g.P(`	}`)
	g.P(`	if x == nil || vv == nil {`)
	g.P(`		return x == vv`)
	g.P(`	}`)
	g.P(`	if x == vv {`)
	g.P(`		return true`)
	g.P(`	}`)
	if len(m.Fields) > 0 {
		p.fastCheckField(gen, g, f, m)
		p.slowCheckField(gen, g, f, m)

		oneofList := make(map[*protogen.Oneof]bool, len(m.Fields))
		for _, field := range m.Fields {
			if field.Oneof != nil {
				oneofList[field.Oneof] = true
			}
		}

		for oneof := range oneofList {
			p.generateOneOf(f, oneof)
		}
	}
	g.P(`	return true`)
	g.P(`}`)
	g.P()
}

func (p *equaler) fastCheckField(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	v := make([]interface{}, 0, 16)
	v = append(v, "		if false ")
	for _, field := range m.Fields {
		if field.Oneof != nil {
			continue
		}

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
	v = append(v, "		if false ")
	for _, field := range m.Fields {
		if field.Oneof != nil {
			continue
		}

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

func (p *equaler) generateOneOf(f *protogen.File, oneof *protogen.Oneof) {
	if len(oneof.Fields) == 0 {
		return
	}
	p.P(`        switch xx:= x.Get`, oneof.GoName, "().(type){")
	for _, field := range oneof.Fields {
		p.P("        case *", field.GoIdent, ":")
		p.P(`    if vv2,ok:=vv.Get`, oneof.GoName, `().(*`, field.GoIdent, `); !ok {`)
		p.P(`    	return false`)
		kind := field.Desc.Kind()
		switch kind {
		case protoreflect.MessageKind:
			// v = append(v, "|| \n!", equal, "(x.", fieldName, ", vv.", fieldName, ")")
			p.P(`    } else if !`, equal, "(vv2.", field.GoName, `, xx.`, field.GoName, `) {`)
		case protoreflect.BytesKind:
			p.P(`    } else if !`, equalBytes, "(vv2.", field.GoName, `, xx.`, field.GoName, `) {`)
			// v = append(v, "|| \n!", equalBytes, "(x.", fieldName, ", vv.", fieldName, ")")
			// p.P(`    } else if vv2.`, field.GoName, ` != xx.`, field.GoName, `{`)
		default:
			p.P(`    } else if vv2.`, field.GoName, ` != xx.`, field.GoName, `{`)
		}
		p.P(`    	return false`)
		p.P(`    }`)
		// if kind == protoreflect.StringKind || kind == protoreflect.BytesKind {
		// 	p.P(`    if vv2,ok:=vv.Get`, oneof.GoName, `().(*`, field.GoIdent, `); !ok {`)
		// 	p.P(`    	return false`)
		// 	p.P(`    } else if vv2.`, field.GoName, ` != xx.`, field.GoName, `{`)
		// 	p.P(`    	return false`)
		// 	// p.P(`    }`)
		// } else if kind == protoreflect.MessageKind {
		// 	p.P(`    if xx.`, field.GoName, ` != nil {`)
		// } else if kind == protoreflect.BoolKind {
		// 	p.P(`    if xx.`, field.GoName, ` {`)
		// } else {
		// 	p.P(`    if xx.`, field.GoName, ` != 0{`)
		// }
		// p.generateEntry(f, "xx."+field.GoName, field.Desc)
		// p.P(`        }`)
	}
	p.P(`        default:`)
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
