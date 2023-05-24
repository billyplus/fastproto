package main

import (
	"flag"
	"fmt"

	"github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal"
	_ "github.com/billyplus/fastproto/cmd/protoc-gen-go-fast/internal/plugins"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion = flag.Bool("version", false, "print the version and exit")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-fast %v\n", internal.Version)
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			internal.GenerateFile(gen, f)
		}
		return nil
	})
}
