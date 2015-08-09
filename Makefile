GO_CROSS = i386-elf-gccgo
LD_CROSS = i386-elf-ld
ASM = nasm -f elf

GOFLAGS_CROSS = -static -Werror -nostdlib

all: kernel.iso

.PHONY: clean
clean:
	rm -f isodir/boot/kernel.bin kernel.iso kernel.sys kernel.img loader.o
	make -C kernel clean
	make -C bootstrap clean
	make -C testapp clean

.PHONY: boot-nogrub
boot-nogrub: kernel.bin
	qemu-system-i386 -kernel isodir/boot/kernel.bin -m 1024

.PHONY: boot-grub
boot-grub: kernel.iso
	qemu-system-x86_64 -cdrom kernel.iso

.PHONY: boot
boot: kernel.img
	qemu-system-x86_64 -vga std -smp 8 -m 256 -drive id=disk,file=kernel.img,if=none -device ahci,id=ahci -device ide-drive,drive=disk,bus=ahci.0 -name "GoOSe" -net nic,model=i82551

.PHONY: bochs
bochs: kernel.iso
	bochs -q

kernel.bin: loader.o bootstrap.a kernel.o
	$(LD_CROSS) -t link.ld --whole-archive -o isodir/boot/kernel.bin loader.o kernel/kernel.o bootstrap/bootstrap.a 

kernel.iso: kernel.bin
	grub-mkrescue -o kernel.iso isodir

kernel.img: kernel/kernel.sys pure64
	bmfs kernel.img initialize 6M Pure64/bmfs_mbr.sys Pure64/pure64.sys kernel.sys

loader.o: loader.s
	$(ASM) -o $@ $<

.PHONY: kernel/kernel.sys
kernel/kernel.sys:
	make -C kernel

.PHONY: kernel.o
kernel.o:
	make -C kernel kernel.o

.PHONY: bootstrap.a
bootstrap.a:
	make -C bootstrap

.PHONY: pure64
pure64:
	cd Pure64 && ./build.sh