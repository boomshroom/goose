### Build params

GO_CROSS =x86_64-elf-gccgo
ASM = nasm -f elf64
GOFLAGS_CROSS = -static -Werror -nostdlib -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2
 
### Targets

all: ../pkg/runtime.o

### Rules
 
../pkg/runtime.o: *.go
	$(GO_CROSS) $(GOFLAGS_CROSS) -o $@ -c *.go

../pkg/runtime_asm.o: runtime.s
	$(ASM) -o $@ $<