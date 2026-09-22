package semantic

import (
	"testing"

	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
	"github.com/stretchr/testify/require"
)

func TestToVarOption(t *testing.T) {
	zType := Type{
		Kind: KindPrimitive,
		Prim: PrimInt32,
	}

	yElType := Type{
		Kind: KindStruct,
		Struct: StructType{
			Fields: []StructField{
				{
					Name: "z",
					Type: zType,
				},
			},
		},
	}

	yType := Type{
		Kind:    KindArray,
		SubType: &yElType,
	}

	xType := Type{
		Kind: KindPointer,
		SubType: &Type{
			Kind: KindStruct,
			Struct: StructType{
				Fields: []StructField{
					{
						Name: "y",
						Type: yType,
					},
				},
			},
		},
	}

	xID := VarID{
		ID:   123,
		Name: "x",
	}

	b := &builder{
		scopes: []*Scope{
			{
				vars: map[cparser.IDEN]Var{
					"x": {
						Type: xType,
						ID:   xID,
					},
				},
			},
		},
	}

	vo := &cparser.SubExpr_VariableOption{
		SubVariableAccess: &cparser.SubVariableAccess{
			V: &cparser.SubVariableAccess_VOption{
				IDEN: "x",
			},
		},
		InnerSubVariableAccess: []*cparser.InnerSubVariableAccess{
			{
				Type: cparser.InnerSubVariableAccessTypeArrow,
				Arrow: &cparser.InnerSubVariableAccess_ArrowOption{
					SubVariableAccess: &cparser.SubVariableAccess{
						V: &cparser.SubVariableAccess_VOption{
							IDEN: "y",
							ArrayIndexAccess: []*cparser.ArrayIndexAccess{
								{
									ArrayIndex: &cparser.ArrayIndexAccess_ArrayIndexOption{
										NUM: "3",
									},
								},
							},
						},
					},
				},
			},
			{
				Type: cparser.InnerSubVariableAccessTypeDot,
				Dot: &cparser.InnerSubVariableAccess_DotOption{
					SubVariableAccess: &cparser.SubVariableAccess{
						V: &cparser.SubVariableAccess_VOption{
							IDEN: "z",
						},
					},
				},
			},
		},
	}

	expected := &Expr{
		Type: zType,
		Field: &FieldExpr{
			Expr: &Expr{
				Type: yElType,
				Index: &IndexExpr{
					Expr: &Expr{
						Type: yType,
						Field: &FieldExpr{
							Expr: &Expr{
								Type: *xType.SubType,
								Deref: &Expr{
									Type: xType,
									Var:  &xID,
								},
							},
							Field: "y",
						},
					},
					Index: 3,
				},
			},
			Field: "z",
		},
	}

	actual, err := b.fromVarOption(vo)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
