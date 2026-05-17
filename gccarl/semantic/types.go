package semantic

import (
	"fmt"

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
	PrimChar
)

func (p PrimitiveType) Size() int {
	switch p {
	case PrimInt32:
		return 4
	case PrimChar:
		return 1
	}

	panic("unset primitive type")
}

type Kind int

const (
	KindVoid Kind = iota
	KindPrimitive
	KindCustom
	KindArray
	KindPointer
)

type Type struct {
	Kind      Kind
	Prim      PrimitiveType
	SubType   *Type
	ArraySize int // todo: array needs subtype
	Custom    TypeName
}

func (t Type) Size() int {
	switch t.Kind {
	case KindPrimitive:
		return t.Prim.Size()
	}
	panic(fmt.Sprintf("unknown type %d", t))
}

func (t Type) Equals(t2 Type) bool {
	if t.Kind != t2.Kind {
		return false
	}

	if t.Prim != t2.Prim {
		return false
	}

	if t.Custom != t2.Custom {
		return false
	}

	if t.SubType == nil && t2.SubType == nil {
		return true
	}

	if t.SubType != nil || t2.SubType != nil {
		return false
	}

	if !t.SubType.Equals(*t2.SubType) {
		return false
	}

	return true
}

func (t Type) IsIntParamType() bool {
	switch t.Kind {
	case KindPrimitive:
		switch t.Prim {
		case PrimInt32:
			return true
		}
	}

	panic("unset primitive type")
}

type Program struct {
	Imports  []*ast.Import
	FuncDefs []*FuncDef
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

type Array struct {
	Vals []*Expr
}

type Expr struct {
	Type     Type
	Add      *AddExpr
	FuncCall *FuncCall
	Literal  Literal
	Var      VarName
	ArrayVar *IndexedVar
	Array    *Array
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
	Char  byte
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
