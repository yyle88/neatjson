[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/neatjson/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/neatjson/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/neatjson)](https://pkg.go.dev/github.com/yyle88/neatjson)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/neatjson/main.svg)](https://coveralls.io/github/yyle88/neatjson?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/neatjson.svg)](https://github.com/yyle88/neatjson/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/neatjson)](https://goreportcard.com/report/github.com/yyle88/neatjson)

# neatjson

`neatjson` makes JSON encoding in Golang clean and simple.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

🎨 **Flexible Indentation Options**: Choose from TAB, spaces (0-4), and custom indentation styles
⚡ **Exception Handling Modes**: Must (panic), Soft (log), and Omit (silent) - pick the mode that fits context
🔄 **Struct & JSON Formatting**: Format both Go structures and raw JSON data (strings/bytes)
📦 **Convenience Packages**: Auto-generated packages with sensible defaults
🛠️ **Type-Safe API**: Clean, chainable interface with compile-time checks

## Installation

```bash
go get github.com/yyle88/neatjson
```

## Usage

### Convert Struct to Formatted JSON

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

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Format Compact JSON Bytes

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

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## Examples

### Convenience Packages

**neatjsons - Must mode:**

```go
import "github.com/yyle88/neatjson/neatjsons"
json := neatjsons.S(data)  // panic on failure
```

**neatjsono - Omit mode:**

```go
import "github.com/yyle88/neatjson/neatjsono"
json := neatjsono.S(data)  // silent on failure
```

**neatjsonz - Soft mode:**

```go
import "github.com/yyle88/neatjson/neatjsonz"
json := neatjsonz.S(data)  // log on failure
```

### Indentation Options

**TAB indentation (default):**

```go
result := neatjson.TAB.Must().S(data)
```

**No indentation (compact):**

```go
result := neatjson.NOI.Must().S(data)
```

**2-space indentation:**

```go
result := neatjson.SP2.Must().S(data)
```

**4-space indentation:**

```go
result := neatjson.SP4.Must().S(data)
```

### Exception Handling Modes

**Must mode - panic on exception:**

```go
result := neatjson.TAB.Must().S(data)
```

**Soft mode - log and give blank:**

```go
result := neatjson.TAB.Soft().S(data)
```

**Omit mode - silent, give blank:**

```go
result := neatjson.TAB.Omit().S(data)
```

### Format Raw JSON Data

**String to String:**

```go
formatted := neatjson.TAB.Must().SxS(`{"compact":"json"}`)
```

**Bytes to Bytes:**

```go
formatted := neatjson.SP2.Must().BxB([]byte(`{"raw":"data"}`))
```

**Bytes to String:**

```go
formatted := neatjson.TAB.Must().SxB(jsonBytes)
```

**String to Bytes:**

```go
formatted := neatjson.SP4.Must().BxS(jsonString)
```

### Convert to JSON Bytes

**Go struct to JSON bytes:**

```go
type User struct {
	Name    string `json:"name"`
	Mailbox string `json:"mailbox"`
}
user := User{Name: "Alice", Mailbox: "alice@example.com"}
jsonBytes := neatjson.SP4.Must().B(user)
```

### Exception Handling with Return Values

**Chain with exception handling:**

```go
result, err := neatjson.TAB.S(complexData)
if err != nil {
	// Handle exception
}
```

## API Reference

### Indentation Constants

| Constant | Description               |
| -------- | ------------------------- |
| `TAB`    | Tab indentation           |
| `SP0`    | No indentation            |
| `SP1`    | 1-space indentation       |
| `SP2`    | 2-space indentation       |
| `SP3`    | 3-space indentation       |
| `SP4`    | 4-space indentation       |
| `NOI`    | Compact JSON (no indent)  |
| `NON`    | Compact JSON (no newline) |

### Neatjson Methods

| Method                    | Description           |
| ------------------------- | --------------------- |
| `S(v)` / `Sjson(v)`       | Struct to JSON string |
| `B(v)` / `Bytes(v)`       | Struct to JSON bytes  |
| `SxS(s)`                  | Reformat JSON string  |
| `BxB(data)`               | Reformat JSON bytes   |
| `SxB(data)` / `B2S(data)` | JSON bytes to string  |
| `BxS(data)` / `S2B(data)` | JSON string to bytes  |

### Handling Modes

| Method   | Description        |
| -------- | ------------------ |
| `Must()` | Panic on exception |
| `Soft()` | Log and give blank |
| `Omit()` | Silent, give blank |

### Design Notes

**Reformat methods keep source input on exception:**

`SxS`/`BxB`/`SxB`/`BxS` pass back source input plus exception when receiving invalid JSON. No data loss.

```go
invalidJSON := `{"name": "test",}`  // trailing comma - invalid
result, e := neatjson.TAB.SxS(invalidJSON)
// e != nil, result == invalidJSON (source input kept safe)
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/neatjson.svg?variant=adaptive)](https://starchart.cc/yyle88/neatjson)
