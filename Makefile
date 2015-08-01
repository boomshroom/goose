GO_CROSS =i686-elf-gccgo
OBJCOPY = i686-elf-objcopy
LD_CROSS = i686-elf-ld
ASM = nasm -f elf

GOFLAGS_CROSS = -static  -Werror -nostdlib -nostartfiles -nodefaultlibs

all: kernel.iso

.PHONY: clean
clean:
	rm -f isodir/boot/kernel.bin kernel.iso
	make -C kernel clean
	make -C bootstrap clean

.PHONY: boot-nogrub
boot-nogrub: kernel.bin
	qemu-system-i386 -kernel isodir/boot/kernel.bin -m 1024

.PHONY: boot
boot: kernel.iso
	qemu-system-x86_64 -cdrom kernel.iso

.PHONY: bochs
bochs: kernel.iso
	bochs -q

kernel.bin: kernel.o bootstrap.a loader.o 
	$(GO_CROSS) $(GOFLAGS_CROSS) -t link.ld -o isodir/boot/kernel.bin loader.o kernel/kernel.o bootstrap/bootstrap.a 
 
kernel.iso: kernel.bin
	grub-mkrescue -o kernel.iso isodir

loader.o: loader.s
	$(ASM) -o $@ $<

.PHONY: kernel.o
kernel.o:
	make -C kernel

.PHONY: bootstrap.a
bootstrap.a:
	make -C bootstrap
