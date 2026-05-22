package ast

type Program struct {
	Imports  []*Import
	FuncDefs []*FuncDef
	Dec      []*Dec
}

type Import struct {
	Name Identifier
}

type RawType struct {
	Type    Identifier
	Sub     *RawType
	Pointer bool
}

type ParamDef struct {
	Type *TypeDef
	Name Identifier
}

type Expr struct {
	Add      *AddExpr
	Val      *Value
	Cast     *Cast
	FuncCall *FuncCall
}

type Cast struct {
	To   Identifier
	Expr *Expr
}

type AddExpr struct {
	Expr1 *Expr
	Expr2 *Expr
}

type Value struct {
	Int  *int
	Var  *Var
	Str  *string
	Char *byte
}

type ArrayDef struct {
	Size    int
	HasSize bool
}

type Var struct {
	Name   Identifier
	Arrays []*ArrayDef
}

type VarAccess struct {
	Name  Identifier
	Index *int
	Deref bool // todo
}

type TypeDef struct {
	Type   *RawType
	Arrays []*ArrayDef
}

type FuncDef struct {
	ReturnType *TypeDef
	Name       Identifier
	Params     []*ParamDef
	Statements []*Statement
	ReturnExpr *Expr
}

type Dec struct {
	Type *TypeDef
	Name Identifier
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
	Array  *ArrayDecAssign
}

type ArrayDecAssign struct {
	Dec   *Dec
	Exprs []*Expr
}

type Assign struct {
	Var  *Var
	Expr *Expr
}

type Identifier string

type FuncCall struct {
	Name Identifier
	Args []*Expr
}
