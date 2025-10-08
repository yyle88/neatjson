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
