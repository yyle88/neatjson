package neatjson

import (
	"bytes"
	"encoding/json"

	"github.com/yyle88/erero"
)

// 这里定义些 常用常量 以便于外部直接使用
var (
	TAB = NewNeatjson("", "\t")
	SP0 = NewNeatjson("", "")
	SP1 = NewNeatjson("", " ")
	SP2 = NewNeatjson("", "  ")
	SP3 = NewNeatjson("", "   ")
	SP4 = NewNeatjson("", "    ")

	NOI = NewNeatjson("", "").setPrettyValue(false) //当使用这个转json时结果不换行

	NON = NOI //换个名字有利于身心健康
)

// Neatjson 该类中实现各种简单的语法糖，以便于把 golang struct 转换为 json string，以便于开发者（我自己）使用
// neat 含义有：整齐的、工整的、简洁的、整洁的
type Neatjson struct {
	prefix string
	indent string
	pretty bool //在转json时是否换行，当为false时不换行，还是用直接转的函数
}

func NewNeatjson(prefix string, indent string) *Neatjson {
	return &Neatjson{
		prefix: prefix,
		indent: indent,
		pretty: true, //默认设置为换行的，因为这个包的主要目的是得到有换行的json，因此默认就是有换行的
	}
}

// Bytes 把结构体转换为字节数组，跟 "json.Marshal" 类似的，只是这是带缩进的
func (N *Neatjson) Bytes(v interface{}) ([]byte, error) {
	if N.useIndentLogic() {
		data, err := json.MarshalIndent(v, N.prefix, N.indent)
		if err != nil {
			return nil, erero.Wro(err)
		}
		return data, nil
	} else {
		data, err := json.Marshal(v)
		if err != nil {
			return nil, erero.Wro(err)
		}
		return data, nil
	}
}

// 结果是否需要换行，因为发现单独写个包处理不换行这件事其实没必要，就在这里使用标识位判定
func (N *Neatjson) useIndentLogic() bool {
	return N.pretty || N.prefix != "" || N.indent != ""
}

// 设置是否需要换行，这里使用非导出函数
func (N *Neatjson) setPrettyValue(pretty bool) *Neatjson {
	N.pretty = pretty
	return N
}

// Sjson 把结构体转换为 string，但是由于 String 这个单词既不利于全局搜索（会干扰IDE索引）也不利于代码重构
// 因此换个个性化的方法名，避免和 interface { String() } 接口重名
func (N *Neatjson) Sjson(v interface{}) (string, error) {
	data, err := N.Bytes(v)
	if err != nil {
		return "", erero.Wro(err)
	}
	return string(data), nil
}

// S 就是 Sjson 方法的简称
func (N *Neatjson) S(v interface{}) (string, error) {
	return N.Sjson(v)
}

// B 就是 Bytes 方法的简称
func (N *Neatjson) B(v interface{}) ([]byte, error) {
	return N.Bytes(v)
}

// SxS 把 json string 转换得到 neat json string
func (N *Neatjson) SxS(s string) (string, error) {
	data, err := N.BxB([]byte(s))
	if err != nil {
		return s, erero.Wro(err)
	}
	return string(data), nil
}

// BxB 把 json bytes 转换得到 neat json bytes
func (N *Neatjson) BxB(data []byte) ([]byte, error) {
	if N.useIndentLogic() {
		var ob bytes.Buffer
		if err := json.Indent(&ob, data, N.prefix, N.indent); err != nil {
			return data, erero.Wro(err)
		}
		return ob.Bytes(), nil
	} else {
		return data, nil
	}
}

// SxB 把 json bytes 转换得到 neat json string. "x" means "from".
// 因为当做 json.Marshal 的时候会把 bytes 转换为 Base64 字符串，而有时这不是我们期望的
// 该函数期望的是把 bytes 转换为带缩进的 json 结果
func (N *Neatjson) SxB(data []byte) (string, error) {
	res, err := N.BxB(data)
	if err != nil {
		return string(data), erero.Wro(err)
	}
	return string(res), nil
}

// BxS 把 json string 转换得到 neat json bytes. "x" means "from".
func (N *Neatjson) BxS(data string) ([]byte, error) {
	res, err := N.BxB([]byte(data))
	if err != nil {
		return []byte(data), erero.Wro(err)
	}
	return res, nil
}

// B2S convert json bytes to neat json string. num 2 means "to".
// 而前面的 SxB 是 get neat json string from json bytes 两种命名看个人爱好吧
func (N *Neatjson) B2S(data []byte) (string, error) {
	return N.SxB(data)
}

// S2B convert json string to neat json bytes. num 2 means "to".
// 我认为把返回写在前面，有利于使用，因为我们在用函数时会更重视它的返回值
// 毕竟返回值是和下文相关的，而下文是我们要开发的，要让注意力去思考下文的需要，而不是去回顾上文的结果，这样有利于提高开发效率
// 因此我更倾向于使用 "x" 操作而非 "2" 操作
func (N *Neatjson) S2B(data string) ([]byte, error) {
	return N.BxS(data)
}
