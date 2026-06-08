package compiler

import (
	"fmt"

	"github.com/carlmango11/gccarl/gccarl/semantic"
)

type Instrs struct {
	instrs []Instr
}

func (i *Instrs) addInstr(format string, args ...any) {
	instr := fmt.Sprintf(format, args...)
	i.instrs = append(i.instrs, Instr(instr))
}

func (i *Instrs) addInstrsIndent(inner *Instrs) {
	for _, instr := range inner.instrs {
		i.instrs = append(i.instrs, "\t"+instr)
	}
}

func (i *Instrs) movInt32ToReg(n int32, to Register) {
	i.addInstr("mov %s, %d", to.Raw(4), n)
}

func (i *Instrs) movByteToReg(b byte, to Register) {
	i.addInstr("mov %s, 0x%X", to.Raw(1), b)
}

func (i *Instrs) movOffsetToReg(s semantic.Size, from Offset, to Register) {
	i.addInstr("mov %s, %s", to.Raw(s), offsetOperand(s, from))
}

func (i *Instrs) movLocToReg(s semantic.Size, from Location, to Register) {
	i.addInstr("mov %s, %s", to.Raw(s), locOperand(s, from))
}

func (i *Instrs) movLocToStack(s semantic.Size, from Location, to Offset) {
	i.addInstr("mov %s, %s", to, locOperand(s, from))
}

func (i *Instrs) movFromReg(s semantic.Size, from Register, to Offset) {
	i.addInstr("mov %s, %s", offsetOperand(s, to), from.Raw(s))
}

func (i *Instrs) cmp(s semantic.Size, reg Register, l2 Location) {
	i.addInstr("cmp %s, %s", reg.Raw(s), locOperand(s, l2))
}

func (i *Instrs) addComment(format string, args ...any) {
	i.addInstr("; "+format, args...)
}

func (i *Instrs) add(s semantic.Size, reg Register, l2 Location) {
	i.addInstr("add %s, %s", reg.Raw(s), locOperand(s, l2))
}

func offsetOperand(s semantic.Size, o Offset) string {
	return fmt.Sprintf("%s [rbp-%d]", typeInstrSize(s), o)
}

func locOperand(s semantic.Size, l Location) string {
	switch l.Type {
	case LTRegister:
		if l.Register == RegUnset {
			panic("unset register in location")
		}

		return string(l.Register.Raw(s))
	case LTOffset:
		return offsetOperand(s, l.Offset)
	}

	panic("invalid")
}
