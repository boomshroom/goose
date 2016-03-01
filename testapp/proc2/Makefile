### Build params

OBJCOPY_CROSS = x86_64-elf-objcopy
GO_CROSS =x86_64-elf-gccgo
GOFLAGS_CROSS = -static  -Werror -nostdlib -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2

ASM = nasm -f elf64

OBJECTS = main.o

### Rules

all: proc2

.PHONY: clean
clean: 
	rm -f proc2 ${OBJECTS}

proc2: ${OBJECTS}
	$(GO_CROSS) $(GOFLAGS_CROSS) -z max-page-size=0x1000 -o $@ main.o ../crt0.o
	cp proc2 ../../isodir/boot/proc2

main.o: main.go
	$(GO_CROSS) $(GOFLAGS_CROSS) -o $@ -c $<