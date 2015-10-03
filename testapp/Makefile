### Build params

OBJCOPY_CROSS = x86_64-elf-objcopy
GO_CROSS =x86_64-elf-gccgo
GOFLAGS_CROSS = -static  -Werror -nostdlib -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2

ASM = nasm -f elf64

OBJECTS = hello.o crt0.o

### Rules

all: main

clean: 
	rm -f app.o main ${OBJECTS}

main: ${OBJECTS}
	$(GO_CROSS) $(GOFLAGS_CROSS) -z max-page-size=0x1000 -o $@ hello.o crt0.o
	cp main ../isodir/boot/main

%.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) -o $@ -c $<

crt0.o: crt0.s
	$(ASM) -o $@ $<