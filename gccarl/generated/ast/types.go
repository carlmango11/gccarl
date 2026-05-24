package ast

type STR string
type IF string
type RETURN string
type CHAR string
type COMMA string
type IDEN string
type RSQUARE string
type CHAR_TYPE string
type SEMI string
type NUM string
type LPAREN string
type RPAREN string
type EQUALS string
type LSQUARE string
type MINUS string
type ASTERISKS string
type LBRACE string
type RBRACE string
type PLUS string
type EEQUALS string
type INT_TYPE string
type ParamsDefParams struct {
	ParamsDef *ParamsDefParams
	ParamsDef []*ParamsDefParams
}

type SubExprValue struct {
	SubExpr *SubExprValue
}

type CompExprAdd struct {
	CompExpr *CompExprAdd
	PLUS PLUS
	CompExpr *CompExprAdd
}

type CompExprSub struct {
	CompExpr *CompExprSub
	MINUS MINUS
	CompExpr *CompExprSub
}

type StatementVarDec struct {
	Statement *StatementVarDec
	Statement *StatementVarDec
}

type ValueStr struct {
	STR STR
}

type VariableVariable struct {
	IDEN IDEN
	Variable []*VariableVariable
}

type ParamsParams struct {
	Params *ParamsParams
	Params []*ParamsParams
}

type CommaParamDefParam struct {
	COMMA COMMA
	CommaParamDef *CommaParamDefParam
}

type ValueVariable struct {
	Value *ValueVariable
}

type ValueCast struct {
	LPAREN LPAREN
	Value *ValueCast
	RPAREN RPAREN
	Value *ValueCast
}

type ValueCompLit struct {
	Value *ValueCompLit
}

type MainMain struct {
	Main []*MainMain
}

type ValueInt struct {
	NUM NUM
}

type CommaExprCommaExpr struct {
	COMMA COMMA
	CommaExpr *CommaExprCommaExpr
}

type ArrayIndexArrayIndex struct {
	LSQUARE LSQUARE
	NUM NUM
	RSQUARE RSQUARE
}

type StatementAssign struct {
	Statement *StatementAssign
	EQUALS EQUALS
	Statement *StatementAssign
}

type StatementExpr struct {
	Statement *StatementExpr
}

type TypeChar struct {
	CHAR_TYPE CHAR_TYPE
}

type TypeCustom struct {
	IDEN IDEN
}

type StatementCommaStatement struct {
	StatementComma *StatementCommaStatement
	SEMI SEMI
}

type ValueChar struct {
	CHAR CHAR
}

type LineControl struct {
	Line *LineControl
}

type LineStatement struct {
	Line *LineStatement
}

type CompositeLiteralArrayVal struct {
	LBRACE LBRACE
	CompositeLiteral *CompositeLiteralArrayVal
	RBRACE RBRACE
}

type TypeInt struct {
	INT_TYPE INT_TYPE
}

type ReturnReturn struct {
	RETURN RETURN
	Return *ReturnReturn
	SEMI SEMI
}

type DecDefDecAssign struct {
	DecDef *DecDefDecAssign
}

type SubExprFuncCall struct {
	IDEN IDEN
	LPAREN LPAREN
	SubExpr *SubExprFuncCall
	RPAREN RPAREN
}

type TypePointer struct {
	ASTERISKS ASTERISKS
	Type *TypePointer
}

type ControlIf struct {
	IF IF
	LPAREN LPAREN
	Control *ControlIf
	RPAREN RPAREN
	LBRACE LBRACE
	Control []*ControlIf
	RBRACE RBRACE
}

type CompExprIsEqual struct {
	CompExpr *CompExprIsEqual
	EEQUALS EEQUALS
	CompExpr *CompExprIsEqual
}

type CompEntriesEntries struct {
	CompEntries *CompEntriesEntries
	CompEntries []*CompEntriesEntries
}

type ExprComp struct {
	Expr *ExprComp
}

type ParamDefParam struct {
	ParamDef *ParamDefParam
	ParamDef *ParamDefParam
}

type SubExprParens struct {
	LPAREN LPAREN
	SubExpr *SubExprParens
	RPAREN RPAREN
}

type StatementDecAssign struct {
	Statement *StatementDecAssign
}

type ExprSubExpr struct {
	Expr *ExprSubExpr
}

type DecAssignStandard struct {
	DecAssign *DecAssignStandard
	DecAssign *DecAssignStandard
	EQUALS EQUALS
	DecAssign *DecAssignStandard
}

type DecDefFuncDef struct {
	DecDef *DecDefFuncDef
	IDEN IDEN
	LPAREN LPAREN
	DecDef *DecDefFuncDef
	RPAREN RPAREN
	LBRACE LBRACE
	DecDef []*DecDefFuncDef
	DecDef *DecDefFuncDef
	RBRACE RBRACE
}


}
type LineType string

const (
	LineTypeControl LineType = "control"
	LineTypeStatement LineType = "statement"
)

type Line struct {
	Type LineType
	Control *LineControl
	Statement *LineStatement
}

type LineControl struct {
	Line *Line
}
type LineStatement struct {
	Line *Line
}
type LineControl struct {	 *
type LineStatement struct {	 *
type ControlType string

const (
	ControlTypeIf ControlType = "if"
)

type Control struct {
	Type ControlType
	If *ControlIf
}

type ControlIf struct {
	IF IF
	LPAREN LPAREN
	Control *Control
	RPAREN RPAREN
	LBRACE LBRACE
	Control []*Control
	RBRACE RBRACE
}
type ControlIf struct {	IF *IF	LPAREN *LPAREN	 *	RPAREN *RPAREN	LBRACE *LBRACE	 []*	RBRACE *RBRACE
type StatementCommaType string

const (
	StatementCommaTypeStatement StatementCommaType = "statement"
)

type StatementComma struct {
	Type StatementCommaType
	Statement *StatementCommaStatement
}

type StatementCommaStatement struct {
	StatementComma *StatementComma
	SEMI SEMI
}
type StatementCommaStatement struct {	 *	SEMI *SEMI
type ArrayIndexType string

const (
	ArrayIndexTypeArrayIndex ArrayIndexType = "array-index"
)

type ArrayIndex struct {
	Type ArrayIndexType
	ArrayIndex *ArrayIndexArrayIndex
}

type ArrayIndexArrayIndex struct {
	LSQUARE LSQUARE
	NUM NUM
	RSQUARE RSQUARE
}
type ArrayIndexArrayIndex struct {	LSQUARE *LSQUARE	NUM *NUM	RSQUARE *RSQUARE
type CompExprType string

const (
	CompExprTypeAdd CompExprType = "add"
	CompExprTypeSub CompExprType = "sub"
	CompExprTypeIsEqual CompExprType = "is-equal"
)

type CompExpr struct {
	Type CompExprType
	Add *CompExprAdd
	Sub *CompExprSub
	IsEqual *CompExprIsEqual
}

type CompExprAdd struct {
	CompExpr *CompExpr
	PLUS PLUS
	CompExpr *CompExpr
}
type CompExprSub struct {
	CompExpr *CompExpr
	MINUS MINUS
	CompExpr *CompExpr
}
type CompExprIsEqual struct {
	CompExpr *CompExpr
	EEQUALS EEQUALS
	CompExpr *CompExpr
}
type CompExprAdd struct {	 *	PLUS *PLUS	 *
type CompExprSub struct {	 *	MINUS *MINUS	 *
type CompExprIsEqual struct {	 *	EEQUALS *EEQUALS	 *
type StatementType string

const (
	StatementTypeDecAssign StatementType = "dec-assign"
	StatementTypeVarDec StatementType = "var-dec"
	StatementTypeAssign StatementType = "assign"
	StatementTypeExpr StatementType = "expr"
)

type Statement struct {
	Type StatementType
	DecAssign *StatementDecAssign
	VarDec *StatementVarDec
	Assign *StatementAssign
	Expr *StatementExpr
}

type StatementDecAssign struct {
	Statement *Statement
}
type StatementVarDec struct {
	Statement *Statement
	Statement *Statement
}
type StatementAssign struct {
	Statement *Statement
	EQUALS EQUALS
	Statement *Statement
}
type StatementExpr struct {
	Statement *Statement
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
	ArrayVal *CompositeLiteralArrayVal
}

type CompositeLiteralArrayVal struct {
	LBRACE LBRACE
	CompositeLiteral *CompositeLiteral
	RBRACE RBRACE
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
	Int *TypeInt
	Char *TypeChar
	Custom *TypeCustom
	Pointer *TypePointer
}

type TypeInt struct {
	INT_TYPE INT_TYPE
}
type TypeChar struct {
	CHAR_TYPE CHAR_TYPE
}
type TypeCustom struct {
	IDEN IDEN
}
type TypePointer struct {
	ASTERISKS ASTERISKS
	Type *Type
}
type TypeInt struct {	INT_TYPE *INT_TYPE
type TypeChar struct {	CHAR_TYPE *CHAR_TYPE
type TypeCustom struct {	IDEN *IDEN
type TypePointer struct {	ASTERISKS *ASTERISKS	 *
type ReturnType string

const (
	ReturnTypeReturn ReturnType = "return"
)

type Return struct {
	Type ReturnType
	Return *ReturnReturn
}

type ReturnReturn struct {
	RETURN RETURN
	Return *Return
	SEMI SEMI
}
type ReturnReturn struct {	RETURN *RETURN	 *	SEMI *SEMI
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
	Int *ValueInt
	Variable *ValueVariable
	Str *ValueStr
	Char *ValueChar
	Cast *ValueCast
	CompLit *ValueCompLit
}

type ValueInt struct {
	NUM NUM
}
type ValueVariable struct {
	Value *Value
}
type ValueStr struct {
	STR STR
}
type ValueChar struct {
	CHAR CHAR
}
type ValueCast struct {
	LPAREN LPAREN
	Value *Value
	RPAREN RPAREN
	Value *Value
}
type ValueCompLit struct {
	Value *Value
}
type ValueInt struct {	NUM *NUM
type ValueVariable struct {	 *
type ValueStr struct {	STR *STR
type ValueChar struct {	CHAR *CHAR
type ValueCast struct {	LPAREN *LPAREN	 *	RPAREN *RPAREN	 *
type ValueCompLit struct {	 *
type CompEntriesType string

const (
	CompEntriesTypeEntries CompEntriesType = "entries"
)

type CompEntries struct {
	Type CompEntriesType
	Entries *CompEntriesEntries
}

type CompEntriesEntries struct {
	CompEntries *CompEntries
	CompEntries []*CompEntries
}
type CompEntriesEntries struct {	 *	 []*
type CommaExprType string

const (
	CommaExprTypeCommaExpr CommaExprType = "comma-expr"
)

type CommaExpr struct {
	Type CommaExprType
	CommaExpr *CommaExprCommaExpr
}

type CommaExprCommaExpr struct {
	COMMA COMMA
	CommaExpr *CommaExpr
}
type CommaExprCommaExpr struct {	COMMA *COMMA	 *
type VariableType string

const (
	VariableTypeVariable VariableType = "variable"
)

type Variable struct {
	Type VariableType
	Variable *VariableVariable
}

type VariableVariable struct {
	IDEN IDEN
	Variable []*Variable
}
type VariableVariable struct {	IDEN *IDEN	 []*
type ExprType string

const (
	ExprTypeSubExpr ExprType = "sub-expr"
	ExprTypeComp ExprType = "comp"
)

type Expr struct {
	Type ExprType
	SubExpr *ExprSubExpr
	Comp *ExprComp
}

type ExprSubExpr struct {
	Expr *Expr
}
type ExprComp struct {
	Expr *Expr
}
type ExprSubExpr struct {	 *
type ExprComp struct {	 *
type DecAssignType string

const (
	DecAssignTypeStandard DecAssignType = "standard"
)

type DecAssign struct {
	Type DecAssignType
	Standard *DecAssignStandard
}

type DecAssignStandard struct {
	DecAssign *DecAssign
	DecAssign *DecAssign
	EQUALS EQUALS
	DecAssign *DecAssign
}
type DecAssignStandard struct {	 *	 *	EQUALS *EQUALS	 *
type ParamsType string

const (
	ParamsTypeParams ParamsType = "params"
)

type Params struct {
	Type ParamsType
	Params *ParamsParams
}

type ParamsParams struct {
	Params *Params
	Params []*Params
}
type ParamsParams struct {	 *	 []*
type MainType string

const (
	MainTypeMain MainType = "main"
)

type Main struct {
	Type MainType
	Main *MainMain
}

type MainMain struct {
	Main []*Main
}
type MainMain struct {	 []*
type DecDefType string

const (
	DecDefTypeFuncDef DecDefType = "func-def"
	DecDefTypeDecAssign DecDefType = "dec-assign"
)

type DecDef struct {
	Type DecDefType
	FuncDef *DecDefFuncDef
	DecAssign *DecDefDecAssign
}

type DecDefFuncDef struct {
	DecDef *DecDef
	IDEN IDEN
	LPAREN LPAREN
	DecDef *DecDef
	RPAREN RPAREN
	LBRACE LBRACE
	DecDef []*DecDef
	DecDef *DecDef
	RBRACE RBRACE
}
type DecDefDecAssign struct {
	DecDef *DecDef
}
type DecDefFuncDef struct {	 *	IDEN *IDEN	LPAREN *LPAREN	 *	RPAREN *RPAREN	LBRACE *LBRACE	 []*	 *	RBRACE *RBRACE
type DecDefDecAssign struct {	 *
type ParamsDefType string

const (
	ParamsDefTypeParams ParamsDefType = "params"
)

type ParamsDef struct {
	Type ParamsDefType
	Params *ParamsDefParams
}

type ParamsDefParams struct {
	ParamsDef *ParamsDef
	ParamsDef []*ParamsDef
}
type ParamsDefParams struct {	 *	 []*
type CommaParamDefType string

const (
	CommaParamDefTypeParam CommaParamDefType = "param"
)

type CommaParamDef struct {
	Type CommaParamDefType
	Param *CommaParamDefParam
}

type CommaParamDefParam struct {
	COMMA COMMA
	CommaParamDef *CommaParamDef
}
type CommaParamDefParam struct {	COMMA *COMMA	 *
type ParamDefType string

const (
	ParamDefTypeParam ParamDefType = "param"
)

type ParamDef struct {
	Type ParamDefType
	Param *ParamDefParam
}

type ParamDefParam struct {
	ParamDef *ParamDef
	ParamDef *ParamDef
}
type ParamDefParam struct {	 *	 *
type SubExprType string

const (
	SubExprTypeValue SubExprType = "value"
	SubExprTypeParens SubExprType = "parens"
	SubExprTypeFuncCall SubExprType = "func-call"
)

type SubExpr struct {
	Type SubExprType
	Value *SubExprValue
	Parens *SubExprParens
	FuncCall *SubExprFuncCall
}

type SubExprValue struct {
	SubExpr *SubExpr
}
type SubExprParens struct {
	LPAREN LPAREN
	SubExpr *SubExpr
	RPAREN RPAREN
}
type SubExprFuncCall struct {
	IDEN IDEN
	LPAREN LPAREN
	SubExpr *SubExpr
	RPAREN RPAREN
}
type SubExprValue struct {	 *
type SubExprParens struct {	LPAREN *LPAREN	 *	RPAREN *RPAREN
type SubExprFuncCall struct {	IDEN *IDEN	LPAREN *LPAREN	 *	RPAREN *RPAREN