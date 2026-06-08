
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
		STR: "\"hello \"",
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

var n62 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n63 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "i",
	},
}

var n66 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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

var n61 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n62,
		Variable: n63,
		EQUALS: "=",
		Expr: n64,
	},
}

var n60 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n61,
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

var n73 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "i",
	},
}

var n72 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n73,
	},
}

var n71 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n72,
	},
}

var n74 = &Operator {
	Type: OperatorTypeLess,
	Less: &Operator_LessOption{
		LESS_THAN: "<",
	},
}

var n77 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "3",
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

var n70 = &CompExpr {
	Type: CompExprTypeCompExpr,
	CompExpr: &CompExpr_CompExprOption{
		SubExpr: n71,
		Operator: n74,
		Expr: n75,
	},
}

var n69 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n70,
	},
}

var n87 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n86 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n87,
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

var n91 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "6",
	},
}

var n90 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n91,
	},
}

var n89 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n90,
	},
}

var n88 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n89,
	},
}

var n83 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n84,
		CommaExpr: []*CommaExpr {
			n88,
		},
	},
}

var n82 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n83,
		RPAREN: ")",
	},
}

var n81 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n82,
	},
}

var n80 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n81,
	},
}

var n79 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n80,
		SEMI: ";",
	},
}

var n78 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n79,
	},
}

var n101 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
	},
}

var n100 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n101,
	},
}

var n99 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n100,
	},
}

var n98 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n99,
	},
}

var n105 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "20",
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

var n102 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n103,
	},
}

var n97 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n98,
		CommaExpr: []*CommaExpr {
			n102,
		},
	},
}

var n96 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n97,
		RPAREN: ")",
	},
}

var n95 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n96,
	},
}

var n94 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n95,
	},
}

var n93 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n94,
		SEMI: ";",
	},
}

var n92 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n93,
	},
}

var n109 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "i",
	},
}

var n114 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "i",
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

var n115 = &Operator {
	Type: OperatorTypePlus,
	Plus: &Operator_PlusOption{
		PLUS: "+",
	},
}

var n118 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
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

var n111 = &CompExpr {
	Type: CompExprTypeCompExpr,
	CompExpr: &CompExpr_CompExprOption{
		SubExpr: n112,
		Operator: n115,
		Expr: n116,
	},
}

var n110 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n111,
	},
}

var n108 = &Statement {
	Type: StatementTypeAssign,
	Assign: &Statement_AssignOption{
		Variable: n109,
		EQUALS: "=",
		Expr: n110,
	},
}

var n107 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n108,
		SEMI: ";",
	},
}

var n106 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n107,
	},
}

var n68 = &Control {
	Type: ControlTypeWhile,
	While: &Control_WhileOption{
		WHILE: "while",
		LPAREN: "(",
		Expr: n69,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n78,
			n92,
			n106,
		},
		RBRACE: "}",
	},
}

var n67 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n68,
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
			n67,
		},
		RBRACE: "}",
	},
}

var n120 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n123 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n125 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n124 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n125,
		},
	},
}

var n122 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n123,
		Variable: n124,
	},
}

var n128 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n129 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n127 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n128,
		Variable: n129,
	},
}

var n126 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n127,
	},
}

var n121 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n122,
		CommaParamDef: []*CommaParamDef {
			n126,
		},
	},
}

var n138 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n137 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n138,
	},
}

var n136 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n137,
	},
}

var n142 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n141 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n142,
	},
}

var n140 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n141,
	},
}

var n139 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n140,
	},
}

var n147 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n146 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n147,
	},
}

var n145 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n146,
	},
}

var n144 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n145,
	},
}

var n143 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n144,
	},
}

var n152 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n151 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n152,
	},
}

var n150 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n151,
	},
}

var n149 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n150,
	},
}

var n148 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n149,
	},
}

var n135 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n136,
		CommaExpr: []*CommaExpr {
			n139,
			n143,
			n148,
		},
	},
}

var n134 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n135,
		RPAREN: ")",
	},
}

var n133 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n134,
	},
}

var n132 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n133,
	},
}

var n131 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n132,
		SEMI: ";",
	},
}

var n130 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n131,
	},
}

var n119 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n120,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n121,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n130,
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
			n119,
		},
	},
}
