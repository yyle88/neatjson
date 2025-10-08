package neatjson

import (
	"testing"

	"github.com/yyle88/rese"
)

// TestNeatjson_SPX_Sjson checks JSON formatting at each space indentation step
// Tests SP0 (no indent), SP1 (1 space), SP2 (2 spaces), SP3 (3 spaces), SP4 (4 spaces)
// Validates that each preset config produces consistent JSON output
//
// TestNeatjson_SPX_Sjson 检查每个空格缩进步骤的 JSON 格式化
// 测试 SP0（无缩进）、SP1（1 空格）、SP2（2 空格）、SP3（3 空格）、SP4（4 空格）
// 验证每个预设配置产生一致的 JSON 输出
func TestNeatjson_SPX_Sjson(t *testing.T) {
	t.Log(rese.C1(SP0.Sjson(caseExample)))
	t.Log(rese.C1(SP1.Sjson(caseExample)))
	t.Log(rese.C1(SP2.Sjson(caseExample)))
	t.Log(rese.C1(SP3.Sjson(caseExample)))
	t.Log(rese.C1(SP4.Sjson(caseExample)))
}
