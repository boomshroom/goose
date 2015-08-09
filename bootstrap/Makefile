### Build params

GO_CROSS = i386-elf-gccgo
AR_CROSS = i386-elf-ar
ASM = nasm -f elf
GOFLAGS_CROSS = -static -Werror -nostdlib
INCLUDE_DIRS = -Ipkg/.

### Sources
CORE_SOURCES = pkg/ptr.o pkg/elf.o pkg/page.o pkg/types.o pkg/goose.o
# Disabled Sources # pkg/color.go.o pkg/color.gox pkg/video.go.o pkg/video.gox pkg/interupts.o pkg/asm.o pkg/asm.go.o pkg/asm.gox pkg/regs.go.o pkg/regs.gox pkg/gdt.go.o pkg/gdt.gox pkg/idt.go.o pkg/idt.gox pkg/pit.go.o pkg/pit.gox pkg/kbd.go.o pkg/kbd.gox 

SOURCE_OBJECTS = $(CORE_SOURCES)

### Rules

all: bootstrap.a

.PHONY: clean
clean:
	rm -f pkg/* bootstrap.a

pkg/%.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) $(INCLUDE_DIRS) -o $@ -c $<

bootstrap.a: $(SOURCE_OBJECTS)
	$(AR_CROSS) rcs bootstrap.a $(SOURCE_OBJECTS)
