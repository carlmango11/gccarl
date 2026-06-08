
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

var n9 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "2",
		RSQUARE: "]",
	},
}

var n8 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n9,
		},
	},
}

var n17 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
	},
}

var n16 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n17,
	},
}

var n15 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n16,
	},
}

var n21 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'E'",
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

var n18 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n19,
	},
}

var n14 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n15,
		CommaExpr: []*CommaExpr {
			n18,
		},
	},
}

var n13 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n14,
		RBRACE: "}",
	},
}

var n12 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n13,
	},
}

var n11 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n12,
	},
}

var n10 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n11,
	},
}

var n6 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n7,
		Variable: n8,
		EQUALS: "=",
		Expr: n10,
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

var n31 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n30 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n31,
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

var n35 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "2",
	},
}

var n34 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n35,
	},
}

var n33 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n34,
	},
}

var n32 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n33,
	},
}

var n27 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n28,
		CommaExpr: []*CommaExpr {
			n32,
		},
	},
}

var n26 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n27,
		RPAREN: ")",
	},
}

var n25 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n26,
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
			n22,
		},
		RBRACE: "}",
	},
}

var n37 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n40 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n42 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n41 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n42,
		},
	},
}

var n39 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n40,
		Variable: n41,
	},
}

var n45 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n46 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n44 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n45,
		Variable: n46,
	},
}

var n43 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n44,
	},
}

var n38 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n39,
		CommaParamDef: []*CommaParamDef {
			n43,
		},
	},
}

var n55 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n54 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n55,
	},
}

var n53 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n54,
	},
}

var n59 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n58 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n59,
	},
}

var n57 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n58,
	},
}

var n56 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n57,
	},
}

var n64 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n63 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n64,
	},
}

var n62 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n63,
	},
}

var n61 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n62,
	},
}

var n60 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n61,
	},
}

var n69 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n68 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n69,
	},
}

var n67 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n68,
	},
}

var n66 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n67,
	},
}

var n65 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n66,
	},
}

var n52 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n53,
		CommaExpr: []*CommaExpr {
			n56,
			n60,
			n65,
		},
	},
}

var n51 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n52,
		RPAREN: ")",
	},
}

var n50 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n51,
	},
}

var n49 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n50,
	},
}

var n48 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n49,
		SEMI: ";",
	},
}

var n47 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n48,
	},
}

var n36 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n37,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n38,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n47,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n36,
		},
	},
}
