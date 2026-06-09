package ast
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

type MainType string

const (
	MainTypeMain MainType = "main"
)

type Main struct {
	Type MainType
	Main *Main_MainOption
}

type ControlType string

const (
	ControlTypeIf ControlType = "if"
	ControlTypeWhile ControlType = "while"
)

type Control struct {
	Type ControlType
	If *Control_IfOption
	While *Control_WhileOption
}

type ParamDefType string

const (
	ParamDefTypeParam ParamDefType = "param"
)

type ParamDef struct {
	Type ParamDefType
	Param *ParamDef_ParamOption
}

type CompExprType string

const (
	CompExprTypeCompExpr CompExprType = "comp-expr"
)

type CompExpr struct {
	Type CompExprType
	CompExpr *CompExpr_CompExprOption
}

type ParamsType string

const (
	ParamsTypeParams ParamsType = "params"
)

type Params struct {
	Type ParamsType
	Params *Params_ParamsOption
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

type StatementCommaType string

const (
	StatementCommaTypeStatement StatementCommaType = "statement"
)

type StatementComma struct {
	Type StatementCommaType
	Statement *StatementComma_StatementOption
}

type CommaParamDefType string

const (
	CommaParamDefTypeParam CommaParamDefType = "param"
)

type CommaParamDef struct {
	Type CommaParamDefType
	Param *CommaParamDef_ParamOption
}

type CommaExprType string

const (
	CommaExprTypeCommaExpr CommaExprType = "comma-expr"
)

type CommaExpr struct {
	Type CommaExprType
	CommaExpr *CommaExpr_CommaExprOption
}

type ArrayIndexType string

const (
	ArrayIndexTypeArrayIndex ArrayIndexType = "array-index"
)

type ArrayIndex struct {
	Type ArrayIndexType
	ArrayIndex *ArrayIndex_ArrayIndexOption
}

type OperatorType string

const (
	OperatorTypeLess OperatorType = "less"
	OperatorTypeEqual OperatorType = "equal"
	OperatorTypePlus OperatorType = "plus"
	OperatorTypeMinus OperatorType = "minus"
)

type Operator struct {
	Type OperatorType
	Less *Operator_LessOption
	Equal *Operator_EqualOption
	Plus *Operator_PlusOption
	Minus *Operator_MinusOption
}

type DecAssignType string

const (
	DecAssignTypeStandard DecAssignType = "standard"
)

type DecAssign struct {
	Type DecAssignType
	Standard *DecAssign_StandardOption
}

type CompositeLiteralType string

const (
	CompositeLiteralTypeArrayVal CompositeLiteralType = "array-val"
)

type CompositeLiteral struct {
	Type CompositeLiteralType
	ArrayVal *CompositeLiteral_ArrayValOption
}

type ElseType string

const (
	ElseTypeElse ElseType = "else"
)

type Else struct {
	Type ElseType
	Else *Else_ElseOption
}

type VariableType string

const (
	VariableTypeVariable VariableType = "variable"
)

type Variable struct {
	Type VariableType
	Variable *Variable_VariableOption
}

type CompEntriesType string

const (
	CompEntriesTypeEntries CompEntriesType = "entries"
)

type CompEntries struct {
	Type CompEntriesType
	Entries *CompEntries_EntriesOption
}

type TypeType string

const (
	TypeTypeInt TypeType = "int"
	TypeTypeChar TypeType = "char"
	TypeTypeVoid TypeType = "void"
	TypeTypeCustom TypeType = "custom"
	TypeTypePointer TypeType = "pointer"
)

type Type struct {
	Type TypeType
	Int *Type_IntOption
	Char *Type_CharOption
	Void *Type_VoidOption
	Custom *Type_CustomOption
	Pointer *Type_PointerOption
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

type BlockOrLineType string

const (
	BlockOrLineTypeBlock BlockOrLineType = "block"
	BlockOrLineTypeLine BlockOrLineType = "line"
)

type BlockOrLine struct {
	Type BlockOrLineType
	Block *BlockOrLine_BlockOption
	Line *BlockOrLine_LineOption
}

type ParamsDefType string

const (
	ParamsDefTypeParams ParamsDefType = "params"
)

type ParamsDef struct {
	Type ParamsDefType
	Params *ParamsDef_ParamsOption
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


type RBRACE string
type NUM string
type LPAREN string
type WHILE string
type CHAR_TYPE string
type LBRACE string
type EQUALS string
type IF string
type LSQUARE string
type RSQUARE string
type LESS_THAN string
type PLUS string
type MINUS string
type CHAR string
type ELSE string
type INT_TYPE string
type ASTERISKS string
type STR string
type RPAREN string
type IDEN string
type RETURN string
type SEMI string
type COMMA string
type EEQUALS string
type VOID string
type Control_IfOption struct {
	IF IF
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
	LBRACE LBRACE
	BlockOrLine *BlockOrLine
	RBRACE RBRACE
	Else *Else
}

type ArrayIndex_ArrayIndexOption struct {
	LSQUARE LSQUARE
	NUM NUM
	RSQUARE RSQUARE
}

type BlockOrLine_LineOption struct {
	Line *Line
}

type Statement_DecAssignOption struct {
	DecAssign *DecAssign
}

type Else_ElseOption struct {
	ELSE ELSE
	BlockOrLine *BlockOrLine
}

type Type_CustomOption struct {
	IDEN IDEN
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

type ParamsDef_ParamsOption struct {
	ParamDef *ParamDef
	CommaParamDef []*CommaParamDef
}

type Value_StrOption struct {
	STR STR
}

type Value_CompLitOption struct {
	CompositeLiteral *CompositeLiteral
}

type Statement_ExprOption struct {
	Expr *Expr
}

type StatementComma_StatementOption struct {
	Statement *Statement
	SEMI SEMI
}

type CommaParamDef_ParamOption struct {
	COMMA COMMA
	ParamDef *ParamDef
}

type CompositeLiteral_ArrayValOption struct {
	LBRACE LBRACE
	CompEntries *CompEntries
	RBRACE RBRACE
}

type Type_IntOption struct {
	INT_TYPE INT_TYPE
}

type SubExpr_ParensOption struct {
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
}

type Control_WhileOption struct {
	WHILE WHILE
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
	LBRACE LBRACE
	Line []*Line
	RBRACE RBRACE
}

type Line_ControlOption struct {
	Control *Control
}

type Type_CharOption struct {
	CHAR_TYPE CHAR_TYPE
}

type DecDef_DecAssignOption struct {
	DecAssign *DecAssign
}

type Value_IntOption struct {
	NUM NUM
}

type SubExpr_ValueOption struct {
	Value *Value
}

type Expr_SubExprOption struct {
	SubExpr *SubExpr
}

type CommaExpr_CommaExprOption struct {
	COMMA COMMA
	Expr *Expr
}

type Operator_PlusOption struct {
	PLUS PLUS
}

type Operator_MinusOption struct {
	MINUS MINUS
}

type Type_VoidOption struct {
	VOID VOID
}

type Type_PointerOption struct {
	ASTERISKS ASTERISKS
	Type *Type
}

type BlockOrLine_BlockOption struct {
	LBRACE LBRACE
	Line []*Line
	RBRACE RBRACE
}

type CompExpr_CompExprOption struct {
	SubExpr *SubExpr
	Operator *Operator
	Expr *Expr
}

type Params_ParamsOption struct {
	Expr *Expr
	CommaExpr []*CommaExpr
}

type Operator_EqualOption struct {
	EEQUALS EEQUALS
}

type CompEntries_EntriesOption struct {
	Expr *Expr
	CommaExpr []*CommaExpr
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

type Statement_AssignOption struct {
	Variable *Variable
	EQUALS EQUALS
	Expr *Expr
}

type ParamDef_ParamOption struct {
	Type *Type
	Variable *Variable
}

type Line_StatementOption struct {
	StatementComma *StatementComma
}

type Operator_LessOption struct {
	LESS_THAN LESS_THAN
}

type DecAssign_StandardOption struct {
	Type *Type
	Variable *Variable
	EQUALS EQUALS
	Expr *Expr
}

type Variable_VariableOption struct {
	IDEN IDEN
	ArrayIndex []*ArrayIndex
}

type Value_VariableOption struct {
	Variable *Variable
}

type Value_CharOption struct {
	CHAR CHAR
}

type Value_CastOption struct {
	LPAREN LPAREN
	Type *Type
	RPAREN RPAREN
	Value *Value
}

type Statement_VarDecOption struct {
	Type *Type
	Variable *Variable
}

type Statement_ReturnOption struct {
	RETURN RETURN
	Expr *Expr
}

type Main_MainOption struct {
	DecDef []*DecDef
}


