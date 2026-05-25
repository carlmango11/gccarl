
package ast

var MainNode = n0

var n2 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n7 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n8 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "x",
	},
}

var n11 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "3",
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
		Variable: n8,
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

var n18 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "1",
		RSQUARE: "]",
	},
}

var n17 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n18,
		},
	},
}

var n26 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
	},
}

var n25 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n26,
	},
}

var n24 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n25,
	},
}

var n23 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n24,
	},
}

var n22 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n23,
		RBRACE: "}",
	},
}

var n21 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n22,
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
		Variable: n17,
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

var n33 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "x",
	},
}

var n32 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n33,
	},
}

var n31 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n32,
	},
}

var n36 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "3",
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

var n30 = &CompExpr {
	Type: CompExprTypeIsEqual,
	IsEqual: &CompExpr_IsEqualOption{
		SubExpr: n31,
		EEQUALS: "==",
		Expr: n34,
	},
}

var n29 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n30,
	},
}

var n45 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n44 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n45,
	},
}

var n43 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n44,
	},
}

var n42 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n43,
	},
}

var n49 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n48 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n49,
	},
}

var n47 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n48,
	},
}

var n46 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n47,
	},
}

var n41 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n42,
		CommaExpr: []*CommaExpr {
			n46,
		},
	},
}

var n40 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n41,
		RPAREN: ")",
	},
}

var n39 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n40,
	},
}

var n38 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n39,
	},
}

var n37 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n38,
		SEMI: ";",
	},
}

var n28 = &Control {
	Type: ControlTypeIf,
	If: &Control_IfOption{
		IF: "if",
		LPAREN: "(",
		Expr: n29,
		RPAREN: ")",
		LBRACE: "{",
		StatementComma: []*StatementComma {
			n37,
		},
		RBRACE: "}",
	},
}

var n27 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n28,
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
			n27,
		},
		RBRACE: "}",
	},
}

var n51 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n54 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n56 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n55 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n56,
		},
	},
}

var n53 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n54,
		Variable: n55,
	},
}

var n59 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n60 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n58 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n59,
		Variable: n60,
	},
}

var n57 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n58,
	},
}

var n52 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n53,
		CommaParamDef: []*CommaParamDef {
			n57,
		},
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

var n73 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n72 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n73,
	},
}

var n71 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n72,
	},
}

var n70 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n71,
	},
}

var n78 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n77 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n78,
	},
}

var n76 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n77,
	},
}

var n75 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n76,
	},
}

var n74 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n75,
	},
}

var n83 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n82 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n83,
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

var n79 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n80,
	},
}

var n66 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n67,
		CommaExpr: []*CommaExpr {
			n70,
			n74,
			n79,
		},
	},
}

var n65 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n66,
		RPAREN: ")",
	},
}

var n64 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n65,
	},
}

var n63 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n64,
	},
}

var n62 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n63,
		SEMI: ";",
	},
}

var n61 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n62,
	},
}

var n50 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n51,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n52,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n61,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n50,
		},
	},
}
