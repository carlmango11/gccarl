
package ast

var MainNode = n0

var n6 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n7 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "age",
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

var n9 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n14 = &Type {
	Type: TypeTypeStruct,
	Struct: &Type_StructOption{
		STRUCT: "struct",
		IDEN: "Person",
	},
}

var n15 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "p",
	},
}

var n21 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'X'",
	},
}

var n20 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n21,
	},
}

var n19 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n20,
	},
}

var n18 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Expr: n19,
	},
}

var n17 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n18,
	},
}

var n16 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n17,
		RBRACE: "}",
	},
}

var n13 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n14,
		VariableDef: n15,
		EQUALS: "=",
		Initialiser: n16,
	},
}

var n12 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n13,
	},
}

var n11 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n12,
		SEMI: ";",
	},
}

var n10 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n11,
	},
}

var n28 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "p",
	},
}

var n30 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "age",
	},
}

var n29 = &InnerSubVariableAccess {
	Type: InnerSubVariableAccessTypeDot,
	Dot: &InnerSubVariableAccess_DotOption{
		FULL_STOP: ".",
		SubVariableAccess: n30,
	},
}

var n27 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n28,
		InnerSubVariableAccess: []*InnerSubVariableAccess {
			n29,
		},
	},
}

var n31 = &Operator {
	Type: OperatorTypeAssign,
	Assign: &Operator_AssignOption{
		EQUALS: "=",
	},
}

var n34 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
	},
}

var n33 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n34,
	},
}

var n32 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n33,
	},
}

var n26 = &CompExpr {
	Type: CompExprTypeCompExpr,
	CompExpr: &CompExpr_CompExprOption{
		SubExpr: n27,
		Operator: n31,
		Expr: n32,
	},
}

var n25 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n26,
	},
}

var n24 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n25,
	},
}

var n23 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n24,
		SEMI: ";",
	},
}

var n22 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n23,
	},
}

var n39 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n41 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n40 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "m",
		ArrayIndexDef: []*ArrayIndexDef {
			n41,
		},
	},
}

var n47 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "p",
	},
}

var n49 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "age",
	},
}

var n48 = &InnerSubVariableAccess {
	Type: InnerSubVariableAccessTypeDot,
	Dot: &InnerSubVariableAccess_DotOption{
		FULL_STOP: ".",
		SubVariableAccess: n49,
	},
}

var n46 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n47,
		InnerSubVariableAccess: []*InnerSubVariableAccess {
			n48,
		},
	},
}

var n45 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n46,
	},
}

var n44 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Expr: n45,
	},
}

var n43 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n44,
	},
}

var n42 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n43,
		RBRACE: "}",
	},
}

var n38 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n39,
		VariableDef: n40,
		EQUALS: "=",
		Initialiser: n42,
	},
}

var n37 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n38,
	},
}

var n36 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n37,
		SEMI: ";",
	},
}

var n35 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n36,
	},
}

var n58 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "m",
	},
}

var n57 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n58,
	},
}

var n56 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n57,
	},
}

var n62 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n61 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n62,
	},
}

var n60 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n61,
	},
}

var n59 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n60,
	},
}

var n55 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n56,
		CommaExpr: []*CommaExpr {
			n59,
		},
	},
}

var n54 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n55,
		RPAREN: ")",
	},
}

var n53 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n54,
	},
}

var n52 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n53,
	},
}

var n51 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n52,
		SEMI: ";",
	},
}

var n50 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n51,
	},
}

var n8 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n9,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n10,
			n22,
			n35,
			n50,
		},
		RBRACE: "}",
	},
}

var n64 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n67 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n69 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n68 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "msg",
		ArrayIndexDef: []*ArrayIndexDef {
			n69,
		},
	},
}

var n66 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n67,
		VariableDef: n68,
	},
}

var n72 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n73 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "len",
	},
}

var n71 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n72,
		VariableDef: n73,
	},
}

var n70 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n71,
	},
}

var n65 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n66,
		CommaParamDef: []*CommaParamDef {
			n70,
		},
	},
}

var n82 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n81 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n82,
	},
}

var n80 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n81,
	},
}

var n86 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n85 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n86,
	},
}

var n84 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n85,
	},
}

var n83 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n84,
	},
}

var n90 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "msg",
	},
}

var n89 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n90,
	},
}

var n88 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n89,
	},
}

var n87 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n88,
	},
}

var n94 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "len",
	},
}

var n93 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n94,
	},
}

var n92 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n93,
	},
}

var n91 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n92,
	},
}

var n79 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n80,
		CommaExpr: []*CommaExpr {
			n83,
			n87,
			n91,
		},
	},
}

var n78 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n79,
		RPAREN: ")",
	},
}

var n77 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n78,
	},
}

var n76 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n77,
	},
}

var n75 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n76,
		SEMI: ";",
	},
}

var n74 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n75,
	},
}

var n63 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n64,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n65,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n74,
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
			n63,
		},
	},
}
