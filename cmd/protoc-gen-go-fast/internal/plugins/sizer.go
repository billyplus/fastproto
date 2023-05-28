package plugin

import (
	"fmt"

	"github.com/billyplus/fastproto"
	"github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal"
	"github.com/billyplus/fastproto/options"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	sizeVarint      = fastproto.ProtoWirePackage.Ident("SizeVarint")
	sizeFixed32     = fastproto.ProtoWirePackage.Ident("SizeFixed32")
	sizeFixed64     = fastproto.ProtoWirePackage.Ident("SizeFixed64")
	sizeTag         = fastproto.ProtoWirePackage.Ident("SizeTag")
	sizeBytes       = fastproto.ProtoWirePackage.Ident("SizeBytes")
	sizeVarintSlice = fastproto.FastProtoPackage.Ident("SizeVarintSlice")
	sizeZigZagSlice = fastproto.FastProtoPackage.Ident("SizeZigZagSlice")
	size            = fastproto.FastProtoPackage.Ident("Size")
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
		for _, field := range m.Fields {
			p.GenerateField(f, field)
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

func (p *sizer) GenerateField(f *protogen.File, field *protogen.Field) {
	repeated := field.Desc.IsList()
	fieldName := field.GoName

	keysize := protowire.SizeTag(field.Desc.Number())
	switch field.Desc.Kind() {
	case protoreflect.EnumKind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`    	l = 0`)
			p.P(`       for _, v := range x.`, fieldName, ` {`)
			p.P(`    		l += `, sizeVarint, `(uint64(v))`)
			p.P(`       }`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(l)`)
			p.P(`    }`)

		} else {
			p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(uint64(x.`, fieldName, `))}`)
		}
	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, sizeVarintSlice, `(x.`, fieldName, `))`)
			p.P(`    }`)
		} else {
			p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(uint64(x.`, fieldName, `))}`)
		}
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, sizeZigZagSlice, `(x.`, fieldName, `))`)
			p.P(`    }`)
		} else {
			p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, ` + `, sizeVarint, `(`, encodeZigZag, `(int64(x.`, fieldName, `)))}`)
		}

	case protoreflect.FloatKind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed32Kind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `)*4)`)
			p.P(`    }`)
		} else {
			p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, `+4}`)
		}
	case protoreflect.DoubleKind,
		protoreflect.Fixed64Kind,
		protoreflect.Sfixed64Kind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `)*8)`)
			p.P(`    }`)
		} else {
			p.P(`    if x.`, fieldName, `!=0 {n += `, keysize, `+8}`)
		}
	case protoreflect.BoolKind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(len(x.`, fieldName, `))`)
			p.P(`    }`)
		} else {
			p.P(`    if x.`, fieldName, ` {n += `, keysize+1, `}`)
		}
	case protoreflect.StringKind, protoreflect.BytesKind:
		if repeated {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			p.P(`        for _, s := range x.`, fieldName, ` {`)
			p.P(`    		n += `, keysize, ` + `, sizeBytes, `(len(s))`)
			p.P(`        }`)
			p.P(`    }`)
		} else {
			p.P(`    l = len(x.`, fieldName, `)`)
			p.P(`    if l > 0 {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(l)`)
			p.P(`    }`)
		}
	case protoreflect.MessageKind:
		if field.Desc.IsMap() {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
			// p.P(`    n += `, keysize)
			p.sizerMap(field)
			p.P(`    }`)
		} else if repeated {
			p.P(`        for _, e := range x.`, fieldName, ` {`)
			p.P(`        	n += `, keysize, ` + `, sizeBytes, `(`, size, `(e))`)
			p.P(`        }`)

		} else {
			p.P(`    if x.`, fieldName, `!=nil {`)
			p.P(`    	n += `, keysize, ` + `, sizeBytes, `(`, size, `(x.`, fieldName, `))`)
			p.P(`    }`)
		}
	}
}

func (p *sizer) sizerMap(field *protogen.Field) {
	keysize := protowire.SizeTag(field.Desc.Number())
	key := field.Desc.MapKey()
	value := field.Desc.MapValue()

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

}

var (
	keySizer = [][]interface{}{
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
)

var (
	valueSizer = [][]interface{}{
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
)
