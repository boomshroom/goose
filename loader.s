;
; Adapted from osdev.orgs Bare Bones tutorial http://wiki.osdev.org/Bare_Bones
;

global loader
global magic
global mbd

; Go compatibility
global __go_runtime_error
global __go_type_hash_identity
global __go_type_equal_identity
global __go_type_hash_error
global __go_type_equal_error
global __go_register_gc_roots
global __kill

global __enable_paging
global __get_pages
global __get_kernel64
global __get_kernel64_size
global __enable_64bit

extern _binary_kernel_bin_start
extern _binary_kernel_bin_size

extern main.main
extern __go_init_main
extern go.video.ErrCode
extern go.types.HashIdent
extern go.types.EqualIdent

bits 32

; Multiboot stuff
MODULEALIGN equ  1<<0
MEMINFO     equ  1<<1
FLAGS       equ  MODULEALIGN | MEMINFO
MAGIC       equ  0x1BADB002
CHECKSUM    equ -(MAGIC + FLAGS)

section .text

align 4
MultiBootHeader:
    dd MAGIC
    dd FLAGS
    dd CHECKSUM

STACKSIZE equ 0x4000  ; Define our stack size at 16k

loader:
	cmp eax, MAGIC + 0x10000000
	je mb_start
	mov [multiboot_table], dword 0
	jmp start
	mb_start:
	mov [multiboot_table], ebx
	start:

	mov esp, stack + STACKSIZE ; Setup stack pointer
	
	;push  eax
	;push  ebx
	
	call __go_init_main
	call main.main

	cli
	
__go_runtime_error:
__go_type_hash_error:
__go_type_equal_error:
__kill:
	hlt
	jmp  __kill

__go_register_gc_roots:
	ret

; Go compatibility
__go_type_hash_identity:
    jmp go.types.HashIdent
__go_type_equal_identity:
    jmp go.types.EqualIdent
	
__enable_paging:
	push ebp
	mov ebp, esp

	mov eax, cr4
	or eax, 0x20
	mov cr4, eax
	
	;mov ecx, 0xC0000080
	;rdmsr
	;or eax, 0x101
	;wrmsr
	
	mov eax, [esp+8]
	mov cr3, eax
	
	mov eax, cr0
	or eax, 0x80000000
	
	cli
	
	mov cr0, eax
	
	mov esp, ebp
	pop ebp

	ret
	
GDT64:                           ; Global Descriptor Table (64-bit).
    .Null: equ $ - GDT64         ; The null descriptor.
    dw 0                         ; Limit (low).
    dw 0                         ; Base (low).
    db 0                         ; Base (middle)
    db 0                         ; Access.
    db 0                         ; Granularity.
    db 0                         ; Base (high).
    .Code: equ $ - GDT64         ; The code descriptor.
    dw 0                         ; Limit (low).
    dw 0                         ; Base (low).
    db 0                         ; Base (middle)
    db 10011000b                 ; Access.
    db 00100000b                 ; Granularity.
    db 0                         ; Base (high).
    .Data: equ $ - GDT64         ; The data descriptor.
    dw 0                         ; Limit (low).
    dw 0                         ; Base (low).
    db 0                         ; Base (middle)
    db 10010000b                 ; Access.
    db 00000000b                 ; Granularity.
    db 0                         ; Base (high).
    .Pointer:                    ; The GDT-pointer.
    dw $ - GDT64 - 1             ; Limit.
    dd GDT64                     ; Base.
    dd 0

__enable_64bit:
	push ebp
	mov ebp, esp
	
	;mov edx, [esp+8]

	mov eax, cr0
	and eax, 01111111111111111111111111111111b
	mov cr0, eax
	
	mov ebx, pages
	mov cr3, ebx
	
	mov ecx, 0xC0000080
	rdmsr
	or eax, 1<<8
	wrmsr
	
	;mov ebx, [esp+8]
	mov eax, cr0
	or eax, 0x80000000
	mov cr0, eax
	
	lgdt [GDT64.Pointer]

	mov eax, pages
	mov ebx, [multiboot_table]
	mov ecx, _binary_kernel_bin_start
	jmp GDT64.Code:0x40000000

__get_kernel64:
	mov eax, _binary_kernel_bin_start
	ret

__get_kernel64_size:
	mov eax, _binary_kernel_bin_size
	ret
	
__get_pages:
	mov eax, pages
	ret

section .bss

align 4096
pages: resb 4096 * 5
stack: resb STACKSIZE   ; Reserve 16k for stack
multiboot_table: resd 1