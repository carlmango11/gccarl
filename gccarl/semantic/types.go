package semantic

import (
	"fmt"
)

type (
	VarName string
	VarID   int
)

type Var struct {
	ID   VarID
	Name VarName
	Type Type
}

type TypeName string
type FuncName string

type PrimitiveType int

const (
	PrimUnset PrimitiveType = iota
	PrimInt32
	PrimInt64
	PrimChar
	PrimUChar
	PrimBool
	PrimFloat32
)

func (p PrimitiveType) Size() Size {
	switch p {
	case PrimInt32:
		return 4
	case PrimChar:
		return 1
	case PrimBool:
		return 1
	}

	panic("unset primitive type")
}

type Kind int

const (
	KindVoid Kind = iota
	KindPrimitive
	KindStruct
	KindArray
	KindPointer
)

type Size int

const (
	Size8  Size = 1
	Size32 Size = 4
	Size64 Size = 8
)

type StructType struct {
	Name   TypeName
	Fields []StructField
}

func (s StructType) Field(name FieldName) (StructField, bool) {
	for _, f := range s.Fields {
		if f.Name == name {
			return f, true
		}
	}

	return StructField{}, false
}

func (s StructType) FieldIndex(name FieldName) int {
	for i, f := range s.Fields {
		if f.Name == name {
			return i
		}
	}

	panic("did not find field")
}

type StructField struct {
	Name FieldName
	Type Type
}

type Type struct {
	Kind      Kind
	Prim      PrimitiveType
	SubType   *Type
	ArraySize int
	Struct    StructType
}

func (t Type) String() string {
	return fmt.Sprintf("[%v/%v]", t.Kind, t.Prim)
}

func (t Type) Size() Size {
	switch t.Kind {
	case KindPrimitive:
		return t.Prim.Size()
	case KindArray:
		return Size(t.ArraySize) * (*t.SubType).Size()
	case KindPointer:
		return 8
	case KindStruct:
		var s Size
		for _, f := range t.Struct.Fields {
			s += f.Type.Size()
		}
		return s
	}
	panic(fmt.Sprintf("unknown type %v", t))
}

func (t Type) Equals(t2 Type) bool {
	if t.Kind != t2.Kind {
		return false
	}

	if t.Prim != t2.Prim {
		return false
	}

	//if t.Struct != t2.Struct {
	//	return false
	//}

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

func (t Type) TakesInitList() bool {
	return t.Kind == KindArray || t.Kind == KindStruct
}

func (t Type) Indexable() bool {
	return t.Kind == KindArray || t.Kind == KindPointer
}

func (t Type) Field(i int) (Type, error) {
	switch t.Kind {
	case KindArray:
		return *t.SubType, nil
	case KindStruct:
		if i >= len(t.Struct.Fields) {
			return Type{}, fmt.Errorf("index out of bounds: %d", i)
		}

		return t.Struct.Fields[i].Type, nil
	}

	panic(fmt.Sprintf("%v does not have %v", t, i))
}

type Program struct {
	Strings  []string
	FuncDefs []*FuncDef
}

type FuncDef struct {
	ReturnType Type
	Name       FuncName
	Params     []Var
	Statements []*Statement
}

type StructDef struct {
	Fields map[FieldName]Type
}

type While struct {
	Condition *Expr
	Statement *Statement
}

type For struct {
	Init      *Statement
	Condition *Statement
	Action    *Statement
	Statement *Statement
}

type Statement struct {
	DeclareInit *InitVar
	Expr        *Expr
	Return      *Expr
	If          *If
	While       *While
	For         *For
	Compound    *Compound
}

type Compound struct {
	Statements []*Statement
}

type If struct {
	Condition     *Expr
	Statement     *Statement
	ElseStatement *Statement
}

type StringID int

type CompareOp int

const (
	OpUnset CompareOp = iota
	OpEquals
	OpNotEquals
	OpLessThan
)

type NumericOp int

const (
	NumOpUnset NumericOp = iota
	NumOpAdd
)

type CompareOpExpr struct {
	Left  *Expr
	Op    CompareOp
	Right *Expr
}

type NumericOpExpr struct {
	Left  *Expr
	Op    NumericOp
	Right *Expr
}

type ListEntry struct {
	Name []FieldName
	Init *Initialiser
}

type CompLiteral struct {
	Entries []*ListEntry
}

type InitList struct {
	Entries []*ListEntry
}

type FieldName string

type FieldExpr struct {
	Expr  *Expr
	Field FieldName
}

type IndexExpr struct {
	Expr  *Expr
	Index int
}

type Expr struct {
	Type Type

	Assign      *Assign
	Add         *AddExpr
	Compare     *CompareOpExpr
	Numeric     *NumericOpExpr
	FuncCall    *FuncCall
	Field       *FieldExpr
	Index       *IndexExpr
	Literal     *Literal
	AddressOf   *Expr
	Var         *Var
	Deref       *Expr
	Cast        *Cast
	CompLiteral *CompLiteral
	StringID    StringID
}

func (e *Expr) Writeable() bool {
	if e.Var != nil {
		return true
	}

	if e.Deref != nil {
		return e.Deref.Writeable()
	}

	return false
}

type IndexedVar struct {
	Name  Var
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

type Deref struct {
	Expr *Expr
}

type VarRead struct {
	Name  VarName
	Index []int
}

type Assign struct {
	To   *Expr
	Expr *Expr
}

type InitVar struct {
	Var         Var
	Initialiser *Initialiser
}

type Initialiser struct {
	Expr *Expr
	List *InitList
}
