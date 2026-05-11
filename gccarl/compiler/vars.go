package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/parser"
)

type Type int

const (
	TypeVoid Type = iota
	TypeInt
)

func (t Type) Size() int {
	switch t {
	case TypeInt:
		return 4
	default:
		panic(fmt.Sprintf("unknown type %d", t))
	}
}

type Offset int

type Var struct {
	name   parser.Identifier
	typ    Type
	offset Offset
}

type LocalVars struct {
	vars []*Var
}

func (lv *LocalVars) Add(name parser.Identifier, t Type) Offset {
	var current Offset

	if len(lv.vars) > 0 {
		current = lv.vars[len(lv.vars)-1].offset
	}

	offset := current + Offset(t.Size())

	lv.vars = append(lv.vars, &Var{
		name:   name,
		typ:    t,
		offset: offset,
	})

	return offset
}

func (lv *LocalVars) Offset(i parser.Identifier) (Offset, bool) {
	for _, v := range lv.vars {
		if v.name == i {
			return v.offset, true
		}
	}

	return 0, false
}

func (lv *LocalVars) Size() int {
	var total int
	for _, v := range lv.vars {
		total += v.typ.Size()
	}

	return total
}
