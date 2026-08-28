package ast
type DecDefType string

const (
	DecDefTypeFuncDef DecDefType = "func-def"
	DecDefTypeDecAssign DecDefType = "dec-assign"
	DecDefTypeTypeDef DecDefType = "type-def"
)

type DecDef struct {
	Type DecDefType
	FuncDef *DecDef_FuncDefOption
	DecAssign *DecDef_DecAssignOption
	TypeDef *DecDef_TypeDefOption
}

type CompExprType string

const (
	CompExprTypeCompExpr CompExprType = "comp-expr"
)

type CompExpr struct {
	Type CompExprType
	CompExpr *CompExpr_CompExprOption
}

type VarDecType string

const (
	VarDecTypeVarDec VarDecType = "var-dec"
)

type VarDec struct {
	Type VarDecType
	VarDec *VarDec_VarDecOption
}

type MainType string

const (
	MainTypeMain MainType = "main"
)

type Main struct {
	Type MainType
	Main *Main_MainOption
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

type ArrayIndexAccessType string

const (
	ArrayIndexAccessTypeArrayIndex ArrayIndexAccessType = "array-index"
)

type ArrayIndexAccess struct {
	Type ArrayIndexAccessType
	ArrayIndex *ArrayIndexAccess_ArrayIndexOption
}

type OperatorType string

const (
	OperatorTypeLess OperatorType = "less"
	OperatorTypeEqual OperatorType = "equal"
	OperatorTypePlus OperatorType = "plus"
	OperatorTypeMinus OperatorType = "minus"
	OperatorTypeAssign OperatorType = "assign"
)

type Operator struct {
	Type OperatorType
	Less *Operator_LessOption
	Equal *Operator_EqualOption
	Plus *Operator_PlusOption
	Minus *Operator_MinusOption
	Assign *Operator_AssignOption
}

type CommaCompEntryType string

const (
	CommaCompEntryTypeE CommaCompEntryType = "e"
)

type CommaCompEntry struct {
	Type CommaCompEntryType
	E *CommaCompEntry_EOption
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

type StructBlockType string

const (
	StructBlockTypeBlock StructBlockType = "block"
)

type StructBlock struct {
	Type StructBlockType
	Block *StructBlock_BlockOption
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

type TypeType string

const (
	TypeTypeInt TypeType = "int"
	TypeTypeChar TypeType = "char"
	TypeTypeVoid TypeType = "void"
	TypeTypeCustom TypeType = "custom"
	TypeTypeStruct TypeType = "struct"
)

type Type struct {
	Type TypeType
	Int *Type_IntOption
	Char *Type_CharOption
	Void *Type_VoidOption
	Custom *Type_CustomOption
	Struct *Type_StructOption
}

type ArrayEntriesType string

const (
	ArrayEntriesTypeEntries ArrayEntriesType = "entries"
)

type ArrayEntries struct {
	Type ArrayEntriesType
	Entries *ArrayEntries_EntriesOption
}

type InnerSubVariableAccessType string

const (
	InnerSubVariableAccessTypeDot InnerSubVariableAccessType = "dot"
	InnerSubVariableAccessTypeArrow InnerSubVariableAccessType = "arrow"
)

type InnerSubVariableAccess struct {
	Type InnerSubVariableAccessType
	Dot *InnerSubVariableAccess_DotOption
	Arrow *InnerSubVariableAccess_ArrowOption
}

type DecAssignType string

const (
	DecAssignTypeStandard DecAssignType = "standard"
)

type DecAssign struct {
	Type DecAssignType
	Standard *DecAssign_StandardOption
}

type CompEntryType string

const (
	CompEntryTypeLabelled CompEntryType = "labelled"
	CompEntryTypeAnon CompEntryType = "anon"
)

type CompEntry struct {
	Type CompEntryType
	Labelled *CompEntry_LabelledOption
	Anon *CompEntry_AnonOption
}

type TypeDefType string

const (
	TypeDefTypeStructDef TypeDefType = "struct-def"
)

type TypeDef struct {
	Type TypeDefType
	StructDef *TypeDef_StructDefOption
}

type VarDecColonType string

const (
	VarDecColonTypeC VarDecColonType = "c"
)

type VarDecColon struct {
	Type VarDecColonType
	C *VarDecColon_COption
}

type ParamsDefType string

const (
	ParamsDefTypeParams ParamsDefType = "params"
)

type ParamsDef struct {
	Type ParamsDefType
	Params *ParamsDef_ParamsOption
}

type SubVariableAccessType string

const (
	SubVariableAccessTypeV SubVariableAccessType = "v"
)

type SubVariableAccess struct {
	Type SubVariableAccessType
	V *SubVariableAccess_VOption
}

type StatementType string

const (
	StatementTypeDecAssign StatementType = "dec-assign"
	StatementTypeVarDec StatementType = "var-dec"
	StatementTypeExpr StatementType = "expr"
	StatementTypeReturn StatementType = "return"
)

type Statement struct {
	Type StatementType
	DecAssign *Statement_DecAssignOption
	VarDec *Statement_VarDecOption
	Expr *Statement_ExprOption
	Return *Statement_ReturnOption
}

type InitialiserType string

const (
	InitialiserTypeExpr InitialiserType = "expr"
	InitialiserTypeList InitialiserType = "list"
)

type Initialiser struct {
	Type InitialiserType
	Expr *Initialiser_ExprOption
	List *Initialiser_ListOption
}

type VarDecCommaType string

const (
	VarDecCommaTypeDecComma VarDecCommaType = "dec-comma"
)

type VarDecComma struct {
	Type VarDecCommaType
	DecComma *VarDecComma_DecCommaOption
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

type ArrayIndexDefType string

const (
	ArrayIndexDefTypeArrayIndex ArrayIndexDefType = "array-index"
)

type ArrayIndexDef struct {
	Type ArrayIndexDefType
	ArrayIndex *ArrayIndexDef_ArrayIndexOption
}

type SubExprType string

const (
	SubExprTypeValue SubExprType = "value"
	SubExprTypeVariable SubExprType = "variable"
	SubExprTypeParens SubExprType = "parens"
	SubExprTypeDeref SubExprType = "deref"
	SubExprTypeAddressOf SubExprType = "address-of"
	SubExprTypeFuncCall SubExprType = "func-call"
)

type SubExpr struct {
	Type SubExprType
	Value *SubExpr_ValueOption
	Variable *SubExpr_VariableOption
	Parens *SubExpr_ParensOption
	Deref *SubExpr_DerefOption
	AddressOf *SubExpr_AddressOfOption
	FuncCall *SubExpr_FuncCallOption
}

type CompEntriesType string

const (
	CompEntriesTypeEntries CompEntriesType = "entries"
)

type CompEntries struct {
	Type CompEntriesType
	Entries *CompEntries_EntriesOption
}

type ParamsType string

const (
	ParamsTypeParams ParamsType = "params"
)

type Params struct {
	Type ParamsType
	Params *Params_ParamsOption
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
	ValueTypeStr ValueType = "str"
	ValueTypeChar ValueType = "char"
	ValueTypeCast ValueType = "cast"
	ValueTypeCompLit ValueType = "comp-lit"
)

type Value struct {
	Type ValueType
	Int *Value_IntOption
	Str *Value_StrOption
	Char *Value_CharOption
	Cast *Value_CastOption
	CompLit *Value_CompLitOption
}

type VariableDefType string

const (
	VariableDefTypePointer VariableDefType = "pointer"
	VariableDefTypeVariable VariableDefType = "variable"
)

type VariableDef struct {
	Type VariableDefType
	Pointer *VariableDef_PointerOption
	Variable *VariableDef_VariableOption
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
type RSQUARE string
type INT_TYPE string
type AMPERSAND string
type LESS_THAN string
type EEQUALS string
type LPAREN string
type NUM string
type PLUS string
type ELSE string
type FULL_STOP string
type RIGHT_ARROW string
type WHILE string
type CHAR string
type STRUCT string
type RPAREN string
type LBRACE string
type IDEN string
type SEMI string
type RETURN string
type IF string
type STR string
type LSQUARE string
type MINUS string
type COMMA string
type CHAR_TYPE string
type VOID string
type ASTERISKS string
type EQUALS string
type StructBlock_BlockOption struct {
	VarDecColon *VarDecColon
	VarDecComma []*VarDecComma
}

type DecDef_TypeDefOption struct {
	TypeDef *TypeDef
}

type Operator_AssignOption struct {
	EQUALS EQUALS
}

type SubExpr_FuncCallOption struct {
	IDEN IDEN
	LPAREN LPAREN
	Params *Params
	RPAREN RPAREN
}

type TypeDef_StructDefOption struct {
	STRUCT STRUCT
	IDEN IDEN
	LBRACE LBRACE
	StructBlock *StructBlock
	RBRACE RBRACE
}

type ArrayIndexDef_ArrayIndexOption struct {
	LSQUARE LSQUARE
	NUM NUM
	RSQUARE RSQUARE
}

type CommaExpr_CommaExprOption struct {
	COMMA COMMA
	Expr *Expr
}

type Type_VoidOption struct {
	VOID VOID
}

type ParamsDef_ParamsOption struct {
	ParamDef *ParamDef
	CommaParamDef []*CommaParamDef
}

type Operator_LessOption struct {
	LESS_THAN LESS_THAN
}

type CompEntry_AnonOption struct {
	Expr *Expr
}

type Statement_ExprOption struct {
	Expr *Expr
}

type Statement_ReturnOption struct {
	RETURN RETURN
	Expr *Expr
}

type SubExpr_DerefOption struct {
	ASTERISKS ASTERISKS
	Expr *Expr
}

type CompEntries_EntriesOption struct {
	CompEntry *CompEntry
	CommaCompEntry []*CommaCompEntry
	COMMA COMMA
}

type VariableDef_VariableOption struct {
	IDEN IDEN
	ArrayIndexDef []*ArrayIndexDef
}

type VarDec_VarDecOption struct {
	Type *Type
	VariableDef *VariableDef
}

type Type_CharOption struct {
	CHAR_TYPE CHAR_TYPE
}

type CompEntry_LabelledOption struct {
	FULL_STOP FULL_STOP
	IDEN IDEN
	EQUALS EQUALS
	Expr *Expr
}

type Value_IntOption struct {
	NUM NUM
}

type Value_CompLitOption struct {
	LPAREN LPAREN
	Type *Type
	ArrayIndexDef []*ArrayIndexDef
	RPAREN RPAREN
	LBRACE LBRACE
	CompEntries *CompEntries
	RBRACE RBRACE
}

type Expr_SubExprOption struct {
	SubExpr *SubExpr
}

type VariableDef_PointerOption struct {
	ASTERISKS ASTERISKS
	VariableDef *VariableDef
}

type Operator_EqualOption struct {
	EEQUALS EEQUALS
}

type CommaCompEntry_EOption struct {
	COMMA COMMA
	CompEntry *CompEntry
}

type CommaParamDef_ParamOption struct {
	COMMA COMMA
	ParamDef *ParamDef
}

type VarDecColon_COption struct {
	VarDec *VarDec
	SEMI SEMI
}

type Statement_VarDecOption struct {
	VarDec *VarDec
}

type Type_StructOption struct {
	STRUCT STRUCT
	IDEN IDEN
}

type SubVariableAccess_VOption struct {
	IDEN IDEN
	ArrayIndexAccess []*ArrayIndexAccess
}

type Initialiser_ListOption struct {
	LBRACE LBRACE
	CompEntries *CompEntries
	RBRACE RBRACE
}

type Control_IfOption struct {
	IF IF
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
	BlockOrLine *BlockOrLine
	Else *Else
}

type Line_ControlOption struct {
	Control *Control
}

type Operator_PlusOption struct {
	PLUS PLUS
}

type StatementComma_StatementOption struct {
	Statement *Statement
	SEMI SEMI
}

type Type_IntOption struct {
	INT_TYPE INT_TYPE
}

type Type_CustomOption struct {
	IDEN IDEN
}

type Initialiser_ExprOption struct {
	Expr *Expr
}

type ArrayIndexAccess_ArrayIndexOption struct {
	LSQUARE LSQUARE
	NUM NUM
	RSQUARE RSQUARE
}

type Operator_MinusOption struct {
	MINUS MINUS
}

type SubExpr_ParensOption struct {
	LPAREN LPAREN
	Expr *Expr
	RPAREN RPAREN
}

type SubExpr_AddressOfOption struct {
	AMPERSAND AMPERSAND
	Expr *Expr
}

type Else_ElseOption struct {
	ELSE ELSE
	BlockOrLine *BlockOrLine
}

type InnerSubVariableAccess_DotOption struct {
	FULL_STOP FULL_STOP
	SubVariableAccess *SubVariableAccess
}

type DecAssign_StandardOption struct {
	Type *Type
	VariableDef *VariableDef
	EQUALS EQUALS
	Initialiser *Initialiser
}

type VarDecComma_DecCommaOption struct {
	COMMA COMMA
	VarDecColon *VarDecColon
}

type SubExpr_ValueOption struct {
	Value *Value
}

type Value_CastOption struct {
	LPAREN LPAREN
	Type *Type
	RPAREN RPAREN
	Value *Value
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

type DecDef_DecAssignOption struct {
	DecAssign *DecAssign
}

type BlockOrLine_LineOption struct {
	Line *Line
}

type InnerSubVariableAccess_ArrowOption struct {
	RIGHT_ARROW RIGHT_ARROW
	SubVariableAccess *SubVariableAccess
}

type Statement_DecAssignOption struct {
	DecAssign *DecAssign
}

type Params_ParamsOption struct {
	Expr *Expr
	CommaExpr []*CommaExpr
}

type ParamDef_ParamOption struct {
	Type *Type
	VariableDef *VariableDef
}

type Value_StrOption struct {
	STR STR
}

type CompExpr_CompExprOption struct {
	SubExpr *SubExpr
	Operator *Operator
	Expr *Expr
}

type Main_MainOption struct {
	DecDef []*DecDef
}

type Line_StatementOption struct {
	StatementComma *StatementComma
}

type ArrayEntries_EntriesOption struct {
	Expr *Expr
	CommaExpr []*CommaExpr
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

type Expr_CompOption struct {
	CompExpr *CompExpr
}

type SubExpr_VariableOption struct {
	SubVariableAccess *SubVariableAccess
	InnerSubVariableAccess []*InnerSubVariableAccess
}

type Value_CharOption struct {
	CHAR CHAR
}

type BlockOrLine_BlockOption struct {
	LBRACE LBRACE
	Line []*Line
	RBRACE RBRACE
}


