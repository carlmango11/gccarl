package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type Size int

const (
	Size64 = 8
)

type Register string

const (
	RegUnset Register = ""
	RegEAX   Register = "eax"
	RegEBX   Register = "ebx"
	RegEDI   Register = "edi"
	RegESI   Register = "esi"
	RegEDX   Register = "edx"
	RegRAX   Register = "rax"
	RegRDI   Register = "rdi"
	RegRSI   Register = "rsi"
	RegRDX   Register = "rdx"
)

type Instr string

type FuncDef struct {
	ReturnType semantic.Type
	Params     []*semantic.ParamDef
}

type Compiler struct {
	funcs  map[semantic.FuncName]*FuncDef
	inFunc bool
}

func New() *Compiler {
	return &Compiler{
		funcs: make(map[semantic.FuncName]*FuncDef),
	}
}

func (c *Compiler) Compile(prog *semantic.Program) ([]byte, error) {
	instrs, err := c.compile(prog)
	if err != nil {
		return nil, err
	}

	var output []byte
	for _, instr := range instrs.instrs {
		output = append(output, instr...)
		output = append(output, '\n')
	}

	return output, nil
}

func (c *Compiler) compile(prog *semantic.Program) (*Instrs, error) {
	full := &Instrs{}
	full.addInstr("section .text")
	full.addInstr("global _start")

	full.addInstr("_start:")

	full.addInstr("\tcall main")
	//c.instrs = append(c.instrs, callPrint...)
	full.addInstr("\tcall exit")
	full.instrs = append(full.instrs, exitRoutine...)
	full.instrs = append(full.instrs, printRoutine...)

	for _, fd := range prog.FuncDefs {
		c.funcs[fd.Name] = &FuncDef{
			ReturnType: fd.ReturnType,
			Params:     fd.Params,
		}

		funcInstrs, err := c.compileFuncDef(fd)
		if err != nil {
			return nil, err
		}

		full.addInstr("%s:", fd.Name)
		full.addInstrsIndent(funcInstrs)
	}

	return full, nil
}

func (c *Compiler) compileStatement(instrs *Instrs, s *semantic.Statement, locals *StackVars) error {
	switch {
	case s.Assign != nil:
		return c.compileAssign(instrs, s.Assign, locals)
	case s.ArrayAssign != nil:
		return c.compileArrayAssign(instrs, s.ArrayAssign, locals)
	case s.Expr != nil:
		_, err := c.compileExpr(instrs, s.Expr, locals)
		return err
	}

	panic("missing statement type")
}

func (c *Compiler) compileArrayAssign(instrs *Instrs, a *semantic.ArrayAssign, locals *StackVars) error {
	startOffset := locals.AddNamed(a.Name, a.Type.Size())

	for i, v := range a.Vals {
		reg, err := c.compileExpr(instrs, v, locals)
		if err != nil {
			return err
		}

		offset := startOffset - Offset(a.Vals[0].Type.Size()*i)
		instrs.addInstr("mov qword [rbp-%d], %s", offset, reg) // TODO needs to change for floats
	}

	return nil
}

func (c *Compiler) compileAssign(instrs *Instrs, a *semantic.Assign, locals *StackVars) error {
	toOffset, ok := locals.Offset(a.Name)
	if !ok {
		return fmt.Errorf("undefined variable: %s", a.Name)
	}

	outputReg, err := c.compileExpr(instrs, a.Expr, locals)
	if err != nil {
		return err
	}

	if outputReg != RegUnset {
		instrs.addInstr(`mov qword [rbp-%d], %s`, toOffset, outputReg)
	} else {
		instrs.addInstr("mov qword [rbp-%d], %s", toOffset, RegRAX)
	}

	return nil
}

//func (c *Compiler) compileVal(v *semantic.Value, locals *Vars, target Register) (ast.RawType, error) {
//	switch {
//	case v.Int != nil:
//		return c.addInstr(nil, "mov %s, %d", target, *v.Int), ast.TypeInt, nil
//	case v.Char != nil:
//		return c.addInstr(nil, "mov %s, 0x%X", target, *v.Char), ast.TypeChar, nil
//	case v.Var != nil:
//		var idx int
//		if v.Var.Index != nil {
//			idx = *v.Var.Index
//		}
//
//		offset, ok := locals.ArrayOffset(v.Var.Name, idx)
//		if !ok {
//			return nil, fmt.Errorf("undefined variable: %s", v.Var.Name)
//		}
//
//		return c.addInstr(nil, "mov %s, [rbp-%d]", target, offset), nil
//	case v.Array != nil:
//		// TODO: types
//		return c.compileArray(v.Array, ast.TypeInt, locals, target)
//	}
//
//	panic("missing value type")
//}

func (c *Compiler) compileExpr(instrs *Instrs, e *semantic.Expr, locals *StackVars) (Register, error) {
	switch {
	case e.FuncCall != nil:
		err := c.functionCall(instrs, e.FuncCall, locals)
		if err != nil {
			return RegUnset, err
		}

		return returnRegister(e.Type), nil
	case e.AddressOf != "":
		offset, ok := locals.Offset(e.AddressOf)
		if !ok {
			return RegUnset, fmt.Errorf("undefined variable: %s", e.Var)
		}

		instrs.addInstr("lea %s, [rbp-%d]", RegRAX, offset)
		return RegRAX, nil

	case e.Literal != nil:
		switch e.Type.Kind {
		//case e.Add != nil:
		//	return c.compileAdd(e.Add, locals, target)
		//case e.Val != nil:
		//	return c.compileVal(e.Val, locals, target)
		case semantic.KindPrimitive:
			switch e.Type.Prim {
			case semantic.PrimInt32:
				instrs.addInstr("mov %s, %d", RegRAX, e.Literal.Int32)
				return RegRAX, nil
			case semantic.PrimChar:
				instrs.addInstr("mov %s, 0x%X", RegRAX, e.Literal.Char)
				return RegRAX, nil
			}
		}
	case e.Var != "":
		offset, ok := locals.Offset(e.Var)
		if !ok {
			return RegUnset, fmt.Errorf("undefined variable: %s", e.Var)
		}

		instrs.addInstr("mov %s, [rbp-%d]", RegRAX, offset)
		return RegRAX, nil
	}

	return "", fmt.Errorf("unknown expr type: %+v", e)
}

//func (c *Compiler) compileAdd(a *semantic.AddExpr, locals *Vars, target Register) (ast.RawType, error) {
//	var all []Instr
//
//	output1, typ1, err := c.compileExpr(a.Expr2, locals, target)
//	if err != nil {
//		return nil, ast.TypeUnset, err
//	}
//
//	all = append(all, output1...)
//
//	output2, typ2, err := c.compileExpr(a.Expr1, locals, RegEBX)
//	if err != nil {
//		return nil, ast.TypeUnset, err
//	}
//
//	all = append(all, output2...)
//
//	return c.addInstr(all, "add %s, %s", target, RegEBX), nil
//}

//func (c *Compiler) compileVarDec(vd *ast.DecAssign, locals *Vars) error {
//	switch typeSize(vd.Type) {
//	case 1:
//		return compile1Assign(vd.Var, vd.Expr, locals)
//	case 8:
//		return compile8Assign(vd.Var, vd.Expr, locals)
//	default:
//		panic(fmt.Sprintf("unhandled assignment type: %s", vd.Type))
//	}
//}

type Instrs struct {
	instrs []Instr
}

func (i *Instrs) addInstr(format string, args ...any) {
	instr := fmt.Sprintf(format, args...)
	i.instrs = append(i.instrs, Instr(instr))
}

func (i *Instrs) addInstrsIndent(inner *Instrs) {
	for _, instr := range inner.instrs {
		i.instrs = append(i.instrs, "\t"+instr)
	}
}

func returnRegister(t semantic.Type) Register {
	switch t.Kind {
	case semantic.KindVoid:
		return RegUnset
	case semantic.KindPrimitive:
		switch t.Prim {
		case semantic.PrimInt32:
			return RegRAX
		}
	case semantic.KindArray:
		panic("array cannot be returned")
	}

	panic("unset type")
}
