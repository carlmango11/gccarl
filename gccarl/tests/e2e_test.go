package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/carlmango11/gccarl/gccarl/compiler"
	"github.com/carlmango11/gccarl/gccarl/generated/cparser"
	"github.com/carlmango11/gccarl/gccarl/semantic"
	"github.com/stretchr/testify/assert"
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

func TestPrograms(t *testing.T) {
	d, err := os.ReadDir("./testdata")
	if err != nil {
		log.Fatal(err)
	}

	exclude := map[string]bool{
		"printf.c": true,
		"for.c":    true,
	}

	for _, fe := range d {
		fileBytes, err := os.ReadFile("./testdata/" + fe.Name())
		if err != nil {
			log.Fatal(err)
		}

		t.Run(fe.Name(), func(t *testing.T) {
			if exclude[fe.Name()] {
				return
			}
			if fe.Name() != "array_equals.c" {
				return
			}

			runTest(t, fileBytes)
		})
	}
}

func runTest(t *testing.T, fileBytes []byte) {
	tree, err := cparser.Parse(string(fileBytes))
	if err != nil {
		assert.NoError(t, err, "cannot parse")
		return
	}

	program, err := semantic.Build(tree)
	if err != nil {
		assert.NoError(t, err, "cannot build semantic program")
		return
	}

	cc := compiler.New()

	c, err := cc.Compile(program)
	if err != nil {
		assert.NoError(t, err, "cannot compile")
		return
	}

	t.Logf("program: \n%v", string(c))

	binary, err := assemble(c)
	if err != nil {
		assert.NoError(t, err, "cannot assemble")
		return
	}

	mu, err := uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
	if err != nil {
		log.Fatal(err)
	}
	defer mu.Close()

	const size = 0x1000 * 100
	const start = 0x1000

	err = mu.MemMap(0x1000, size)
	if err != nil {
		log.Fatalf("cannot map memory: %v", err)
	}

	err = mu.RegWrite(uc.X86_REG_RSP, start+size)
	if err != nil {
		log.Fatalf("cannot set stack pointer: %v", err)
	}

	err = mu.MemWrite(start, binary)
	if err != nil {
		log.Fatalf("cannot write: %v", err)
	}

	_, err = mu.HookAdd(uc.HOOK_INSN, func(mu uc.Unicorn) {
		number, err := mu.RegRead(uc.X86_REG_RAX)
		if err != nil {
			log.Fatalf("cannot read syscall number: %v", err)
		}
		if number != 9 {
			return
		}

		expected, err := mu.RegRead(uc.X86_REG_RDI)
		if err != nil {
			log.Fatalf("cannot read syscall number: %v", err)
		}

		actual, err := mu.RegRead(uc.X86_REG_RSI)
		if err != nil {
			log.Fatalf("cannot read syscall number: %v", err)
		}

		assert.Equal(t, expected, actual)
	}, 1, 0, uc.X86_INS_SYSCALL)
	if err != nil {
		log.Fatalf("cannot hook syscall: %v", err)
	}

	//_, err = mu.HookAdd(uc.HOOK_CODE, func(mu uc.Unicorn, addr uint64, size uint32) {
	//	code, err := mu.MemRead(addr, uint64(size))
	//	if err != nil {
	//		log.Fatalf("cannot read instruction at %#x: %v", addr, err)
	//	}
	//	rax, _ := mu.RegRead(uc.X86_REG_RAX)
	//	rsp, _ := mu.RegRead(uc.X86_REG_RSP)
	//	rbp, _ := mu.RegRead(uc.X86_REG_RBP)
	//	log.Printf("RIP=%#x bytes=% x RAX=%#x RSP=%#x RBP=%#x", addr, code, rax, rsp, rbp)
	//}, 1, 0)
	//if err != nil {
	//	log.Fatalf("cannot hook instruction trace: %v", err)
	//}

	_, err = mu.HookAdd(uc.HOOK_MEM_INVALID,
		func(mu uc.Unicorn, access int, addr uint64, size int, value int64) bool {
			rip, _ := mu.RegRead(uc.X86_REG_RIP)
			rsp, _ := mu.RegRead(uc.X86_REG_RSP)

			log.Printf(
				"memory fault: access=%d addr=%#x size=%d value=%#x RIP=%#x RSP=%#x",
				access, addr, size, value, rip, rsp,
			)
			return false // Leave the access unhandled so Unicorn reports the error.
		}, 1, 0)
	if err != nil {
		log.Fatal(err)
	}

	if err := mu.Start(start, start+uint64(len(binary))); err != nil {
		log.Printf("program: \n%v", string(c))
		log.Fatalf("cannot start: %v", err)
	}
}

func assemble(source []byte) ([]byte, error) {
	dir, err := os.MkdirTemp("", "gccarl-nasm-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	asmPath := filepath.Join(dir, "test.asm")
	binPath := filepath.Join(dir, "test.bin")
	source = append([]byte("BITS 64\nORG 0x1000\n"), source...)
	if err := os.WriteFile(asmPath, source, 0600); err != nil {
		return nil, err
	}

	cmd := exec.Command("nasm", "-f", "bin", "-o", binPath, asmPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("nasm: %w\n%s", err, output)
	}

	return os.ReadFile(binPath)
}
