// Code generated using sure/cls_stub_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: sure_stub.gen_test.go:44 -> sures.TestGenSoft
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package neatjsons

import "github.com/yyle88/neatjson"

func Bytes(v interface{}) []byte {
	return neatjson.TAB.Soft().Bytes(v)
}
func Sjson(v interface{}) string {
	return neatjson.TAB.Soft().Sjson(v)
}
func S(v interface{}) string {
	return neatjson.TAB.Soft().S(v)
}
func B(v interface{}) []byte {
	return neatjson.TAB.Soft().B(v)
}
func SxS(s string) string {
	return neatjson.TAB.Soft().SxS(s)
}
func BxB(data []byte) []byte {
	return neatjson.TAB.Soft().BxB(data)
}
func SxB(data []byte) string {
	return neatjson.TAB.Soft().SxB(data)
}
func BxS(data string) []byte {
	return neatjson.TAB.Soft().BxS(data)
}
func B2S(data []byte) string {
	return neatjson.TAB.Soft().B2S(data)
}
func S2B(data string) []byte {
	return neatjson.TAB.Soft().S2B(data)
}
