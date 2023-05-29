package plugin

import (
	"fmt"

	"github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal"
	"github.com/billyplus/fastproto/goimport"
	"github.com/billyplus/fastproto/options"
	"github.com/billyplus/fastproto/protohelper"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	consumeFixed32 = goimport.ProtoWirePackage.Ident("ConsumeFixed32")
	consumeFixed64 = goimport.ProtoWirePackage.Ident("ConsumeFixed64")
	consumeVarint  = goimport.ProtoWirePackage.Ident("ConsumeVarint")
	consumeBytes   = goimport.ProtoWirePackage.Ident("ConsumeBytes")
	consumeMessage = goimport.FastProtoHelperPackage.Ident("ConsumeMessage")
	consumeTag     = goimport.ProtoWirePackage.Ident("ConsumeTag")

	float32frombits = goimport.MathPackage.Ident("Float32frombits")
	float64frombits = goimport.MathPackage.Ident("Float64frombits")
	decodeZigZag    = goimport.ProtoWirePackage.Ident("DecodeZigZag")

	parseError     = goimport.ProtoWirePackage.Ident("ParseError")
	calcListLength = goimport.FastProtoHelperPackage.Ident("CalcListLength")
	skip           = goimport.FastProtoHelperPackage.Ident("Skip")
)

func init() {
	internal.RegisterPlugin(newDecoder())
}

type decoder struct {
	*protogen.GeneratedFile
	plugin *protogen.Plugin
}

func newDecoder() internal.Plugin {
	return &decoder{}
}

func (p *decoder) Name() string {
	return "decoder"
}

func (p *decoder) Init() {
}

func (p *decoder) GenerateMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *protogen.File, idx int, m *protogen.Message) {
	if !options.IsUnmarshaler(f.Desc, m.Desc) {
		return
	}

	p.GeneratedFile = g
	p.plugin = gen
	p.generateXxxReset(f, idx, m)
	// p.generateFillMessageInfo(f, idx, m)
	p.generateUnmarshaler(f, idx, m)
}

func (p *decoder) generateXxxReset(f *protogen.File, idx int, m *protogen.Message) {
	p.P(`func (x *`, m.GoIdent.GoName, `) XxxReset() {`)
	p.P("	*x = ", m.GoIdent.GoName, "{}")
	p.P(`}`)
	p.P(``)
}

// func (p *decoder) generateFillMessageInfo(f *protogen.File, idx int, m *protogen.Message) {
// 	p.P(`func (x *`, m.GoIdent.GoName, `) FillMessageInfo() {`)
// 	p.P(`	if x != nil{`)
// 	p.P(`		ms := `, fastproto.ProtoImplPackage.Ident("X.MessageStateOf"), "(", fastproto.ProtoImplPackage.Ident("Pointer"), "(x))")
// 	p.P(`		if ms.LoadMessageInfo() == nil {`)
// 	p.P(`			mi := &`, fastproto.MessageTypesVarName(f), `[`, idx, `]`)
// 	p.P(`			ms.StoreMessageInfo(mi)`)
// 	p.P(`		}`)
// 	p.P(`	}`)
// 	p.P(`}`)
// 	p.P(``)
// }

func (p *decoder) generateUnmarshaler(f *protogen.File, idx int, m *protogen.Message) {
	p.P(`func (x *`, m.GoIdent.GoName, `) Unmarshal(data []byte) error {`)

	if len(m.Fields) > 0 {
		p.P(`	for len(data) > 0 {`)
		p.P(`		num, wireType, n := `, consumeTag, `(data)`)
		p.P(`		if n < 0 {`)
		p.P(`			return `, parseError, "(n)")
		p.P(`		}`)
		// p.P(`		if wireType == 4 {`)
		// p.P(`			return fmt.Errorf("proto: `, m.Desc.Name(), `: wiretype end group for non-group")`)
		// p.P(`		}`)
		p.P(`		if num <= 0 {`)
		p.P(`			return fmt.Errorf("proto: `, m.Desc.Name(), `: illegal tag %d (wire type %d)", num, wireType)`)
		p.P(`	   	}`)
		p.P(`   	data = data[n:]`)

		p.P(`		switch num {`)

		for _, field := range m.Fields {
			p.generateField(f, field)
		}

		p.P(`		default:`)
		p.P(`		    _, n = `, consumeBytes, "(data)")
		p.P(`		    if n < 0 {`)
		p.P(`				return `, parseError, "(n)")
		p.P(`			}`)
		p.P(`   		x.unknownFields = append(x.unknownFields, data[:n]...)`)
		p.P(`   		data = data[n:]`)
		p.P(`		}`)
		p.P(`   }`)
	} else {
		// p.P(`		_, n = `, fastproto.ProtoWirePackage.Ident("ConsumeBytes"), "(data)")
		p.P(`		if len(data) > 0 {`)
		p.P(`   		x.unknownFields = append(x.unknownFields, data[:]...)`)
		p.P(`		}`)
	}

	p.P(`   return nil`)
	p.P(`}`)

	p.P()
}

func (p *decoder) generateField(f *protogen.File, field *protogen.Field) {
	fieldNumber := field.Desc.Number()
	p.P(fmt.Sprintf(`		case %d:`, fieldNumber))
	kind := field.Desc.Kind()
	wireType := protohelper.KindToType(kind)
	dec := valueDecoder[kind]
	switch kind {
	case protoreflect.StringKind, protoreflect.BytesKind:
		p.genStringField(f, wireType, field, dec)
	case protoreflect.MessageKind:
		if field.Desc.IsMap() {
			p.genMap(f, 2, field)
		} else {
			p.genMessage(f, wireType, field, dec)
		}
	default:
		p.genField(f, wireType, field, dec)
	}
}

func (p *decoder) genField(f *protogen.File, wireType protowire.Type, field *protogen.Field, method protogen.GoIdent) {

	if field.Desc.IsList() {
		p.genList(f, wireType, field, method)
	} else {
		kind := field.Desc.Kind()
		p.P(`		if wireType != `, wireType, ` {`)
		p.P(`			return fmt.Errorf("proto: wrong wireType = %d for field `, field.GoName, `", wireType)`)
		p.P(`		}`)
		p.P(`		v, n := `, method, `(data)`)
		p.P(`		if n < 0 { return `, parseError, `(n)}`)
		p.P(`   	data = data[n:]`)
		switch kind {
		case protoreflect.BoolKind:
			p.P(`		x.`, field.GoName, ` = bool(v!=0)`)
		case protoreflect.FloatKind:
			p.P(`   	x.`, field.GoName, " = ", float32frombits, "(v)")
		case protoreflect.DoubleKind:
			p.P(`   	x.`, field.GoName, " = ", float64frombits, "(v)")
		case protoreflect.Sint32Kind,
			protoreflect.Sint64Kind:
			p.P(`   	x.`, field.GoName, " = ", protohelper.GoTypeOfField(field.Desc), "(", decodeZigZag, "(v))")

		default:
			p.P(`   	x.`, field.GoName, " = ", protohelper.GoTypeOfField(field.Desc), "(v)")
		}
	}
}

func (p *decoder) genStringField(f *protogen.File, wireType protowire.Type, field *protogen.Field, method protogen.GoIdent) {
	p.P(`		if wireType != `, wireType, ` {`)
	p.P(`			return fmt.Errorf("proto: wrong wireType = %d for field `, field.GoName, `", wireType)`)
	p.P(`		}`)
	p.P(`		v, n := `, method, `(data)`)
	p.P(`		if n < 0 { return `, parseError, `(n)}`)
	p.P(`   	data = data[n:]`)
	if field.Desc.IsList() {
		if field.Desc.Kind() == protoreflect.StringKind {
			p.P(`		x.`, field.GoName, ` = append(x.`, field.GoName, `, string(v))`)
		} else {
			p.P(`		x.`, field.GoName, " = append(x.", field.GoName, "[:], append([]byte{}, v...))")
		}
	} else {
		if field.Desc.Kind() == protoreflect.StringKind {
			p.P(`		x.`, field.GoName, ` = string(v)`)
		} else {
			p.P(`		x.`, field.GoName, " = append(x.", field.GoName, "[:0], v...)")
		}
	}
}

func (p *decoder) genList(f *protogen.File, wireType protowire.Type, field *protogen.Field, method protogen.GoIdent) {
	kind := field.Desc.Kind()
	p.P(`		if wireType == `, wireType, ` {`)
	if kind == protoreflect.BoolKind {
		p.P(`			v, n := `, method, `(data)`)
	} else {
		p.P(`			v, n := `, method, `(data)`)
	}
	p.P(`			if n < 0 { return `, parseError, `(n)}`)
	p.P(`   		data = data[n:]`)
	switch kind {
	case protoreflect.BoolKind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, `, v != 0)`)
	case protoreflect.FloatKind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", float32frombits, `(v))`)
	case protoreflect.DoubleKind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", float64frombits, `(v))`)
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", protohelper.GoTypeOfField(field.Desc), "(", decodeZigZag, `(v)))`)
	default:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", protohelper.GoTypeOfField(field.Desc), `(v))`)
	}
	p.P(`		} else if wireType == 2 {`)
	p.P(`			msglen, n := `, calcListLength, `(data)`)
	p.P(`			if n < 0 { return `, parseError, `(n)}`)
	p.P(`   		data = data[n:]`)
	switch kind {
	case protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind,
		protoreflect.EnumKind:
		p.P(`		    elementCount := 0`)
		p.P(`		    for _, i := range data[:msglen] {`)
		p.P(`		        if i < 128 {`)
		p.P(`		            elementCount++`)
		p.P(`		         }`)
		p.P(`		     }`)
	case protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind, protoreflect.FloatKind:
		p.P(`		    elementCount := msglen / 4`)
	case protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind, protoreflect.DoubleKind:
		p.P(`		    elementCount := msglen / 8`)
	case protoreflect.BoolKind:
		p.P(`		    elementCount := msglen`)
	}
	p.P(`		    if elementCount != 0 && len(x.`, field.GoName, `) == 0 {`)
	p.P(`		        x.`, field.GoName, ` = make([]`, protohelper.GoTypeOfField(field.Desc), `, 0, elementCount)`)
	p.P(`		    }`)
	p.P(`		    for elementCount > 0 {`)
	p.P(`				v, n := `, method, `(data)`)
	p.P(`				if n < 0 { return `, parseError, `(n)}`)
	p.P(`   			data = data[n:]`)
	p.P(`   			elementCount--`)
	switch kind {
	case protoreflect.BoolKind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, `, v != 0)`)
	case protoreflect.FloatKind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", float32frombits, `(v))`)
	case protoreflect.DoubleKind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", float64frombits, `(v))`)
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", protohelper.GoTypeOfField(field.Desc), "(", decodeZigZag, `(v)))`)
	default:
		p.P(`		        x.`, field.GoName, ` = append(x.`, field.GoName, ",", protohelper.GoTypeOfField(field.Desc), `(v))`)
	}
	p.P(`		    }`)
	p.P(`		} else {`)
	p.P(`			return fmt.Errorf("proto: wrong wireType = %d for field `, field.GoName, `", wireType)`)
	p.P(`		}`)
}

func (p *decoder) genMessage(f *protogen.File, wireType protowire.Type, field *protogen.Field, method protogen.GoIdent) {
	p.P(`		if wireType != `, wireType, ` {`)
	p.P(`			return fmt.Errorf("proto: wrong wireType = %d for field `, field.GoName, `", wireType)`)
	p.P(`		}`)

	if field.Desc.IsList() {
		p.P(`		v := &`, p.QualifiedGoIdent(field.Message.GoIdent), "{}")
		p.P(`		if n, err := `, method, `(data, v); err != nil {`)
		p.P(`		    return err`)
		p.P(`		} else {`)
		p.P(`			data = data[n:]`)
		p.P(`		}`)

		p.P(`		x.`, field.GoName, ` = append(x.`, field.GoName, `, v)`)
	} else {
		p.P(`		if x.`, field.GoName, ` == nil {`)
		p.P(`		    x.`, field.GoName, ` = &`, p.QualifiedGoIdent(field.Message.GoIdent), "{}")
		p.P(`		}`)
		p.P(`		if n, err := `, method, `(data, x.`, field.GoName, `); err != nil {`)
		p.P(`		    return err`)
		p.P(`		} else {`)
		p.P(`			data = data[n:]`)
		p.P(`		}`)
		// p.P(`		msglen, n := `, protobuf.FastProtoPackage.Ident("CalcListLength"), `(data)`)
		// p.P(`		if n < 0 { return `, protobuf.ProtoWirePackage.Ident("ParseError"), `(n)}`)
		// p.P(`		data = data[n:]`)
		// p.P(`		if err := x.`, field.GoName, `.Unmarshal(data[:msglen]); err != nil {`)
		// p.P(`		    return err`)
		// p.P(`		} else {`)
		// p.P(`			data = data[msglen:]`)
		// p.P(`		}`)
	}
}

func (p *decoder) genMap(f *protogen.File, wireType uint64, field *protogen.Field) {
	key := field.Desc.MapKey()
	value := field.Desc.MapValue()

	keyKind := protohelper.GoTypeOfField(key)
	valKind := protohelper.GoTypeOfField(value)

	p.P(`		if wireType != `, wireType, ` {`)
	p.P(`			return fmt.Errorf("proto: wrong wireType = %d for field `, field.GoName, `", wireType)`)
	p.P(`		}`)
	p.P(`		msglen, n := `, calcListLength, `(data)`)
	p.P(`		if n < 0 { return `, parseError, `(n)}`)
	p.P(`   	data = data[n:]`)
	p.P(`		if x.`, field.GoName, ` == nil {`)
	p.P(`		    x.`, field.GoName, ` = make(map[`, keyKind, `]`, valKind, `)`)
	p.P(`		}`)
	p.P(`		var mapkey `, keyKind)
	p.P(`		var mapvalue `, valKind)
	p.P(`		for msglen > 0 {`)
	// p.P(`		    entryPreIndex := r.idx`)
	p.P(`			subNum, subWireType, n := `, consumeTag, `(data)`)
	p.P(`			if n < 0 { return `, parseError, `(n)}`)
	// p.P(`   		data = data[n:]`)
	// p.P(`		    msglen -= n`)
	p.P(`   		data, msglen = data[n:], msglen-n`)
	p.P(`		    if subNum == 1 {`)
	p.generateEntry(f, "mapkey", key)
	p.P(`		    } else if subNum == 2 {`)
	p.generateEntry(f, "mapvalue", value)
	p.P(`		    } else {`)
	p.P(`		        if skippy, err := `, skip, `(data); err!=nil{`)
	p.P(`		            return err`)
	p.P(`		        } else {`)
	p.P(`		        	data = data[skippy:]`)
	p.P(`		        	msglen -= skippy`)
	p.P(`		        }`)
	p.P(`		    }`)
	p.P(`		}`)
	p.P(`		x.`, field.GoName, `[mapkey] = mapvalue`)
}

// func (p *decoder) genMapKeyVal(vname string, field protoreflect.FieldDescriptor) {
// 	kind := field.Kind()
// 	dec := valueDecoder[kind]
// 	wireType := fastproto.KindToType(kind)

// 	p.P(`				if subWireType != `, wireType, ` {`)
// 	p.P(`					return fmt.Errorf("proto: wrong wireType = %d for field `, field.Name(), `", subWireType)`)
// 	p.P(`				}`)

// 	switch field.Kind() {
// 	case protoreflect.MessageKind:
// 		p.P(`		        `, vname, ` = &`, field.Message().Name(), "{}")
// 		tmp := append([]interface{}{`				n, err := `}, dec...)
// 		tmp = append(tmp, "(data, ", vname, ")")
// 		p.P(tmp...)
// 		// p.P(`				n, err := `, dec, "(data, ", vname, ")")
// 		p.P(`		  		if err != nil { return err }`)
// 	case protoreflect.Int32Kind,
// 		protoreflect.Int64Kind,
// 		protoreflect.Uint32Kind,
// 		protoreflect.Uint64Kind,
// 		protoreflect.BoolKind,
// 		protoreflect.EnumKind:
// 	default:
// 		p.P(`				v, n := `, consumeVarint, `(data)`)
// 	}
// 	p.P(`				if n < 0 { return `, fastproto.ProtoWirePackage.Ident("ParseError"), `(n)}`)
// 	p.P(`   			data = data[n:]`)
// 	p.P(`		        msglen -= n`)
// 	switch field.Kind() {
// 	case protoreflect.MessageKind:
// 	case protoreflect.StringKind:
// 		p.P(`			`, vname, ` = string(v)`)
// 	case protoreflect.BoolKind:
// 		p.P(`			`, vname, ` = bool(v!=0)`)
// 	default:
// 		p.P(`			`, vname, " = ", fastproto.GoTypeOfField(field), `(v)`)
// 	}
// }

func (p *decoder) generateEntry(f *protogen.File, fieldName string, field protoreflect.FieldDescriptor) {
	kind := field.Kind()
	dec := valueDecoder[kind]
	wireType := protohelper.KindToType(kind)
	p.P(`				if subWireType != `, wireType, ` {`)
	p.P(`					return fmt.Errorf("proto: wrong wireType = %d for field `, field.Name(), `", subWireType)`)
	p.P(`				}`)

	switch kind {
	case protoreflect.MessageKind:
		p.P(`		        `, fieldName, ` = &`, field.Message().Name(), "{}")
		p.P(`				n, err := `, dec, "(data, ", fieldName, ")")
		p.P(`		  		if err != nil { return err }`)
	default:
		p.P(`		v, n := `, dec, `(data)`)
		p.P(`		if n < 0 { return `, parseError, `(n)}`)
	}
	p.P(`   	data, msglen = data[n:], msglen-n`)
	// p.P(`   	msglen -= n`)

	switch kind {
	case protoreflect.BoolKind:
		p.P(`		`, fieldName, ` = bool(v!=0)`)
	case protoreflect.FloatKind:
		p.P(`   	`, fieldName, " = ", float32frombits, "(v)")
	case protoreflect.DoubleKind:
		p.P(`   	`, fieldName, " = ", float64frombits, "(v)")
	case protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		p.P(`   	`, fieldName, " = ", protohelper.GoTypeOfField(field), "(", decodeZigZag, "(v))")
	case protoreflect.MessageKind:
		// skip
	default:
		p.P(`   	`, fieldName, " = ", protohelper.GoTypeOfField(field), "(v)")
	}
}

var (
	valueDecoder = []protogen.GoIdent{
		protoreflect.Int32Kind:    consumeVarint,
		protoreflect.Int64Kind:    consumeVarint,
		protoreflect.FloatKind:    consumeFixed32,
		protoreflect.DoubleKind:   consumeFixed64,
		protoreflect.Uint32Kind:   consumeVarint,
		protoreflect.Uint64Kind:   consumeVarint,
		protoreflect.Sint32Kind:   consumeVarint,
		protoreflect.Sint64Kind:   consumeVarint,
		protoreflect.Fixed32Kind:  consumeFixed32,
		protoreflect.Fixed64Kind:  consumeFixed64,
		protoreflect.Sfixed32Kind: consumeFixed32,
		protoreflect.Sfixed64Kind: consumeFixed64,
		protoreflect.BoolKind:     consumeVarint,
		protoreflect.EnumKind:     consumeVarint,
		protoreflect.StringKind:   consumeBytes,
		protoreflect.BytesKind:    consumeBytes,
		protoreflect.MessageKind:  consumeMessage,
	}
)
