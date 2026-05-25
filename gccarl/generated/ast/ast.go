
package ast

var MainNode = n0

var n2 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
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
		NUM: "1",
		RSQUARE: "]",
	},
}

var n8 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n9,
		},
	},
}

var n17 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "' '",
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

var n14 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n15,
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

var n26 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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

var n30 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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

var n35 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
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

var n31 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n32,
	},
}

var n39 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n38 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n39,
	},
}

var n37 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n38,
	},
}

var n36 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n37,
	},
}

var n23 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n24,
		CommaExpr: []*CommaExpr {
			n27,
			n31,
			n36,
		},
	},
}

var n22 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n23,
		RPAREN: ")",
	},
}

var n21 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n22,
	},
}

var n20 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n21,
	},
}

var n19 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n20,
		SEMI: ";",
	},
}

var n18 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n19,
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
			n18,
		},
		RBRACE: "}",
	},
}

var n41 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n46 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n48 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "1",
		RSQUARE: "]",
	},
}

var n47 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n48,
		},
	},
}

var n56 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "' '",
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

var n53 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n54,
	},
}

var n52 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n53,
		RBRACE: "}",
	},
}

var n51 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n52,
	},
}

var n50 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n51,
	},
}

var n49 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n50,
	},
}

var n45 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n46,
		Variable: n47,
		EQUALS: "=",
		Expr: n49,
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

var n65 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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

var n69 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "0",
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

var n74 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
	},
}

var n73 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n74,
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

var n78 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n77 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n78,
	},
}

var n76 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n77,
	},
}

var n75 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n76,
	},
}

var n62 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n63,
		CommaExpr: []*CommaExpr {
			n66,
			n70,
			n75,
		},
	},
}

var n61 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
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

var n83 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n85 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "3",
		RSQUARE: "]",
	},
}

var n84 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n85,
		},
	},
}

var n93 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
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

var n97 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'E'",
	},
}

var n96 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n97,
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

var n101 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'S'",
	},
}

var n100 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n101,
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

var n90 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		Expr: n91,
		CommaExpr: []*CommaExpr {
			n94,
			n98,
		},
	},
}

var n89 = &CompositeLiteral {
	Type: CompositeLiteralTypeArrayVal,
	ArrayVal: &CompositeLiteral_ArrayValOption{
		LBRACE: "{",
		CompEntries: n90,
		RBRACE: "}",
	},
}

var n88 = &Value {
	Type: ValueTypeCompLit,
	CompLit: &Value_CompLitOption{
		CompositeLiteral: n89,
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

var n82 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n83,
		Variable: n84,
		EQUALS: "=",
		Expr: n86,
	},
}

var n81 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n82,
	},
}

var n80 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n81,
		SEMI: ";",
	},
}

var n79 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n80,
	},
}

var n109 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		NUM: "0",
		RSQUARE: "]",
	},
}

var n108 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "c",
		ArrayIndex: []*ArrayIndex {
			n109,
		},
	},
}

var n107 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n108,
	},
}

var n106 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n107,
	},
}

var n112 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'x'",
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

var n105 = &CompExpr {
	Type: CompExprTypeIsEqual,
	IsEqual: &CompExpr_IsEqualOption{
		SubExpr: n106,
		EEQUALS: "==",
		Expr: n110,
	},
}

var n104 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n105,
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

var n125 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "3",
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

var n117 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n118,
		CommaExpr: []*CommaExpr {
			n122,
		},
	},
}

var n116 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "print",
		LPAREN: "(",
		Params: n117,
		RPAREN: ")",
	},
}

var n115 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n116,
	},
}

var n114 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n115,
	},
}

var n113 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n114,
		SEMI: ";",
	},
}

var n103 = &Control {
	Type: ControlTypeIf,
	If: &Control_IfOption{
		IF: "if",
		LPAREN: "(",
		Expr: n104,
		RPAREN: ")",
		LBRACE: "{",
		StatementComma: []*StatementComma {
			n113,
		},
		RBRACE: "}",
	},
}

var n102 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n103,
	},
}

var n40 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n41,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n42,
			n57,
			n79,
			n102,
		},
		RBRACE: "}",
	},
}

var n127 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n130 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n132 = &ArrayIndex {
	Type: ArrayIndexTypeArrayIndex,
	ArrayIndex: &ArrayIndex_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n131 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
		ArrayIndex: []*ArrayIndex {
			n132,
		},
	},
}

var n129 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n130,
		Variable: n131,
	},
}

var n135 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n136 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n134 = &ParamDef {
	Type: ParamDefTypeParam,
	Param: &ParamDef_ParamOption{
		Type: n135,
		Variable: n136,
	},
}

var n133 = &CommaParamDef {
	Type: CommaParamDefTypeParam,
	Param: &CommaParamDef_ParamOption{
		COMMA: ",",
		ParamDef: n134,
	},
}

var n128 = &ParamsDef {
	Type: ParamsDefTypeParams,
	Params: &ParamsDef_ParamsOption{
		ParamDef: n129,
		CommaParamDef: []*CommaParamDef {
			n133,
		},
	},
}

var n145 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n144 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n145,
	},
}

var n143 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n144,
	},
}

var n149 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "1",
	},
}

var n148 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n149,
	},
}

var n147 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n148,
	},
}

var n146 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n147,
	},
}

var n154 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "msg",
	},
}

var n153 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n154,
	},
}

var n152 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n153,
	},
}

var n151 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n152,
	},
}

var n150 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n151,
	},
}

var n159 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "len",
	},
}

var n158 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n159,
	},
}

var n157 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n158,
	},
}

var n156 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n157,
	},
}

var n155 = &CommaExpr {
	Type: CommaExprTypeCommaExpr,
	CommaExpr: &CommaExpr_CommaExprOption{
		COMMA: ",",
		Expr: n156,
	},
}

var n142 = &Params {
	Type: ParamsTypeParams,
	Params: &Params_ParamsOption{
		Expr: n143,
		CommaExpr: []*CommaExpr {
			n146,
			n150,
			n155,
		},
	},
}

var n141 = &SubExpr {
	Type: SubExprTypeFuncCall,
	FuncCall: &SubExpr_FuncCallOption{
		IDEN: "do_syscall",
		LPAREN: "(",
		Params: n142,
		RPAREN: ")",
	},
}

var n140 = &Expr {
	Type: ExprTypeSubExpr,
	SubExpr: &Expr_SubExprOption{
		SubExpr: n141,
	},
}

var n139 = &Statement {
	Type: StatementTypeExpr,
	Expr: &Statement_ExprOption{
		Expr: n140,
	},
}

var n138 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n139,
		SEMI: ";",
	},
}

var n137 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n138,
	},
}

var n126 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n127,
		IDEN: "print",
		LPAREN: "(",
		ParamsDef: n128,
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n137,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n40,
			n126,
		},
	},
}
