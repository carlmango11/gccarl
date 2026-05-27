
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

var n36 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "0",
		RSQUARE: "]",
	},
}

var n35 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n36,
		},
	},
}

var n34 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n35,
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

var n31 = &Return {
	Type: ReturnTypeReturn,
	Return: &Return_ReturnOption{
		RETURN: "return",
		Expr: n32,
		SEMI: ";",
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
		},
		Return: n31,
		RBRACE: "}",
	},
}

var n38 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n43 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n45 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "3",
		RSQUARE: "]",
	},
}

var n44 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n45,
		},
	},
}

var n53 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
	},
}

var n52 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n53,
	},
}

var n51 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n52,
	},
}

var n57 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'E'",
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

var n54 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n55,
	},
}

var n61 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'S'",
	},
}

var n60 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n61,
	},
}

var n59 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n60,
	},
}

var n58 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n59,
	},
}

var n50 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n51,
		CommaExpr: []*CommaExpr {
			n54,
			n58,
		},
	},
}

var n49 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n50,
		RBRACE: "}",
	},
}

var n48 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n49,
	},
}

var n47 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n48,
	},
}

var n46 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n47,
	},
}

var n42 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n43,
		Variable: n44,
		EQUALS: "=",
		Expr: n46,
	},
}

var n41 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n42,
	},
}

var n40 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n41,
		SEMI: ";",
	},
}

var n39 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n40,
	},
}

var n66 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n67 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "read",
	},
}

var n69 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "read_char",
		LPAREN: "(",
		RPAREN: ")",
	},
}

var n68 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n69,
	},
}

var n65 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n66,
		Variable: n67,
		EQUALS: "=",
		Expr: n68,
	},
}

var n64 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n65,
	},
}

var n63 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n64,
		SEMI: ";",
	},
}

var n62 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n63,
	},
}

var n76 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "read",
	},
}

var n75 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n76,
	},
}

var n74 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n75,
	},
}

var n79 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'x'",
	},
}

var n78 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n79,
	},
}

var n77 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n78,
	},
}

var n73 = &CompExpr {
	Type: CompExprTypeIsEqual,
	IsEqual: &CompExpr_IsEqualOption{
		SubExpr: n74,
		EEQUALS: "==",
		Expr: n77,
	},
}

var n72 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n73,
	},
}

var n88 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n87 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n88,
	},
}

var n86 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n87,
	},
}

var n85 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n86,
	},
}

var n92 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "3",
	},
}

var n91 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n92,
	},
}

var n90 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n91,
	},
}

var n89 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n90,
	},
}

var n84 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n85,
		CommaExpr: []*CommaExpr {
			n89,
		},
	},
}

var n83 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n84,
		RPAREN: ")",
	},
}

var n82 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n83,
	},
}

var n81 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n82,
	},
}

var n80 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n81,
		SEMI: ";",
	},
}

var n71 = &Control {
	Type: ControlTypeIf,
	If: &Control_IfOption{
		IF: "if",
		LPAREN: "(",
		Expr: n72,
		RPAREN: ")",
		LBRACE: "{",
		StatementComma: []*StatementComma {
			n80,
		},
		RBRACE: "}",
	},
}

var n70 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n71,
	},
}

var n37 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n38,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n39,
			n62,
			n70,
		},
		RBRACE: "}",
	},
}

var n94 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n97 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n99 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n98 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n99,
		},
	},
}

var n96 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n97,
		Variable: n98,
	},
}

var n102 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n103 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n101 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n102,
		Variable: n103,
	},
}

var n100 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n101,
	},
}

var n95 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n96,
		CommaParamDef: []*CommaParamDef {
			n100,
		},
	},
}

var n112 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n111 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n112,
	},
}

var n110 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n111,
	},
}

var n116 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n115 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n116,
	},
}

var n114 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n115,
	},
}

var n113 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n114,
	},
}

var n121 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n120 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n121,
	},
}

var n119 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n120,
	},
}

var n118 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n119,
	},
}

var n117 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n118,
	},
}

var n126 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n125 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n126,
	},
}

var n124 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n125,
	},
}

var n123 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n124,
	},
}

var n122 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n123,
	},
}

var n109 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n110,
		CommaExpr: []*CommaExpr {
			n113,
			n117,
			n122,
		},
	},
}

var n108 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n109,
		RPAREN: ")",
	},
}

var n107 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n108,
	},
}

var n106 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n107,
	},
}

var n105 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n106,
		SEMI: ";",
	},
}

var n104 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n105,
	},
}

var n93 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n94,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n95,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n104,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n37,
			n93,
		},
	},
}
