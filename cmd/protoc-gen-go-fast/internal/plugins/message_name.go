package plugin

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func init() {
	// internal.RegisterPlugin(NewMessageName())
}

type messageName struct {
	*protogen.GeneratedFile
}

// func NewMessageName() internal.Plugin {
// 	return &messageName{}
// }

func (p *messageName) Name() string {
	return "namer"
}

func (p *messageName) Init() {
}

func (p *messageName) GenerateMessage(g *protogen.GeneratedFile, f *protogen.File, m *protogen.Message) {
	for _, msg := range f.Messages {
		g.P(fmt.Sprintf(`func (m *%s)MessageName() string { return "%s"}`, msg.GoIdent.GoName, msg.GoIdent.GoName))
		g.P()
	}
}
