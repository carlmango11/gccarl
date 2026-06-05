package compiler

import "github.com/carlmango11/gccarl/gccarl/semantic"

type Register int

const (
	RegUnset Register = iota
	RegA
	RegB
	RegD
	RegSI
	RegR10
)

func (r Register) Raw(s semantic.Size) RawReg {
	raw, ok := sizedReg[r][s]
	if !ok {
		panic("unknown register type")
	}

	return raw
}

type RawReg string

const (
	RawUnset RawReg = ""
	RawEAX   RawReg = "eax"
	RawEBX   RawReg = "ebx"
	RawEDI   RawReg = "edi"
	RawESI   RawReg = "esi"
	RawEDX   RawReg = "edx"
	RawR10D  RawReg = "r10d"
	RawRAX   RawReg = "rax"
	RawRBX   RawReg = "rbx"
	RawRDI   RawReg = "rdi"
	RawRSI   RawReg = "rsi"
	RawRDX   RawReg = "rdx"
	RawR10   RawReg = "r10"
	RawAL    RawReg = "al"
	RawDIL   RawReg = "dil"
	RawSIL   RawReg = "sil"
	RawDL    RawReg = "dl"
	RawR10B  RawReg = "r10b"
)

var paramReg = []Register{
	RegD, RegSI, RegD, RegR10,
}

var sizedReg = map[Register]map[semantic.Size]RawReg{
	RegA: {
		1: RawAL,
		4: RawEAX,
		8: RawRAX,
	},
	RegR10: {
		1: RawR10B,
		4: RawR10D,
		8: RawR10,
	},
	RegD: {
		1: RawRDX,
		4: RawEDX,
		8: RawDL,
	},
	RegSI: {
		1: RawRSI,
		4: RawESI,
		8: RawSIL,
	},
}
