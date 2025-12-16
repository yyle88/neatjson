// Code generated using sure/sure_cls_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: sure_cls.gen_test.go:40 -> neatjson.TestGen
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package neatjson

import "github.com/yyle88/sure"

type Neatjson_Must struct{ N *Neatjson }

func (N *Neatjson) Must() *Neatjson_Must {
	return &Neatjson_Must{N: N}
}
func (T *Neatjson_Must) Bytes(v interface{}) (res []byte) {
	res, err1 := T.N.Bytes(v)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) Sjson(v interface{}) (res string) {
	res, err1 := T.N.Sjson(v)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) S(v interface{}) (res string) {
	res, err1 := T.N.S(v)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) B(v interface{}) (res []byte) {
	res, err1 := T.N.B(v)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) SxS(s string) (res string) {
	res, err1 := T.N.SxS(s)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) BxB(data []byte) (res []byte) {
	res, err1 := T.N.BxB(data)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) SxB(data []byte) (res string) {
	res, err1 := T.N.SxB(data)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) BxS(data string) (res []byte) {
	res, err1 := T.N.BxS(data)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) B2S(data []byte) (res string) {
	res, err1 := T.N.B2S(data)
	sure.Must(err1)
	return res
}
func (T *Neatjson_Must) S2B(data string) (res []byte) {
	res, err1 := T.N.S2B(data)
	sure.Must(err1)
	return res
}

type Neatjson_Soft struct{ N *Neatjson }

func (N *Neatjson) Soft() *Neatjson_Soft {
	return &Neatjson_Soft{N: N}
}
func (T *Neatjson_Soft) Bytes(v interface{}) (res []byte) {
	res, err1 := T.N.Bytes(v)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) Sjson(v interface{}) (res string) {
	res, err1 := T.N.Sjson(v)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) S(v interface{}) (res string) {
	res, err1 := T.N.S(v)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) B(v interface{}) (res []byte) {
	res, err1 := T.N.B(v)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) SxS(s string) (res string) {
	res, err1 := T.N.SxS(s)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) BxB(data []byte) (res []byte) {
	res, err1 := T.N.BxB(data)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) SxB(data []byte) (res string) {
	res, err1 := T.N.SxB(data)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) BxS(data string) (res []byte) {
	res, err1 := T.N.BxS(data)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) B2S(data []byte) (res string) {
	res, err1 := T.N.B2S(data)
	sure.Soft(err1)
	return res
}
func (T *Neatjson_Soft) S2B(data string) (res []byte) {
	res, err1 := T.N.S2B(data)
	sure.Soft(err1)
	return res
}

type Neatjson_Omit struct{ N *Neatjson }

func (N *Neatjson) Omit() *Neatjson_Omit {
	return &Neatjson_Omit{N: N}
}
func (T *Neatjson_Omit) Bytes(v interface{}) (res []byte) {
	res, err1 := T.N.Bytes(v)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) Sjson(v interface{}) (res string) {
	res, err1 := T.N.Sjson(v)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) S(v interface{}) (res string) {
	res, err1 := T.N.S(v)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) B(v interface{}) (res []byte) {
	res, err1 := T.N.B(v)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) SxS(s string) (res string) {
	res, err1 := T.N.SxS(s)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) BxB(data []byte) (res []byte) {
	res, err1 := T.N.BxB(data)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) SxB(data []byte) (res string) {
	res, err1 := T.N.SxB(data)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) BxS(data string) (res []byte) {
	res, err1 := T.N.BxS(data)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) B2S(data []byte) (res string) {
	res, err1 := T.N.B2S(data)
	sure.Omit(err1)
	return res
}
func (T *Neatjson_Omit) S2B(data string) (res []byte) {
	res, err1 := T.N.S2B(data)
	sure.Omit(err1)
	return res
}

type Compactjson_Must struct{ N *Compactjson }

func (N *Compactjson) Must() *Compactjson_Must {
	return &Compactjson_Must{N: N}
}
func (T *Compactjson_Must) Bytes(v interface{}) (res []byte) {
	res, err1 := T.N.Bytes(v)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) Sjson(v interface{}) (res string) {
	res, err1 := T.N.Sjson(v)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) S(v interface{}) (res string) {
	res, err1 := T.N.S(v)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) B(v interface{}) (res []byte) {
	res, err1 := T.N.B(v)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) SxS(s string) (res string) {
	res, err1 := T.N.SxS(s)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) BxB(data []byte) (res []byte) {
	res, err1 := T.N.BxB(data)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) SxB(data []byte) (res string) {
	res, err1 := T.N.SxB(data)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) BxS(data string) (res []byte) {
	res, err1 := T.N.BxS(data)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) B2S(data []byte) (res string) {
	res, err1 := T.N.B2S(data)
	sure.Must(err1)
	return res
}
func (T *Compactjson_Must) S2B(data string) (res []byte) {
	res, err1 := T.N.S2B(data)
	sure.Must(err1)
	return res
}

type Compactjson_Soft struct{ N *Compactjson }

func (N *Compactjson) Soft() *Compactjson_Soft {
	return &Compactjson_Soft{N: N}
}
func (T *Compactjson_Soft) Bytes(v interface{}) (res []byte) {
	res, err1 := T.N.Bytes(v)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) Sjson(v interface{}) (res string) {
	res, err1 := T.N.Sjson(v)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) S(v interface{}) (res string) {
	res, err1 := T.N.S(v)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) B(v interface{}) (res []byte) {
	res, err1 := T.N.B(v)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) SxS(s string) (res string) {
	res, err1 := T.N.SxS(s)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) BxB(data []byte) (res []byte) {
	res, err1 := T.N.BxB(data)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) SxB(data []byte) (res string) {
	res, err1 := T.N.SxB(data)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) BxS(data string) (res []byte) {
	res, err1 := T.N.BxS(data)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) B2S(data []byte) (res string) {
	res, err1 := T.N.B2S(data)
	sure.Soft(err1)
	return res
}
func (T *Compactjson_Soft) S2B(data string) (res []byte) {
	res, err1 := T.N.S2B(data)
	sure.Soft(err1)
	return res
}

type Compactjson_Omit struct{ N *Compactjson }

func (N *Compactjson) Omit() *Compactjson_Omit {
	return &Compactjson_Omit{N: N}
}
func (T *Compactjson_Omit) Bytes(v interface{}) (res []byte) {
	res, err1 := T.N.Bytes(v)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) Sjson(v interface{}) (res string) {
	res, err1 := T.N.Sjson(v)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) S(v interface{}) (res string) {
	res, err1 := T.N.S(v)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) B(v interface{}) (res []byte) {
	res, err1 := T.N.B(v)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) SxS(s string) (res string) {
	res, err1 := T.N.SxS(s)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) BxB(data []byte) (res []byte) {
	res, err1 := T.N.BxB(data)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) SxB(data []byte) (res string) {
	res, err1 := T.N.SxB(data)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) BxS(data string) (res []byte) {
	res, err1 := T.N.BxS(data)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) B2S(data []byte) (res string) {
	res, err1 := T.N.B2S(data)
	sure.Omit(err1)
	return res
}
func (T *Compactjson_Omit) S2B(data string) (res []byte) {
	res, err1 := T.N.S2B(data)
	sure.Omit(err1)
	return res
}
