
package ast

var MainNode = n0

var n6 = &Type {
	Type: TypeTypeCustom,
	Custom: &Type_CustomOption{
		IDEN: "Inner",
	},
}

var n7 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "i",
	},
}

var n5 = &VarDec {
	Type: VarDecTypeVarDec,
	VarDec: &VarDec_VarDecOption{
		Type: n6,
		VariableDef: n7,
	},
}

var n4 = &VarDecColon {
	Type: VarDecColonTypeC,
	C: &VarDecColon_COption{
		VarDec: n5,
		SEMI: ";",
	},
}

var n3 = &StructBlock {
	Type: StructBlockTypeBlock,
	Block: &StructBlock_BlockOption{
		VarDecColon: n4,
	},
}

var n2 = &TypeDef {
	Type: TypeDefTypeStructDef,
	StructDef: &TypeDef_StructDefOption{
		STRUCT: "struct",
		IDEN: "Person",
		LBRACE: "{",
		StructBlock: n3,
		RBRACE: "}",
	},
}

var n1 = &DecDef {
	Type: DecDefTypeTypeDef,
	TypeDef: &DecDef_TypeDefOption{
		TypeDef: n2,
	},
}

var n13 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n14 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "y",
	},
}

var n12 = &VarDec {
	Type: VarDecTypeVarDec,
	VarDec: &VarDec_VarDecOption{
		Type: n13,
		VariableDef: n14,
	},
}

var n11 = &VarDecColon {
	Type: VarDecColonTypeC,
	C: &VarDecColon_COption{
		VarDec: n12,
		SEMI: ";",
	},
}

var n10 = &StructBlock {
	Type: StructBlockTypeBlock,
	Block: &StructBlock_BlockOption{
		VarDecColon: n11,
	},
}

var n9 = &TypeDef {
	Type: TypeDefTypeStructDef,
	StructDef: &TypeDef_StructDefOption{
		STRUCT: "struct",
		IDEN: "Inner",
		LBRACE: "{",
		StructBlock: n10,
		RBRACE: "}",
	},
}

var n8 = &DecDef {
	Type: DecDefTypeTypeDef,
	TypeDef: &DecDef_TypeDefOption{
		TypeDef: n9,
	},
}

var n16 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n21 = &Type {
	Type: TypeTypeStruct,
	Struct: &Type_StructOption{
		STRUCT: "struct",
		IDEN: "Person",
	},
}

var n22 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "p",
	},
}

var n26 = &Type {
	Type: TypeTypeCustom,
	Custom: &Type_CustomOption{
		IDEN: "Person",
	},
}

var n31 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "4",
	},
}

var n30 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n31,
	},
}

var n29 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n30,
	},
}

var n28 = &CompEntry {
	Type: CompEntryTypeLabelled,
	Labelled: &CompEntry_LabelledOption{
		FULL_STOP: ".",
		IDEN: "i",
		EQUALS: "=",
		Expr: n29,
	},
}

var n36 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "45",
	},
}

var n35 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n36,
	},
}

var n34 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n35,
	},
}

var n33 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Expr: n34,
	},
}

var n32 = &CommaCompEntry {
	Type: CommaCompEntryTypeE,
	E: &CommaCompEntry_EOption{
		COMMA: ",",
		CompEntry: n33,
	},
}

var n27 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n28,
		CommaCompEntry: []*CommaCompEntry {
			n32,
		},
		COMMA: ",",
	},
}

var n25 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		LPAREN: "(",
		Type: n26,
		RPAREN: ")",
		LBRACE: "{",
		CompEntries: n27,
		RBRACE: "}",
	},
}

var n24 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n25,
	},
}

var n23 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n24,
	},
}

var n20 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n21,
		VariableDef: n22,
		EQUALS: "=",
		Expr: n23,
	},
}

var n19 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n20,
	},
}

var n18 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n19,
		SEMI: ";",
	},
}

var n17 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n18,
	},
}

var n15 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n16,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n17,
		},
		RBRACE: "}",
	},
}

var n38 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n41 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n43 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n42 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "msg",
		ArrayIndexDef: []*ArrayIndexDef {
			n43,
		},
	},
}

var n40 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n41,
		VariableDef: n42,
	},
}

var n46 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n47 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "len",
	},
}

var n45 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n46,
		VariableDef: n47,
	},
}

var n44 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n45,
	},
}

var n39 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n40,
		CommaParamDef: []*CommaParamDef {
			n44,
		},
	},
}

var n56 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n55 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n56,
	},
}

var n54 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n55,
	},
}

var n60 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n59 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n60,
	},
}

var n58 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n59,
	},
}

var n57 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n58,
	},
}

var n66 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "msg",
	},
}

var n65 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		SubVariableAccess: n66,
	},
}

var n64 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n65,
	},
}

var n63 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n64,
	},
}

var n62 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n63,
	},
}

var n61 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n62,
	},
}

var n72 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "len",
	},
}

var n71 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		SubVariableAccess: n72,
	},
}

var n70 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n71,
	},
}

var n69 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n70,
	},
}

var n68 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n69,
	},
}

var n67 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n68,
	},
}

var n53 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n54,
		CommaExpr: []*CommaExpr {
			n57,
			n61,
			n67,
		},
	},
}

var n52 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n53,
		RPAREN: ")",
	},
}

var n51 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n52,
	},
}

var n50 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n51,
	},
}

var n49 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n50,
		SEMI: ";",
	},
}

var n48 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n49,
	},
}

var n37 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n38,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n39,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n48,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n8,
			n15,
			n37,
		},
	},
}
