package preprocessor

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/tokens"
)

type DirectiveType int

const (
	DirNone DirectiveType = iota
	DirIf
	DirElif
	DirElse
	DirEndIf
	DirDefine
	DirUndefine
	DirIfDef
	DirIfNotDef
	DirInclude
)

func (d DirectiveType) IsIf() bool {
	switch d {
	case DirIf:
		return true
	default:
		return false
	}
}

func toDirType(t *tokens.Token) (DirectiveType, error) {
	switch t.Name {
	case "IF":
		return DirIf, nil
	case "ELIF":
		return DirElif, nil
	case "IF_DEF":
		return DirIfDef, nil
	case "IF_NOT_DEF":
		return DirIfNotDef, nil
	case "ELSE":
		return DirElse, nil
	case "DEFINE":
		return DirDefine, nil
	case "UNDEFINE":
		return DirUndefine, nil
	case "END_IF":
		return DirEndIf, nil
	case "INCLUDE":
		return DirInclude, nil
	}

	return 0, fmt.Errorf("%v is not a valid directive", t.Val)
}
