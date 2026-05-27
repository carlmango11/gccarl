
package ast

var MainNode = n0

var n2 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n6 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n8 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "1",
		RSQUARE: "]",
	},
}

var n7 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n8,
		},
	},
}

var n5 = &Statement {
	Type: StatementTypeVarDec,
	VarDec: &Statement_VarDecOption{
		Type: n6,
		Variable: n7,
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

var n17 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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

var n26 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
	},
}

var n25 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n26,
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

var n22 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n23,
	},
}

var n30 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
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

var n27 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n28,
	},
}

var n14 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n15,
		CommaExpr: []*CommaExpr {
			n18,
			n22,
			n27,
		},
	},
}

var n13 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n14,
		RPAREN: ")",
	},
}

var n12 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n13,
	},
}

var n11 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n12,
	},
}

var n10 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n11,
		SEMI: ";",
	},
}

var n9 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n10,
	},
}

var n38 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "0",
		RSQUARE: "]",
	},
}

var n37 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n38,
		},
	},
}

var n36 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n37,
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

var n33 = &Statement {
	Type: StatementTypeReturn,
	Return: &Statement_ReturnOption{
		RETURN: "return",
		Expr: n34,
	},
}

var n32 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n33,
		SEMI: ";",
	},
}

var n31 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n32,
	},
}

var n1 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n2,
		IDEN: "read_char",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n3,
			n9,
			n31,
		},
		RBRACE: "}",
	},
}

var n40 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
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
		NUM: "3",
		RSQUARE: "]",
	},
}

var n46 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n47,
		},
	},
}

var n55 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
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
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'E'",
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

var n63 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'S'",
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

var n52 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n53,
		CommaExpr: []*CommaExpr {
			n56,
			n60,
		},
	},
}

var n51 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n52,
		RBRACE: "}",
	},
}

var n50 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n51,
	},
}

var n49 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n50,
	},
}

var n48 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n49,
	},
}

var n44 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n45,
		Variable: n46,
		EQUALS: "=",
		Expr: n48,
	},
}

var n43 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n44,
	},
}

var n42 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n43,
		SEMI: ";",
	},
}

var n41 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n42,
	},
}

var n68 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n69 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "read",
	},
}

var n71 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "read_char",
		LPAREN: "(",
		RPAREN: ")",
	},
}

var n70 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n71,
	},
}

var n67 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n68,
		Variable: n69,
		EQUALS: "=",
		Expr: n70,
	},
}

var n66 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n67,
	},
}

var n65 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n66,
		SEMI: ";",
	},
}

var n64 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n65,
	},
}

var n78 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "read",
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

var n81 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'x'",
	},
}

var n80 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n81,
	},
}

var n79 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n80,
	},
}

var n75 = &CompExpr {
	Type: CompExprTypeIsEqual,
	IsEqual: &CompExpr_IsEqualOption{
		SubExpr: n76,
		EEQUALS: "==",
		Expr: n79,
	},
}

var n74 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n75,
	},
}

var n91 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n90 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n91,
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

var n95 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "3",
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

var n87 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n88,
		CommaExpr: []*CommaExpr {
			n92,
		},
	},
}

var n86 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n87,
		RPAREN: ")",
	},
}

var n85 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n86,
	},
}

var n84 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n85,
	},
}

var n83 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n84,
		SEMI: ";",
	},
}

var n82 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n83,
	},
}

var n73 = &Control {
	Type: ControlTypeIf,
	If: &Control_IfOption{
		IF: "if",
		LPAREN: "(",
		Expr: n74,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n82,
		},
		RBRACE: "}",
	},
}

var n72 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n73,
	},
}

var n39 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n40,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n41,
			n64,
			n72,
		},
		RBRACE: "}",
	},
}

var n97 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n100 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n102 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n101 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n102,
		},
	},
}

var n99 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n100,
		Variable: n101,
	},
}

var n105 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n106 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n104 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n105,
		Variable: n106,
	},
}

var n103 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n104,
	},
}

var n98 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n99,
		CommaParamDef: []*CommaParamDef {
			n103,
		},
	},
}

var n115 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n114 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n115,
	},
}

var n113 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n114,
	},
}

var n119 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n118 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n119,
	},
}

var n117 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n118,
	},
}

var n116 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n117,
	},
}

var n124 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n123 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n124,
	},
}

var n122 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n123,
	},
}

var n121 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n122,
	},
}

var n120 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n121,
	},
}

var n129 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n128 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n129,
	},
}

var n127 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n128,
	},
}

var n126 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n127,
	},
}

var n125 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n126,
	},
}

var n112 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n113,
		CommaExpr: []*CommaExpr {
			n116,
			n120,
			n125,
		},
	},
}

var n111 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n112,
		RPAREN: ")",
	},
}

var n110 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n111,
	},
}

var n109 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n110,
	},
}

var n108 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n109,
		SEMI: ";",
	},
}

var n107 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n108,
	},
}

var n96 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n97,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n98,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n107,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n39,
			n96,
		},
	},
}
