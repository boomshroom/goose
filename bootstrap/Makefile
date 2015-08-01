### Build params

CC_CROSS = i686-elf-gcc
LD_CROSS = i686-elf-ld
GO_CROSS = i686-elf-gccgo
AR_CROSS = i686-elf-ar
OBJCOPY = i686-elf-objcopy
PREPROC = $(CC_CROSS) -E -x c -P
CC = gcc
LD = ld
ASM = nasm -f elf
CFLAGS_CROSS = -Werror -nostdlib -fno-builtin -nostartfiles -nodefaultlibs
GOFLAGS_CROSS = -static  -Werror -nostdlib -nostartfiles -nodefaultlibs
INCLUDE_DIRS = -Ipkg/.

### Sources
CORE_SOURCES = pkg/ptr.go.o pkg/ptr.gox pkg/elf.go.o pkg/elf.gox pkg/page.go.o pkg/page.gox pkg/types.go.o pkg/types.gox pkg/goose.go.o
# Disabled Sources # pkg/color.go.o pkg/color.gox pkg/video.go.o pkg/video.gox pkg/interupts.o pkg/asm.o pkg/asm.go.o pkg/asm.gox pkg/regs.go.o pkg/regs.gox pkg/gdt.go.o pkg/gdt.gox pkg/idt.go.o pkg/idt.gox pkg/pit.go.o pkg/pit.gox pkg/kbd.go.o pkg/kbd.gox 

SOURCE_OBJECTS = $(CORE_SOURCES)

### Rules

all: bootstrap.a

.PHONY: clean
clean:
	rm -f $(SOURCE_OBJECTS) bootstrap.o

pkg/%.o: %.s
	$(ASM) $(INCLUDE_DIRS) -o $@ $<

pkg/%.gox: pkg/%.go.o
	$(OBJCOPY) -j .go_export $< $@

pkg/%.go.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) $(INCLUDE_DIRS) -o $@ -c $<

bootstrap.a: $(SOURCE_OBJECTS)
	$(AR_CROSS) rcs bootstrap.a $(SOURCE_OBJECTS)
