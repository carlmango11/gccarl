package compiler

import (
	"iter"
	"sort"

	"github.com/carlmango11/gccarl/gccarl/semantic"
)

var assertRoutine = []Instr{
	"assert:",
	"\tpush rbp",
	"\tmov rbp, rsp",
	"\tmov rax, 9",
	"\tsyscall",
	"\tpop rbp",
	"\tret",
}

func (c *Compiler) compileFuncDef(f *semantic.FuncDef) (*Instrs, error) {
	funcInstrs := &Instrs{}
	funcInstrs.addInstr("push rbp")
	funcInstrs.addInstr("mov rbp, rsp")

	locals := newStackVars()
	body := &Instrs{}

	for name, typ := range sortParams(f.Locals) {
		offset := locals.AddNamed(name, typ.Size())
		body.addComment("var %s = %d", name, offset)
	}

	c.handleParamsDef(body, f.Params, locals)

	for _, l := range f.Lines {
		err := c.compileLine(body, l, locals)
		if err != nil {
			return nil, err
		}
	}

	stackSize := locals.Size()
	if stackSize > 0 {
		funcInstrs.addInstr("sub rsp, %d", stackSize)
	}

	funcInstrs.instrs = append(funcInstrs.instrs, body.instrs...)

	funcInstrs.addInstr("mov rsp, rbp")
	funcInstrs.addInstr("pop rbp")
	funcInstrs.addInstr("ret")

	return funcInstrs, nil
}

func sortParams(locals map[semantic.VarName]semantic.Type) iter.Seq2[semantic.VarName, semantic.Type] {
	vals := make([]string, 0, len(locals))
	for name := range locals {
		vals = append(vals, string(name))
	}
	sort.Strings(vals)

	return func(yield func(semantic.VarName, semantic.Type) bool) {
		for _, val := range vals {
			nextKey := semantic.VarName(val)
			nextVal := locals[nextKey]

			if !yield(nextKey, nextVal) {
				return
			}
		}
	}
}

func (c *Compiler) compileControl(instrs *Instrs, control *semantic.Control, locals *StackVars) error {
	switch {
	case control.If != nil:
		return c.compileIf(instrs, control.If, locals)
	case control.While != nil:
		return c.compileWhile(instrs, control.While, locals)
	}

	panic("invalid control")
}

func (c *Compiler) handleParamsDef(instrs *Instrs, ps []*semantic.ParamDef, locals *StackVars) {
	for i, p := range ps {
		instrs.addComment("receive %s", p.Name)

		offset := locals.AddNamed(p.Name, p.Type.Size())
		instrs.movFromReg(p.Type.Size(), paramReg[i], offset)
	}
}

func (c *Compiler) functionCall(instrs *Instrs, fc *semantic.FuncCall, locals *StackVars) (Location, error) {
	paramOffsets := make([]Offset, len(fc.Args))

	for i, expr := range fc.Args {
		loc, err := c.compileExpr(instrs, expr, locals)
		if err != nil {
			return Location{}, err
		}

		switch loc.Type {
		case LTRegister:
			offset := locals.Add(expr.Type.Size())
			instrs.movFromReg(expr.Type.Size(), loc.Register, offset)

			paramOffsets[i] = offset
		case LTOffset:
			paramOffsets[i] = loc.Offset
		case LTLabel:
			offset := locals.Add(8) // pointer
			instrs.movLocToReg(8, loc, RegA)
			instrs.movFromReg(8, RegA, offset)

			paramOffsets[i] = offset
		}
	}

	for i, expr := range fc.Args {
		instrs.addComment("move arg %d for call to %s", i, fc.Func)

		if expr.Type.Kind == semantic.KindArray {
			instrs.addInstr("lea %s, [rbp-%d]", paramReg[i].Raw(8), paramOffsets[i])
			continue
		}

		size := fc.Args[i].Type.Size()
		instrs.movOffsetToReg(size, paramOffsets[i], paramReg[i])
	}

	instrs.addInstr("call %s", fc.Func)

	retType := c.funcs[fc.Func].ReturnType
	if retType.Kind == semantic.KindVoid {
		return Location{}, nil
	}

	return regLocation(RegA), nil
}

func regLocation(reg Register) Location {
	return Location{
		Type:     LTRegister,
		Register: reg,
	}
}
