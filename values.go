package neatjson

// Common formatting constants allowing quick JSON formatting with different indentation
// These pre-configured instances give convenient access to standard formatting options
//
// 常用的格式化常量，用于快速 JSON 格式化，支持不同的缩进方式
// 这些预配置的实例提供了便捷的标准格式化选项
var (
	TAB = NewNeatjson("", "\t")   // TAB uses tab (\t) in indentation // TAB 使用制表符缩进
	SP0 = NewNeatjson("", "")     // SP0 uses no indentation // SP0 不使用缩进
	SP1 = NewNeatjson("", " ")    // SP1 uses 1 space for indentation // SP1 使用 1 个空格缩进
	SP2 = NewNeatjson("", "  ")   // SP2 uses 2 spaces for indentation // SP2 使用 2 个空格缩进
	SP3 = NewNeatjson("", "   ")  // SP3 uses 3 spaces for indentation // SP3 使用 3 个空格缩进
	SP4 = NewNeatjson("", "    ") // SP4 uses 4 spaces for indentation // SP4 uses 4 个空格缩进
)

// Compact JSON constants using Compactjson type
// These produce truly compact JSON without any whitespace
//
// 紧凑 JSON 常量，使用 Compactjson 类型
// 生成真正紧凑的 JSON，不包含任何空白
var (
	NOI = NewCompactjson() // NOI produces compact JSON without line breaks (No Indent) // NOI 生成紧凑的 JSON 不换行（无缩进）
	NON = NewCompactjson() // NON produces compact JSON without line breaks (No Indent) // NON 生成紧凑的 JSON 不换行（无缩进）
)
