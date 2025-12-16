// Package neatjson provides elegant JSON formatting with flexible indentation
// Compactjson handles JSON compression to produce truly compact output
//
// neatjson 包提供优雅的 JSON 格式化，支持灵活的缩进控制
// Compactjson 处理 JSON 压缩，生成真正紧凑的输出
package neatjson

import (
	"bytes"
	"encoding/json"

	"github.com/yyle88/erero"
)

// Compactjson implements JSON compression without any whitespace or newlines
// Produces truly compact JSON output for minimal size
// Provides same method interface as Neatjson for consistency
//
// Compactjson 实现 JSON 压缩，不包含任何空白或换行
// 生成真正紧凑的 JSON 输出，体积最小化
// 提供与 Neatjson 相同的方法接口，保持一致性
type Compactjson struct{}

// NewCompactjson creates Compactjson instance for JSON compression
// Returns instance that produces compact JSON without whitespace
//
// NewCompactjson 创建用于 JSON 压缩的 Compactjson 实例
// 返回生成无空白紧凑 JSON 的实例
func NewCompactjson() *Compactjson {
	return &Compactjson{}
}

// Bytes converts struct to compact JSON bytes
// Uses json.Marshal to produce minimal JSON output
//
// Bytes 将结构体转换为紧凑的 JSON 字节数组
// 使用 json.Marshal 生成最小化的 JSON 输出
func (C *Compactjson) Bytes(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return data, nil
}

// Sjson converts struct to compact JSON string
// Wraps Bytes method and converts result to string
//
// Sjson 将结构体转换为紧凑的 JSON 字符串
// 封装 Bytes 方法并将结果转换为字符串
func (C *Compactjson) Sjson(v interface{}) (string, error) {
	data, err := C.Bytes(v)
	if err != nil {
		return "", erero.Wro(err)
	}
	return string(data), nil
}

// S is shorthand for Sjson
//
// S 是 Sjson 方法的简写
func (C *Compactjson) S(v interface{}) (string, error) {
	return C.Sjson(v)
}

// B is shorthand for Bytes
//
// B 是 Bytes 方法的简写
func (C *Compactjson) B(v interface{}) ([]byte, error) {
	return C.Bytes(v)
}

// SxS converts JSON string to compact JSON string
// Uses json.Compact to remove all unnecessary whitespace
//
// SxS 将 JSON 字符串转换为紧凑的 JSON 字符串
// 使用 json.Compact 移除所有不必要的空白
func (C *Compactjson) SxS(s string) (string, error) {
	data, err := C.BxB([]byte(s))
	if err != nil {
		return s, erero.Wro(err)
	}
	return string(data), nil
}

// BxB converts JSON bytes to compact JSON bytes
// Uses json.Compact to produce truly compact output
//
// BxB 将 JSON 字节转换为紧凑的 JSON 字节
// 使用 json.Compact 生成真正紧凑的输出
func (C *Compactjson) BxB(data []byte) ([]byte, error) {
	var result bytes.Buffer
	if err := json.Compact(&result, data); err != nil {
		return data, erero.Wro(err)
	}
	return result.Bytes(), nil
}

// SxB converts JSON bytes to compact string (x means from)
// Returns compact string, input data as string on errors
//
// SxB 将 JSON 字节转换为紧凑字符串（x 表示 from）
// 返回紧凑字符串，错误时返回输入数据字符串
func (C *Compactjson) SxB(data []byte) (string, error) {
	res, err := C.BxB(data)
	if err != nil {
		return string(data), erero.Wro(err)
	}
	return string(res), nil
}

// BxS converts JSON string to compact bytes (x means from)
// Returns compact bytes, input data as bytes on errors
//
// BxS 将 JSON 字符串转换为紧凑字节（x 表示 from）
// 返回紧凑字节，错误时返回输入数据字节
func (C *Compactjson) BxS(data string) ([]byte, error) {
	res, err := C.BxB([]byte(data))
	if err != nil {
		return []byte(data), erero.Wro(err)
	}
	return res, nil
}

// B2S converts bytes to compact JSON string (2 means to)
//
// B2S 将字节转换为紧凑 JSON 字符串（2 表示 to）
func (C *Compactjson) B2S(data []byte) (string, error) {
	return C.SxB(data)
}

// S2B converts string to compact JSON bytes (2 means to)
//
// S2B 将字符串转换为紧凑 JSON 字节（2 表示 to）
func (C *Compactjson) S2B(data string) ([]byte, error) {
	return C.BxS(data)
}
