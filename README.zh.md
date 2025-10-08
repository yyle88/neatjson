[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/neatjson/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/neatjson/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/neatjson)](https://pkg.go.dev/github.com/yyle88/neatjson)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/neatjson/main.svg)](https://coveralls.io/github/yyle88/neatjson?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/neatjson.svg)](https://github.com/yyle88/neatjson/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/neatjson)](https://goreportcard.com/report/github.com/yyle88/neatjson)

# neatjson

`neatjson` 让 Golang 中的 JSON 编码变得简洁易用。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

🎨 **灵活的缩进选项**: 可选择 TAB、空格 (0-4) 或自定义缩进样式
⚡ **错误处理模式**: Must (panic)、Soft (log) 或 Omit (silent) - 根据上下文选择适合的模式
🔄 **双向格式化**: 支持格式化 Go 结构体和原始 JSON 数据（字符串/字节）
📦 **便捷封装包**: 自动生成的包，具有合理的默认值，便于快速开发
🛠️ **类型安全 API**: 简洁的链式接口，具备编译时安全检查

## 安装

```bash
go get github.com/yyle88/neatjson
```

## 使用方法

### 结构体转格式化 JSON

```go
package main

import (
	"fmt"

	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	// Convert struct to formatted JSON string
	data := map[string]any{
		"name":  "example",
		"count": 42,
	}

	result := neatjsons.S(data)
	fmt.Println(result)
}
```

⬆️ **源码：** [源码](internal/demos/demo1x/main.go)

### 格式化紧凑 JSON 字节

```go
package main

import (
	"fmt"

	"github.com/yyle88/neatjson/neatjsons"
)

func main() {
	// Format compact JSON bytes to readable string
	jsonBytes := []byte(`{"name":"test","age":30,"active":true}`)

	result := neatjsons.SxB(jsonBytes)
	fmt.Println(result)
}
```

⬆️ **源码：** [源码](internal/demos/demo2x/main.go)

## 示例

### 便捷封装包

**neatjsons - Must 模式：**
```go
import "github.com/yyle88/neatjson/neatjsons"
json := neatjsons.S(data)  // 出错时 panic
```

**neatjsono - Omit 模式：**
```go
import "github.com/yyle88/neatjson/neatjsono"
json := neatjsono.S(data)  // 出错时静默
```

**neatjsonz - Soft 模式：**
```go
import "github.com/yyle88/neatjson/neatjsonz"
json := neatjsonz.S(data)  // 出错时记录日志
```

### 缩进选项

**TAB 缩进（默认）：**
```go
result := neatjson.TAB.Must().S(data)
```

**无缩进（紧凑格式）：**
```go
result := neatjson.NOI.Must().S(data)
```

**2 空格缩进：**
```go
result := neatjson.SP2.Must().S(data)
```

**4 空格缩进：**
```go
result := neatjson.SP4.Must().S(data)
```

### 错误处理模式

**Must 模式 - 出错时 panic：**
```go
result := neatjson.TAB.Must().S(data)
```

**Soft 模式 - 出错时记录日志并返回空值：**
```go
result := neatjson.TAB.Soft().S(data)
```

**Omit 模式 - 出错时静默返回空值：**
```go
result := neatjson.TAB.Omit().S(data)
```

### 格式化原始 JSON 数据

**字符串到字符串：**
```go
formatted := neatjson.TAB.Must().SxS(`{"compact":"json"}`)
```

**字节到字节：**
```go
formatted := neatjson.SP2.Must().BxB([]byte(`{"raw":"data"}`))
```

**字节到字符串：**
```go
formatted := neatjson.TAB.Must().SxB(jsonBytes)
```

**字符串到字节：**
```go
formatted := neatjson.SP4.Must().BxS(jsonString)
```

### 转换为 JSON 字节

**Go 结构体转 JSON 字节：**
```go
type User struct {
	Name    string `json:"name"`
	Mailbox string `json:"mailbox"`
}
user := User{Name: "Alice", Mailbox: "alice@example.com"}
jsonBytes := neatjson.SP4.Must().B(user)
```

### 带返回值的错误处理

**链式调用并处理错误：**
```go
result, err := neatjson.TAB.S(complexData)
if err != nil {
	// 处理错误
}
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/neatjson.svg?variant=adaptive)](https://starchart.cc/yyle88/neatjson)
