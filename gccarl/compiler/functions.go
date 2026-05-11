package compiler

import (
	"fmt"

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

type ParamDef struct {
	Type Type
	Name parser.Identifier
}

type Function struct {
	Name       parser.Identifier
	Params     []*ParamDef
	Statements []*parser.Node
	ReturnType Type
	ReturnExpr *parser.Node
}

func toFunction(node *parser.Node) (*Function, error) {
	retType, err := toType(node.Values[0])
	if err != nil {
		return nil, err
	}

	f := &Function{
		Name:       node.Values[1].Identifier,
		ReturnType: retType,
	}

	if node.Values[3].Node != nil && node.Values[3].Node.Name == "params" {
		for i, v := range node.Values[3].Node.Values {
			paramNode := v.Node
			if i > 0 {
				paramNode = v.Node.Values[1].Node
			}

			p, err := parseParamDef(paramNode)
			if err != nil {
				return nil, err
			}

			f.Params = append(f.Params, p)
		}
	}

	for i := 4; i < len(node.Values); i++ {
		n := node.Values[i].Node
		if n == nil {
			continue
		}

		switch n.Name {
		case "return":
			f.ReturnExpr = n.Values[1].Node
		case "statement":
			f.Statements = append(f.Statements, n.Values[0].Node)
		}
	}

	return f, nil
}

func parseParamDef(n *parser.Node) (*ParamDef, error) {
	paramType, err := toType(n.Values[0])
	if err != nil {
		return nil, err
	}

	return &ParamDef{
		Type: paramType,
		Name: n.Values[1].Identifier,
	}, nil
}

func compileFuncDef(node *parser.Node) ([]Instr, error) {
	f, err := toFunction(node)
	if err != nil {
		return nil, err
	}

	funcInstrs := addInstr(nil, "%s:", f.Name)

	bodyInstrs := addInstr(nil, "push rbp")
	bodyInstrs = addInstr(bodyInstrs, "mov rbp, rsp")

	var paramsSize int
	for _, p := range f.Params {
		paramsSize += p.Type.Size()
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
		returnInstrs, err := handleExpr(f.ReturnExpr, locals, RegEAX)
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
	offset := locals.Add(name, TypeInt)

	if n >= len(intParamRegisters) {
		panic("not implemented")
	}

	reg := intParamRegisters[n]

	return addInstr(instrs, "mov dword [rbp-%d], %s", offset, reg), nil
}

func handleParamsDef(ps []*ParamDef, locals *LocalVars) ([]Instr, error) {
	var instrs []Instr

	var standardC int
	var err error

	for _, p := range ps {
		switch p.Type {
		case TypeInt:
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

func functionCall(vs []*parser.Value, locals *LocalVars) ([]Instr, error) {
	name := vs[0].Identifier
	params := vs[2].Node

	var instrs []Instr

	for i, p := range params.Values {
		exprNode := p.Node
		if i > 0 {
			exprNode = p.Node.Values[1].Node
		}

		if i >= len(intParamRegisters) {
			panic("not implemented")
		}

		reg := intParamRegisters[i]

		exprInstrs, err := handleExpr(exprNode, locals, reg)
		if err != nil {
			return nil, err
		}

		instrs = append(instrs, exprInstrs...)
		//instrs = addInstr(instrs, "mov %s, %s", reg, RegEAX)
	}

	instrs = addInstr(instrs, "call %s", name)
	return instrs, nil
}
