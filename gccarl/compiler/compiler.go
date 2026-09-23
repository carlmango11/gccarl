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

type Instr string
type DataLabel string

type FuncDef struct {
	ReturnType semantic.Type
	Params     []semantic.Var
}

type Compiler struct {
	funcs  map[semantic.FuncName]*FuncDef
	labelC int

	globals map[semantic.VarID]Offset // TODO relative offset?
	locals  *StackVars
	reg     map[Register]bool
}

func New() *Compiler {
	return &Compiler{
		reg: make(map[Register]bool),
		funcs: map[semantic.FuncName]*FuncDef{
			"assert": {
				ReturnType: semantic.Type{
					Kind: semantic.KindPrimitive,
					Prim: semantic.PrimInt64,
				},
			},
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
	full.addInstr("\thlt")

	full.instrs = append(full.instrs, assertRoutine...)

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

func (c *Compiler) compileStatement(instrs *Instrs, s *semantic.Statement) error {
	if s == nil {
		// TODO this is a bit of a cop-out but variable declarations generate a nil statement because there' nothing to
		// be done
		return nil
	}

	switch {
	case s.Expr != nil:
		_, err := c.compileExpr(instrs, s.Expr)
		return err
	case s.Return != nil:
		loc, err := c.compileExpr(instrs, s.Return)
		if err != nil {
			return err
		}

		instrs.movLocToReg(s.Return.Type.Size(), loc, RegA)

		return nil
	case s.DeclareInit != nil:
		err := c.compileDecInit(instrs, s.DeclareInit)
		if err != nil {
			return err
		}

		return nil
	case s.If != nil:
		return c.compileIf(instrs, s.If)
	case s.While != nil:
		return c.compileWhile(instrs, s.While)
	case s.For != nil:
		return c.compileFor(instrs, s.For)
	case s.Compound != nil:
		for _, cs := range s.Compound.Statements {
			err := c.compileStatement(instrs, cs)
			if err != nil {
				return err
			}
		}

		return nil
	}

	panic("missing statement type")
}

func (c *Compiler) compileDecInit(instrs *Instrs, a *semantic.InitVar) error {
	instrs.addComment("START declare and assign %v", a.Var.Name)
	defer instrs.addComment("END declare and assign %v", a.Var.Name)

	offset := c.varOffset(a.Var)

	switch {
	case a.Initialiser.List != nil:
		return c.compileInitialiser(instrs, a.Var.Type, offset, a.Initialiser)
	default:
		reg, err := c.compileExprToReg(instrs, a.Initialiser.Expr)
		if err != nil {
			return err
		}

		instrs.movFromReg(a.Var.Type.Size(), reg, offset)

		return nil
	}
}

func (c *Compiler) compileInitialiser(instrs *Instrs, typ semantic.Type, to Offset, init *semantic.Initialiser) error {
	switch {
	case init.Expr != nil:
		loc, err := c.compileExpr(instrs, init.Expr)
		if err != nil {
			return err
		}

		return c.movExprToOffset(instrs, typ, loc, to)

	case init.List != nil:
		var i int

		for _, entry := range init.List.Entries {
			if len(entry.Name) > 0 {
				// jump to this
				i = fieldNameIndex(typ, entry.Name)
			}

			// get offset of i
			fieldOffset, fieldType := fieldByIndex(typ, i)
			i++

			err := c.compileInitialiser(instrs, fieldType, to-fieldOffset, entry.Init)
			if err != nil {
				return err
			}
		}

		return nil
	default:
		panic("invalid initialiser type")
	}
}

func fieldByIndex(typ semantic.Type, i int) (Offset, semantic.Type) {
	switch typ.Kind {
	case semantic.KindArray:
		return Offset(typ.SubType.Size()) * Offset(i), *typ.SubType
	case semantic.KindStruct:
		var offset Offset

		for j, f := range typ.Struct.Fields {
			if j == i {
				return offset, f.Type
			}

			offset += Offset(f.Type.Size())
		}

		panic("never reached i")
	}

	panic("invalid kind")
}

func fieldNameIndex(typ semantic.Type, names []semantic.FieldName) int {
	if len(names) > 1 {
		panic("impl")
	}

	for i, f := range typ.Struct.Fields {
		if f.Name == names[0] {
			return i
		}
	}

	panicf("field %s not found", names[0])
	return 0
}

func (c *Compiler) compileAssign(instrs *Instrs, to, e *semantic.Expr) error {
	if to.Deref != nil {
		panic("handle deref")
	}

	toLoc, err := c.compileExpr(instrs, to)
	if err != nil {
		return err
	}

	if toLoc.Type != LTOffset {
		panicf("cannot write to %v", toLoc.Type)
	}

	reg, err := c.compileExprToReg(instrs, e)
	if err != nil {
		return err
	}

	instrs.movFromReg(e.Type.Size(), reg, toLoc.Offset)

	return nil
}

func (c *Compiler) compileExprToReg(instrs *Instrs, e *semantic.Expr) (Register, error) {
	loc, err := c.compileExpr(instrs, e)
	if err != nil {
		return RegUnset, err
	}

	switch loc.Type {
	case LTRegister:
		return loc.Register, nil
	case LTOffset:
		instrs.movLocToReg(e.Type.Size(), loc, RegA)
		return RegA, nil
	default:
		panic("invalid")
	}
}

func (c *Compiler) compileExpr(instrs *Instrs, e *semantic.Expr) (Location, error) {
	switch {
	case e.Assign != nil:
		err := c.compileAssign(instrs, e.Assign.To, e.Assign.Expr)
		return Location{}, err // TODO: return Loc
	case e.AddressOf != nil:
		return c.compileAddressOf(instrs, e.AddressOf)
	case e.Compare != nil:
		return c.compileCompare(instrs, e.Compare)
	case e.Numeric != nil:
		return c.compileNumeric(instrs, e.Numeric)
	case e.FuncCall != nil:
		return c.functionCall(instrs, e.FuncCall)
	case e.StringID != 0:
		return Location{
			Type:  LTLabel,
			Label: c.stringLabel(e.StringID),
		}, nil

	case e.Literal != nil:
		switch e.Type.Kind {
		case semantic.KindPrimitive:
			switch e.Type.Prim {
			case semantic.PrimInt32:
				instrs.movInt32ToReg(e.Literal.Int32, RegA)
				return regLocation(RegA), nil
			case semantic.PrimChar:
				instrs.movByteToReg(e.Literal.Char, RegA)
				return regLocation(RegA), nil // todo return lit
			}
		}
	case e.Deref != nil:
		return c.compileDeref(instrs, e.Deref)
	case e.Var != nil:
		return c.varExprLoc(*e.Var)
	case e.Field != nil:
		return c.compileFieldExpr(instrs, e.Field)
	case e.Index != nil:
		return c.compileIndexExpr(instrs, e.Index)
	}

	panic(fmt.Sprintf("unknown expr type: %+v", e))
}

func (c *Compiler) compileIndexExpr(instrs *Instrs, f *semantic.IndexExpr) (Location, error) {
	loc, err := c.compileExpr(instrs, f.Expr)
	if err != nil {
		return Location{}, err
	}

	offset := Offset(f.Expr.Type.SubType.Size()) * Offset(f.Index)

	switch loc.Type {
	case LTOffset:
		loc.Offset -= offset
	case LTRegister:
		instrs.addN(8, loc.Register, int(offset))
	default:
		panic("invalid location")
	}

	return loc, nil
}

func (c *Compiler) compileFieldExpr(instrs *Instrs, f *semantic.FieldExpr) (Location, error) {
	loc, err := c.compileExpr(instrs, f.Expr)
	if err != nil {
		return Location{}, err
	}

	offset := fieldNameOffset(f.Expr.Type, f.Field)

	switch loc.Type {
	case LTOffset:
		loc.Offset -= offset
	case LTRegister:
		instrs.addN(8, loc.Register, int(offset))
	default:
		panic("invalid location")
	}

	return loc, nil
}

func (c *Compiler) varOffset(v semantic.Var) Offset {
	// todo global

	return c.locals.Offset(v)
}

func (c *Compiler) varExprLoc(v semantic.Var) (Location, error) {
	offset := c.varOffset(v)

	return offsetLoc(offset), nil
}

func (c *Compiler) stringLabel(id semantic.StringID) DataLabel {
	return DataLabel(fmt.Sprintf("str_%d", id))
}

func (c *Compiler) compileIf(instrs *Instrs, ifs *semantic.If) error {
	loc, err := c.compileExpr(instrs, ifs.Condition)
	if err != nil {
		return err
	}

	instrs.movLocToReg(ifs.Condition.Type.Size(), loc, RegA)
	instrs.addInstr("mov %s, 1", RawRBX)

	instrs.addInstr("cmp %s, %s", RawRAX, RawRBX)

	endIfLabel := c.newLabel("end_if")
	elseLabel := c.newLabel("else")

	instrs.addInstr("jne %s", elseLabel)

	err = c.compileStatement(instrs, ifs.Statement)
	if err != nil {
		return err
	}

	// finished if clause, jump to end
	instrs.addInstr("jmp %s", endIfLabel)

	// else
	instrs.addInstr("%s:", elseLabel)

	err = c.compileStatement(instrs, ifs.ElseStatement)
	if err != nil {
		return err
	}

	instrs.addInstr("%v:", endIfLabel)

	return nil
}

func (c *Compiler) newLabel(prefix string) string {
	c.labelC++
	return fmt.Sprintf("%v_%d", prefix, c.labelC)
}

func (c *Compiler) compileCompare(instrs *Instrs, e *semantic.CompareOpExpr) (Location, error) {
	rightLoc, err := c.compileExpr(instrs, e.Right)
	if err != nil {
		return Location{}, err
	}

	instrs.movLocToReg(e.Left.Type.Size(), rightLoc, RegD)

	leftLoc, err := c.compileExpr(instrs, e.Left)
	if err != nil {
		return Location{}, err
	}

	r10 := RegR10.Raw(e.Left.Type.Size())
	instrs.addInstr("mov %s, 1", r10)

	instrs.cmp(e.Left.Type.Size(), RegD, leftLoc)

	jump := c.newLabel("jump")

	switch e.Op {
	case semantic.OpLessThan:
		instrs.addInstr("jl %s", jump)
	case semantic.OpEquals:
		instrs.addInstr("je %s", jump)
	default:
		panic("missing op")
	}

	instrs.addInstr("mov %s, 0", r10)
	instrs.addInstr("%s:", jump)

	return Location{
		Type:     LTRegister,
		Register: RegR10,
	}, nil
}

func (c *Compiler) compileNumeric(instrs *Instrs, n *semantic.NumericOpExpr) (Location, error) {
	rightLoc, err := c.compileExpr(instrs, n.Right)
	if err != nil {
		return Location{}, err
	}

	instrs.movLocToReg(n.Left.Type.Size(), rightLoc, RegD)

	leftLoc, err := c.compileExpr(instrs, n.Left)
	if err != nil {
		return Location{}, err
	}

	switch n.Op {
	case semantic.NumOpAdd:
		instrs.add(n.Left.Type.Size(), RegD, leftLoc)
	default:
		panic("missing op")
	}

	return Location{
		Type:     LTRegister,
		Register: RegD,
	}, nil
}

func (c *Compiler) compileWhile(instrs *Instrs, w *semantic.While) error {
	repeat := c.newLabel("repeat")
	instrs.addInstr("%v:", repeat)

	loc, err := c.compileExpr(instrs, w.Condition)
	if err != nil {
		return err
	}

	instrs.movLocToReg(w.Condition.Type.Size(), loc, RegA)
	instrs.addInstr("mov %s, 1", RawRBX)

	instrs.addInstr("cmp %s, %s", RawRAX, RawRBX)

	skip := c.newLabel("skip")
	instrs.addInstr("jne %s", skip)

	err = c.compileStatement(instrs, w.Statement)
	if err != nil {
		return err
	}

	instrs.addComment("start again")
	instrs.addInstr("jmp %s", repeat)

	instrs.addInstr("%v:", skip)
	return nil
}

func (c *Compiler) compileFor(instrs *Instrs, f *semantic.For) error {
	c.compileStatement(instrs, f.Init)

	repeat := c.newLabel("repeat")
	instrs.addInstr("%v:", repeat)

	//loc, err := c.compileExpr(instrs, w.Condition)
	//if err != nil {
	//	return err
	//}
	//
	//instrs.movLocToReg(w.Condition.Type.Size(), loc, RegA)
	//instrs.addInstr("mov %s, 1", RawRBX)
	//
	//instrs.addInstr("cmp %s, %s", RawRAX, RawRBX)
	//
	//skip := c.newLabel("skip")
	//instrs.addInstr("jne %s", skip)
	//
	//for _, l := range w.Lines {
	//	err := c.compileLine(instrs, l)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//instrs.addComment("start again")
	//instrs.addInstr("jmp %s", repeat)
	//
	//instrs.addInstr("%v:", skip)
	return nil
}

func (c *Compiler) compileAddressOf(instrs *Instrs, e *semantic.Expr) (Location, error) {
	loc, err := c.compileExpr(instrs, e)
	if err != nil {
		return Location{}, err
	}

	if loc.Type != LTOffset {
		return Location{}, fmt.Errorf("address of %s is not offset", loc.Type)
	}

	instrs.addInstr("lea %s, [rbp-%d] ; addressOf", RawRAX, loc.Offset)

	//if addr.IsStack() {
	//	instrs.addInstr("lea %s, [rbp-%d] ; addressOf", RegA.Raw(typ.Size()), addr.stack)
	//} else {
	//	panic("impl")
	//	//instrs.addInstr("lea %s, [rel %s] ; addressOf", RegA.Raw(typ.Size()), addr.label)
	//}

	return Location{
		Type:     LTRegister,
		Register: RegA,
	}, nil
}

func (c *Compiler) compileDeref(instrs *Instrs, e *semantic.Expr) (Location, error) {
	loc, err := c.compileExpr(instrs, e)
	if err != nil {
		return Location{}, err
	}

	// loc contains the address of the thing to be deref'd
	instrs.movLocToReg(e.Type.Size(), loc, RegA)

	// I know it's a pointer so we use full register
	instrs.movFromAddressToReg(RawRAX, RawRAX)

	return Location{
		Type:     LTRegister,
		Register: RegA,
	}, nil
}

func (c *Compiler) movExprToOffset(instrs *Instrs, typ semantic.Type, loc Location, to Offset) error {
	switch loc.Type {
	case LTRegister:
		instrs.movFromReg(typ.Size(), loc.Register, to)
		return nil
	case LTOffset:
		left := typ.Size()
		from := loc.Offset

		for left > 0 {
			thisMove := min(8, left)

			instrs.movOffsetToReg(thisMove, from, RegA)
			instrs.movFromReg(thisMove, RegA, to)

			from += Offset(thisMove)
			to += Offset(thisMove)
			left -= thisMove
		}

		return nil
	default:
		panic("invalid loc type")
	}
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

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}
