package semantic

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/ast"
)

type VarName ast.Identifier

type TypeName string
type FuncName string

type PrimitiveType int

const (
	PrimUnset PrimitiveType = iota
	PrimInt32
	PrimChar
	PrimBool
	PrimFloat32
)

func (p PrimitiveType) Size() int {
	switch p {
	case PrimInt32:
		return 8
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
	ArraySize int
	Custom    TypeName
}

func (t Type) String() string {
	return fmt.Sprintf("[%v/%v]", t.Kind, t.Prim)
}

func (t Type) Size() int {
	switch t.Kind {
	case KindPrimitive:
		return t.Prim.Size()
	case KindArray:
		return t.ArraySize * (*t.SubType).Size()
	case KindPointer:
		return 8
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

	if t.SubType == nil || t2.SubType == nil {
		return false
	}

	if !t.SubType.Equals(*t2.SubType) {
		return false
	}

	return true
}

func (t Type) IsIntParamType() bool {
	fl := t.Kind == KindPrimitive && t.Prim == PrimFloat32
	return !fl
}

type Program struct {
	Strings  []string
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
	Lines      []*Line
	ReturnExpr *Expr
}

type Statement struct {
	If     *If
	Assign *Assign
	Expr   *Expr
}

type If struct {
	Condition  *Expr
	Statements []*Statement
}

type Dec struct {
	Name VarName
	Type *Type
}

type StringID int

type IsEqual struct {
	Left  *Expr
	Right *Expr
}

type Line struct {
	Statement *Statement
	Control   *Control
}

type Control struct {
	If *If
}

type Expr struct {
	Type        Type
	Add         *AddExpr
	IsEqual     *IsEqual
	FuncCall    *FuncCall
	Literal     *Literal
	Var         VarName
	AddressOf   VarName
	ArrayVar    *IndexedVar
	Cast        *Cast
	CompLiteral []*Expr
	StringID    StringID
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
	Name VarName
	Expr *Expr
}

type ParamDef struct {
	Type Type
	Name VarName
}
