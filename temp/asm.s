section .text
bits 32

global inportb
inportb:
	push ebp
	mov	ebp, esp
	
	sub	esp, 4
	mov eax, 0
	mov edx, 0
	mov	dx, word [ebp+8]
	in al, dx
	
	leave
	ret

global outportb
outportb:
	push ebp
	mov	ebp, esp
	
	mov edx, 0
	mov eax, 0
	mov	dx, word [ebp+8]
	mov	al, byte [ebp+12]
	out dx, al
	
	leave
	ret
	
global io_wait
io_wait:
	mov eax, 0
	out 0x80, al
	ret
	

global enable_ints
enable_ints:
	sti
	ret