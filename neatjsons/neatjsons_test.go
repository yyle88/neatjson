package neatjsons_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestBytes(t *testing.T) {
	input := map[string]any{"key": "value", "num": 123}
	expected := `{
	"key": "value",
	"num": 123
}`
	result := string(neatjsons.Bytes(input))
	require.Equal(t, expected, result)
}

func TestSjson(t *testing.T) {
	input := map[string]any{"key": "value", "num": 123}
	expected := `{
	"key": "value",
	"num": 123
}`
	result := neatjsons.Sjson(input)
	require.Equal(t, expected, result)
}

func TestS(t *testing.T) {
	input := map[string]any{"a": "abc", "n": 123}
	expected := `{
	"a": "abc",
	"n": 123
}`
	result := neatjsons.S(input)
	require.Equal(t, expected, result)
}

func TestB(t *testing.T) {
	input := map[string]any{"a": "abc", "n": 123}
	expected := `{
	"a": "abc",
	"n": 123
}`
	result := string(neatjsons.B(input))
	require.Equal(t, expected, result)
}

func TestSxS(t *testing.T) {
	input := `{"a":"abc","n":123}`
	expected := `{
	"a": "abc",
	"n": 123
}`
	result := neatjsons.SxS(input)
	require.Equal(t, expected, result)
}

func TestBxB(t *testing.T) {
	input := []byte(`{"a":"abc","n":123}`)
	expected := []byte(`{
	"a": "abc",
	"n": 123
}`)
	result := neatjsons.BxB(input)
	require.Equal(t, string(expected), string(result))
}

func TestSxB(t *testing.T) {
	input := []byte(`{"a":"abc","n":123}`)
	expected := `{
	"a": "abc",
	"n": 123
}`
	result := neatjsons.SxB(input)
	require.Equal(t, expected, result)
}

func TestBxS(t *testing.T) {
	input := `{"a":"abc","n":123}`
	expected := []byte(`{
	"a": "abc",
	"n": 123
}`)
	result := neatjsons.BxS(input)
	require.Equal(t, string(expected), string(result))
}

func TestB2S(t *testing.T) {
	input := []byte(`{"key":"value","num":123}`)
	expected := `{
	"key": "value",
	"num": 123
}`
	result := neatjsons.B2S(input)
	require.Equal(t, expected, result)
}

func TestS2B(t *testing.T) {
	input := `{"key":"value","num":123}`
	expected := []byte(`{
	"key": "value",
	"num": 123
}`)
	result := neatjsons.S2B(input)
	require.Equal(t, string(expected), string(result))
}
