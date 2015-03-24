### Build params

CC_CROSS = x86_64-elf-gcc
LD_CROSS = x86_64-elf-ld
GO_CROSS =x86_64-elf-gccgo
OBJCOPY = x86_64-elf-objcopy
PREPROC = $(CC_CROSS) -E -x c -P
CC = gcc
LD = ld
ASM = nasm -f elf64
CFLAGS_CROSS = -Werror -nostdlib -fno-builtin -nostartfiles -nodefaultlibs -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2
GOFLAGS_CROSS = -static  -Werror -nostdlib -nostartfiles -nodefaultlibs -mcmodel=large -mno-red-zone -mno-mmx -mno-sse -mno-sse2
INCLUDE_DIRS = -Ipkg/.

### Sources

CORE_SOURCES = pkg/ptr.go.o pkg/ptr.gox pkg/color.go.o pkg/color.gox pkg/video.go.o pkg/video.gox pkg/segment.go.o pkg/segment.gox pkg/gdt.go.o pkg/gdt.gox pkg/types.go.o pkg/types.gox pkg/loader.o pkg/goose.go.o

SOURCE_OBJECTS = $(CORE_SOURCES)
 
### Targets

all: kernel.iso

clean:
	rm -f $(SOURCE_OBJECTS) $(TEST_EXECS) kernel64.bin kernel64.o

### Rules

pkg/%.o: %.s
	$(ASM) $(INCLUDE_DIRS) -o $@ $<

pkg/%.gox: pkg/%.go.o
	$(OBJCOPY) -j .go_export $< $@

pkg/%.go.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) $(INCLUDE_DIRS) -o $@ -c $<

kernel64.bin: $(SOURCE_OBJECTS)
	$(LD_CROSS) -T link.ld -z max-page-size=0x1000 -o kernel64.bin $(SOURCE_OBJECTS)
 
kernel64.o: kernel64.bin
	i686-elf-objcopy -B i386 -I binary -O elf32-i386 kernel64.bin kernel64.o