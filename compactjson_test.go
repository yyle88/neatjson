package neatjson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompactjson_Bytes(t *testing.T) {
	data := map[string]interface{}{
		"a": "abc",
		"n": 123,
	}
	res, err := NewCompactjson().Bytes(data)
	require.NoError(t, err)
	require.Contains(t, string(res), `"a":"abc"`)
	require.Contains(t, string(res), `"n":123`)
	require.NotContains(t, string(res), "\n")
	t.Log(string(res))
}

func TestCompactjson_Sjson(t *testing.T) {
	data := map[string]interface{}{
		"a": "abc",
		"n": 123,
	}
	res, err := NewCompactjson().Sjson(data)
	require.NoError(t, err)
	require.Contains(t, res, `"a":"abc"`)
	require.Contains(t, res, `"n":123`)
	require.NotContains(t, res, "\n")
	t.Log(res)
}

func TestCompactjson_SxS(t *testing.T) {
	// Test compacting formatted JSON
	formattedJSON := `{
	"a": "abc",
	"n": 123
}`
	res, err := NewCompactjson().SxS(formattedJSON)
	require.NoError(t, err)
	require.Equal(t, `{"a":"abc","n":123}`, res)
	require.NotContains(t, res, "\n")
	require.NotContains(t, res, "\t")
	t.Log(res)
}

func TestCompactjson_BxB(t *testing.T) {
	// Test compacting formatted JSON bytes
	formattedJSON := []byte(`{
	"a": "abc",
	"n": 123
}`)
	res, err := NewCompactjson().BxB(formattedJSON)
	require.NoError(t, err)
	require.Equal(t, `{"a":"abc","n":123}`, string(res))
	require.NotContains(t, string(res), "\n")
	require.NotContains(t, string(res), "\t")
	t.Log(string(res))
}

func TestCompactjson_SxB(t *testing.T) {
	formattedJSON := []byte(`{
	"name": "test",
	"age": 30
}`)
	res, err := NewCompactjson().SxB(formattedJSON)
	require.NoError(t, err)
	require.Contains(t, res, `"name":"test"`)
	require.Contains(t, res, `"age":30`)
	require.NotContains(t, res, "\n")
	require.NotContains(t, res, "\t")
	t.Log(res)
}

func TestCompactjson_BxS(t *testing.T) {
	formattedJSON := `{
	"name": "test",
	"age": 30
}`
	res, err := NewCompactjson().BxS(formattedJSON)
	require.NoError(t, err)
	require.Contains(t, string(res), `"name":"test"`)
	require.Contains(t, string(res), `"age":30`)
	require.NotContains(t, string(res), "\n")
	require.NotContains(t, string(res), "\t")
	t.Log(string(res))
}

// TestCompactjson_BxB_InvalidJSON checks that BxB returns original input on invalid JSON
// This ensures the design rule: compact methods return original input on exception
//
// TestCompactjson_BxB_InvalidJSON 检查 BxB 在无效 JSON 时返回原始输入
// 确保设计规则：压缩方法在异常时返回原始输入
func TestCompactjson_BxB_InvalidJSON(t *testing.T) {
	// Invalid JSON: missing closing brace
	// 无效 JSON：缺少右大括号
	invalidJSON := []byte(`{"name": "test", "age": 30`)

	res, err := NewCompactjson().BxB(invalidJSON)
	require.Error(t, err)
	// Must return original input, not nil or blank
	// 必须返回原始输入，而不是 nil 或空白
	require.Equal(t, invalidJSON, res)
	t.Log("Returned original input on exception:", string(res))
}

// TestCompactjson_SxS_InvalidJSON checks that SxS returns original input on invalid JSON
// This ensures the design rule: compact methods return original input on exception
//
// TestCompactjson_SxS_InvalidJSON 检查 SxS 在无效 JSON 时返回原始输入
// 确保设计规则：压缩方法在异常时返回原始输入
func TestCompactjson_SxS_InvalidJSON(t *testing.T) {
	// Invalid JSON: trailing comma
	// 无效 JSON：尾随逗号
	invalidJSON := `{"name": "test",}`

	res, err := NewCompactjson().SxS(invalidJSON)
	require.Error(t, err)
	// Must return original input, not blank string
	// 必须返回原始输入，而不是空字符串
	require.Equal(t, invalidJSON, res)
	t.Log("Returned original input on exception:", res)
}
