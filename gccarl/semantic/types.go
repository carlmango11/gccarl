package program

import (
	"github.com/carlmango11/gccarl/gccarl/ast"
	"github.com/carlmango11/gccarl/gccarl/parser"
)

type VarName parser.Identifier

type TypeName string
type FuncName string

type PrimitiveType int

const (
	PrimUnset PrimitiveType = iota
	PrimInt32
)

type Kind int

const (
	KindVoid Kind = iota
	KindPrimitive
	KindCustom
	KindArray
	KindPointer
)

type Type struct {
	Kind    Kind
	Prim    PrimitiveType
	SubType *Type
	Custom  TypeName
}

func (t Type) Equals(t2 Type) bool {
	if t.Kind != t2.Kind {
		return false
	}

	if t.Prim != t2.Prim {
		return false
	}

	if !t.SubType.Equals(*t2.SubType) {
		return false
	}

	if t.Custom != t2.Custom {
		return false
	}

	return true
}

type Program struct {
	Imports  []*ast.Import
	FuncDefs []*ast.FuncDef
	Dec      []*ast.Dec
}

type FuncDef struct {
	ReturnType Type
	Name       FuncName
	Params     []*ParamDef
	//TypeDefs []
	Locals     map[VarName]Type
	Statements []*Statement
	ReturnExpr *Expr
}

type Statement struct {
	Assign *Assign
	Expr   *Expr
}

type Dec struct {
	Name VarName
	Type *Type
}

type Expr struct {
	Type     Type
	Add      *AddExpr
	FuncCall *FuncCall
	Literal  Literal
	Var      VarName
	ArrayVar *IndexedVar
	Cast     *Cast
}

type IndexedVar struct {
	Name  VarName
	Index int
}

type Cast struct {
	To   Type
	Expr *Expr
}

type Literal struct {
	Int32 int32
}

type AddExpr struct {
	Expr1 *Expr
	Expr2 *Expr
}

type FuncCall struct {
	Func FuncName
	Args []*Expr
}

type Assign struct {
	Name  VarName
	Index int
	Expr  *Expr
}

type ParamDef struct {
	Type Type
	Name VarName
}
