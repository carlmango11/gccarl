package compiler

import (
	"github.com/carlmango11/gccarl/gccarl/ast"
	"github.com/carlmango11/gccarl/gccarl/parser"
)

type Offset int

type Var struct {
	name   parser.Identifier
	typ    ast.Type
	offset Offset
}

type LocalVars struct {
	vars []*Var
}

func (lv *LocalVars) Add(name parser.Identifier, t ast.Type) Offset {
	var current Offset

	if len(lv.vars) > 0 {
		current = lv.vars[len(lv.vars)-1].offset
	}

	offset := current + Offset(typeSize(t))

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
		total += typeSize(v.typ)
	}

	return total
}
