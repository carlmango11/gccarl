package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/ast"
	"github.com/carlmango11/gccarl/gccarl/parser"
)

var exitRoutine = []Instr{
	"exit:",
	"push rbp",
	"mov rbp, rsp",
	"mov rax, 60",
	"syscall",
	"pop rbp",
	"ret",
}

var callPrint = []Instr{
	"mov rdi, 1",
	"lea rsi, [rel msg]",
	"mov rdx, msg_size",
	"call print",
}

var printRoutine = []Instr{
	"print:",
	"push rbp",
	"mov rbp, rsp",
	"mov rax, 1",
	"syscall",
	"pop rbp",
	"ret",
}

var data = []Instr{
	"section .data",
	`msg db "Hello, world!", 0xA`,
	"msg_size equ $ - msg",
}

func compileFuncDef(f *ast.FuncDef) ([]Instr, error) {
	funcInstrs := addInstr(nil, "%s:", f.Name)

	bodyInstrs := addInstr(nil, "push rbp")
	bodyInstrs = addInstr(bodyInstrs, "mov rbp, rsp")

	var paramsSize int
	for _, p := range f.Params {
		paramsSize += typeSize(p.Type)
	}

	if paramsSize > 0 {
		bodyInstrs = addInstr(bodyInstrs, "sub rsp, %d", paramsSize)
	}

	locals := &LocalVars{}

	paramsInstrs, err := handleParamsDef(f.Params, locals)
	if err != nil {
		return nil, err
	}

	bodyInstrs = append(bodyInstrs, paramsInstrs...)

	var allStatementInstrs []Instr
	for _, s := range f.Statements {
		statementInstrs, err := compileStatement(s, locals)
		if err != nil {
			return nil, err
		}

		for _, instr := range statementInstrs {
			allStatementInstrs = append(allStatementInstrs, instr)
		}
	}

	bodyInstrs = addInstr(bodyInstrs, "sub rsp, %d", locals.Size())

	bodyInstrs = append(bodyInstrs, allStatementInstrs...)

	if f.ReturnExpr != nil {
		returnInstrs, err := compileExpr(f.ReturnExpr, locals, RegEAX)
		if err != nil {
			return nil, err
		}

		bodyInstrs = append(bodyInstrs, returnInstrs...)
	}

	bodyInstrs = addInstr(bodyInstrs, "mov rsp, rbp")
	bodyInstrs = addInstr(bodyInstrs, "pop rbp")
	bodyInstrs = addInstr(bodyInstrs, "ret")

	for _, instr := range bodyInstrs {
		funcInstrs = append(funcInstrs, "\t"+instr)
	}

	return funcInstrs, nil
}

var paramRegisters = []Register{
	RegRDI, RegRSI, RegRDX,
}

var intParamRegisters = []Register{
	RegEDI, RegESI, RegEDX,
}

func handleIntParam(instrs []Instr, n int, name parser.Identifier, locals *LocalVars) ([]Instr, error) {
	offset := locals.Add(name, ast.TypeInt)

	if n >= len(intParamRegisters) {
		panic("not implemented")
	}

	reg := intParamRegisters[n]

	return addInstr(instrs, "mov dword [rbp-%d], %s", offset, reg), nil
}

func handleParamsDef(ps []*ast.ParamDef, locals *LocalVars) ([]Instr, error) {
	var instrs []Instr

	var standardC int
	var err error

	for _, p := range ps {
		switch p.Type {
		case ast.TypeInt:
			instrs, err = handleIntParam(instrs, standardC, p.Name, locals)
			if err != nil {
				return nil, err
			}

			standardC++
		default:
			return nil, fmt.Errorf("unknown type: %v", p.Type)
		}
	}

	return instrs, nil
}

func functionCall(fc *ast.FuncCall, locals *LocalVars) ([]Instr, error) {
	var instrs []Instr

	for i, expr := range fc.Params {
		if i >= len(intParamRegisters) {
			panic("not implemented")
		}

		reg := intParamRegisters[i]

		exprInstrs, err := compileExpr(expr, locals, reg)
		if err != nil {
			return nil, err
		}

		instrs = append(instrs, exprInstrs...)
	}

	instrs = addInstr(instrs, "call %s", fc.Name)
	return instrs, nil
}
