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

type Compiler struct {
	globalVars *Vars
	instrs     []Instr
	inFunc     bool
}

func New() *Compiler {
	return &Compiler{
		globalVars: NewVars(),
	}
}

func (c *Compiler) Compile(prog *semantic.Program) ([]byte, error) {
	err := c.compile(prog)
	if err != nil {
		return nil, err
	}

	var output []byte
	for _, instr := range c.instrs {
		output = append(output, instr...)
		output = append(output, '\n')
	}

	return output, nil
}

func (c *Compiler) compile(prog *semantic.Program) error {
	return c.compileMain(prog)
}

func (c *Compiler) compileMain(prog *semantic.Program) error {
	c.addInstr("section .text")
	c.addInstr("global _start")

	c.addInstr("_start:")

	c.addInstr("\tcall main")
	//c.instrs = append(c.instrs, callPrint...)
	c.addInstr("\tcall exit")
	c.instrs = append(c.instrs, exitRoutine...)
	c.instrs = append(c.instrs, printRoutine...)

	for _, fd := range prog.FuncDefs {
		err := c.compileFuncDef(fd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) compileStatement(s *semantic.Statement, locals *Vars) error {
	switch {
	case s.Assign != nil:
		return c.compileAssign(s.Assign, locals)
	case s.Expr != nil:
		_, _, err := c.compileExpr(s.Expr, locals)
		return err
	}

	panic("missing statement type")
}

func (c *Compiler) compileAssign(a *semantic.Assign, locals *Vars) error {
	toOffset, ok := locals.Offset(a.Name)
	if !ok {
		return fmt.Errorf("undefined variable: %s", a.Name)
	}

	outputReg, _, err := c.compileExpr(a.Expr, locals)
	if err != nil {
		return err
	}

	if outputReg != RegUnset {
		c.addInstr(`mov qword [rbp-%d], %s`, toOffset, outputReg)
	} else {
		c.addInstr("mov qword [rbp-%d], %s", toOffset, RegRAX)
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

func (c *Compiler) compileArray(a *semantic.Array, locals *Vars) error {
	if len(a.Vals) == 0 {
		panic("handle")
	}

	size := a.Vals[0].Type.Prim.Size() * len(a.Vals)
	startOffset := locals.Add("", size)

	for i, v := range a.Vals {
		reg, _, err := c.compileExpr(v, locals)
		if err != nil {
			return err
		}

		if reg != RegUnset {
			offset := startOffset - Offset(a.Vals[0].Type.Size()*i)
			c.addInstr("mov qword [rbp-%d], %s", offset, reg)
		} else {
			panic("impl")
		}
	}

	c.addInstr("lea %s, [rbp-%d]", RegRAX, startOffset)
	return nil
}

func (c *Compiler) compileExpr(e *semantic.Expr, locals *Vars) (Register, Offset, error) {
	switch e.Type.Kind {
	//case e.Add != nil:
	//	return c.compileAdd(e.Add, locals, target)
	//case e.Val != nil:
	//	return c.compileVal(e.Val, locals, target)
	//case e.FuncCall != nil:
	//	return functionCall(e.FuncCall, locals)
	case semantic.KindPrimitive:
		switch e.Type.Prim {
		case semantic.PrimInt32:
			c.addInstr("mov %s, %s", RegRAX, e.Literal.Int32)
			return RegRAX, 0, nil
		case semantic.PrimChar:
			c.addInstr("mov %s, 0x%X", RegRAX, e.Literal.Char)
			return RegRAX, 0, nil
		}
	case semantic.KindArray:
		// RAX
		err := c.compileArray(e.Array, locals)
		if err != nil {
			return "", 0, err
		}

		return RegUnset, 0, nil
	}

	return "", 0, fmt.Errorf("unknown expr type: %+v", e)
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

func (c *Compiler) addInstr(format string, args ...any) {
	if c.inFunc {
		format = "\t" + format
	}

	i := fmt.Sprintf(format, args...)
	c.instrs = append(c.instrs, Instr(i))
}
