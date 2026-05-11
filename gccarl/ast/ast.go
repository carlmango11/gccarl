package ast

import (
	"github.com/carlmango11/gccarl/gccarl/parser"
)

type Program struct {
	Imports  []*Import
	FuncDefs []*FuncDef
	Dec      []*Declaration
}

type Import struct {
	Name parser.Identifier
}

type Declaration struct {
}

type Type int

const (
	TypeVoid Type = iota
	TypeInt
)

type ParamDef struct {
	Type Type
	Name parser.Identifier
}

type Expr struct {
	Add *AddExpr
	Val *Value
}

type AddExpr struct {
	Val  *Value
	Expr *Expr
}

type Value struct {
	Int int
	Var *Var
}

type Var struct {
	Name  parser.Identifier
	Index *int
}

type FuncDef struct {
	Name       parser.Identifier
	Params     []*ParamDef
	Statements []*Statement
	ReturnType Type
	ReturnExpr *Expr
}

type Statement struct {
	VarDec   *VarDec
	Assign   *Assign
	FuncCall *FuncCall
}

type VarDec struct {
	Type Type
	Var  *Var
	Expr *Expr
}

type Assign struct {
	Var  *Var
	Expr *Expr
}

type FuncCall struct {
	Name   parser.Identifier
	Params []*Expr
}
