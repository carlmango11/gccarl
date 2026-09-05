
package ast

var MainNode = n0

var n5 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n6 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "name",
	},
}

var n4 = &VarDec {
	Type: VarDecTypeVarDec,
	VarDec: &VarDec_VarDecOption{
		Type: n5,
		VariableDef: n6,
	},
}

var n3 = &VarDecColon {
	Type: VarDecColonTypeC,
	C: &VarDecColon_COption{
		VarDec: n4,
		SEMI: ";",
	},
}

var n9 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n10 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "age",
	},
}

var n8 = &VarDec {
	Type: VarDecTypeVarDec,
	VarDec: &VarDec_VarDecOption{
		Type: n9,
		VariableDef: n10,
	},
}

var n7 = &VarDecColon {
	Type: VarDecColonTypeC,
	C: &VarDecColon_COption{
		VarDec: n8,
		SEMI: ";",
	},
}

var n2 = &TypeDef {
	Type: TypeDefTypeStructDef,
	StructDef: &TypeDef_StructDefOption{
		STRUCT: "struct",
		IDEN: "Person",
		LBRACE: "{",
		VarDecColon: []*VarDecColon {
			n3,
			n7,
		},
		RBRACE: "}",
	},
}

var n1 = &DecDef {
	Type: DecDefTypeTypeDef,
	TypeDef: &DecDef_TypeDefOption{
		TypeDef: n2,
	},
}

var n12 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n17 = &Type {
	Type: TypeTypeStruct,
	Struct: &Type_StructOption{
		STRUCT: "struct",
		IDEN: "Person",
	},
}

var n19 = &ArrayIndexDef {
	Type: ArrayIndexDefTypeArrayIndex,
	ArrayIndex: &ArrayIndexDef_ArrayIndexOption{
		LSQUARE: "[",
		RSQUARE: "]",
	},
}

var n18 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "p",
		ArrayIndexDef: []*ArrayIndexDef {
			n19,
		},
	},
}

var n29 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'X'",
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

var n26 = &Initialiser {
	Type: InitialiserTypeExpr,
	Expr: &Initialiser_ExprOption{
		Expr: n27,
	},
}

var n25 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Initialiser: n26,
	},
}

var n35 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "17",
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

var n32 = &Initialiser {
	Type: InitialiserTypeExpr,
	Expr: &Initialiser_ExprOption{
		Expr: n33,
	},
}

var n31 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Initialiser: n32,
	},
}

var n30 = &CommaCompEntry {
	Type: CommaCompEntryTypeE,
	E: &CommaCompEntry_EOption{
		COMMA: ",",
		CompEntry: n31,
	},
}

var n24 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n25,
		CommaCompEntry: []*CommaCompEntry {
			n30,
		},
	},
}

var n23 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n24,
		RBRACE: "}",
	},
}

var n22 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Initialiser: n23,
	},
}

var n44 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'Y'",
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

var n41 = &Initialiser {
	Type: InitialiserTypeExpr,
	Expr: &Initialiser_ExprOption{
		Expr: n42,
	},
}

var n40 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Initialiser: n41,
	},
}

var n50 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "12",
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

var n47 = &Initialiser {
	Type: InitialiserTypeExpr,
	Expr: &Initialiser_ExprOption{
		Expr: n48,
	},
}

var n46 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Initialiser: n47,
	},
}

var n45 = &CommaCompEntry {
	Type: CommaCompEntryTypeE,
	E: &CommaCompEntry_EOption{
		COMMA: ",",
		CompEntry: n46,
	},
}

var n39 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n40,
		CommaCompEntry: []*CommaCompEntry {
			n45,
		},
	},
}

var n38 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n39,
		RBRACE: "}",
	},
}

var n37 = &CompEntry {
	Type: CompEntryTypeAnon,
	Anon: &CompEntry_AnonOption{
		Initialiser: n38,
	},
}

var n36 = &CommaCompEntry {
	Type: CommaCompEntryTypeE,
	E: &CommaCompEntry_EOption{
		COMMA: ",",
		CompEntry: n37,
	},
}

var n21 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n22,
		CommaCompEntry: []*CommaCompEntry {
			n36,
		},
		COMMA: ",",
	},
}

var n20 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n21,
		RBRACE: "}",
	},
}

var n16 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n17,
		VariableDef: n18,
		EQUALS: "=",
		Initialiser: n20,
	},
}

var n15 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n16,
	},
}

var n14 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n15,
		SEMI: ";",
	},
}

var n13 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n14,
	},
}

var n11 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n12,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n13,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n11,
		},
	},
}
