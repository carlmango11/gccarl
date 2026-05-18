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

var callPrint = []Instr{
	"\tmov rdi, 1",
	"\tlea rsi, [rel msg]",
	"\tmov rdx, msg_size",
	"\tcall print",
}

var printRoutine = []Instr{
	"print:",
	"\tpush rbp",
	"\tmov rbp, rsp",
	"\tmov rax, 1",
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

	funcInstrs.addInstr("%s:", f.Name)

	c.inFunc = true
	defer func() {
		c.inFunc = false
	}()

	funcInstrs.addInstr("push rbp")
	funcInstrs.addInstr("mov rbp, rsp")

	locals := newStackVars()
	body := &Instrs{}

	err := c.handleParamsDef(body, f.Params, locals)
	if err != nil {
		return nil, err
	}

	for _, s := range f.Statements {
		err := c.compileStatement(body, s, locals)
		if err != nil {
			return nil, err
		}
	}

	if f.ReturnExpr != nil {
		_, err := c.compileExpr(body, f.ReturnExpr, locals)
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

var paramRegisters = []Register{
	RegRDI, RegRSI, RegRDX,
}

var intParamRegisters = []Register{
	RegEDI, RegESI, RegEDX,
}

func (c *Compiler) handleIntParam(instrs *Instrs, n int, name semantic.VarName, locals *StackVars) {
	offset := locals.AddNamed(name, Size64)

	if n >= len(paramRegisters) {
		panic("not implemented")
	}

	reg := paramRegisters[n]

	instrs.addInstr("mov qword [rbp-%d], %s", offset, reg)
}

func (c *Compiler) handleParamsDef(instrs *Instrs, ps []*semantic.ParamDef, locals *StackVars) error {
	var standardC int

	for _, p := range ps {
		if p.Type.IsIntParamType() {
			c.handleIntParam(instrs, standardC, p.Name, locals)
			standardC++
		}
	}

	return nil
}

func (c *Compiler) functionCall(instrs *Instrs, fc *semantic.FuncCall, locals *StackVars) error {
	paramOffsets := make([]Offset, len(fc.Args))

	for i, expr := range fc.Args {
		resultReg, err := c.compileExpr(instrs, expr, locals)
		if err != nil {
			return err
		}

		offset := locals.Add(8) // TODO: other sizes
		paramOffsets[i] = offset

		instrs.addInstr("mov [rbp-%d], %s", offset, resultReg)
	}

	for i := range len(fc.Args) {
		if i >= len(paramRegisters) {
			panic("not implemented")
		}

		argReg := paramRegisters[i]

		instrs.addInstr("mov %s, [rbp-%d]", argReg, paramOffsets[i])
	}

	instrs.addInstr("call %s", fc.Func)
	return nil
}
