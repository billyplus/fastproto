package plugin

// import (
// 	"ddz/tools/gen/internal"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"google.golang.org/protobuf/compiler/protogen"
// )

// func init() {
// 	internal.RegisterPlugin(NewRoute())
// }

// type route struct {
// 	*protogen.GeneratedFile
// }

// func NewRoute() internal.Plugin {
// 	return &route{}
// }

// func (p *route) Name() string {
// 	return "router"
// }

// func (p *route) Init() {
// }

// func (p *route) Generate(g *protogen.GeneratedFile, f *protogen.File) {
// 	genEnum(g, f, f.Enums)
// }

// func genEnum(g *protogen.GeneratedFile, f *protogen.File, e []*protogen.Enum) {
// 	for _, enum := range e {
// 		enumName := enum.GoIdent.GoName
// 		isServer := strings.HasPrefix(enumName, "SR")
// 		isClient := strings.HasPrefix(enumName, "CR")
// 		if (isClient || isServer) && len(enumName) > 2 {
// 			num, err := strconv.Atoi(enumName[2:])
// 			if err != nil {
// 				fmt.Printf("Failed to gen Enum[%s]: %v\n", enumName, err)
// 			}

// 			for _, val := range enum.Values {
// 				// log.Debug().Interface("comments", val.Comments).Msg("enum")
// 				if strings.HasPrefix(val.GoIdent.GoName, "None") {
// 					continue
// 				}
// 				msgName := strings.Split(val.GoIdent.GoName, "_")[1]
// 				route := uint32(num)<<10 | uint32(val.Desc.Number())
// 				g.P(fmt.Sprintf(`func (m *%s)Route() uint32 { return %d}`, msgName, route))
// 				g.P()
// 			}
// 		}
// 	}
// }
