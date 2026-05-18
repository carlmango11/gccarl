package main

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/compiler"
	"github.com/carlmango11/gccarl/gccarl/semantic"
)

func main() {
	prog := &semantic.Program{
		FuncDefs: []*semantic.FuncDef{
			{
				ReturnType: semantic.Type{
					Kind: semantic.KindArray,
				},
				Name: "main",
				Locals: map[semantic.VarName]semantic.Type{
					"x": {
						Kind: semantic.KindArray,
						SubType: &semantic.Type{
							Kind: semantic.KindPrimitive,
							Prim: semantic.PrimChar,
						},
						ArraySize: 2,
					},
				},
				Statements: []*semantic.Statement{
					{
						ArrayAssign: &semantic.ArrayAssign{
							Name: "x",
							Type: semantic.Type{
								Kind: semantic.KindArray,
								SubType: &semantic.Type{
									Kind: semantic.KindPrimitive,
									Prim: semantic.PrimChar,
								},
								ArraySize: 2,
							},
							Vals: []*semantic.Expr{
								{
									Type: semantic.Type{
										Kind: semantic.KindPrimitive,
										Prim: semantic.PrimChar,
									},
									Literal: &semantic.Literal{
										Char: 'H',
									},
								},
								{
									Type: semantic.Type{
										Kind: semantic.KindPrimitive,
										Prim: semantic.PrimChar,
									},
									Literal: &semantic.Literal{
										Char: 'i',
									},
								},
							},
						},
					},
					{
						Expr: &semantic.Expr{
							Type: semantic.Type{
								Kind: semantic.KindVoid,
							},
							FuncCall: &semantic.FuncCall{
								Func: "print",
								Args: []*semantic.Expr{
									{
										Type: semantic.Type{
											Kind: semantic.KindPrimitive,
											Prim: semantic.PrimInt32,
										},
										Literal: &semantic.Literal{
											Int32: 1,
										},
									},
									{
										Type: semantic.Type{
											Kind: semantic.KindPointer,
											SubType: &semantic.Type{
												Kind: semantic.KindPrimitive,
												Prim: semantic.PrimChar,
											},
										},
										AddressOf: "x",
									},
									{
										Type: semantic.Type{
											Kind: semantic.KindPrimitive,
											Prim: semantic.PrimInt32,
										},
										Literal: &semantic.Literal{
											Int32: 2,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	c := compiler.New()
	output, err := c.Compile(prog)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(output))
}
