# neatjson

`neatjson` 使得在 Golang 中使用 "encoding/json" 更加简洁和方便。

## 英文文档

[English README](README.md)

## 安装

```bash
go get github.com/yyle88/neatjson
```

## 特性

- 将 Go 结构体转换为带缩进的 JSON 字符串。
- 格式化原始字符串和字节数组中的 JSON 数据。

## 用法

以下是如何将 Go 数据结构格式化为带缩进的 JSON 字符串的示例：

```go
package main

import (
	"fmt"
	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	arg := map[string]any{"a": "abc", "n": 123}
	res := neatjsons.S(arg)

	fmt.Println(res)
}
```

输出：

```json
{
	"a": "abc",
	"n": 123
}
```

---

## 许可

`neatjson` 是一个开源项目，发布于 MIT 许可证下。有关更多信息，请参阅 [LICENSE](LICENSE) 文件。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
