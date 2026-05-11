package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/parser"
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

func Compile(n *parser.Node) ([]byte, error) {
	instrs, err := compile(n)
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

func compile(n *parser.Node) ([]Instr, error) {
	return compileMain(n.Values)
}

func compileMain(vs []*parser.Value) ([]Instr, error) {
	var instrs []Instr

	instrs = addInstr(instrs, "section .text")
	instrs = addInstr(instrs, "global _start")

	instrs = addInstr(instrs, "_main:")
	instrs = addInstr(instrs, "\tcall main")

	//locals := &LocalVars{}

	for _, v := range vs {
		switch v.Node.Name {
		case "include":
		case "func-def":
			funcInstrs, err := compileFuncDef(v.Node)
			if err != nil {
				return nil, err
			}

			instrs = append(instrs, funcInstrs...)
		}
	}

	return instrs, nil
}

func compileStatement(n *parser.Node, locals *LocalVars) ([]Instr, error) {
	switch n.Name {
	case "dec-assign":
		return compileDecAssign(n.Values[0].Node.Values, locals)
	case "assign":
		return assign(n.Values, locals)
	case "func-call":
		return functionCall(n.Values, locals)
	}

	return nil, fmt.Errorf("unknown node: %s", n.Name)
}

func assign(vs []*parser.Value, locals *LocalVars) ([]Instr, error) {
	toOffset, ok := locals.Offset(vs[0].Identifier)
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", vs[0].Identifier)
	}

	output, err := handleExpr(vs[2].Node, locals, RegEAX)
	if err != nil {
		return nil, err
	}

	return addInstr(output, `mov dword [rbp-%d], %s`, toOffset, RegEAX), nil
}

func loadValue(n *parser.Node, locals *LocalVars, target Register) ([]Instr, error) {
	if len(n.Values) != 1 {
		return nil, fmt.Errorf("invalid number of values: %d", len(n.Values))
	}

	val := n.Values[0]

	switch n.Name {
	case "int":
		return addInstr(nil, "mov %s, %d", target, int(val.Number)), nil
	case "variable":
		offset, ok := locals.Offset(val.Identifier)
		if !ok {
			return nil, fmt.Errorf("undefined variable: %s", val.Identifier)
		}

		return addInstr(nil, "mov %s, [rbp-%d]", target, offset), nil
	default:
		return nil, fmt.Errorf("unknown variable: %s", n.Name)
	}
}

func handleExpr(n *parser.Node, locals *LocalVars, target Register) ([]Instr, error) {
	switch n.Name {
	case "add":
		return compileAdd(n.Values, locals, target)
	case "value":
		return loadValue(n.Values[0].Node, locals, target)
	default:
		return nil, fmt.Errorf("unknown expr type: %s", n.Name)
	}
}

func compileAdd(vs []*parser.Value, locals *LocalVars, target Register) ([]Instr, error) {
	var all []Instr

	output, err := handleExpr(vs[2].Node, locals, target)
	if err != nil {
		return nil, err
	}

	all = append(all, output...)

	output, err = loadValue(vs[0].Node, locals, RegEBX)
	if err != nil {
		return nil, err
	}

	all = append(all, output...)

	return addInstr(all, "add %s, %s", target, RegEBX), nil
}

func compileDecAssign(vs []*parser.Value, locals *LocalVars) ([]Instr, error) {
	typ, err := toType(vs[0])
	if err != nil {
		return nil, err
	}

	switch typ {
	case TypeInt:
		return compileIntAssignment(vs, locals)
	default:
		return nil, fmt.Errorf("unknown type: %v", typ)
	}
}

func compileIntAssignment(vs []*parser.Value, locals *LocalVars) ([]Instr, error) {
	name := vs[1].Identifier
	val := vs[3].Number

	offset := locals.Add(name, TypeInt)

	return addInstr(nil, "mov dword [rbp-%d], %d", offset, int(val)), nil
}

func addInstr(instrs []Instr, format string, args ...any) []Instr {
	i := fmt.Sprintf(format, args...)
	return append(instrs, Instr(i))
}

func toType(v *parser.Value) (Type, error) {
	if v.Node == nil {
		return TypeVoid, fmt.Errorf("expected type, got %+v", v)
	}

	switch v.Node.Values[0].Literal {
	case "int":
		return TypeInt, nil
	default:
		return TypeVoid, fmt.Errorf("unknown type: %s", v.Node.Values[0].Literal)
	}
}
