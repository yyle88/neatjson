// Package neatjson provides elegant JSON formatting with flexible indentation
// Implements convenient syntax helpers converting Go structures to formatted JSON
// Supports TAB/spaces indentation and Must/Soft/Omit handling modes
// Designed with testing, debugging, and logging in mind
//
// neatjson 包提供优雅的 JSON 格式化，支持灵活的缩进控制
// 实现便捷的语法辅助，将 Go 结构体转换为格式化的 JSON
// 支持 TAB/空格缩进和 Must/Soft/Omit 处理模式
// 专为测试、调试和日志场景设计
package neatjson

import (
	"bytes"
	"encoding/json"

	"github.com/yyle88/erero"
)

// Neatjson implements syntax helpers to convert Go structs into JSON
// Provides neat indented JSON with customizable formatting options
// Supports struct marshaling and JSON reformatting
//
// Neatjson 实现将 Go 结构体转换为 JSON 的语法辅助
// 提供整齐的带缩进 JSON，具有可自定义格式化选项
// 支持结构体编组和 JSON 重新格式化
type Neatjson struct {
	prefix string // Prefix for JSON indentation // 结果 JSON 的缩进前缀
	indent string // Indentation characters // 缩进字符
	active bool   // Enable line breaks when converting to JSON // 是否使用缩进的转换函数
}

// NewNeatjson creates Neatjson with specified prefix and indent
// Returns instance with line breaks enabled as default setting
//
// NewNeatjson 根据指定的前缀和缩进字符创建 Neatjson 实例
// 返回默认启用换行的实例
func NewNeatjson(prefix string, indent string) *Neatjson {
	return &Neatjson{
		prefix: prefix,
		indent: indent,
		active: true, // Default to using line breaks because the main purpose of this package is to create indented JSON // 默认启用换行，因为这个包的主要目的就是生成带缩进的 JSON
	}
}

// notNeatjson returns Neatjson with line breaks disabled
// Used inside package code to create NON and NOI constants
//
// notNeatjson 返回禁用换行的 Neatjson 实例
// 用于包内代码创建 NON 和 NOI 常量
func notNeatjson() *Neatjson {
	return NewNeatjson("", "").withActive(false)
}

// Bytes converts struct to JSON bytes with indentation when configured
// Like json.Marshal but supports custom formatting
// Auto selects MarshalIndent vs Marshal based on settings
//
// Bytes 将结构体转换为 JSON 字节数组，配置时支持缩进
// 类似 json.Marshal，但支持自定义格式化
// 根据配置自动选择 MarshalIndent 与 Marshal
func (N *Neatjson) Bytes(v interface{}) ([]byte, error) {
	if N.needIndent() {
		data, err := json.MarshalIndent(v, N.prefix, N.indent)
		if err != nil {
			return nil, erero.Wro(err)
		}
		return data, nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return data, nil
}

// needIndent checks if indentation is needed
// Returns true when active flag is set, prefix/indent configured
//
// needIndent 检查是否需要缩进
// 当 active 标志开启、前缀/缩进已配置时返回 true
func (N *Neatjson) needIndent() bool {
	return N.active || N.prefix != "" || N.indent != ""
}

// withActive sets if line breaks are enabled
// Returns instance allowing method chaining
//
// withActive 设置是否启用换行
// 返回实例以支持方法链式调用
func (N *Neatjson) withActive(active bool) *Neatjson {
	N.active = active
	return N
}

// Sjson converts struct to JSON string with desired formatting
// Wraps Bytes method and converts result to string
//
// Sjson 将结构体转换为格式化的 JSON 字符串
// 封装 Bytes 方法并将结果转换为字符串
func (N *Neatjson) Sjson(v interface{}) (string, error) {
	data, err := N.Bytes(v)
	if err != nil {
		return "", erero.Wro(err)
	}
	return string(data), nil
}

// S is shorthand for Sjson
//
// S 是 Sjson 方法的简写
func (N *Neatjson) S(v interface{}) (string, error) {
	return N.Sjson(v)
}

// B is shorthand for Bytes
//
// B 是 Bytes 方法的简写
func (N *Neatjson) B(v interface{}) ([]byte, error) {
	return N.Bytes(v)
}

// SxS converts JSON string to formatted JSON string
// Reformats existing JSON using configured indentation
// Returns source string on errors, preserving input
//
// SxS 将 JSON 字符串转换为格式化的 JSON 字符串
// 根据配置的缩进重新格式化现有 JSON
// 错误时返回源字符串，保留输入
func (N *Neatjson) SxS(s string) (string, error) {
	data, err := N.BxB([]byte(s))
	if err != nil {
		return s, erero.Wro(err)
	}
	return string(data), nil
}

// BxB converts JSON bytes to formatted JSON bytes
// Reformats existing JSON using json.Indent when configured
// Returns source bytes unchanged when indentation disabled
//
// BxB 将 JSON 字节转换为格式化的 JSON 字节
// 配置时使用 json.Indent 重新格式化现有 JSON
// 禁用缩进时返回未更改的源字节
func (N *Neatjson) BxB(data []byte) ([]byte, error) {
	if N.needIndent() {
		var result bytes.Buffer
		if err := json.Indent(&result, data, N.prefix, N.indent); err != nil {
			return data, erero.Wro(err)
		}
		return result.Bytes(), nil
	}
	return data, nil
}

// SxB converts JSON bytes to string (x means from)
// Returns formatted string, input data as string on errors
//
// SxB 将 JSON 字节转换为字符串（x 表示 from）
// 返回格式化的字符串，错误时返回输入数据字符串
func (N *Neatjson) SxB(data []byte) (string, error) {
	res, err := N.BxB(data)
	if err != nil {
		return string(data), erero.Wro(err)
	}
	return string(res), nil
}

// BxS converts JSON string to bytes (x means from)
// Returns formatted bytes, input data as bytes on errors
//
// BxS 将 JSON 字符串转换为字节（x 表示 from）
// 返回格式化的字节，错误时返回输入数据字节
func (N *Neatjson) BxS(data string) ([]byte, error) {
	res, err := N.BxB([]byte(data))
	if err != nil {
		return []byte(data), erero.Wro(err)
	}
	return res, nil
}

// B2S converts bytes to JSON string (2 means to)
//
// B2S 将字节转换为 JSON 字符串（2 表示 to）
func (N *Neatjson) B2S(data []byte) (string, error) {
	return N.SxB(data)
}

// S2B converts string to JSON bytes (2 means to)
//
// S2B 将字符串转换为 JSON 字节（2 表示 to）
func (N *Neatjson) S2B(data string) ([]byte, error) {
	return N.BxS(data)
}
