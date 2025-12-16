// Code generated using sure/cls_stub_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: sure_stub.gen_test.go:70 -> sures.TestGenMust
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package neatjsonm

import "github.com/yyle88/neatjson"

func Bytes(v interface{}) []byte {
	return neatjson.TAB.Must().Bytes(v)
}
func Sjson(v interface{}) string {
	return neatjson.TAB.Must().Sjson(v)
}
func S(v interface{}) string {
	return neatjson.TAB.Must().S(v)
}
func B(v interface{}) []byte {
	return neatjson.TAB.Must().B(v)
}
func SxS(s string) string {
	return neatjson.TAB.Must().SxS(s)
}
func BxB(data []byte) []byte {
	return neatjson.TAB.Must().BxB(data)
}
func SxB(data []byte) string {
	return neatjson.TAB.Must().SxB(data)
}
func BxS(data string) []byte {
	return neatjson.TAB.Must().BxS(data)
}
func B2S(data []byte) string {
	return neatjson.TAB.Must().B2S(data)
}
func S2B(data string) []byte {
	return neatjson.TAB.Must().S2B(data)
}
