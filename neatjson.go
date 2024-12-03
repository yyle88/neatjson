package neatjson

import (
	"bytes"
	"encoding/json"

	"github.com/yyle88/erero"
)

// Neatjson implements various syntactic sugar to convert Go structs into JSON strings.
// The purpose of this package is to provide neat, indented JSON representations.
// Neatjson 实现了一些简单的语法糖，用于将 Go 的结构体转换为 JSON 字符串。
// neat 的含义包括：整齐的、工整的、简洁的、整洁的。
type Neatjson struct {
	prefix string // Prefix for JSON indentation. // 结果 JSON 的缩进前缀。
	indent string // Indentation characters. // 缩进字符。
	active bool   // Whether to use line breaks when converting to JSON. // 是否使用缩进的转换函数。
}

// NewNeatjson creates a new instance of Neatjson with the specified prefix and indent.
// NewNeatjson 根据指定的前缀和缩进字符创建一个 Neatjson 实例。
func NewNeatjson(prefix string, indent string) *Neatjson {
	return &Neatjson{
		prefix: prefix,
		indent: indent,
		active: true, // Default to using line breaks because the main purpose of this package is to create indented JSON. // 默认启用换行，因为这个包的主要目的就是生成带缩进的 JSON。
	}
}

// notNeatjson returns a Neatjson instance that disables line breaks.
// notNeatjson 返回一个禁用换行的 Neatjson 实例。
func notNeatjson() *Neatjson {
	return NewNeatjson("", "").withActive(false)
}

// Bytes converts a struct to a JSON byte array, similar to "json.Marshal" but supports indentation.
// Bytes 将结构体转换为 JSON 字节数组，与 "json.Marshal" 类似，但支持缩进。
func (N *Neatjson) Bytes(v interface{}) ([]byte, error) {
	if N.needIndent() { // Check if indentation is required. // 检查是否需要缩进。
		data, err := json.MarshalIndent(v, N.prefix, N.indent)
		if err != nil {
			return nil, erero.Wro(err)
		}
		return data, nil
	} else { // Use standard JSON marshaling without indentation. // 使用标准的 JSON 转换，不带缩进。
		data, err := json.Marshal(v)
		if err != nil {
			return nil, erero.Wro(err)
		}
		return data, nil
	}
}

// needIndent checks whether line breaks and indentation are required.
// needIndent 检查是否需要换行和缩进。
func (N *Neatjson) needIndent() bool {
	return N.active || N.prefix != "" || N.indent != ""
}

// withActive sets whether line breaks are enabled for JSON conversion.
// withActive 设置 JSON 转换时是否启用换行。
func (N *Neatjson) withActive(active bool) *Neatjson {
	N.active = active
	return N
}

// Sjson converts a struct to a JSON string with the desired formatting.
// Sjson 将结构体转换为符合要求格式的 JSON 字符串。
func (N *Neatjson) Sjson(v interface{}) (string, error) {
	data, err := N.Bytes(v)
	if err != nil {
		return "", erero.Wro(err)
	}
	return string(data), nil
}

// S is a shorthand for Sjson.
// S 是 Sjson 方法的简写。
func (N *Neatjson) S(v interface{}) (string, error) {
	return N.Sjson(v)
}

// B is a shorthand for Bytes.
// B 是 Bytes 方法的简写。
func (N *Neatjson) B(v interface{}) ([]byte, error) {
	return N.Bytes(v)
}

// SxS converts a JSON string into a neat JSON string with the desired formatting.
// SxS 将 JSON 字符串转换为符合要求格式的整齐 JSON 字符串。
func (N *Neatjson) SxS(s string) (string, error) {
	data, err := N.BxB([]byte(s))
	if err != nil {
		return s, erero.Wro(err)
	}
	return string(data), nil
}

// BxB converts a JSON byte array into a neat JSON byte array with the desired formatting.
// BxB 将 JSON 字节数组转换为符合要求格式的整齐 JSON 字节数组。
func (N *Neatjson) BxB(data []byte) ([]byte, error) {
	if N.needIndent() { // Use json.Indent for formatting when indentation is needed. // 在需要缩进时使用 json.Indent 格式化。
		var result bytes.Buffer
		if err := json.Indent(&result, data, N.prefix, N.indent); err != nil {
			return data, erero.Wro(err)
		}
		return result.Bytes(), nil
	} else {
		return data, nil
	}
}

// SxB converts a JSON byte array into a neat JSON string. "x" means "from".
// SxB 将 JSON 字节数组转换为整齐的 JSON 字符串。"x" 表示 "from"。
func (N *Neatjson) SxB(data []byte) (string, error) {
	res, err := N.BxB(data)
	if err != nil {
		return string(data), erero.Wro(err)
	}
	return string(res), nil
}

// BxS converts a JSON string into a neat JSON byte array. "x" means "from".
// BxS 将 JSON 字符串转换为整齐的 JSON 字节数组。"x" 表示 "from"。
func (N *Neatjson) BxS(data string) ([]byte, error) {
	res, err := N.BxB([]byte(data))
	if err != nil {
		return []byte(data), erero.Wro(err)
	}
	return res, nil
}

// B2S converts JSON bytes into a neat JSON string. "2" means "to".
// B2S 将 JSON 字节数组转换为整齐的 JSON 字符串。"2" 表示 "to"。
func (N *Neatjson) B2S(data []byte) (string, error) {
	return N.SxB(data)
}

// S2B converts a JSON string into a neat JSON byte array. "2" means "to".
// S2B 将 JSON 字符串转换为整齐的 JSON 字节数组。"2" 表示 "to"。
func (N *Neatjson) S2B(data string) ([]byte, error) {
	return N.BxS(data)
}
