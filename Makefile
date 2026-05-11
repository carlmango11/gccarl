BUILDDIR := build
LD := x86_64-elf-ld

build:
	mkdir -p build

build/gccarl: build gccarl/
	go build -o build/gccarl gccarl/cmd/gccarl/main.go

build/basic.asm: build/gccarl
	build/gccarl -o build/basic.asm samples/basic.c

build/basic.o: build/basic.asm
	nasm -f elf64 build/basic.asm -o build/basic.o

build/basic: build/basic.o
	#clang -arch x86_64 build/basic.o -o build/basic
	$(LD) -o $(BUILDDIR)/basic $(BUILDDIR)/basic.o

all: build/basic

PHONY: clean

clean:
	rm -rf build