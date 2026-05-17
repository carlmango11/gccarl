package compiler

import (
	"fmt"

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

func (c *Compiler) compileFuncDef(f *semantic.FuncDef) error {
	c.addInstr("%s:", f.Name)

	c.inFunc = true
	defer func() {
		c.inFunc = false
	}()

	c.addInstr("push rbp")
	c.addInstr("mov rbp, rsp")

	var stackSize int
	for _, p := range f.Params {
		switch p.Type.Kind {
		case semantic.KindPrimitive:
			stackSize += p.Type.Prim.Size()
		default:
			panic(fmt.Sprintf("unknown type %v", p.Type))
		}
	}

	locals := NewVars()

	for name, typ := range f.Locals {
		switch typ.Kind {
		case semantic.KindPrimitive:
			locals.Add(name, typ.Prim.Size())
			stackSize += typ.Prim.Size()
		case semantic.KindArray:
			locals.Add(name, typ.Prim.Size())
			stackSize += typ.Prim.Size() // TODO wrong
		default:
			panic(fmt.Sprintf("unknown type %v", typ))
		}
	}

	if stackSize > 0 {
		c.addInstr("sub rsp, %d", stackSize)
	}

	err := c.handleParamsDef(f.Params, locals)
	if err != nil {
		return err
	}

	for _, s := range f.Statements {
		err := c.compileStatement(s, locals)
		if err != nil {
			return err
		}
	}

	if f.ReturnExpr != nil {
		_, _, err := c.compileExpr(f.ReturnExpr, locals)
		if err != nil {
			return err
		}
	}

	c.addInstr("mov rsp, rbp")
	c.addInstr("pop rbp")
	c.addInstr("ret")

	return nil
}

var paramRegisters = []Register{
	RegRDI, RegRSI, RegRDX,
}

var intParamRegisters = []Register{
	RegEDI, RegESI, RegEDX,
}

func (c *Compiler) handleIntParam(n int, name semantic.VarName, locals *Vars) {
	offset := locals.Add(name, Size64)

	if n >= len(intParamRegisters) {
		panic("not implemented")
	}

	reg := intParamRegisters[n]

	c.addInstr("mov qword [rbp-%d], %s", offset, reg)
}

func (c *Compiler) handleParamsDef(ps []*semantic.ParamDef, locals *Vars) error {
	var standardC int

	for _, p := range ps {
		if p.Type.IsIntParamType() {
			c.handleIntParam(standardC, p.Name, locals)
			standardC++
		}
	}

	return nil
}

func (c *Compiler) functionCall(fc *semantic.FuncCall, locals *Vars) error {
	for i, expr := range fc.Args {
		resultReg, _, err := c.compileExpr(expr, locals)
		if err != nil {
			return err
		}

		if i >= len(paramRegisters) {
			panic("not implemented")
		}

		argReg := paramRegisters[i]

		c.addInstr("mov %s, %s", argReg, resultReg)
	}

	c.addInstr("call %s", fc.Func)
	return nil
}
