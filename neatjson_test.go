package neatjson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// exampleType represents a sample structure in JSON formatting tests
// Used as test data to check formatting operations
//
// exampleType 代表用于 JSON 格式化测试的示例结构
// 用作测试数据以检查格式化操作
type exampleType struct {
	A string `json:"a"` // String field // 字符串字段
	N int    `json:"n"` // Int field // 整数字段
}

// caseExample is the test data instance used across multiple test cases
// Provides consistent input to check JSON formatting actions
//
// caseExample 是跨多个测试用例使用的测试数据实例
// 为检查 JSON 格式化操作提供一致的输入
var caseExample = exampleType{
	A: "abc",
	N: 123,
}

// TestNeatjson_Bytes checks JSON formatting to byte slice with 2-space indentation
// Tests the Bytes method using SP2 (2-space indent) config
//
// TestNeatjson_Bytes 检查使用 2 空格缩进格式化 JSON 到字节切片
// 测试使用 SP2（2 空格缩进）配置的 Bytes 方法
func TestNeatjson_Bytes(t *testing.T) {
	res, err := SP2.Bytes(caseExample)
	require.NoError(t, err)
	t.Log(string(res))
}

// TestNeatjson_Sjson checks JSON formatting to string with tab indentation
// Tests the Sjson method using TAB config producing standard JSON output
//
// TestNeatjson_Sjson 检查使用制表符缩进格式化 JSON 到字符串
// 测试使用 TAB 配置的 Sjson 方法，生成标准 JSON 输出
func TestNeatjson_Sjson(t *testing.T) {
	res, err := TAB.Sjson(caseExample)
	require.NoError(t, err)
	t.Log(res)
}

// TestNeatjson_NOI_Sjson checks JSON string formatting without indentation
// Tests the NOI (no indent) config producing compact JSON output
//
// TestNeatjson_NOI_Sjson 检查不带缩进的 JSON 字符串格式化
// 测试 NOI（无缩进）配置，生成紧凑的 JSON 输出
func TestNeatjson_NOI_Sjson(t *testing.T) {
	res, err := NOI.Sjson(caseExample)
	require.NoError(t, err)
	t.Log(res)
}

// TestNeatjson_SxS checks string-to-string JSON reformatting with 4-space indentation
// Tests the SxS method converting existing JSON strings to formatted versions
//
// TestNeatjson_SxS 检查使用 4 空格缩进的字符串到字符串 JSON 重新格式化
// 测试 SxS 方法，将现有 JSON 字符串转换为格式化版本
func TestNeatjson_SxS(t *testing.T) {
	arg := `{"a": "abc","n": 123}`
	res, err := SP4.SxS(arg)
	require.NoError(t, err)
	t.Log(res)
}

// TestNeatjson_NON_SxS checks JSON string normalization without newlines
// Tests the NON (no newline) config producing single-line JSON output
//
// TestNeatjson_NON_SxS 检查不带换行符的 JSON 字符串规范化
// 测试 NON（无换行）配置，生成单行 JSON 输出
func TestNeatjson_NON_SxS(t *testing.T) {
	arg := `{"a": "abc","n": 123}`
	res, err := NON.SxS(arg)
	require.NoError(t, err)
	t.Log(res)
}

// TestNeatjson_BxB checks byte-to-byte JSON reformatting with 1-space indentation
// Tests the BxB method converting byte slices to formatted JSON bytes
//
// TestNeatjson_BxB 检查使用 1 空格缩进的字节到字节 JSON 重新格式化
// 测试 BxB 方法，将字节切片转换为格式化的 JSON 字节
func TestNeatjson_BxB(t *testing.T) {
	arg := []byte(`{"a": "abc","n": 123}`)
	res, err := SP1.BxB(arg)
	require.NoError(t, err)
	t.Log(string(res))
}

// TestNeatjson_BxB_InvalidJSON checks that BxB returns original input on invalid JSON
// This ensures the design rule: reformat methods return original input on exception
//
// TestNeatjson_BxB_InvalidJSON 检查 BxB 在无效 JSON 时返回原始输入
// 确保设计规则：重新格式化方法在异常时返回原始输入
func TestNeatjson_BxB_InvalidJSON(t *testing.T) {
	// Invalid JSON: missing closing brace
	// 无效 JSON：缺少右大括号
	invalidJSON := []byte(`{"name": "test", "age": 30`)

	res, err := TAB.BxB(invalidJSON)
	require.Error(t, err)
	// Must return original input, not nil or blank
	// 必须返回原始输入，而不是 nil 或空白
	require.Equal(t, invalidJSON, res)
	t.Log("Returned original input on exception:", string(res))
}

// TestNeatjson_SxS_InvalidJSON checks that SxS returns original input on invalid JSON
// This ensures the design rule: reformat methods return original input on exception
//
// TestNeatjson_SxS_InvalidJSON 检查 SxS 在无效 JSON 时返回原始输入
// 确保设计规则：重新格式化方法在异常时返回原始输入
func TestNeatjson_SxS_InvalidJSON(t *testing.T) {
	// Invalid JSON: trailing comma
	// 无效 JSON：尾随逗号
	invalidJSON := `{"name": "test",}`

	res, err := SP2.SxS(invalidJSON)
	require.Error(t, err)
	// Must return original input, not blank string
	// 必须返回原始输入，而不是空字符串
	require.Equal(t, invalidJSON, res)
	t.Log("Returned original input on exception:", res)
}
