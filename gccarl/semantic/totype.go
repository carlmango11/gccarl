package semantic

import (
	"fmt"
	"strconv"

	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
)

type RawType struct {
	typ     *cparser.Type
	pointer *RawType
	arrays  []int
}

// TODO: if I changed the grammar a bit I wouldn't need 2 versions of this func
func (b *builder) paramTypeToRaw(typ *cparser.Type, args *cparser.TypeArgs) (*RawType, error) {
	switch args.Type {
	case cparser.TypeArgsTypePointer:
		inner, err := b.paramTypeToRaw(typ, args.Pointer.TypeArgs)
		if err != nil {
			return nil, err
		}

		return &RawType{
			typ:     typ,
			pointer: inner,
		}, nil
	case cparser.TypeArgsTypeArray:
		var sizes []int
		for _, a := range args.Array.ArrayIndexDef {
			if a.ArrayIndex.NUM == "" {
				return nil, fmt.Errorf("array index must be non-empty")
			}

			var err error
			size, err := strconv.Atoi(string(a.ArrayIndex.NUM))
			if err != nil {
				return nil, err
			}

			sizes = append(sizes, size)
		}

		return &RawType{
			typ:    typ,
			arrays: sizes,
		}, nil
	}

	panic("unreachable")
}

func (b *builder) varDefToRaw(typ *cparser.Type, v *cparser.VariableDef) (*RawType, error) {
	switch v.Type {
	case cparser.VariableDefTypePointer:
		inner, err := b.varDefToRaw(typ, v.Pointer.VariableDef)
		if err != nil {
			return nil, err
		}

		return &RawType{
			typ:     typ,
			pointer: inner,
		}, nil
	case cparser.VariableDefTypeVariable:
		var sizes []int
		for _, a := range v.Variable.ArrayIndexDef {
			size := -1

			if a.ArrayIndex.NUM != "" {
				var err error
				size, err = strconv.Atoi(string(a.ArrayIndex.NUM))
				if err != nil {
					return nil, err
				}
			}

			sizes = append(sizes, size)
		}

		return &RawType{
			typ:    typ,
			arrays: sizes,
		}, nil
	}

	panic("invalid type")
}

func (b *builder) toType(raw *RawType, isParam bool) (Type, error) {
	if raw.pointer != nil {
		sub, err := b.toType(raw.pointer, false) // no need for array decay check anymore
		if err != nil {
			return Type{}, err
		}

		return Type{
			Kind:    KindPointer,
			SubType: &sub,
		}, nil
	}

	return b.toVarType(raw.typ, raw.arrays, isParam)
}

func (b *builder) toVarType(typ *cparser.Type, arrs []int, isParam bool) (Type, error) {
	if len(arrs) > 0 {
		if isParam {
			// arrays decay to pointers in params
			return Type{
				Kind: KindPointer,
				Prim: astTypeToPrim(typ),
			}, nil
		}

		size := arrs[0]

		sub, err := b.toVarType(typ, arrs[1:], false)
		if err != nil {
			return Type{}, err
		}

		return Type{
			Kind:      KindArray,
			SubType:   &sub,
			ArraySize: size,
		}, nil
	}

	switch typ.Type {
	case cparser.TypeTypeStruct:
		st, ok := b.structType(typ.Struct.IDEN)
		if !ok {
			panicf("%v is not defined", typ.Struct.IDEN)
		}

		return st, nil
	case cparser.TypeTypeCustom:
		t, ok := b.customType(typ.Custom.IDEN)
		if !ok {
			panicf("%v is not defined", typ.Custom.IDEN)
		}

		return t, nil
	}

	return Type{
		Kind: KindPrimitive,
		Prim: astTypeToPrim(typ),
	}, nil
}

func (b *builder) structType(n cparser.IDEN) (Type, bool) {
	for i := len(b.scopes) - 1; i >= 0; i-- {
		s, ok := b.scopes[i].structs[n]
		if ok {
			return s, true
		}
	}

	return Type{}, false
}

func (b *builder) customType(n cparser.IDEN) (Type, bool) {
	for i := len(b.scopes) - 1; i >= 0; i-- {
		s, ok := b.scopes[i].types[n]
		if ok {
			return s, true
		}
	}

	return Type{}, false
}
