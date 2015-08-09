GO_CROSS = i386-elf-gccgo
LD_CROSS = i386-elf-ld
ASM = nasm -f elf

GOFLAGS_CROSS = -static -Werror -nostdlib

all: kernel.iso

.PHONY: clean
clean:
	rm -f isodir/boot/kernel.bin kernel.iso loader.o
	make -C kernel clean
	make -C bootstrap clean
	make -C testapp clean

.PHONY: boot-nogrub
boot-nogrub: kernel.bin
	qemu-system-i386 -kernel isodir/boot/kernel.bin -m 1024

.PHONY: boot
boot: kernel.iso
	qemu-system-x86_64 -cdrom kernel.iso

.PHONY: bochs
bochs: kernel.iso
	bochs -q

kernel.bin: loader.o bootstrap.a kernel.o
	$(LD_CROSS) -t link.ld --whole-archive -o isodir/boot/kernel.bin loader.o kernel/kernel.o bootstrap/bootstrap.a 

kernel.iso: kernel.bin
	grub-mkrescue -o kernel.iso isodir

loader.o: loader.s
	$(ASM) -o $@ $<

.PHONY: kernel.o
kernel.o:
	make -C kernel kernel.o

.PHONY: bootstrap.a
bootstrap.a:
	make -C bootstrap