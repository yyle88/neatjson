package neatjson

import (
	"testing"

	"github.com/yyle88/rese"
)

func TestNeatjson_SPX_Sjson(t *testing.T) {
	t.Log(rese.C1(SP0.Sjson(caseExample)))
	t.Log(rese.C1(SP1.Sjson(caseExample)))
	t.Log(rese.C1(SP2.Sjson(caseExample)))
	t.Log(rese.C1(SP3.Sjson(caseExample)))
	t.Log(rese.C1(SP4.Sjson(caseExample)))
}
