package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/ast"
)

type Register string

const (
	RegEAX Register = "eax"
	RegEBX Register = "ebx"
	RegEDI Register = "edi"
	RegESI Register = "esi"
	RegEDX Register = "edx"
	RegRDI Register = "rdi"
	RegRSI Register = "rsi"
	RegRDX Register = "rdx"
)

type Instr string

func Compile(prog *ast.Program) ([]byte, error) {
	instrs, err := compile(prog)
	if err != nil {
		return nil, err
	}

	var output []byte
	for _, instr := range instrs {
		output = append(output, instr...)
		output = append(output, '\n')
	}

	return output, nil
}

func compile(prog *ast.Program) ([]Instr, error) {
	return compileMain(prog)
}

func compileMain(prog *ast.Program) ([]Instr, error) {
	var instrs []Instr

	instrs = addInstr(instrs, "section .text")
	instrs = addInstr(instrs, "global _start")

	instrs = addInstr(instrs, "_start:")

	instrs = append(instrs, callPrint...)

	instrs = addInstr(instrs, "\tcall main")
	instrs = addInstr(instrs, "\tcall exit")
	instrs = append(instrs, exitRoutine...)
	instrs = append(instrs, printRoutine...)

	//locals := &LocalVars{}

	for _, fd := range prog.FuncDefs {
		funcInstrs, err := compileFuncDef(fd)
		if err != nil {
			return nil, err
		}

		instrs = append(instrs, funcInstrs...)
	}

	instrs = append(instrs, data...)

	return instrs, nil
}

func compileStatement(s *ast.Statement, locals *LocalVars) ([]Instr, error) {
	switch {
	case s.VarDec != nil:
		return compileDecAssign(s.VarDec, locals)
	case s.Assign != nil:
		return compileAssign(s.Assign, locals)
	case s.FuncCall != nil:
		return functionCall(s.FuncCall, locals)
	}

	panic("missing statement type")
}

func compileAssign(a *ast.Assign, locals *LocalVars) ([]Instr, error) {
	if a.Var.Index != nil {
		panic("impl")
	}

	toOffset, ok := locals.Offset(a.Var.Name)
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", a.Var.Name)
	}

	output, err := compileExpr(a.Expr, locals, RegEAX)
	if err != nil {
		return nil, err
	}

	return addInstr(output, `mov dword [rbp-%d], %s`, toOffset, RegEAX), nil
}

func loadValue(v *ast.Value, locals *LocalVars, target Register) ([]Instr, error) {
	if v.Var != nil {
		if v.Var.Index != nil {
			panic("impl")
		}

		offset, ok := locals.Offset(v.Var.Name)
		if !ok {
			return nil, fmt.Errorf("undefined variable: %s", v.Var.Name)
		}

		return addInstr(nil, "mov %s, [rbp-%d]", target, offset), nil
	}

	return addInstr(nil, "mov %s, %d", target, v.Int), nil
}

func compileExpr(e *ast.Expr, locals *LocalVars, target Register) ([]Instr, error) {
	switch {
	case e.Add != nil:
		return compileAdd(e.Add, locals, target)
	case e.Val != nil:
		return loadValue(e.Val, locals, target)
	default:
		return nil, fmt.Errorf("unknown expr type: %+v", e)
	}
}

func compileAdd(a *ast.AddExpr, locals *LocalVars, target Register) ([]Instr, error) {
	var all []Instr

	output, err := compileExpr(a.Expr, locals, target)
	if err != nil {
		return nil, err
	}

	all = append(all, output...)

	output, err = loadValue(a.Val, locals, RegEBX)
	if err != nil {
		return nil, err
	}

	all = append(all, output...)

	return addInstr(all, "add %s, %s", target, RegEBX), nil
}

func compileDecAssign(vd *ast.VarDec, locals *LocalVars) ([]Instr, error) {
	switch vd.Type {
	case ast.TypeInt:
		return compileIntAssignment(vd.Var, vd.Expr, locals)
	default:
		panic(fmt.Sprintf("unhandled assignment type: %s", vd.Type))
	}
}

func compileIntAssignment(v *ast.Var, e *ast.Expr, locals *LocalVars) ([]Instr, error) {
	if v.Index != nil {
		panic("unhandled index")
	}

	instrs, err := compileExpr(e, locals, RegEAX)
	if err != nil {
		return nil, err
	}

	offset := locals.Add(v.Name, ast.TypeInt)

	return addInstr(instrs, "mov dword [rbp-%d], %s", offset, RegEAX), nil
}

func addInstr(instrs []Instr, format string, args ...any) []Instr {
	i := fmt.Sprintf(format, args...)
	return append(instrs, Instr(i))
}

func typeSize(t ast.Type) int {
	switch t {
	case ast.TypeInt:
		return 4
	default:
		panic(fmt.Sprintf("unknown type %d", t))
	}
}
