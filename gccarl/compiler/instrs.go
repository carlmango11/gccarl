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
	i.addInstr("mov %s, %d", to, n)
}

func (i *Instrs) movByteToReg(b byte, to Register) {
	i.addInstr("mov %s, 0x%X", to, b)
}

func (i *Instrs) movOffsetToReg(s semantic.Size, from Offset, to Register) {
	i.addInstr("mov %s, %s", to, offsetOperand(s, from))
}

func (i *Instrs) movLocToReg(s semantic.Size, from Location, to Register) {
	i.addInstr("mov %s, %s", to, locOperand(s, from))
}

func (i *Instrs) movFromReg(s semantic.Size, from Register, to Offset) {
	i.addInstr("mov %s, %s", offsetOperand(s, to), from)
}

func (i *Instrs) cmp(t semantic.Size, reg Register, l2 Location) {
	i.addInstr("cmp %s, %s", reg, locOperand(t, l2))
}

func (i *Instrs) addComment(format string, args ...any) {
	i.addInstr("; "+format, args...)
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

		return string(l.Register)
	case LTOffset:
		return offsetOperand(s, l.Offset)
	}

	panic("invalid")
}
