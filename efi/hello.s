format pe64 dll efi
entry main

section '.text' code executable readable

include 'uefi.inc'

main:

	InitializeLib
	jc .end
	
	uefi_call_wrapper ConOut, OutputString, ConOut, _hello
	
	;uefi_call_wrapper BootServices, OpenProtocol, efi_handler, protocaluuid, loaded_image, efi_handler, 0, 2
	
	;mov ebx, [loaded_image.LoadOptionsSize]
	;mov edx, []
	
	mov			dword [memmapdescsize], 48
	uefi_call_wrapper	BootServices, GetMemoryMap, memmapsize, memmapbuff, memmapkey, memmapdescsize, memmapdescver
	cmp			dword [memmapdescsize], 0
	jnz			@f
	mov			dword [memmapdescsize], 48
@@:	clc
	cmp			rax, EFI_SUCCESS
	je			@f
	stc
@@:

	uefi_call_wrapper	BootServices, ExitBootServices
	
	.halt:
	jmp .halt
	
.end:	mov eax, EFI_SUCCESS
	retn
	
section '.data' data readable writeable

_hello	du 'Hello World',13,10,0

;loaded_image EFI_LOADED_IMAGE_PROTOCOL

;protocaluuid db EFI_LOADED_IMAGE_PROTOCOL_UUID

MEMMAP_BUFFSIZE equ 0x10

memmapsize:	dq			MEMMAP_BUFFSIZE
memmapkey:	dq			0
memmapdescsize:	dq			0
memmapdescver:	dq			0
memmapbuff:	rb			MEMMAP_BUFFSIZE

section '.reloc' fixups data discardable