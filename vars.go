package neatjson

// 这里定义些 常用常量 以便于外部直接使用
var (
	TAB = NewNeatjson("", "\t")
	SP0 = NewNeatjson("", "")
	SP1 = NewNeatjson("", " ")
	SP2 = NewNeatjson("", "  ")
	SP3 = NewNeatjson("", "   ")
	SP4 = NewNeatjson("", "    ")

	NOI = notNeatjson() //在转json时结果不换行
	NON = notNeatjson() //在转json时结果不换行
)
