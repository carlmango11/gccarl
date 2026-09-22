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
	vars map[semantic.VarID]*Var
	size semantic.Size
}

func newStackVars() *StackVars {
	return &StackVars{
		vars: make(map[semantic.VarID]*Var),
	}
}

func (lv *StackVars) Add(size semantic.Size) Offset {
	offset := Offset(lv.size + size)

	lv.size += size

	return offset
}

func (lv *StackVars) AddNamed(name semantic.VarID, size semantic.Size) Offset {
	offset := Offset(lv.size + size)

	lv.vars[name] = &Var{
		address: Address{
			stack: offset,
		},
	}

	lv.size += size

	return offset
}

func (lv *StackVars) AddLabelled(name semantic.VarID, label DataLabel) {
	lv.vars[name] = &Var{
		address: Address{
			label: label,
		},
	}
}

func (lv *StackVars) Offset(id semantic.VarID) (Offset, bool) {
	v, ok := lv.vars[id]
	if !ok {
		return 0, false
	}

	return v.address.stack, true
}

func (lv *StackVars) Address(id semantic.VarID) (Address, bool) {
	v, ok := lv.vars[id]
	if !ok {
		return Address{}, false
	}

	return v.address, true
}

func (lv *StackVars) Size() semantic.Size {
	return lv.size
}

func fieldNameOffset(t semantic.Type, f semantic.FieldName) Offset {
	if t.Kind != semantic.KindStruct {
		panic(fmt.Sprintf("reading %v from %v", f, t.Kind))
	}

	var offset Offset
	for _, field := range t.Struct.Fields {
		if field.Name == f {
			break
		}

		offset += Offset(field.Type.Size())
	}

	return offset
}
