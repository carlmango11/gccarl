
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
		},
		RBRACE: "}",
	},
}

var n34 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
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

var n36 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n37,
		VariableDef: n38,
	},
}

var n42 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n43 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "len",
	},
}

var n41 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n42,
		VariableDef: n43,
	},
}

var n40 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n41,
	},
}

var n35 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n36,
		CommaParamDef: []*CommaParamDef {
			n40,
		},
	},
}

var n52 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n51 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n52,
	},
}

var n50 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n51,
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

var n53 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n54,
	},
}

var n61 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "msg",
	},
}

var n60 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n61,
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

var n66 = &VariableAccess {
	Type: VariableAccessTypeVariable,
	Variable: &VariableAccess_VariableOption{
		IDEN: "len",
	},
}

var n65 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		VariableAccess: n66,
	},
}

var n64 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n65,
	},
}

var n63 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n64,
	},
}

var n62 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n63,
	},
}

var n49 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n50,
		CommaExpr: []*CommaExpr {
			n53,
			n57,
			n62,
		},
	},
}

var n48 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n49,
		RPAREN: ")",
	},
}

var n47 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n48,
	},
}

var n46 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n47,
	},
}

var n45 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n46,
		SEMI: ";",
	},
}

var n44 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n45,
	},
}

var n33 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n34,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n35,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n44,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n33,
		},
	},
}
