package neatjsons

import (
	"testing"
)

func TestS(t *testing.T) {
	t.Log(S(map[string]any{"a": "abc", "n": 123}))
}

func TestB(t *testing.T) {
	t.Log(string(B(map[string]any{"a": "abc", "n": 123})))
}
