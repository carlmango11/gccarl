BUILDDIR := build
SRC := $(wildcard gccarl/*)
LD := x86_64-elf-ld

$(BUILDDIR)/basic: $(BUILDDIR)/basic.o
	$(LD) -o $(BUILDDIR)/basic $(BUILDDIR)/basic.o

$(BUILDDIR):
	mkdir -p $(BUILDDIR)

$(BUILDDIR)/gccarl: $(BUILDDIR) $(SRC)
	go build -o $(BUILDDIR)/gccarl gccarl/cmd/gccarl/main.go

$(BUILDDIR)/parser: $(BUILDDIR) $(SRC)
	go build -o $(BUILDDIR)/gccarl gccarl/cmd/parser/main.go

$(BUILDDIR)/basic.asm: $(BUILDDIR)/gccarl
	$(BUILDDIR)/gccarl -o $(BUILDDIR)/basic.asm samples/basic.c

$(BUILDDIR)/basic.o: $(BUILDDIR)/basic.asm
	nasm -f elf64 $(BUILDDIR)/basic.asm -o $(BUILDDIR)/basic.o

all: $(BUILDDIR)/basic

PHONY: clean

clean:
	rm -rf $(BUILDDIR)