package sures

import (
	"testing"

	"github.com/yyle88/neatjson"
	"github.com/yyle88/runpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGenSoft(t *testing.T) {
	param := cls_stub_gen.NewParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Soft()")

	sourceDIRPath := runpath.PARENT.Join("./../../../neatjson")
	t.Log(sourceDIRPath)

	targetSrcPath := runpath.PARENT.Join("./../../neatjsons/neatjsons.go")
	t.Log(targetSrcPath)

	cfg := &cls_stub_gen.Config{
		SrcRoot:       sourceDIRPath,
		TargetPkgName: "neatjsons",
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		TargetSrcPath: targetSrcPath,
		CanCreateFile: false,
	}
	cls_stub_gen.Gen(cfg, param)
}

func TestGenMust(t *testing.T) {
	param := cls_stub_gen.NewParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Must()")

	sourceDIRPath := runpath.PARENT.Join("./../../../neatjson")
	t.Log(sourceDIRPath)

	targetSrcPath := runpath.PARENT.Join("./../../neatjsonm/neatjsonm.go")
	t.Log(targetSrcPath)

	cfg := &cls_stub_gen.Config{
		SrcRoot:       sourceDIRPath,
		TargetPkgName: "neatjsonm",
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		TargetSrcPath: targetSrcPath,
		CanCreateFile: false,
	}
	cls_stub_gen.Gen(cfg, param)
}

func TestGenOmit(t *testing.T) {
	param := cls_stub_gen.NewParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Omit()")

	sourceDIRPath := runpath.PARENT.Join("./../../../neatjson")
	t.Log(sourceDIRPath)

	targetSrcPath := runpath.PARENT.Join("./../../neatjsono/neatjsono.go")
	t.Log(targetSrcPath)

	cfg := &cls_stub_gen.Config{
		SrcRoot:       sourceDIRPath,
		TargetPkgName: "neatjsono",
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		TargetSrcPath: targetSrcPath,
		CanCreateFile: false,
	}
	cls_stub_gen.Gen(cfg, param)
}
