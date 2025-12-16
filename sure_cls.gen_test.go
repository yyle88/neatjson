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

// TestGen generates handling adapters (Must, Soft, Omit) providing Neatjson convenience methods
// Uses sure framework to auto-generate three modes with consistent naming
// Regenerates sure_cls.gen.go when core API changes, keeping adapters aligned
//
// TestGen 为 Neatjson 生成处理适配（Must、Soft、Omit），提供便捷方法
// 使用 sure 框架自动生成三种模式，采用一致的命名
// 当核心 API 变化时重新生成 sure_cls.gen.go，保持适配对齐
func TestGen(t *testing.T) {
	// Configure class generation using underscore as junction and suffix naming
	// 配置类生成，使用下划线作为连接点和后缀命名
	options := sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path()).
		WithNewClassNameParts("_").
		WithNamingPatternType(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE).
		MoreErrorHandlingModes(sure.MUST, sure.SOFT, sure.OMIT)

	// Setup generation config with package info and output path
	// 设置生成配置，包括包信息和输出路径
	config := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: options,
		PackageName:     syntaxgo.CurrentPackageName(),
		ImportOptions:   syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:      runtestpath.SrcPath(t),
	}

	// Generate wrap classes supplying each handling mode
	// 为每种处理模式生成包装类
	sure_cls_gen.GenerateClasses(config, Neatjson{}, Compactjson{})
}
