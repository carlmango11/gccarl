
package ast

var MainNode = n0

var n2 = &Type {
	Type: TypeTypeVoid,
	Void: &Type_VoidOption{
		VOID: "void",
	},
}

var n5 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n7 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n6 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "buf",
		ArrayIndex: []*ArrayIndex {
			n7,
		},
	},
}

var n4 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n5,
		Variable: n6,
	},
}

var n3 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n4,
	},
}

var n16 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
	},
}

var n15 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n16,
	},
}

var n14 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n15,
	},
}

var n20 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
	},
}

var n19 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n20,
	},
}

var n18 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n19,
	},
}

var n17 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n18,
	},
}

var n25 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "buf",
	},
}

var n24 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n25,
	},
}

var n23 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n24,
	},
}

var n22 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n23,
	},
}

var n21 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n22,
	},
}

var n29 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "20",
	},
}

var n28 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n29,
	},
}

var n27 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n28,
	},
}

var n26 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n27,
	},
}

var n13 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n14,
		CommaExpr: []*CommaExpr {
			n17,
			n21,
			n26,
		},
	},
}

var n12 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n13,
		RPAREN: ")",
	},
}

var n11 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n12,
	},
}

var n10 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n11,
	},
}

var n9 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n10,
		SEMI: ";",
	},
}

var n8 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n9,
	},
}

var n1 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n2,
		IDEN: "read_str",
		LPAREN: "(",
		ParamsDef: n3,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n8,
		},
		RBRACE: "}",
	},
}

var n31 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n36 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n38 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n37 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n38,
		},
	},
}

var n41 = &Value {
	Type: ValueTypeStr,
	Str: &Value_StrOption{
		STR: "\"hello\"",
	},
}

var n40 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n41,
	},
}

var n39 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n40,
	},
}

var n35 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n36,
		Variable: n37,
		EQUALS: "=",
		Expr: n39,
	},
}

var n34 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n35,
	},
}

var n33 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n34,
		SEMI: ";",
	},
}

var n32 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n33,
	},
}

var n45 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n47 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "20",
		RSQUARE: "]",
	},
}

var n46 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n47,
		},
	},
}

var n44 = &Statement {
	Type: StatementTypeVarDec,
	VarDec: &Statement_VarDecOption{
		Type: n45,
		Variable: n46,
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

var n57 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
	},
}

var n56 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n57,
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

var n53 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n54,
	},
}

var n52 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "read_str",
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

var n67 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n66 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n67,
	},
}

var n65 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n66,
	},
}

var n64 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n65,
	},
}

var n71 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "5",
	},
}

var n70 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n71,
	},
}

var n69 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n70,
	},
}

var n68 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n69,
	},
}

var n63 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n64,
		CommaExpr: []*CommaExpr {
			n68,
		},
	},
}

var n62 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n63,
		RPAREN: ")",
	},
}

var n61 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n62,
	},
}

var n60 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n61,
	},
}

var n59 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n60,
		SEMI: ";",
	},
}

var n58 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n59,
	},
}

var n81 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
	},
}

var n80 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n81,
	},
}

var n79 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n80,
	},
}

var n78 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n79,
	},
}

var n85 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "5",
	},
}

var n84 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n85,
	},
}

var n83 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n84,
	},
}

var n82 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n83,
	},
}

var n77 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n78,
		CommaExpr: []*CommaExpr {
			n82,
		},
	},
}

var n76 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n77,
		RPAREN: ")",
	},
}

var n75 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n76,
	},
}

var n74 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n75,
	},
}

var n73 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n74,
		SEMI: ";",
	},
}

var n72 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n73,
	},
}

var n30 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n31,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n32,
			n42,
			n48,
			n58,
			n72,
		},
		RBRACE: "}",
	},
}

var n87 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n90 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n92 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n91 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n92,
		},
	},
}

var n89 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n90,
		Variable: n91,
	},
}

var n95 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n96 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n94 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n95,
		Variable: n96,
	},
}

var n93 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n94,
	},
}

var n88 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n89,
		CommaParamDef: []*CommaParamDef {
			n93,
		},
	},
}

var n105 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n104 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n105,
	},
}

var n103 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n104,
	},
}

var n109 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n108 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n109,
	},
}

var n107 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n108,
	},
}

var n106 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n107,
	},
}

var n114 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n113 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n114,
	},
}

var n112 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n113,
	},
}

var n111 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n112,
	},
}

var n110 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n111,
	},
}

var n119 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n118 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n119,
	},
}

var n117 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n118,
	},
}

var n116 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n117,
	},
}

var n115 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n116,
	},
}

var n102 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n103,
		CommaExpr: []*CommaExpr {
			n106,
			n110,
			n115,
		},
	},
}

var n101 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n102,
		RPAREN: ")",
	},
}

var n100 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n101,
	},
}

var n99 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n100,
	},
}

var n98 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n99,
		SEMI: ";",
	},
}

var n97 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n98,
	},
}

var n86 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n87,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n88,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n97,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n30,
			n86,
		},
	},
}
