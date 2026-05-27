package compiler

import (
	"github.com/carlmango11/gccarl/gccarl/semantic"
)

var exitRoutine = []Instr{
	"exit:",
	"\tpush rbp",
	"\tmov rbp, rsp",
	"\tmov rax, 60",
	"\tsyscall",
	"\tpop rbp",
	"\tret",
}

var data = []Instr{
	"section .data",
	`msg db "Hello, world!", 0xA`,
	"msg_size equ $ - msg",
}

func (c *Compiler) compileFuncDef(f *semantic.FuncDef) (*Instrs, error) {
	funcInstrs := &Instrs{}
	funcInstrs.addInstr("push rbp")
	funcInstrs.addInstr("mov rbp, rsp")

	locals := newStackVars()
	body := &Instrs{}

	for name, typ := range f.Locals {
		locals.AddNamed(name, typ.Size())
	}

	c.handleParamsDef(body, f.Params, locals)

	for _, l := range f.Lines {
		if l.Statement != nil {
			err := c.compileStatement(body, l.Statement, locals)
			if err != nil {
				return nil, err
			}
		} else if l.Control != nil {
			err := c.compileControl(body, l.Control, locals)
			if err != nil {
				return nil, err
			}
		}
	}

	if f.ReturnExpr != nil {
		loc, err := c.compileExpr(body, f.ReturnExpr, locals)
		if err != nil {
			return nil, err
		}

		retReg := returnRegister(f.ReturnExpr.Type)
		body.movLocToReg(f.ReturnExpr.Type.Size(), loc, retReg)
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

func (c *Compiler) compileControl(instrs *Instrs, control *semantic.Control, locals *StackVars) error {
	switch {
	case control.If != nil:
		return c.compileIf(instrs, control.If, locals)
	}

	panic("invalid control")
}

func (c *Compiler) handleParamsDef(instrs *Instrs, ps []*semantic.ParamDef, locals *StackVars) {
	for i, p := range ps {
		offset := locals.AddNamed(p.Name, p.Type.Size())
		instrs.movFromReg(p.Type.Size(), paramReg[p.Type.Size()][i], offset)
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
		}
	}

	for i := range len(fc.Args) {
		instrs.addComment("move arg %d for call to %s", i, fc.Func)

		size := fc.Args[i].Type.Size()
		instrs.movOffsetToReg(size, paramOffsets[i], paramReg[size][i])
	}

	instrs.addInstr("call %s", fc.Func)

	retType := c.funcs[fc.Func].ReturnType
	if retType.Kind == semantic.KindVoid {
		return Location{}, nil
	}

	returnReg := returnRegister(retType)
	return regLocation(returnReg), nil
}

func regLocation(reg Register) Location {
	return Location{
		Type:     LTRegister,
		Register: reg,
	}
}
