package neatjson

import (
	"testing"

	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := sure_cls_gen.NewGenParam(runpath.PARENT.Path())
	param.SetSubClassNamePartWords("_")
	param.SetSubClassNameStyleEnum(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE)
	param.SetSureEnum(sure.MUST, sure.SOFT, sure.OMIT)

	cfg := &sure_cls_gen.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions(),
		SrcPath:       runtestpath.SrcPath(t),
	}
	sure_cls_gen.Gen(cfg, Neatjson{})
}
