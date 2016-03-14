### Build params

GO_CROSS =x86_64-elf-gccgo
ASM = nasm -f elf64
GOFLAGS_CROSS = -static -Werror -nostdlib -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2 -fno-exceptions
INCLUDE_DIRS = -Ipkg/.

### Sources

CORE_SOURCES = pkg/loader.s.o pkg/interupts.s.o pkg/asm.s.o pkg/asm.o pkg/capability.o pkg/segment.o pkg/page.o pkg/proc.o pkg/gdt.o pkg/runtime.o pkg/runtime_asm.o pkg/color.o pkg/video.o pkg/vbe.o pkg/vga.o pkg/tables.o pkg/elf.o pkg/idt.o pkg/kbd.o pkg/syscall.o pkg/goose.o
# 
SOURCE_OBJECTS = $(CORE_SOURCES)
 
### Targets

all: kernel.bin

.PHONY: clean
clean:
	rm -f $(SOURCE_OBJECTS) kernel64.bin kernel64.o

### Rules

pkg/%.s.o: %.s
	$(ASM) -o $@ $<

.PHONY: pkg/runtime.o
pkg/runtime.o:
	make -C runtime

.PHONY: pkg/runtime_asm.o
pkg/runtime_asm.o:
	make -C runtime ../pkg/runtime_asm.o

pkg/%.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) $(INCLUDE_DIRS) -o $@ -c $<

kernel.bin: $(SOURCE_OBJECTS)
	$(GO_CROSS) $(GOFLAGS_CROSS)  -T link.ld -z max-page-size=0x1000 -o kernel.bin $(SOURCE_OBJECTS)
 
kernel.o: kernel.bin
	i386-elf-objcopy -B i386 -I binary -O elf32-i386 kernel.bin kernel.o

.PHONY: ../testapp/app.o
../testapp/app.o:
	make -C ../testapp
