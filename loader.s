;
; Adapted from osdev.org's Bare Bones tutorial http://wiki.osdev.org/Bare_Bones
;

global loader
global magic
global mbd

; Go compatibility
global __go_register_gc_roots
global __go_runtime_error
global __go_type_hash_identity
global __go_type_equal_identity
global __go_type_hash_error
global __go_type_equal_error
global __go_print_string
global __go_print_uint64
global __go_print_bool
global __go_print_nl

global __load_idt
global __load_gdt
global __generic_isr
global __test_int
global __arbitrary_convert
global __call
global __reload_segments
global __enable_paging
global __kernel_start
global __kernel_end
global __kernel_size

extern kernel_start
extern kernel_end
extern kernel_size

extern go.kernel.Kmain
extern go.kernel.init
extern go.video.ErrCode
extern go.video.Print
extern go.video.PrintUint
extern go.video.PrintBool
extern go.video.NL
extern go.types.HashIdent
extern go.types.EqualIdent
extern go.types.HashError
extern go.types.EqualError

extern go.idt.IDT
extern go.gdt.GDT
extern go.idt.ISR

; Multiboot stuff
MODULEALIGN equ  1<<0
MEMINFO     equ  1<<1
FLAGS       equ  MODULEALIGN | MEMINFO
MAGIC       equ  0x1BADB002
CHECKSUM    equ -(MAGIC + FLAGS)

section .text

bits 32

align 4
    dd MAGIC
    dd FLAGS
    dd CHECKSUM

STACKSIZE equ 0x4000  ; Define our stack size at 16k

loader:
    mov  esp, stack + STACKSIZE ; Setup stack pointer

    mov  [magic], eax
    mov  [mbd], ebx


    ;call go.kernel.init
    call go.kernel.Kmain   ; Jump to Go's kernel.Kmain

    ;cli
.hang:
    hlt
    jmp  .hang

; Go compatibility
__go_runtime_error:
    jmp go.video.ErrCode
__go_type_hash_identity:
    jmp go.types.HashIdent
__go_type_equal_identity:
    jmp go.types.EqualIdent
__go_type_hash_error:
    jmp go.types.HashError
__go_type_equal_error:
    jmp go.types.EqualError
__go_print_string:
    jmp go.video.Print
__go_print_uint64:
    jmp go.video.PrintUint
__go_print_bool:
    jmp go.video.PrintBool
__go_print_nl:
    jmp go.video.NL
    
__load_idt:
    lidt [go.idt.IDT]
    ret
    
gdtr dw 0
	dd 0
    
__load_gdt:
    cli
    mov eax, [esp+4]
    mov [gdtr+2], eax
    mov ax, [esp+8]
    mov [gdtr], ax
    lgdt [gdtr]
    ret
    
__reload_segments:
	jmp 0x0008:reload_cs
    reload_cs:
	mov ax, 0x10
	mov ds, ax
	mov es, ax
	mov fs, ax
	mov gs, ax
	mov ss, ax
	ret
    
__call:
	push ebp
	mov ebp, esp
	;sub esp, 0
	mov eax,  dword [ebp+8]
	push dword [ebp+12]
	call eax
	leave
	ret
	
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
	mov cr0, eax
	
	mov esp, ebp
	pop ebp
	ret
    
__arbitrary_convert: ret

__kernel_start:
	mov eax, kernel_start
	ret
	
__kernel_end:
	mov eax, kernel_end
	ret

section .bss

align 4
stack: resb STACKSIZE   ; Reserve 16k for stack
magic: resd 1
mbd:   resd 1
