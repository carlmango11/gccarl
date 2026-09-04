
package ast

var MainNode = n0

var n6 = &Type {
	Type: TypeTypeChar,
	Char: &Type_CharOption{
		CHAR_TYPE: "char",
	},
}

var n7 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "age",
	},
}

var n5 = &VarDec {
	Type: VarDecTypeVarDec,
	VarDec: &VarDec_VarDecOption{
		Type: n6,
		VariableDef: n7,
	},
}

var n4 = &VarDecColon {
	Type: VarDecColonTypeC,
	C: &VarDecColon_COption{
		VarDec: n5,
		SEMI: ";",
	},
}

var n3 = &StructBlock {
	Type: StructBlockTypeBlock,
	Block: &StructBlock_BlockOption{
		VarDecColon: n4,
	},
}

var n2 = &TypeDef {
	Type: TypeDefTypeStructDef,
	StructDef: &TypeDef_StructDefOption{
		STRUCT: "struct",
		IDEN: "Person",
		LBRACE: "{",
		StructBlock: n3,
		RBRACE: "}",
	},
}

var n1 = &DecDef {
	Type: DecDefTypeTypeDef,
	TypeDef: &DecDef_TypeDefOption{
		TypeDef: n2,
	},
}

var n9 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n14 = &Type {
	Type: TypeTypeStruct,
	Struct: &Type_StructOption{
		STRUCT: "struct",
		IDEN: "Person",
	},
}

var n15 = &VariableDef {
	Type: VariableDefTypeVariable,
	Variable: &VariableDef_VariableOption{
		IDEN: "p",
	},
}

var n20 = &EntryLabelField {
	Type: EntryLabelFieldTypeC,
	C: &EntryLabelField_COption{
		FULL_STOP: ".",
		IDEN: "age",
	},
}

var n21 = &EntryLabelField {
	Type: EntryLabelFieldTypeC,
	C: &EntryLabelField_COption{
		FULL_STOP: ".",
		IDEN: "year",
	},
}

var n19 = &EntryLabel {
	Type: EntryLabelTypeL,
	L: &EntryLabel_LOption{
		EntryLabelField: n20,
	},
}

var n25 = &Value {
	Type: ValueTypeChar,
	Char: &Value_CharOption{
		CHAR: "'X'",
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

var n22 = &Initialiser {
	Type: InitialiserTypeExpr,
	Expr: &Initialiser_ExprOption{
		Expr: n23,
	},
}

var n18 = &CompEntry {
	Type: CompEntryTypeLabelled,
	Labelled: &CompEntry_LabelledOption{
		EntryLabel: n19,
		EQUALS: "=",
		Initialiser: n22,
	},
}

var n17 = &CompEntries {
	Type: CompEntriesTypeEntries,
	Entries: &CompEntries_EntriesOption{
		CompEntry: n18,
		COMMA: ",",
	},
}

var n16 = &Initialiser {
	Type: InitialiserTypeList,
	List: &Initialiser_ListOption{
		LBRACE: "{",
		CompEntries: n17,
		RBRACE: "}",
	},
}

var n13 = &DecAssign {
	Type: DecAssignTypeStandard,
	Standard: &DecAssign_StandardOption{
		Type: n14,
		VariableDef: n15,
		EQUALS: "=",
		Initialiser: n16,
	},
}

var n12 = &Statement {
	Type: StatementTypeDecAssign,
	DecAssign: &Statement_DecAssignOption{
		DecAssign: n13,
	},
}

var n11 = &StatementComma {
	Type: StatementCommaTypeStatement,
	Statement: &StatementComma_StatementOption{
		Statement: n12,
		SEMI: ";",
	},
}

var n10 = &Line {
	Type: LineTypeStatement,
	Statement: &Line_StatementOption{
		StatementComma: n11,
	},
}

var n8 = &DecDef {
	Type: DecDefTypeFuncDef,
	FuncDef: &DecDef_FuncDefOption{
		Type: n9,
		IDEN: "main",
		LPAREN: "(",
		RPAREN: ")",
		LBRACE: "{",
		Line: []*Line {
			n10,
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
			n8,
		},
	},
}
