package sures

import (
	"testing"

	"github.com/yyle88/neatjson"
	"github.com/yyle88/runpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGenSoft(t *testing.T) {
	stubParam := cls_stub_gen.NewStubParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Soft()")

	sourceRootPath := runpath.PARENT.Join("./../../../neatjson")
	t.Log(sourceRootPath)

	outputPath := runpath.PARENT.Join("./../../neatjsons/neatjsons.go")
	t.Log(outputPath)

	config := &cls_stub_gen.StubGenConfig{
		SourceRootPath:    sourceRootPath,
		TargetPackageName: "neatjsons",
		ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:        outputPath,
		AllowFileCreation: false,
	}
	cls_stub_gen.GenerateStubs(config, stubParam)
}

func TestGenMust(t *testing.T) {
	stubParam := cls_stub_gen.NewStubParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Must()")

	sourceRootPath := runpath.PARENT.Join("./../../../neatjson")
	t.Log(sourceRootPath)

	outputPath := runpath.PARENT.Join("./../../neatjsonm/neatjsonm.go")
	t.Log(outputPath)

	config := &cls_stub_gen.StubGenConfig{
		SourceRootPath:    sourceRootPath,
		TargetPackageName: "neatjsonm",
		ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:        outputPath,
		AllowFileCreation: false,
	}
	cls_stub_gen.GenerateStubs(config, stubParam)
}

func TestGenOmit(t *testing.T) {
	stubParam := cls_stub_gen.NewStubParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Omit()")

	sourceRootPath := runpath.PARENT.Join("./../../../neatjson")
	t.Log(sourceRootPath)

	outputPath := runpath.PARENT.Join("./../../neatjsono/neatjsono.go")
	t.Log(outputPath)

	config := &cls_stub_gen.StubGenConfig{
		SourceRootPath:    sourceRootPath,
		TargetPackageName: "neatjsono",
		ImportOptions:     syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:        outputPath,
		AllowFileCreation: false,
	}
	cls_stub_gen.GenerateStubs(config, stubParam)
}
