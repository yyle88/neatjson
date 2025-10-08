// Package sures contains code generation tests creating neatjson convenience packages
// Auto generates wrap packages (neatjsons, neatjsonm, neatjsono) with different handling modes
// Uses sure framework stub generation creating typed wrappers with consistent API patterns
// These tests regenerate the convenience packages when the core API changes
//
// sures 包包含创建 neatjson 便捷包的代码生成测试
// 自动生成不同处理模式的包装包（neatjsons、neatjsonm、neatjsono）
// 使用 sure 框架的 stub 生成，创建具有一致 API 模式的类型化包装
// 当核心 API 发生变化时，这些测试会重新生成便捷包
package sures

import (
	"testing"

	"github.com/yyle88/neatjson"
	"github.com/yyle88/runpath"
	"github.com/yyle88/sure/cls_stub_gen"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

// TestGenSoft generates the neatjsons package with Soft handling mode
// Creates wrap functions returning zero-value results on errors without panicking
// Updates neatjsons/neatjsons.go with auto-generated code allowing convenient usage
//
// TestGenSoft 生成带有 Soft 处理模式的 neatjsons 包
// 创建在错误时返回零值结果而不 panic 的包装函数
// 更新 neatjsons/neatjsons.go，生成便捷使用的自动代码
func TestGenSoft(t *testing.T) {
	stubParam := cls_stub_gen.NewStubParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Soft()")

	sourceRootPath := runpath.PARENT.UpTo(3, "neatjson")
	t.Log(sourceRootPath)

	outputPath := runpath.PARENT.UpTo(2, "neatjsons/neatjsons.go")
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

// TestGenMust generates the neatjsonm package with Must handling mode
// Creates wrap functions that panic on errors providing fast-exit action
// Updates neatjsonm/neatjsonm.go with auto-generated code suited to main operations
//
// TestGenMust 生成带有 Must 处理模式的 neatjsonm 包
// 创建在错误时 panic 的包装函数，提供快速退出动作
// 更新 neatjsonm/neatjsonm.go，生成用于主要操作的自动代码
func TestGenMust(t *testing.T) {
	stubParam := cls_stub_gen.NewStubParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Must()")

	sourceRootPath := runpath.PARENT.UpTo(3, "neatjson")
	t.Log(sourceRootPath)

	outputPath := runpath.PARENT.UpTo(2, "neatjsonm/neatjsonm.go")
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

// TestGenOmit generates the neatjsono package with Omit handling mode
// Creates wrap functions that skip errors without logging, without panicking
// Updates neatjsono/neatjsono.go with auto-generated code suited to side logging
//
// TestGenOmit 生成带有 Omit 处理模式的 neatjsono 包
// 创建跳过错误的包装函数，不记录日志也不 panic
// 更新 neatjsono/neatjsono.go，生成用于辅助日志的自动代码
func TestGenOmit(t *testing.T) {
	stubParam := cls_stub_gen.NewStubParam(&neatjson.Neatjson_Soft{}, "neatjson.TAB.Omit()")

	sourceRootPath := runpath.PARENT.UpTo(3, "neatjson")
	t.Log(sourceRootPath)

	outputPath := runpath.PARENT.UpTo(2, "neatjsono/neatjsono.go")
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
