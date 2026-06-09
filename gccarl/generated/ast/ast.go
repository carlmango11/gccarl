
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

var n46 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "ans",
	},
}

var n48 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "read_char",
		LPAREN: "(",
		RPAREN: ")",
	},
}

var n47 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n48,
	},
}

var n44 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n45,
		Variable: n46,
		EQUALS: "=",
		Expr: n47,
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

var n55 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "ans",
	},
}

var n54 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n55,
	},
}

var n53 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n54,
	},
}

var n56 = &Operator {
	Type: OperatorTypeEqual,
	Equal: &Operator_EqualOption{
		EEQUALS: "==",
	},
}

var n59 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'1'",
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

var n52 = &CompExpr {
	Type: CompExprTypeCompExpr,
	CompExpr: &CompExpr_CompExprOption{
		SubExpr: n53,
		Operator: n56,
		Expr: n57,
	},
}

var n51 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n52,
	},
}

var n69 = &Value {
	Type: ValueTypeStr,
	Str: &Value_StrOption{
		STR: "\"was 1\"",
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
		NUM: "5",
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

var n66 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n67,
		CommaExpr: []*CommaExpr {
			n70,
		},
	},
}

var n65 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
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

var n60 = &BlockOrLine {
	Type: BlockOrLineTypeLine,
	Line: &BlockOrLine_LineOption{
		Line: n61,
	},
}

var n84 = &Value {
	Type: ValueTypeStr,
	Str: &Value_StrOption{
		STR: "\"was 2\"",
	},
}

var n83 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n84,
	},
}

var n82 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n83,
	},
}

var n88 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "5",
	},
}

var n87 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n88,
	},
}

var n86 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n87,
	},
}

var n85 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n86,
	},
}

var n81 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n82,
		CommaExpr: []*CommaExpr {
			n85,
		},
	},
}

var n80 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n81,
		RPAREN: ")",
	},
}

var n79 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n80,
	},
}

var n78 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n79,
	},
}

var n77 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n78,
		SEMI: ";",
	},
}

var n76 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n77,
	},
}

var n75 = &BlockOrLine {
	Type: BlockOrLineTypeLine,
	Line: &BlockOrLine_LineOption{
		Line: n76,
	},
}

var n74 = &Else {
	Type: ElseTypeElse,
	Else: &Else_ElseOption{
		ELSE: "else",
		BlockOrLine: n75,
	},
}

var n50 = &Control {
	Type: ControlTypeIf,
	If: &Control_IfOption{
		IF: "if",
		LPAREN: "(",
		Expr: n51,
		RPAREN: ")",
		BlockOrLine: n60,
		Else: n74,
	},
}

var n49 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n50,
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
			n49,
		},
		RBRACE: "}",
	},
}

var n90 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n93 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n95 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n94 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n95,
		},
	},
}

var n92 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n93,
		Variable: n94,
	},
}

var n98 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n99 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n97 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n98,
		Variable: n99,
	},
}

var n96 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n97,
	},
}

var n91 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n92,
		CommaParamDef: []*CommaParamDef {
			n96,
		},
	},
}

var n108 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n107 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n108,
	},
}

var n106 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n107,
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

var n109 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n110,
	},
}

var n117 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n116 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n117,
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

var n122 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n121 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n122,
	},
}

var n120 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n121,
	},
}

var n119 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n120,
	},
}

var n118 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n119,
	},
}

var n105 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n106,
		CommaExpr: []*CommaExpr {
			n109,
			n113,
			n118,
		},
	},
}

var n104 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n105,
		RPAREN: ")",
	},
}

var n103 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n104,
	},
}

var n102 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n103,
	},
}

var n101 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n102,
		SEMI: ";",
	},
}

var n100 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n101,
	},
}

var n89 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n90,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n91,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n100,
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
			n89,
		},
	},
}
