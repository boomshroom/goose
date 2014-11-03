### Build params

CC_CROSS = i686-elf-gcc
LD_CROSS = i686-elf-ld
GO_CROSS = i686-elf-gccgo
OBJCOPY = i686-elf-objcopy
PREPROC = $(CC_CROSS) -E -x c -P
CC = gcc
LD = ld
ASM = nasm -f elf
CFLAGS_CROSS = -Werror -nostdlib -fno-builtin -nostartfiles -nodefaultlibs
GOFLAGS_CROSS = -static  -Werror -nostdlib -nostartfiles -nodefaultlibs 
INCLUDE_DIRS = -I.

### Sources

CORE_SOURCES = loader.o interupts.o asm.o asm.go.o asm.gox regs.go.o regs.gox ptr.go.o ptr.gox color.go.o color.gox video.go.o video.gox gdt.go.o gdt.gox idt.go.o idt.gox pit.go.o pit.gox kbd.go.o kbd.gox goose.go.o

SOURCE_OBJECTS = $(CORE_SOURCES)
 
### Targets

all: kernel.iso

clean:
	rm -f $(SOURCE_OBJECTS) $(TEST_EXECS) kernel.bin isodir/boot/kernel.bin kernel.iso

boot-nogrub: kernel.bin
	qemu-system-i386 -kernel kernel.bin -m 1024

boot: kernel.iso
	qemu-system-i386 -cdrom kernel.iso

### Rules

%.o: %.s
	$(ASM) $(INCLUDE_DIRS) -o $@ $<

%.gox: %.go.o
	$(OBJCOPY) -j .go_export $< $@

%.go.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) $(INCLUDE_DIRS) -o $@ -c $<

kernel.bin: $(SOURCE_OBJECTS)
	$(LD_CROSS) -T link$(PAGE).ld -o kernel.bin $(SOURCE_OBJECTS)
 
kernel.iso: kernel.bin
	cp kernel.bin isodir/boot/kernel.bin
	grub-mkrescue -o kernel.iso isodir
