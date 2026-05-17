package ast

import (
	"github.com/carlmango11/gccarl/gccarl/parser"
)

type Program struct {
	Imports  []*Import
	FuncDefs []*FuncDef
	Dec      []*Dec
}

type Import struct {
	Name parser.Identifier
}

type RawType struct {
	Type    parser.Identifier
	Pointer *RawType
}

type ParamDef struct {
	Type *TypeDef
	Name parser.Identifier
}

type Expr struct {
	Add      *AddExpr
	Val      *Value
	Cast     *Cast
	FuncCall *FuncCall
}

type Cast struct {
	To   parser.Identifier
	Expr *Expr
}

type AddExpr struct {
	Expr1 *Expr
	Expr2 *Expr
}

type Value struct {
	Int   *int
	Var   *Var
	Str   *string
	Char  *byte
	Array *Array
}

type Array struct {
	Entries []*Expr
}

type Var struct {
	Name    parser.Identifier
	Indexed bool
	Index   int
}

type VarAccess struct {
	Name  parser.Identifier
	Index *int
	Deref bool // todo
}

type TypeDef struct {
	Type  *RawType
	Array bool
}

type FuncDef struct {
	ReturnType *TypeDef
	Name       parser.Identifier
	Params     []*ParamDef
	Statements []*Statement
	ReturnExpr *Expr
}

type Dec struct {
	Type *TypeDef
	Size int
	Name parser.Identifier
}

type Statement struct {
	Dec       *Dec
	DecAssign *DecAssign
	Assign    *Assign
	Expr      *Expr
}

type DecAssign struct {
	Dec    *Dec
	Assign *Assign
}

type Assign struct {
	Var  *Var
	Expr *Expr
}

type FuncCall struct {
	Name   parser.Identifier
	Params []*Expr
}
