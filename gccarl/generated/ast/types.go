package ast
type CompExprType string

const (
	CompExprTypeAdd CompExprType = "add"
	CompExprTypeSub CompExprType = "sub"
	CompExprTypeIsEqual CompExprType = "is-equal"
)

type CompExpr struct {
	Type CompExprType
	Add *CompExpr
	Sub *CompExpr
	IsEqual *CompExpr
}

type CompExprAdd struct {	 *	PLUS *PLUS	 *
type CompExprSub struct {	 *	MINUS *MINUS	 *
type CompExprIsEqual struct {	 *	EEQUALS *EEQUALS	 *
type DecAssignType string

const (
	DecAssignTypeStandard DecAssignType = "standard"
)

type DecAssign struct {
	Type DecAssignType
	Standard *DecAssign
}

type DecAssignStandard struct {	 *	 *	EQUALS *EQUALS	 *
type CompEntriesType string

const (
	CompEntriesTypeEntries CompEntriesType = "entries"
)

type CompEntries struct {
	Type CompEntriesType
	Entries *CompEntries
}

type CompEntriesEntries struct {	 *	 []*
type ParamsType string

const (
	ParamsTypeParams ParamsType = "params"
)

type Params struct {
	Type ParamsType
	Params *Params
}

type ParamsParams struct {	 *	 []*
type ControlType string

const (
	ControlTypeIf ControlType = "if"
)

type Control struct {
	Type ControlType
	If *Control
}

type ControlIf struct {	IF *IF	LPAREN *LPAREN	 *	RPAREN *RPAREN	LBRACE *LBRACE	 []*	RBRACE *RBRACE
type ReturnType string

const (
	ReturnTypeReturn ReturnType = "return"
)

type Return struct {
	Type ReturnType
	Return *Return
}

type ReturnReturn struct {	RETURN *RETURN	 *	SEMI *SEMI
type ArrayIndexType string

const (
	ArrayIndexTypeArrayIndex ArrayIndexType = "array-index"
)

type ArrayIndex struct {
	Type ArrayIndexType
	ArrayIndex *ArrayIndex
}

type ArrayIndexArrayIndex struct {	LSQUARE *LSQUARE	NUM *NUM	RSQUARE *RSQUARE
type ExprType string

const (
	ExprTypeSubExpr ExprType = "sub-expr"
	ExprTypeComp ExprType = "comp"
)

type Expr struct {
	Type ExprType
	SubExpr *Expr
	Comp *Expr
}

type ExprSubExpr struct {	 *
type ExprComp struct {	 *
type StatementType string

const (
	StatementTypeDecAssign StatementType = "dec-assign"
	StatementTypeVarDec StatementType = "var-dec"
	StatementTypeAssign StatementType = "assign"
	StatementTypeExpr StatementType = "expr"
)

type Statement struct {
	Type StatementType
	DecAssign *Statement
	VarDec *Statement
	Assign *Statement
	Expr *Statement
}

type StatementDecAssign struct {	 *
type StatementVarDec struct {	 *	 *
type StatementAssign struct {	 *	EQUALS *EQUALS	 *
type StatementExpr struct {	 *
type CompositeLiteralType string

const (
	CompositeLiteralTypeArrayVal CompositeLiteralType = "array-val"
)

type CompositeLiteral struct {
	Type CompositeLiteralType
	ArrayVal *CompositeLiteral
}

type CompositeLiteralArrayVal struct {	LBRACE *LBRACE	 *	RBRACE *RBRACE
type TypeType string

const (
	TypeTypeInt TypeType = "int"
	TypeTypeChar TypeType = "char"
	TypeTypeCustom TypeType = "custom"
	TypeTypePointer TypeType = "pointer"
)

type Type struct {
	Type TypeType
	Int *Type
	Char *Type
	Custom *Type
	Pointer *Type
}

type TypeInt struct {	INT_TYPE *INT_TYPE
type TypeChar struct {	CHAR_TYPE *CHAR_TYPE
type TypeCustom struct {	IDEN *IDEN
type TypePointer struct {	ASTERISKS *ASTERISKS	 *
type MainType string

const (
	MainTypeMain MainType = "main"
)

type Main struct {
	Type MainType
	Main *Main
}

type MainMain struct {	 []*
type LineType string

const (
	LineTypeControl LineType = "control"
	LineTypeStatement LineType = "statement"
)

type Line struct {
	Type LineType
	Control *Line
	Statement *Line
}

type LineControl struct {	 *
type LineStatement struct {	 *
type StatementCommaType string

const (
	StatementCommaTypeStatement StatementCommaType = "statement"
)

type StatementComma struct {
	Type StatementCommaType
	Statement *StatementComma
}

type StatementCommaStatement struct {	 *	SEMI *SEMI
type ParamsDefType string

const (
	ParamsDefTypeParams ParamsDefType = "params"
)

type ParamsDef struct {
	Type ParamsDefType
	Params *ParamsDef
}

type ParamsDefParams struct {	 *	 []*
type CommaExprType string

const (
	CommaExprTypeCommaExpr CommaExprType = "comma-expr"
)

type CommaExpr struct {
	Type CommaExprType
	CommaExpr *CommaExpr
}

type CommaExprCommaExpr struct {	COMMA *COMMA	 *
type VariableType string

const (
	VariableTypeVariable VariableType = "variable"
)

type Variable struct {
	Type VariableType
	Variable *Variable
}

type VariableVariable struct {	IDEN *IDEN	 []*
type SubExprType string

const (
	SubExprTypeValue SubExprType = "value"
	SubExprTypeParens SubExprType = "parens"
	SubExprTypeFuncCall SubExprType = "func-call"
)

type SubExpr struct {
	Type SubExprType
	Value *SubExpr
	Parens *SubExpr
	FuncCall *SubExpr
}

type SubExprValue struct {	 *
type SubExprParens struct {	LPAREN *LPAREN	 *	RPAREN *RPAREN
type SubExprFuncCall struct {	IDEN *IDEN	LPAREN *LPAREN	 *	RPAREN *RPAREN
type DecDefType string

const (
	DecDefTypeFuncDef DecDefType = "func-def"
	DecDefTypeDecAssign DecDefType = "dec-assign"
)

type DecDef struct {
	Type DecDefType
	FuncDef *DecDef
	DecAssign *DecDef
}

type DecDefFuncDef struct {	 *	IDEN *IDEN	LPAREN *LPAREN	 *	RPAREN *RPAREN	LBRACE *LBRACE	 []*	 *	RBRACE *RBRACE
type DecDefDecAssign struct {	 *
type ValueType string

const (
	ValueTypeInt ValueType = "int"
	ValueTypeVariable ValueType = "variable"
	ValueTypeStr ValueType = "str"
	ValueTypeChar ValueType = "char"
	ValueTypeCast ValueType = "cast"
	ValueTypeCompLit ValueType = "comp-lit"
)

type Value struct {
	Type ValueType
	Int *Value
	Variable *Value
	Str *Value
	Char *Value
	Cast *Value
	CompLit *Value
}

type ValueInt struct {	NUM *NUM
type ValueVariable struct {	 *
type ValueStr struct {	STR *STR
type ValueChar struct {	CHAR *CHAR
type ValueCast struct {	LPAREN *LPAREN	 *	RPAREN *RPAREN	 *
type ValueCompLit struct {	 *
type CommaParamDefType string

const (
	CommaParamDefTypeParam CommaParamDefType = "param"
)

type CommaParamDef struct {
	Type CommaParamDefType
	Param *CommaParamDef
}

type CommaParamDefParam struct {	COMMA *COMMA	 *
type ParamDefType string

const (
	ParamDefTypeParam ParamDefType = "param"
)

type ParamDef struct {
	Type ParamDefType
	Param *ParamDef
}

type ParamDefParam struct {	 *	 *