package compiler

import (
	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type Offset int

type Var struct {
	offset Offset
}

type StackVars struct {
	vars map[semantic.VarName]*Var
	size int
}

func newStackVars() *StackVars {
	return &StackVars{
		vars: make(map[semantic.VarName]*Var),
	}
}

func (lv *StackVars) Add(size int) Offset {
	offset := Offset(lv.size + size)

	lv.size += size

	return offset
}

func (lv *StackVars) AddNamed(name semantic.VarName, size int) Offset {
	offset := Offset(lv.size + size)

	lv.vars[name] = &Var{
		offset: offset,
	}

	lv.size += size

	return offset
}

func (lv *StackVars) Offset(id semantic.VarName) (Offset, bool) {
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

func (lv *StackVars) Size() int {
	return lv.size
}
