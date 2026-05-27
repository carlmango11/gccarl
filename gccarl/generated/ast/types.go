package ast
type MainType string

const (
	MainTypeMain MainType = "main"
)

type Main struct {
	Type MainType
	Main *Main_MainOption
}

type ElseType string

const (
	ElseTypeElse ElseType = "else"
)

type Else struct {
	Type ElseType
	Else *Else_ElseOption
}

type StatementCommaType string

const (
	StatementCommaTypeStatement StatementCommaType = "statement"
)

type StatementComma struct {
	Type StatementCommaType
	Statement *StatementComma_StatementOption
}

type CommaExprType string

const (
	CommaExprTypeCommaExpr CommaExprType = "comma-expr"
)

type CommaExpr struct {
	Type CommaExprType
	CommaExpr *CommaExpr_CommaExprOption
}

type TypeType string

const (
	TypeTypeInt TypeType = "int"
	TypeTypeChar TypeType = "char"
	TypeTypeCustom TypeType = "custom"
	TypeTypePointer TypeType = "pointer"
)

type Type struct {
	Type TypeType
	Int *Type_IntOption
	Char *Type_CharOption
	Custom *Type_CustomOption
	Pointer *Type_PointerOption
}

type LineType string

const (
	LineTypeControl LineType = "control"
	LineTypeStatement LineType = "statement"
)

type Line struct {
	Type LineType
	Control *Line_ControlOption
	Statement *Line_StatementOption
}

type ElseIfType string

const (
	ElseIfTypeElseIf ElseIfType = "else-if"
)

type ElseIf struct {
	Type ElseIfType
	ElseIf *ElseIf_ElseIfOption
}

type ParamsDefType string

const (
	ParamsDefTypeParams ParamsDefType = "params"
)

type ParamsDef struct {
	Type ParamsDefType
	Params *ParamsDef_ParamsOption
}

type VariableType string

const (
	VariableTypeVariable VariableType = "variable"
)

type Variable struct {
	Type VariableType
	Variable *Variable_VariableOption
}

type SubExprType string

const (
	SubExprTypeValue SubExprType = "value"
	SubExprTypeParens SubExprType = "parens"
	SubExprTypeFuncCall SubExprType = "func-call"
)

type SubExpr struct {
	Type SubExprType
	Value *SubExpr_ValueOption
	Parens *SubExpr_ParensOption
	FuncCall *SubExpr_FuncCallOption
}

type CompExprType string

const (
	CompExprTypeAdd CompExprType = "add"
	CompExprTypeSub CompExprType = "sub"
	CompExprTypeIsEqual CompExprType = "is-equal"
)

type CompExpr struct {
	Type CompExprType
	Add *CompExpr_AddOption
	Sub *CompExpr_SubOption
	IsEqual *CompExpr_IsEqualOption
}

type CompositeLiteralType string

const (
	CompositeLiteralTypeArrayVal CompositeLiteralType = "array-val"
)

type CompositeLiteral struct {
	Type CompositeLiteralType
	ArrayVal *CompositeLiteral_ArrayValOption
}

type CompEntriesType string

const (
	CompEntriesTypeEntries CompEntriesType = "entries"
)

type CompEntries struct {
	Type CompEntriesType
	Entries *CompEntries_EntriesOption
}

type DecDefType string

const (
	DecDefTypeFuncDef DecDefType = "func-def"
	DecDefTypeDecAssign DecDefType = "dec-assign"
)

type DecDef struct {
	Type DecDefType
	FuncDef *DecDef_FuncDefOption
	DecAssign *DecDef_DecAssignOption
}

type ControlType string

const (
	ControlTypeIf ControlType = "if"
)

type Control struct {
	Type ControlType
	If *Control_IfOption
}

type ParamDefType string

const (
	ParamDefTypeParam ParamDefType = "param"
)

type ParamDef struct {
	Type ParamDefType
	Param *ParamDef_ParamOption
}

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
	Int *Value_IntOption
	Variable *Value_VariableOption
	Str *Value_StrOption
	Char *Value_CharOption
	Cast *Value_CastOption
	CompLit *Value_CompLitOption
}

type ArrayIndexType string

const (
	ArrayIndexTypeArrayIndex ArrayIndexType = "array-index"
)

type ArrayIndex struct {
	Type ArrayIndexType
	ArrayIndex *ArrayIndex_ArrayIndexOption
}

type ExprType string

const (
	ExprTypeSubExpr ExprType = "sub-expr"
	ExprTypeComp ExprType = "comp"
)

type Expr struct {
	Type ExprType
	SubExpr *Expr_SubExprOption
	Comp *Expr_CompOption
}

type StatementType string

const (
	StatementTypeDecAssign StatementType = "dec-assign"
	StatementTypeVarDec StatementType = "var-dec"
	StatementTypeAssign StatementType = "assign"
	StatementTypeExpr StatementType = "expr"
	StatementTypeReturn StatementType = "return"
)

type Statement struct {
	Type StatementType
	DecAssign *Statement_DecAssignOption
	VarDec *Statement_VarDecOption
	Assign *Statement_AssignOption
	Expr *Statement_ExprOption
	Return *Statement_ReturnOption
}

type ParamsType string

const (
	ParamsTypeParams ParamsType = "params"
)

type Params struct {
	Type ParamsType
	Params *Params_ParamsOption
}

type CommaParamDefType string

const (
	CommaParamDefTypeParam CommaParamDefType = "param"
)

type CommaParamDef struct {
	Type CommaParamDefType
	Param *CommaParamDef_ParamOption
}

type DecAssignType string

const (
	DecAssignTypeStandard DecAssignType = "standard"
)

type DecAssign struct {
	Type DecAssignType
	Standard *DecAssign_StandardOption
}


type RPAREN string
type LSQUARE string
type COMMA string
type INT_TYPE string
type IDEN string
type ASTERISKS string
type IF string
type MINUS string
type STR string
type CHAR string
type EQUALS string
type RBRACE string
type NUM string
type RSQUARE string
type RETURN string
type SEMI string
type CHAR_TYPE string
type LPAREN string
type ELSE string
type LBRACE string
type PLUS string
type EEQUALS string
type ArrayIndex_ArrayIndexOption struct {
	LSQUARE LSQUARE
	NUM NUM
	RSQUARE RSQUARE
}

type Else_ElseOption struct {
	ELSE ELSE
	ElseIf *ElseIf
	LBRACE LBRACE
	Line []*Line
	RBRACE RBRACE
}

type ParamsDef_ParamsOption struct {
	ParamDef *ParamDef
	CommaParamDef []*CommaParamDef
}

type SubExpr_ValueOption struct {
	Value *Value
}

type CompositeLiteral_ArrayValOption struct {
	LBRACE LBRACE
	CompEntries *CompEntries
	RBRACE RBRACE
}

type Value_CastOption struct {
	LPAREN LPAREN
	Type *Type
	RPAREN RPAREN
	Value *Value
}

type Type_CharOption struct {
	CHAR_TYPE CHAR_TYPE
}

type ElseIf_ElseIfOption struct {
	IF IF
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
}

type Value_CompLitOption struct {
	CompositeLiteral *CompositeLiteral
}

type DecAssign_StandardOption struct {
	Type *Type
	Variable *Variable
	EQUALS EQUALS
	Expr *Expr
}

type Line_StatementOption struct {
	StatementComma *StatementComma
}

type DecDef_DecAssignOption struct {
	DecAssign *DecAssign
}

type Control_IfOption struct {
	IF IF
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
	LBRACE LBRACE
	Line []*Line
	RBRACE RBRACE
	Else *Else
}

type Value_StrOption struct {
	STR STR
}

type Expr_CompOption struct {
	CompExpr *CompExpr
}

type SubExpr_FuncCallOption struct {
	IDEN IDEN
	LPAREN LPAREN
	Params *Params
	RPAREN RPAREN
}

type Value_CharOption struct {
	CHAR CHAR
}

type Expr_SubExprOption struct {
	SubExpr *SubExpr
}

type Statement_VarDecOption struct {
	Type *Type
	Variable *Variable
}

type Type_IntOption struct {
	INT_TYPE INT_TYPE
}

type Type_PointerOption struct {
	ASTERISKS ASTERISKS
	Type *Type
}

type SubExpr_ParensOption struct {
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
}

type CompExpr_SubOption struct {
	SubExpr *SubExpr
	MINUS MINUS
	Expr *Expr
}

type CompEntries_EntriesOption struct {
	Expr *Expr
	CommaExpr []*CommaExpr
}

type Value_IntOption struct {
	NUM NUM
}

type Statement_DecAssignOption struct {
	DecAssign *DecAssign
}

type Statement_AssignOption struct {
	Variable *Variable
	EQUALS EQUALS
	Expr *Expr
}

type Params_ParamsOption struct {
	Expr *Expr
	CommaExpr []*CommaExpr
}

type CommaExpr_CommaExprOption struct {
	COMMA COMMA
	Expr *Expr
}

type Variable_VariableOption struct {
	IDEN IDEN
	ArrayIndex []*ArrayIndex
}

type Statement_ExprOption struct {
	Expr *Expr
}

type Statement_ReturnOption struct {
	RETURN RETURN
	Expr *Expr
}

type Type_CustomOption struct {
	IDEN IDEN
}

type CompExpr_AddOption struct {
	SubExpr *SubExpr
	PLUS PLUS
	Expr *Expr
}

type ParamDef_ParamOption struct {
	Type *Type
	Variable *Variable
}

type CommaParamDef_ParamOption struct {
	COMMA COMMA
	ParamDef *ParamDef
}

type Main_MainOption struct {
	DecDef []*DecDef
}

type StatementComma_StatementOption struct {
	Statement *Statement
	SEMI SEMI
}

type Line_ControlOption struct {
	Control *Control
}

type CompExpr_IsEqualOption struct {
	SubExpr *SubExpr
	EEQUALS EEQUALS
	Expr *Expr
}

type DecDef_FuncDefOption struct {
	Type *Type
	IDEN IDEN
	LPAREN LPAREN
	ParamsDef *ParamsDef
	RPAREN RPAREN
	LBRACE LBRACE
	Line []*Line
	RBRACE RBRACE
}

type Value_VariableOption struct {
	Variable *Variable
}


