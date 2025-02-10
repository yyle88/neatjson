[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/neatjson/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/neatjson/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/neatjson)](https://pkg.go.dev/github.com/yyle88/neatjson)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/neatjson/master.svg)](https://coveralls.io/github/yyle88/neatjson?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/neatjson.svg)](https://github.com/yyle88/neatjson/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/neatjson)](https://goreportcard.com/report/github.com/yyle88/neatjson)

# neatjson

`neatjson` make it neat to use "encoding/json" in golang.

## CHINESE README

[中文说明](README.zh.md)

## Installation

```bash
go get github.com/yyle88/neatjson
```

## Features

- Convert Go structures to indented JSON strings.
- Format JSON data from raw strings, byte arrays.

## Usage

Here's an example of how to format a Go data structure into a indented JSON string:

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

Output:

```json
{
	"a": "abc",
	"n": 123
}
```

---

## License

`neatjson` is open-source and released under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `neatjson`!** 🎉

Give me stars. Thank you!!!

## Starring

[![starring](https://starchart.cc/yyle88/neatjson.svg?variant=adaptive)](https://starchart.cc/yyle88/neatjson)
