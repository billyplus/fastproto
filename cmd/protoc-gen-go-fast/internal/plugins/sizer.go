package plugin

import (
	"fmt"

	"github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal"
	"github.com/billyplus/fastproto/goimport"
	"github.com/billyplus/fastproto/options"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	sizeVarint      = goimport.ProtoWirePackage.Ident("SizeVarint")
	sizeFixed32     = goimport.ProtoWirePackage.Ident("SizeFixed32")
	sizeFixed64     = goimport.ProtoWirePackage.Ident("SizeFixed64")
	sizeTag         = goimport.ProtoWirePackage.Ident("SizeTag")
	sizeBytes       = goimport.ProtoWirePackage.Ident("SizeBytes")
	sizeVarintSlice = goimport.FastProtoHelperPackage.Ident("SizeVarintSlice")
	sizeZigZagSlice = goimport.FastProtoHelperPackage.Ident("SizeZigZagSlice")
	size            = goimport.FastProtoPackage.Ident("Size")
)

func init() {
	internal.RegisterPlugin(newSizer())
}

type sizer struct {
	*protogen.GeneratedFile
}

func newSizer() internal.Plugin {
	return &sizer{}
}

func (p *sizer) Name() string {
	return "sizer"
}

func (p *sizer) Init() {
}

func (p *sizer) GenerateMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, idx int, m *protogen.Message) {
	if !options.IsSizer(f.Desc, m.Desc) {
		return
	}

	p.GeneratedFile = g
	p.P(fmt.Sprintf(`func (x *%s) Size() (n int) {`, m.GoIdent.GoName))
	if len(m.Fields) > 0 {
		p.P(`    if x == nil {`)
		p.P(`        return 0`)
		p.P(`    }`)
		p.P(`    var l int`)
		p.P(`    _ = l`)
		oneofList := make(map[*protogen.Oneof]bool, len(m.Fields))
		for _, field := range m.Fields {
			if field.Oneof != nil {
				oneofList[field.Oneof] = true
				continue
			}
			p.generateField(f, field)
		}

		for oneof := range oneofList {
			p.generateOneOf(f, oneof)
		}
		p.P(`    if x.unknownFields != nil {`)
		p.P(`        n += len(x.unknownFields)`)
		p.P(`    }`)
		p.P(`    x.sizeCache = int32(n)`)
	}
	p.P(`    return`)
	p.P(`}`)
	p.P()
}

func (p *sizer) generateField(f *protogen.File, field *protogen.Field) {
	// repeated := field.Desc.IsList()

	if field.Desc.IsList() {
		p.generateList(f, field)
	} else if field.Desc.IsMap() {
		// p.P(`    n += `, keysize)
		p.generateMap(field)
	} else {
		p.generateEntry("x."+field.GoName, field.Desc)
	}

	// keysize := protowire.SizeTag(field.Desc.Number())
	// switch field.Desc.Kind() {
	// case protoreflect.EnumKind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`    	l = 0`)
	// 		p.P(`       for _, v := range x.`, fieldName, ` {`)
	// 		p.P(`    		l += `, sizeVarint, `(uint64(v))`)
	// 		p.P(`       }`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(l)`)
	// 		p.P(`    }`)

	// 	} else {
	// 		p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(uint64(x.`, fieldName, `))}`)
	// 	}
	// case protoreflect.Int32Kind,
	// 	protoreflect.Int64Kind,
	// 	protoreflect.Uint32Kind,
	// 	protoreflect.Uint64Kind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, sizeVarintSlice, `(x.`, fieldName, `))`)
	// 		p.P(`    }`)
	// 	} else {
	// 		p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(uint64(x.`, fieldName, `))}`)
	// 	}
	// case protoreflect.Sint32Kind,
	// 	protoreflect.Sint64Kind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, sizeZigZagSlice, `(x.`, fieldName, `))`)
	// 		p.P(`    }`)
	// 	} else {
	// 		p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(`, encodeZigZag, `(int64(x.`, fieldName, `)))}`)
	// 	}

	// case protoreflect.FloatKind,
	// 	protoreflect.Fixed32Kind,
	// 	protoreflect.Sfixed32Kind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `)*4)`)
	// 		p.P(`    }`)
	// 	} else {
	// 		p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, `+4}`)
	// 	}
	// case protoreflect.DoubleKind,
	// 	protoreflect.Fixed64Kind,
	// 	protoreflect.Sfixed64Kind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `)*8)`)
	// 		p.P(`    }`)
	// 	} else {
	// 		p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, `+8}`)
	// 	}
	// case protoreflect.BoolKind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `))`)
	// 		p.P(`    }`)
	// 	} else {
	// 		p.P(`    if x.`, fieldName, ` {n += `, keysize+1, `}`)
	// 	}
	// case protoreflect.StringKind, protoreflect.BytesKind:
	// 	if repeated {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		p.P(`        for _, s := range x.`, fieldName, ` {`)
	// 		p.P(`    		n += `, keysize, ` + `, sizeBytes, `(len(s))`)
	// 		p.P(`        }`)
	// 		p.P(`    }`)
	// 	} else {
	// 		p.P(`    l = len(x.`, fieldName, `)`)
	// 		p.P(`    if l > 0 {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(l)`)
	// 		p.P(`    }`)
	// 	}
	// case protoreflect.MessageKind:
	// 	if field.Desc.IsMap() {
	// 		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	// 		// p.P(`    n += `, keysize)
	// 		p.sizerMap(field)
	// 		p.P(`    }`)
	// 	} else if repeated {
	// 		p.P(`        for _, e := range x.`, fieldName, ` {`)
	// 		p.P(`        	n += `, keysize, ` + `, sizeBytes, `(`, size, `(e))`)
	// 		p.P(`        }`)

	// 	} else {
	// 		p.P(`    if x.`, fieldName, `!=nil {`)
	// 		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, size, `(x.`, fieldName, `))`)
	// 		p.P(`    }`)
	// 	}
	// }
}

func (p *sizer) generateList(f *protogen.File, field *protogen.Field) {
	// repeated := field.Desc.IsList()
	fieldName := field.GoName

	keysize := protowire.SizeTag(field.Desc.Number())
	switch field.Desc.Kind() {
	case protoreflect.EnumKind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`    	l = 0`)
		p.P(`       for _, v := range x.`, fieldName, ` {`)
		p.P(`    		l += `, sizeVarint, `(uint64(v))`)
		p.P(`       }`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(l)`)
		p.P(`    }`)
	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, sizeVarintSlice, `(x.`, fieldName, `))`)
		p.P(`    }`)
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, sizeZigZagSlice, `(x.`, fieldName, `))`)
		p.P(`    }`)
	case protoreflect.FloatKind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed32Kind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `)*4)`)
		p.P(`    }`)
	case protoreflect.DoubleKind,
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `)*8)`)
		p.P(`    }`)
	case protoreflect.BoolKind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `))`)
		p.P(`    }`)
	case protoreflect.StringKind, protoreflect.BytesKind:
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
		p.P(`        for _, s := range x.`, fieldName, ` {`)
		p.P(`    		n += `, keysize, ` + `, sizeBytes, `(len(s))`)
		p.P(`        }`)
		p.P(`    }`)
	case protoreflect.MessageKind:
		p.P(`        for _, e := range x.`, fieldName, ` {`)
		p.P(`        	n += `, keysize, ` + `, sizeBytes, `(`, size, `(e))`)
		p.P(`        }`)
	}
}

func (p *sizer) generateEntry(fieldName string, field protoreflect.FieldDescriptor) {
	keysize := protowire.SizeTag(field.Number())
	switch field.Kind() {
	case protoreflect.EnumKind:
		p.P(`    if `, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(uint64(`, fieldName, `))}`)
	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind:
		p.P(`    if `, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(uint64(`, fieldName, `))}`)
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`    if `, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(`, encodeZigZag, `(int64(`, fieldName, `)))}`)
	case protoreflect.FloatKind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed32Kind:
		p.P(`    if `, fieldName, `!=0 {n += `, keysize, `+4}`)
	case protoreflect.DoubleKind,
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind:
		p.P(`    if `, fieldName, `!=0 {n += `, keysize, `+8}`)
	case protoreflect.BoolKind:
		p.P(`    if `, fieldName, ` {n += `, keysize+1, `}`)
	case protoreflect.StringKind, protoreflect.BytesKind:
		p.P(`    l = len(`, fieldName, `)`)
		p.P(`    if l > 0 {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(l)`)
		p.P(`    }`)
	case protoreflect.MessageKind:
		p.P(`    if `, fieldName, `!=nil {`)
		p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, size, `(`, fieldName, `))`)
		p.P(`    }`)
	}
}

func (p *sizer) generateMap(field *protogen.Field) {
	keysize := protowire.SizeTag(field.Desc.Number())
	key := field.Desc.MapKey()
	value := field.Desc.MapValue()

	p.P(`    if len(x.`, field.GoName, `) > 0 {`)
	p.P(`        for k, v := range x.`, field.GoName, ` {`)
	p.P(`            _, _ = k, v`)
	tmp := append([]interface{}{`            l = `}, keySizer[key.Kind()]...)

	if value.Kind() == protoreflect.MessageKind {
		p.P(`               sz := `, size, `(v)`)
		tmp = append(tmp, "+ sz + 1 + ", sizeVarint, `(uint64(sz))`)
	} else {
		tmp = append(tmp, " + ")
		tmp = append(tmp, valueSizer[value.Kind()]...)
	}
	p.P(tmp...)

	p.P(`        	n += `, keysize, ` + l  + `, sizeVarint, `(uint64(l))`)
	p.P(`        }`)
	p.P(`    }`)
}

func (p *sizer) generateOneOf(f *protogen.File, oneof *protogen.Oneof) {
	if len(oneof.Fields) == 0 {
		return
	}
	p.P(`        switch vv:= x.Get`, oneof.GoName, "().(type){")
	for _, field := range oneof.Fields {
		p.P("        case *", field.GoIdent, ":")
		p.generateEntry("vv."+field.GoName, field.Desc)
	}
	p.P(`        default:`)
	p.P(`        }`)
}

var keySizer = [][]interface{}{
	protoreflect.Int32Kind:    {"1 +", sizeVarint, "(uint64(k))"},
	protoreflect.Int64Kind:    {"1 +", sizeVarint, "(uint64(k))"},
	protoreflect.FloatKind:    {"5"},
	protoreflect.DoubleKind:   {"9"},
	protoreflect.Uint32Kind:   {"1 +", sizeVarint, "(uint64(k))"},
	protoreflect.Uint64Kind:   {"1 +", sizeVarint, "(uint64(k))"},
	protoreflect.Sint32Kind:   {"1 +", sizeVarint, "(", encodeZigZag, "(int64(k)))"},
	protoreflect.Sint64Kind:   {"1 +", sizeVarint, "(", encodeZigZag, "(k))"},
	protoreflect.Fixed32Kind:  {"5"},
	protoreflect.Fixed64Kind:  {"9"},
	protoreflect.Sfixed32Kind: {"5"},
	protoreflect.Sfixed64Kind: {"9"},
	protoreflect.BoolKind:     {"2"},
	protoreflect.StringKind:   {"1 + ", sizeBytes, "(len(k))"},
}

var valueSizer = [][]interface{}{
	protoreflect.Int32Kind:    {"1 +", sizeVarint, "(uint64(v))"},
	protoreflect.Int64Kind:    {"1 +", sizeVarint, "(uint64(v))"},
	protoreflect.FloatKind:    {"5"},
	protoreflect.DoubleKind:   {"9"},
	protoreflect.Uint32Kind:   {"1 +", sizeVarint, "(uint64(v))"},
	protoreflect.Uint64Kind:   {"1 +", sizeVarint, "(uint64(v))"},
	protoreflect.Sint32Kind:   {"1 +", sizeVarint, "(", encodeZigZag, "(int64(v)))"},
	protoreflect.Sint64Kind:   {"1 +", sizeVarint, "(", encodeZigZag, "(v))"},
	protoreflect.Fixed32Kind:  {"5"},
	protoreflect.Fixed64Kind:  {"9"},
	protoreflect.Sfixed32Kind: {"5"},
	protoreflect.Sfixed64Kind: {"9"},
	protoreflect.BoolKind:     {"2"},
	protoreflect.StringKind:   {"1 + ", sizeBytes, "(len(v))"},
	protoreflect.BytesKind:    {"1 + ", sizeBytes, "(len(v))"},
	protoreflect.EnumKind:     {"1 +", sizeVarint, "(uint64(v))"},
}
