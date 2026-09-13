package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type Offset int

type Var struct {
	address Address
}

type Address struct {
	stack Offset
	label DataLabel
}

func (a Address) IsStack() bool {
	return a.label == ""
}

type StackVars struct {
	vars map[semantic.VarName]*Var
	size semantic.Size
}

func newStackVars() *StackVars {
	return &StackVars{
		vars: make(map[semantic.VarName]*Var),
	}
}

func (lv *StackVars) Add(size semantic.Size) Offset {
	offset := Offset(lv.size + size)

	lv.size += size

	return offset
}

func (lv *StackVars) AddNamed(name semantic.VarName, size semantic.Size) Offset {
	offset := Offset(lv.size + size)

	lv.vars[name] = &Var{
		address: Address{
			stack: offset,
		},
	}

	lv.size += size

	return offset
}

func (lv *StackVars) AddLabelled(name semantic.VarName, label DataLabel) {
	lv.vars[name] = &Var{
		address: Address{
			label: label,
		},
	}
}

func (lv *StackVars) Offset(id semantic.VarName) (Offset, bool) {
	v, ok := lv.vars[id]
	if !ok {
		return 0, false
	}

	return v.address.stack, true
}

func (lv *StackVars) Address(id semantic.VarName) (Address, bool) {
	v, ok := lv.vars[id]
	if !ok {
		return Address{}, false
	}

	return v.address, true
}

func (lv *StackVars) Size() semantic.Size {
	return lv.size
}

func fieldNameOffset(t semantic.Type, fs []semantic.VarRead) Offset {
	if len(fs) == 0 {
		return 0
	}

	f := fs[0]

	if t.Kind != semantic.KindStruct {
		panic(fmt.Sprintf("reading %v from %v", f, t.Kind))
	}

	var fieldType semantic.Type

	var offset Offset
	for _, field := range t.Struct.Fields {
		if field.Name == f.Name {
			fieldType = field.Type
			break
		}

		offset += Offset(field.Type.Size())
	}

	offset += indexOffset(fieldType, f)

	return offset + fieldNameOffset(fieldType, fs[1:])
}

func indexOffset(t semantic.Type, f semantic.VarRead) Offset {
	var offset Offset

	for _, i := range f.Index {
		t = *t.SubType
		offset += Offset(i) * Offset(t.Size())
	}

	return offset
}
