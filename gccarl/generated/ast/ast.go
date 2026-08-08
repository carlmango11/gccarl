
package ast

var MainNode = n0

var n2 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n7 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n8 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "c",
	},
}

var n11 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
	},
}

var n10 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n11,
	},
}

var n9 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n10,
	},
}

var n6 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n7,
		VariableDef: n8,
		EQUALS: "=",
		Expr: n9,
	},
}

var n5 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n6,
	},
}

var n4 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n5,
		SEMI: ";",
	},
}

var n3 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n4,
	},
}

var n16 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n18 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "p",
	},
}

var n17 = &VariableDef {
	Type: VariableDefTypePointer,
	Pointer: &VariableDef_PointerOption{
		ASTERISKS: "*",
		VariableDef: n18,
	},
}

var n23 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "c",
	},
}

var n22 = &VariableAccess {
	Type: VariableAccessTypeAddressOf,
	AddressOf: &VariableAccess_AddressOfOption{
		AMPERSAND: "&",
		VariableAccess: n23,
	},
}

var n21 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n22,
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

var n15 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n16,
		VariableDef: n17,
		EQUALS: "=",
		Expr: n19,
	},
}

var n14 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n15,
	},
}

var n13 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n14,
		SEMI: ";",
	},
}

var n12 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n13,
	},
}

var n27 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "c",
	},
}

var n32 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "p",
	},
}

var n31 = &VariableAccess {
	Type: VariableAccessTypeDeref,
	Deref: &VariableAccess_DerefOption{
		ASTERISKS: "*",
		VariableAccess: n32,
	},
}

var n30 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n31,
	},
}

var n29 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n30,
	},
}

var n28 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n29,
	},
}

var n26 = &Statement {
	Type: StatementTypeAssign,
	Assign: &Statement_AssignOption{
		VariableAccess: n27,
		EQUALS: "=",
		Expr: n28,
	},
}

var n25 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n26,
		SEMI: ";",
	},
}

var n24 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n25,
	},
}

var n37 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n39 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "1",
		RSQUARE: "]",
	},
}

var n38 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "msg",
		ArrayIndexDef: []*ArrayIndexDef {
			n39,
		},
	},
}

var n48 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "c",
	},
}

var n47 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n48,
	},
}

var n46 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n47,
	},
}

var n45 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n46,
	},
}

var n44 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n45,
	},
}

var n43 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n44,
		RBRACE: "}",
	},
}

var n42 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n43,
	},
}

var n41 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n42,
	},
}

var n40 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n41,
	},
}

var n36 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n37,
		VariableDef: n38,
		EQUALS: "=",
		Expr: n40,
	},
}

var n35 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n36,
	},
}

var n34 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n35,
		SEMI: ";",
	},
}

var n33 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n34,
	},
}

var n58 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "msg",
	},
}

var n57 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n58,
	},
}

var n56 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n57,
	},
}

var n55 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n56,
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

var n54 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n55,
		CommaExpr: []*CommaExpr {
			n59,
		},
	},
}

var n53 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n54,
		RPAREN: ")",
	},
}

var n52 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n53,
	},
}

var n51 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n52,
	},
}

var n50 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n51,
		SEMI: ";",
	},
}

var n49 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n50,
	},
}

var n1 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n2,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n3,
			n12,
			n24,
			n33,
			n49,
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

var n91 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "msg",
	},
}

var n90 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n91,
	},
}

var n89 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n90,
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

var n96 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "len",
	},
}

var n95 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n96,
	},
}

var n94 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n95,
	},
}

var n93 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n94,
	},
}

var n92 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n93,
	},
}

var n79 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n80,
		CommaExpr: []*CommaExpr {
			n83,
			n87,
			n92,
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
			n63,
		},
	},
}
