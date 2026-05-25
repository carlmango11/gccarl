
package ast

var MainNode = n0

var n2 = &Type {
	Type: TypeTypeInt,
	Int: &Type_IntOption{
		INT_TYPE: "int",
	},
}

var n9 = &Variable {
	Type: VariableTypeVariable,
	Variable: &Variable_VariableOption{
		IDEN: "x",
	},
}

var n8 = &Value {
	Type: ValueTypeVariable,
	Variable: &Value_VariableOption{
		Variable: n9,
	},
}

var n7 = &SubExpr {
	Type: SubExprTypeValue,
	Value: &SubExpr_ValueOption{
		Value: n8,
	},
}

var n12 = &Value {
	Type: ValueTypeInt,
	Int: &Value_IntOption{
		NUM: "4",
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

var n6 = &CompExpr {
	Type: CompExprTypeIsEqual,
	IsEqual: &CompExpr_IsEqualOption{
		SubExpr: n7,
		EEQUALS: "==",
		Expr: n10,
	},
}

var n5 = &Expr {
	Type: ExprTypeComp,
	Comp: &Expr_CompOption{
		CompExpr: n6,
	},
}

var n4 = &Control {
	Type: ControlTypeIf,
	If: &Control_IfOption{
		IF: "if",
		LPAREN: "(",
		Expr: n5,
		RPAREN: ")",
		LBRACE: "{",
		RBRACE: "}",
	},
}

var n3 = &Line {
	Type: LineTypeControl,
	Control: &Line_ControlOption{
		Control: n4,
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
		},
		RBRACE: "}",
	},
}

var n0 = &Main {
	Type: MainTypeMain,
	Main: &Main_MainOption{
		DecDef: []*DecDef {
			n1,
		},
	},
}
