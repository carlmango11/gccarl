package compiler

import (
	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type Offset int

type Var struct {
	name   semantic.VarName
	offset Offset
}

type Vars struct {
	vars map[semantic.VarName]*Var
	size int
}

func NewVars() *Vars {
	return &Vars{
		vars: make(map[semantic.VarName]*Var),
	}
}

func (lv *Vars) Add(name semantic.VarName, size int) Offset {
	offset := Offset(lv.size + size)

	lv.vars[name] = &Var{
		name:   name,
		offset: offset,
	}

	lv.size += size

	return offset
}

func (lv *Vars) Offset(id semantic.VarName) (Offset, bool) {
	v, ok := lv.vars[id]
	if !ok {
		return 0, false
	}

	return v.offset, true
}

//func (lv *Vars) ArrayOffset(id semantic.VarName, idx int) (Offset, bool) {
//	v, ok := lv.vars[id]
//	if !ok {
//		return 0, false
//	}
//
//	size := typeSize(v.typ)
//	arrayOffset := size * idx
//	offset := v.offset + Offset(arrayOffset)
//
//	return offset, true
//}

func (lv *Vars) Size() int {
	return lv.size
}
