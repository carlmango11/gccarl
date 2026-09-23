package compiler

import (
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
	c.locals = newStackVars()

	funcInstrs := &Instrs{}
	funcInstrs.addInstr("push rbp")
	funcInstrs.addInstr("mov rbp, rsp")

	body := &Instrs{}

	c.handleParamsDef(body, f.Params)

	for _, l := range f.Statements {
		err := c.compileStatement(body, l)
		if err != nil {
			return nil, err
		}
	}

	stackSize := c.locals.Size()
	if stackSize > 0 {
		funcInstrs.addInstr("sub rsp, %d", stackSize)
	}

	funcInstrs.instrs = append(funcInstrs.instrs, body.instrs...)

	funcInstrs.addInstr("mov rsp, rbp")
	funcInstrs.addInstr("pop rbp")
	funcInstrs.addInstr("ret")

	return funcInstrs, nil
}

func (c *Compiler) handleParamsDef(instrs *Instrs, ps []semantic.Var) {
	for i, p := range ps {
		instrs.addComment("receive %s", p.Name)

		offset := c.varOffset(p)
		instrs.movFromReg(p.Type.Size(), paramReg[i], offset)
	}
}

func (c *Compiler) functionCall(instrs *Instrs, fc *semantic.FuncCall) (Location, error) {
	paramOffsets := make([]Offset, len(fc.Args))

	for i, expr := range fc.Args {
		loc, err := c.compileExpr(instrs, expr)
		if err != nil {
			return Location{}, err
		}

		switch loc.Type {
		case LTRegister:
			offset := c.locals.Add(expr.Type.Size())
			instrs.movFromReg(expr.Type.Size(), loc.Register, offset)

			paramOffsets[i] = offset
		case LTOffset:
			paramOffsets[i] = loc.Offset
		case LTLabel:
			offset := c.locals.Add(8) // pointer
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
