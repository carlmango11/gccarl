package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type LocationType int

const (
	LTUnset LocationType = iota
	LTOffset
	LTRegister
	LTLabel
)

type Location struct {
	Type     LocationType
	Offset   Offset
	Register Register
	Label    DataLabel
}

func raxSized(s semantic.Size) Register {
	switch s {
	case semantic.Size64:
		return RegRAX
	case semantic.Size8:
		return RegAL
	case semantic.Size32:
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
		funcs: map[semantic.FuncName]*FuncDef{
			"do_syscall": {
				ReturnType: semantic.Type{
					Kind: semantic.KindPrimitive,
					Prim: semantic.PrimInt64,
				},
			},
		},
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
	}

	for _, fd := range prog.FuncDefs {
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

func (c *Compiler) compileLine(instrs *Instrs, l *semantic.Line, locals *StackVars) error {
	switch {
	case l.Control != nil:
		return c.compileControl(instrs, l.Control, locals)
	case l.Statement != nil:
		return c.compileStatement(instrs, l.Statement, locals)
	}
	panic("invalid")
}

func (c *Compiler) compileStatement(instrs *Instrs, s *semantic.Statement, locals *StackVars) error {
	switch {
	case s.Assign != nil:
		return c.compileAssign(instrs, s.Assign, locals)
	case s.Expr != nil:
		_, err := c.compileExpr(instrs, s.Expr, locals)
		return err
	case s.Return != nil:
		loc, err := c.compileExpr(instrs, s.Return, locals)
		if err != nil {
			return err
		}

		retReg := returnRegister(s.Return.Type)
		instrs.movLocToReg(s.Return.Type.Size(), loc, retReg)

		return nil
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
			reg, err := c.compileExprToReg(instrs, v, locals)
			if err != nil {
				return err
			}

			o := Offset(v.Type.Size()) * Offset(i)
			offset := startOffset - o

			instrs.movFromReg(a.Expr.Type.SubType.Size(), reg, offset)
		}
	}

	return nil
}

func (c *Compiler) compileStandardAssign(instrs *Instrs, a *semantic.Assign, locals *StackVars) error {
	toOffset, ok := locals.Offset(a.Name)
	if !ok {
		return fmt.Errorf("undefined variable: %s", a.Name)
	}

	reg, err := c.compileExprToReg(instrs, a.Expr, locals)
	if err != nil {
		return err
	}

	instrs.movFromReg(a.Expr.Type.Size(), reg, toOffset)

	return nil
}

func (c *Compiler) compileExprToReg(instrs *Instrs, e *semantic.Expr, locals *StackVars) (Register, error) {
	loc, err := c.compileExpr(instrs, e, locals)
	if err != nil {
		return RegUnset, err
	}

	switch loc.Type {
	case LTRegister:
		return loc.Register, nil
	case LTOffset:
		instrs.movLocToReg(e.Type.Size(), loc, accumulator(e.Type))
		return accumulator(e.Type), nil
	default:
		panic("invalid")
	}
}

func (c *Compiler) compileExpr(instrs *Instrs, e *semantic.Expr, locals *StackVars) (Location, error) {
	switch {
	case e.IsEqual != nil:
		err := c.compileIsEqual(instrs, e.IsEqual, locals)
		return Location{}, err
	case e.FuncCall != nil:
		return c.functionCall(instrs, e.FuncCall, locals)
	case e.AddressOf != "":
		addr, ok := locals.Address(e.AddressOf)
		if !ok {
			return Location{}, fmt.Errorf("undefined variable: %s", e.Var)
		}

		if addr.IsStack() {
			instrs.addInstr("lea %s, [rbp-%d] ; addressOf", RegRAX, addr.stack)
		} else {
			instrs.addInstr("lea %s, [rel %s] ; addressOf", RegRAX, addr.label)
		}

		return Location{
			Type:     LTRegister,
			Register: RegRAX,
		}, nil

	//case e.StringID != 0:
	//	instrs.addInstr("lea %s, [rel %s]", RegRAX, c.stringLabel(e.StringID))
	//	return Location{}, nil
	//	return RegRAX, nil

	case e.Literal != nil:
		switch e.Type.Kind {
		case semantic.KindPrimitive:
			switch e.Type.Prim {
			case semantic.PrimInt32:
				instrs.movInt32ToReg(e.Literal.Int32, RegEAX)
				return regLocation(RegEAX), nil
			case semantic.PrimChar:
				instrs.movByteToReg(e.Literal.Char, RegAL)
				return regLocation(RegAL), nil // todo return lit
			}
		}
	case e.Var != "":
		addr, ok := locals.Address(e.Var)
		if !ok {
			return Location{}, fmt.Errorf("undefined variable: %s", e.Var)
		}

		return offsetLoc(addr.stack), nil
	case e.ArrayVar != nil:
		offset, ok := locals.Offset(e.ArrayVar.Name)
		if !ok {
			return Location{}, fmt.Errorf("undefined variable: %s", e.Var)
		}

		offset += Offset(e.ArrayVar.Index) * Offset(e.Type.Size())

		return Location{
			Type:   LTOffset,
			Offset: offset,
		}, nil
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

	for _, l := range ifs.Lines {
		err := c.compileLine(instrs, l, locals)
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
	leftExprLoc, err := c.compileExpr(instrs, e.Left, locals)
	if err != nil {
		return err
	}

	size := e.Left.Type.Size()

	if leftExprLoc.Type == LTRegister {
		rightLoc, err := c.compileExpr(instrs, e.Right, locals)
		if err != nil {
			return err
		}

		instrs.cmp(size, leftExprLoc.Register, rightLoc)
		return nil
	}

	rightReg, err := c.compileExprToReg(instrs, e.Right, locals)
	if err != nil {
		return err
	}

	instrs.cmp(size, rightReg, leftExprLoc)

	return nil
}

func returnRegister(t semantic.Type) Register {
	switch t.Kind {
	case semantic.KindVoid:
		return RegUnset
	case semantic.KindPrimitive:
		switch t.Prim {
		case semantic.PrimInt32:
			return RegEAX
		case semantic.PrimInt64:
			return RegRAX
		case semantic.PrimChar:
			return RegAL
		}
	case semantic.KindArray:
		panic("array cannot be returned")
	}

	panic("unset type")
}

func accumulator(t semantic.Type) Register {
	switch t.Kind {
	case semantic.KindPrimitive:
		switch t.Prim {
		case semantic.PrimChar:
			return RegAL
		case semantic.PrimInt32:
			return RegEAX
		case semantic.PrimInt64:
			return RegRAX
		}
	}

	panic("unset")
}

func typeInstrSize(s semantic.Size) string {
	switch s {
	case 1:
		return "byte"
	case 2:
		return "word"
	case 4:
		return "dword"
	case 8:
		return "qword"
	}

	panic("invalid")
}

func offsetLoc(o Offset) Location {
	return Location{
		Type:   LTOffset,
		Offset: o,
	}
}

func locAccumulator(t semantic.Type) Location {
	return Location{
		Type:     LTRegister,
		Register: accumulator(t),
	}
}
