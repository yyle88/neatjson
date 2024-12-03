package neatjson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type exampleType struct {
	A string `json:"a"`
	N int    `json:"n"`
}

var caseExample = exampleType{
	A: "abc",
	N: 123,
}

func TestNeatjson_Bytes(t *testing.T) {
	res, err := SP2.Bytes(caseExample)
	require.NoError(t, err)
	t.Log(string(res))
}

func TestNeatjson_Sjson(t *testing.T) {
	res, err := TAB.Sjson(caseExample)
	require.NoError(t, err)
	t.Log(res)
}

func TestNeatjson_NOI_Sjson(t *testing.T) {
	res, err := NOI.Sjson(caseExample)
	require.NoError(t, err)
	t.Log(res)
}

func TestNeatjson_SxS(t *testing.T) {
	arg := `{"a": "abc","n": 123}`
	res, err := SP4.SxS(arg)
	require.NoError(t, err)
	t.Log(res)
}

func TestNeatjson_NON_SxS(t *testing.T) {
	arg := `{"a": "abc","n": 123}`
	res, err := NON.SxS(arg)
	require.NoError(t, err)
	t.Log(res)
}

func TestNeatjson_BxB(t *testing.T) {
	arg := []byte(`{"a": "abc","n": 123}`)
	res, err := SP1.BxB(arg)
	require.NoError(t, err)
	t.Log(string(res))
}
