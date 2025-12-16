// Code generated using sure/cls_stub_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: sure_stub.gen_test.go:96 -> sures.TestGenOmit
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package neatjsono

import "github.com/yyle88/neatjson"

func Bytes(v interface{}) []byte {
	return neatjson.TAB.Omit().Bytes(v)
}
func Sjson(v interface{}) string {
	return neatjson.TAB.Omit().Sjson(v)
}
func S(v interface{}) string {
	return neatjson.TAB.Omit().S(v)
}
func B(v interface{}) []byte {
	return neatjson.TAB.Omit().B(v)
}
func SxS(s string) string {
	return neatjson.TAB.Omit().SxS(s)
}
func BxB(data []byte) []byte {
	return neatjson.TAB.Omit().BxB(data)
}
func SxB(data []byte) string {
	return neatjson.TAB.Omit().SxB(data)
}
func BxS(data string) []byte {
	return neatjson.TAB.Omit().BxS(data)
}
func B2S(data []byte) string {
	return neatjson.TAB.Omit().B2S(data)
}
func S2B(data string) []byte {
	return neatjson.TAB.Omit().S2B(data)
}
