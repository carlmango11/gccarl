package compiler

import "github.com/carlmango11/gccarl/gccarl/semantic"

type Register string

const (
	RegUnset Register = ""
	RegEAX   Register = "eax"
	RegEBX   Register = "ebx"
	RegEDI   Register = "edi"
	RegESI   Register = "esi"
	RegEDX   Register = "edx"
	RegR10D  Register = "r10d"
	RegRAX   Register = "rax"
	RegRBX   Register = "rbx"
	RegRDI   Register = "rdi"
	RegRSI   Register = "rsi"
	RegRDX   Register = "rdx"
	RegR10   Register = "r10"
	RegAL    Register = "al"
	RegDIL   Register = "dil"
	RegSIL   Register = "sil"
	RegDL    Register = "dl"
	RegR10B  Register = "r10b"
)

var paramReg = map[semantic.Size][]Register{
	1: {RegDIL, RegSIL, RegDL, RegR10B},
	4: {RegEDI, RegESI, RegEDX, RegR10D},
	8: {RegRDI, RegRSI, RegRDX, RegR10},
}
