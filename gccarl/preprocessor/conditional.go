package preprocessor

import (
	"fmt"
	"io"
)

func (r *reader) processIfDef(positiveDef bool) error {
	name, err := r.nextToken()
	if err != nil {
		return err
	}

	skip := len(r.ifs) > 0 && r.ifs[len(r.ifs)-1].skip
	executed := true // pretend we executed if skipping

	if !skip {
		_, defined := r.macros[name.Val]
		executed = defined == positiveDef
		skip = !executed
	}

	r.ifs = append(r.ifs, &ifState{
		executed: executed,
		skip:     skip,
	})

	r.directive.finished = true
	return nil
}

func (r *reader) processIf() error {
	exprTokens, err := r.parseRest()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("unexpected end of if expression")
		}

		return err
	}

	skip := len(r.ifs) > 0 && r.ifs[len(r.ifs)-1].skip
	executed := true // pretend we executed if skipping

	if !skip {
		executed, err = r.evalExpr(exprTokens)
		if err != nil {
			return err
		}

		skip = !executed
	}

	r.ifs = append(r.ifs, &ifState{
		executed: executed,
		skip:     skip,
	})

	r.directive.finished = true
	return nil
}

func (r *reader) processEndIf() error {
	r.ifs = r.ifs[:len(r.ifs)-1]

	r.directive.finished = true
	return nil
}

//func asdf() {
//case DirIf, DirIfNotDef, DirIfDef:
//	skip := len(r.ifs) > 0 && r.ifs[len(r.ifs)-1].skip
//	executed := true // pretend we executed if skipping
//
//	if !skip {
//	executed = r.eval(s.dir)
//	skip = !executed
//	}
//
//	r.ifs = append(r.ifs, &ifState{
//	executed: executed,
//	skip:     skip,
//	})
//
//	continue
//	case DirElif:
//	curr := r.ifs[len(ifs)-1]
//
//	if curr.executed {
//	curr.skip = true
//	continue
//	}
//
//	curr.executed = r.eval(s.dir)
//	curr.skip = !curr.executed
//
//	continue
//
//	case DirElse:
//	curr := r.ifs[len(ifs)-1]
//
//	if curr.executed {
//	curr.skip = true
//	continue
//	}
//
//	curr.executed = true
//	curr.skip = false
//
//	continue
//
//	case DirEndIf:
//	r.ifs = r.ifs[:len(r.ifs)-1]
//	continue
//
//	}
