package compiler

import (
	"testing"

	"github.com/carlmango11/gccarl/gccarl/semantic"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	tcs := []struct {
		prog     *semantic.Program
		expected []Instr
	}{
		{
			prog: &semantic.Program{
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
									Prim: semantic.PrimInt32,
								},
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
										ArraySize: 1,
									},
									Vals: []*semantic.Expr{
										{
											Type: semantic.Type{
												Kind: semantic.KindPrimitive,
												Prim: semantic.PrimChar,
											},
											Literal: &semantic.Literal{
												Int32: 74,
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
										Args: []*semantic.Expr{},
									},
								},
							},
						},
					},
				},
			},
			expected: []Instr{
				"mov dword [rbp-4], 5",
				"mov dword [rbp-8], 3",
				"mov rax, [rbp-8]",
				"mov rbx, [rbp-4]",
				"add rax, rbx",
				"mov dword [rbp-8], rax",
			},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			c := New()
			err := c.compile(tc.prog)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, c.instrs)
		})
	}
}
