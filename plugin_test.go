package protocplugin_test

import (
	"fmt"
	protocplugin "github.com/Jumpaku/protoc-plugin-lib"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
)

func ExampleRun() {
	handle := func(
		req *pluginpb.CodeGeneratorRequest,
		files map[string]*protocplugin.File,
	) ([]*protocplugin.GeneratedFile, error) {
		out := []*protocplugin.GeneratedFile{}
		for _, f := range req.FileToGenerate {
			in := files[f]
			out = append(out, &protocplugin.GeneratedFile{
				Name:    string(in.Desc.Path()) + ".dump",
				Content: string(in.Desc.FullName()),
			})
		}
		return out, nil
	}
	err := protocplugin.Run(os.Stdin, os.Stdout, handle)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
