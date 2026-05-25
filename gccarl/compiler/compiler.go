package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type Size int

const (
	Size8  = 1
	Size32 = 4
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
	RegR10   Register = "r10"
	RegAL    Register = "al"
)

func raxSized(s Size) Register {
	switch s {
	case Size64:
		return RegRAX
	case Size8:
		return RegAL
	case Size32:
		return RegEAX
	default:
		panic(fmt.Sprintf("raxSized: unknown size: %d", s))
	}
}

type Instr string
type DataLabel string

type FuncDef struct {
	ReturnType semantic.Type
	Params     []*semantic.ParamDef
}

type Compiler struct {
	funcs  map[semantic.FuncName]*FuncDef
	labelC int
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

	c.addDataSection(prog, full)

	full.addInstr("section .text")
	full.addInstr("global _start")

	full.addInstr("_start:")

	full.addInstr("\tcall main")
	full.addInstr("\tcall exit")

	full.instrs = append(full.instrs, exitRoutine...)

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

func (c *Compiler) addDataSection(prog *semantic.Program, full *Instrs) {
	full.addInstr("section .data")

	for id, val := range prog.Strings {
		label := DataLabel(fmt.Sprintf("str_%d", id+1))
		full.addInstr(`%s db "%s", 0xA`, label, val)
	}
}

func (c *Compiler) compileStatement(instrs *Instrs, s *semantic.Statement, locals *StackVars) error {
	switch {
	case s.If != nil:
		return c.compileIf(instrs, s.If, locals)
	case s.Assign != nil:
		return c.compileAssign(instrs, s.Assign, locals)
	case s.Expr != nil:
		_, err := c.compileExpr(instrs, s.Expr, locals)
		return err
	}

	panic("missing statement type")
}

func (c *Compiler) compileAssign(instrs *Instrs, a *semantic.Assign, locals *StackVars) error {
	switch a.Expr.Type.Kind {
	case semantic.KindArray:
		return c.compileArrayAssign(instrs, a, locals)
	default:
		return c.compileStandardAssign(instrs, a, locals)
	}
}

func (c *Compiler) compileArrayAssign(instrs *Instrs, a *semantic.Assign, locals *StackVars) error {
	switch {
	case a.Expr.StringID != 0:
		locals.AddLabelled(a.Name, c.stringLabel(a.Expr.StringID))
	default:
		startOffset, ok := locals.Offset(a.Name)
		if !ok {
			return fmt.Errorf("undefined variable %s", a.Name)
		}

		for i, v := range a.Expr.CompLiteral {
			reg, err := c.compileExpr(instrs, v, locals)
			if err != nil {
				return err
			}

			offset := startOffset - Offset(v.Type.Size()*i)
			instrs.addInstr("mov %s [rbp-%d], %s", typeInstrSize(v.Type), offset, reg) // TODO needs to change for floats
		}
	}

	return nil
}

func typeInstrSize(t semantic.Type) string {
	switch t.Kind {
	case semantic.KindPrimitive:
		switch t.Prim {
		case semantic.PrimChar:
			return "byte"
		}
	}

	return "qword"
}

func (c *Compiler) compileStandardAssign(instrs *Instrs, a *semantic.Assign, locals *StackVars) error {
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

func (c *Compiler) compileExpr(instrs *Instrs, e *semantic.Expr, locals *StackVars) (Register, error) {
	switch {
	case e.IsEqual != nil:
		err := c.compileIsEqual(instrs, e.IsEqual, locals)
		return RegUnset, err
	case e.FuncCall != nil:
		err := c.functionCall(instrs, e.FuncCall, locals)
		if err != nil {
			return RegUnset, err
		}

		return returnRegister(e.Type), nil
	case e.AddressOf != "":
		addr, ok := locals.Address(e.AddressOf)
		if !ok {
			return RegUnset, fmt.Errorf("undefined variable: %s", e.Var)
		}

		if addr.IsStack() {
			instrs.addInstr("lea %s, [rbp-%d] ; addressOf", RegRAX, addr.stack)
		} else {
			instrs.addInstr("lea %s, [rel %s] ; addressOf", RegRAX, addr.label)
		}

		return RegRAX, nil

	case e.StringID != 0:
		instrs.addInstr("lea %s, [rel %s]", RegRAX, c.stringLabel(e.StringID))
		return RegRAX, nil

	case e.Literal != nil:
		switch e.Type.Kind {
		case semantic.KindPrimitive:
			switch e.Type.Prim {
			case semantic.PrimInt32:
				instrs.addInstr("mov %s, %d", RegRAX, e.Literal.Int32)
				return RegRAX, nil
			case semantic.PrimChar:
				instrs.addInstr("mov byte %s, 0x%X", RegAL, e.Literal.Char)
				return RegAL, nil
			}
		}
	case e.Var != "":
		addr, ok := locals.Address(e.Var)
		if !ok {
			return RegUnset, fmt.Errorf("undefined variable: %s", e.Var)
		}

		instrs.moveVarRAX(e.Type, addr)
		return RegRAX, nil
	}

	panic(fmt.Sprintf("unknown expr type: %+v", e))
}

func (c *Compiler) stringLabel(id semantic.StringID) DataLabel {
	return DataLabel(fmt.Sprintf("str_%d", id))
}

func (c *Compiler) compileIf(instrs *Instrs, ifs *semantic.If, locals *StackVars) error {
	_, err := c.compileExpr(instrs, ifs.Condition, locals)
	if err != nil {
		return err
	}

	skip := c.newLabel("skip")
	instrs.addInstr("jne %s", skip)

	for _, s := range ifs.Statements {
		err := c.compileStatement(instrs, s, locals)
		if err != nil {
			return err
		}
	}

	instrs.addInstr("%v:", skip)

	return nil
}

func (c *Compiler) newLabel(prefix string) string {
	c.labelC++
	return fmt.Sprintf("%v_%d", prefix, c.labelC)
}

func (c *Compiler) compileIsEqual(instrs *Instrs, e *semantic.IsEqual, locals *StackVars) error {
	reg, err := c.compileExpr(instrs, e.Left, locals)
	if err != nil {
		return err
	}

	offset := locals.Add(8)
	instrs.addInstr("mov %s [rbp-%d], %s", typeInstrSize(e.Left.Type), offset, reg)

	reg, err = c.compileExpr(instrs, e.Right, locals)
	if err != nil {
		return err
	}

	instrs.addInstr("cmp [rbp-%d], %s", offset, reg)
	return nil
}

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

func (i *Instrs) moveVarRAX(t semantic.Type, a Address) {
	if a.IsStack() {
		i.addInstr("mov %s, [rbp-%d]", raxSized(Size(t.Size())), a.stack) // todo shitty types
	} else {
		panic("unreachable")
	}
}

func (i *Instrs) addComment(format string, args ...any) {
	i.addInstr("; "+format, args...)
}

func returnRegister(t semantic.Type) Register {
	switch t.Kind {
	case semantic.KindVoid:
		return RegUnset
	case semantic.KindPrimitive:
		switch t.Prim {
		case semantic.PrimInt32, semantic.PrimChar:
			return RegRAX
		}
	case semantic.KindArray:
		panic("array cannot be returned")
	}

	panic("unset type")
}
