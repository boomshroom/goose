### Build params

OBJCOPY_CROSS = x86_64-elf-objcopy
GO_CROSS =x86_64-elf-gccgo
GOFLAGS_CROSS = -static  -Werror -nostdlib -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2

ASM = nasm -f elf64

OBJECTS = hello.o crt0.o

### Rules

all: app.o

clean: 
	rm -f app.o hello ${OBJECTS}

app.o: hello
	$(OBJCOPY_CROSS) -B i386 -I binary -O elf64-x86-64 hello app.o

hello: ${OBJECTS}
	$(GO_CROSS) $(GOFLAGS_CROSS) -z max-page-size=0x1000 -o $@ hello.o crt0.o

%.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) -o $@ -c $<

crt0.o: crt0.s
	$(ASM) -o $@ $<