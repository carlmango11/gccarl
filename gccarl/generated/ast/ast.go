
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

var n26 = &Type {
	Type: TypeTypeStruct,
	Struct: &Type_StructOption{
		STRUCT: "struct",
		IDEN: "Person",
	},
}

var n28 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "y",
	},
}

var n27 = &VariableDef {
	Type: VariableDefTypePointer,
	Pointer: &VariableDef_PointerOption{
		ASTERISKS: "*",
		VariableDef: n28,
	},
}

var n25 = &VarDec {
	Type: VarDecTypeVarDec,
	VarDec: &VarDec_VarDecOption{
		Type: n26,
		VariableDef: n27,
	},
}

var n24 = &Statement {
	Type: StatementTypeVarDec,
	VarDec: &Statement_VarDecOption{
		VarDec: n25,
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

var n35 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "y",
	},
}

var n34 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n35,
	},
}

var n36 = &Operator {
	Type: OperatorTypeAssign,
	Assign: &Operator_AssignOption{
		EQUALS: "=",
	},
}

var n41 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "p",
	},
}

var n40 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n41,
	},
}

var n39 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n40,
	},
}

var n38 = &SubExpr {
	Type: SubExprTypeAddressOf,
	AddressOf: &SubExpr_AddressOfOption{
		AMPERSAND: "&",
		Expr: n39,
	},
}

var n37 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n38,
	},
}

var n33 = &CompExpr {
	Type: CompExprTypeCompExpr,
	CompExpr: &CompExpr_CompExprOption{
		SubExpr: n34,
		Operator: n36,
		Expr: n37,
	},
}

var n32 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n33,
	},
}

var n31 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n32,
	},
}

var n30 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n31,
		SEMI: ";",
	},
}

var n29 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n30,
	},
}

var n46 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n48 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n47 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "m",
		ArrayIndexDef: []*ArrayIndexDef {
			n48,
		},
	},
}

var n54 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "y",
	},
}

var n56 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "age",
	},
}

var n55 = &InnerSubVariableAccess {
	Type: InnerSubVariableAccessTypeArrow,
	Arrow: &InnerSubVariableAccess_ArrowOption{
		RIGHT_ARROW: "->",
		SubVariableAccess: n56,
	},
}

var n53 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n54,
		InnerSubVariableAccess: []*InnerSubVariableAccess {
			n55,
		},
	},
}

var n52 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n53,
	},
}

var n51 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Expr: n52,
	},
}

var n50 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n51,
	},
}

var n49 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n50,
		RBRACE: "}",
	},
}

var n45 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n46,
		VariableDef: n47,
		EQUALS: "=",
		Initialiser: n49,
	},
}

var n44 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n45,
	},
}

var n43 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n44,
		SEMI: ";",
	},
}

var n42 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n43,
	},
}

var n65 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "m",
	},
}

var n64 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n65,
	},
}

var n63 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n64,
	},
}

var n69 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n68 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n69,
	},
}

var n67 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n68,
	},
}

var n66 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n67,
	},
}

var n62 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n63,
		CommaExpr: []*CommaExpr {
			n66,
		},
	},
}

var n61 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n62,
		RPAREN: ")",
	},
}

var n60 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n61,
	},
}

var n59 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n60,
	},
}

var n58 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n59,
		SEMI: ";",
	},
}

var n57 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n58,
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
			n29,
			n42,
			n57,
		},
		RBRACE: "}",
	},
}

var n71 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n74 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n76 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n75 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "msg",
		ArrayIndexDef: []*ArrayIndexDef {
			n76,
		},
	},
}

var n73 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n74,
		VariableDef: n75,
	},
}

var n79 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n80 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "len",
	},
}

var n78 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n79,
		VariableDef: n80,
	},
}

var n77 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n78,
	},
}

var n72 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n73,
		CommaParamDef: []*CommaParamDef {
			n77,
		},
	},
}

var n89 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n88 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n89,
	},
}

var n87 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n88,
	},
}

var n93 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n92 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n93,
	},
}

var n91 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n92,
	},
}

var n90 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n91,
	},
}

var n97 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "msg",
	},
}

var n96 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n97,
	},
}

var n95 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n96,
	},
}

var n94 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n95,
	},
}

var n101 = &SubVariableAccess {
	Type: SubVariableAccessTypeV,
	V: &SubVariableAccess_VOption{
		IDEN: "len",
	},
}

var n100 = &SubExpr {
	Type: SubExprTypeVariable,
	Variable: &SubExpr_VariableOption{
		SubVariableAccess: n101,
	},
}

var n99 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n100,
	},
}

var n98 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n99,
	},
}

var n86 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n87,
		CommaExpr: []*CommaExpr {
			n90,
			n94,
			n98,
		},
	},
}

var n85 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n86,
		RPAREN: ")",
	},
}

var n84 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n85,
	},
}

var n83 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n84,
	},
}

var n82 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n83,
		SEMI: ";",
	},
}

var n81 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n82,
	},
}

var n70 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n71,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n72,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n81,
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
			n70,
		},
	},
}
