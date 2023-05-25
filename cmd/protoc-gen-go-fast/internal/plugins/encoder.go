package plugin

import (
	"github.com/billyplus/fastproto"
	"github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	encodeZigZag        = fastproto.ProtoWirePackage.Ident("EncodeZigZag")
	encodeTag           = fastproto.ProtoWirePackage.Ident("EncodeTag")
	appendVarint        = fastproto.ProtoWirePackage.Ident("AppendVarint")
	appendTag           = fastproto.ProtoWirePackage.Ident("AppendTag")
	appendString        = fastproto.ProtoWirePackage.Ident("AppendString")
	appendBytes         = fastproto.ProtoWirePackage.Ident("AppendBytes")
	appendFixed32       = fastproto.ProtoWirePackage.Ident("AppendFixed32")
	appendFixed64       = fastproto.ProtoWirePackage.Ident("AppendFixed64")
	bool2Int            = fastproto.FastProtoPackage.Ident("Bool2Int")
	appendToSizedBuffer = fastproto.FastProtoPackage.Ident("AppendToSizedBuffer")
)

func init() {
	internal.RegisterPlugin(newEncoder())
}

type encoder struct {
	*protogen.GeneratedFile
}

func newEncoder() internal.Plugin {
	return &encoder{}
}

func (p *encoder) Name() string {
	return "encoder"
}

func (p *encoder) Init() {
}

func (p *encoder) GenerateMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, idx int, m *protogen.Message) {
	p.GeneratedFile = g

	g.P(`func (x *`, m.GoIdent.GoName, `) MarshalTo(data []byte) (n int, err error) {`)
	if len(m.Fields) == 0 {
		g.P(`	return 0, nil`)
	} else {
		p.P("    maxN := cap(data)")
		p.P("    data, err = x.AppendToSizedBuffer(data[:0])")
		p.P(`    if maxN < len(data) {`)
		p.P(`    	return 0, fmt.Errorf("Not enough space for message(`, m.GoIdent.GoName, `)")`)
		p.P(`    }`)
		g.P(`	return len(data), nil`)
	}
	g.P(`}`)
	g.P()
	g.P(`func (x *`, m.GoIdent.GoName, `) Marshal() ([]byte, error) {`)
	if len(m.Fields) == 0 {
		g.P(`	return []byte{}, nil`)
	} else {
		g.P(`	data := make([]byte, 0, `, size, `(x))`)
		g.P(`	data, err := x.AppendToSizedBuffer(data)`)
		g.P(`	if err != nil {`)
		g.P(`		return nil, err`)
		g.P(`	}`)
		g.P(`	return data, nil`)
	}
	g.P(`}`)
	g.P()
	p.genMarshalToSizedBuffer(f, m)
	g.P()
}

func (p *encoder) genMarshalToSizedBuffer(f *protogen.File, m *protogen.Message) {
	p.P(`func (x *`, m.GoIdent.GoName, `) AppendToSizedBuffer(data []byte) (ret []byte, err error) {`)
	if len(m.Fields) > 0 {
		for _, field := range m.Fields {
			p.generateField(f, field)
		}
	}
	p.P(`    return data, nil`)
	p.P("}")
}

func (p *encoder) generateField(f *protogen.File, field *protogen.Field) {
	repeated := field.Desc.IsList()
	fieldName := field.GoName

	kind := field.Desc.Kind()

	if repeated || kind == protoreflect.StringKind || kind == protoreflect.BytesKind {
		p.P(`    if len(x.`, fieldName, `) > 0 {`)
	} else if kind == protoreflect.MessageKind {
		if field.Desc.IsMap() {
			p.P(`    if len(x.`, fieldName, `) > 0 {`)
		} else {
			p.P(`    if x.`, fieldName, ` != nil {`)
		}
	} else if kind == protoreflect.BoolKind {
		p.P(`    if x.`, fieldName, ` {`)
	} else {
		p.P(`    if x.`, fieldName, ` != 0{`)
	}

	if repeated {
		p.generateSliceField(f, field)
	} else if field.Desc.IsMap() {
		p.generateMapField(f, field)
	} else {
		p.generateEntry(f, "x."+fieldName, field.Desc)
	}
	p.P(`    }`)
}

func (p *encoder) generateSliceField(f *protogen.File, field *protogen.Field) {
	kind := field.Desc.Kind()
	fieldName := field.GoName
	fieldNumber := field.Desc.Number()
	wireType := fastproto.KindToType(field.Desc.Kind())

	switch kind {
	case protoreflect.BytesKind,
		protoreflect.StringKind:
	case protoreflect.MessageKind:
	default:
		p.P(`        data = `, appendVarint, `(data, `, protowire.EncodeTag(fieldNumber, protowire.BytesType), `)`)
	}

	switch kind {
	case protoreflect.EnumKind:
		p.P(`        sz := 0`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	sz += `, sizeVarint, `(uint64(v))`)
		p.P(`        }`)
		p.P(`        data = `, appendVarint, `(data, uint64(sz))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendVarint, `(data, uint64(v))`)
		p.P(`        }`)
	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind:
		p.P(`        data = `, appendVarint, `(data, uint64(`, sizeVarintSlice, `(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendVarint, `(data, uint64(v))`)
		p.P(`        }`)
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`        data = `, appendVarint, `(data, uint64(`, sizeZigZagSlice, `(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendVarint, `(data, `, encodeZigZag, `(int64(v)))`)
		p.P(`        }`)
	case protoreflect.FloatKind:
		p.P(`        data = `, appendVarint, `(data, uint64(4 * len(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendFixed32, `(data, uint32(math.Float32bits(v)))`)
		p.P(`        }`)
	case protoreflect.DoubleKind:
		p.P(`        data = `, appendVarint, `(data, uint64(8 * len(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendFixed64, `(data, uint64(math.Float64bits(v)))`)
		p.P(`        }`)
	case protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		p.P(`        data = `, appendVarint, `(data, uint64(4 * len(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendFixed32, `(data, uint32(v))`)
		p.P(`        }`)
	case protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		p.P(`        data = `, appendVarint, `(data, uint64(8 * len(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendFixed64, `(data, uint64(v))`)
		p.P(`        }`)
	case protoreflect.BoolKind:
		p.P(`        data = `, appendVarint, `(data, uint64(len(x.`, fieldName, `)))`)
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = append(data, uint8(`, bool2Int, `(v)))`)
		p.P(`        }`)
	case protoreflect.BytesKind,
		protoreflect.StringKind:
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendVarint, `(data, `, protowire.EncodeTag(fieldNumber, wireType), `)`)
		p.P(`        	data = append(`, appendVarint, `(data, uint64(len(v))), v...)`)
		p.P(`        }`)
	case protoreflect.MessageKind:
		p.P(`        for _, v := range x.`, fieldName, ` {`)
		p.P(`        	data = `, appendVarint, `(data, `, protowire.EncodeTag(fieldNumber, wireType), `)`)
		p.P(`        	sz := `, size, `(v)`)
		p.P(`        	data = `, appendVarint, `(data, uint64(sz))`)
		p.P(`        	if sz > 0 {`)
		p.P(`        		data, err = `, appendToSizedBuffer, `(data, v)`)
		p.P(`        		if err != nil {`)
		p.P(`        			return nil, err`)
		p.P(`        		}`)
		p.P(`        	}`)
		p.P(`        }`)
	}
}

func (p *encoder) generateMapField(f *protogen.File, field *protogen.Field) {
	fieldNumber := field.Desc.Number()
	wireType := fastproto.KindToType(field.Desc.Kind())

	key := field.Desc.MapKey()
	value := field.Desc.MapValue()

	p.P(`        for k, v := range x.`, field.GoName, ` {`)
	p.P(`        	data = `, appendVarint, `(data, `, protowire.EncodeTag(fieldNumber, wireType), `)`)
	p.P(`            _, _ = k, v`)
	p.generateMapEntrySize(f, key, value)
	p.generateEntry(f, "k", key)
	p.generateEntry(f, "v", value)
	p.P(`        }`)

}

func (p *encoder) generateMapEntrySize(f *protogen.File, key, value protoreflect.FieldDescriptor) {
	tmp := append([]interface{}{`            l := `}, keySizer[key.Kind()]...)

	if value.Kind() == protoreflect.MessageKind {
		tmp = append(tmp, " + 1 + ", sizeBytes, `(`, size, `(v))`)
	} else {
		tmp = append(tmp, " + ")
		tmp = append(tmp, valueSizer[value.Kind()]...)
	}
	p.P(tmp...)
	p.P(`        data = `, appendVarint, `(data, uint64(l))`)

}

func (p *encoder) generateEntry(f *protogen.File, fieldName string, entryField protoreflect.FieldDescriptor) {
	kind := entryField.Kind()
	wireType := fastproto.KindToType(kind)

	fieldNumber := entryField.Number()
	p.P(`        data = `, appendVarint, `(data, `, protowire.EncodeTag(fieldNumber, wireType), `)`)

	switch kind {
	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind,
		protoreflect.EnumKind:
		p.P(`        data = `, appendVarint, `(data, uint64(`, fieldName, `))`)
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`        data = `, appendVarint, `(data, `, encodeZigZag, `(int64(`, fieldName, `)))`)
	case protoreflect.FloatKind:
		p.P(`        data = `, appendFixed32, `(data, uint32(math.Float32bits(`, fieldName, `)))`)
	case protoreflect.DoubleKind:
		p.P(`        data = `, appendFixed64, `(data, uint64(math.Float64bits(`, fieldName, `)))`)
	case protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		p.P(`        data = `, appendFixed32, `(data, uint32(`, fieldName, `))`)
	case protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		p.P(`        data = `, appendFixed64, `(data, uint64(`, fieldName, `))`)
	case protoreflect.BoolKind:
		p.P(`        data = append(data, uint8(`, bool2Int, `(`, fieldName, `)))`)
	case protoreflect.BytesKind,
		protoreflect.StringKind:
		p.P(`        data = `, appendVarint, `(data, uint64(len(`, fieldName, `)))`)
		p.P(`        data = append(data, `, fieldName, `...)`)
	case protoreflect.MessageKind:
		p.P(`        sz := uint64(`, size, `(`, fieldName, `))`)
		p.P(`        data = `, appendVarint, `(data, uint64(sz))`)
		p.P(`        if sz > 0 {`)
		p.P(`        	data, err = `, appendToSizedBuffer, `(data, `, fieldName, `)`)
		p.P(`        	if err != nil {`)
		p.P(`        		return nil, err`)
		p.P(`        	}`)
		p.P(`        }`)
	}
}
