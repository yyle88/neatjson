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
