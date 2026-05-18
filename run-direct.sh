go run gccarl/tools/compile/compile.go > build/direct.asm
nasm -f elf64 build/direct.asm -o build/direct.o
x86_64-elf-ld -o build/direct build/direct.o
docker build -t direct .
docker run --rm --platform linux/amd64 direct