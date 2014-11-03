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
;global __go_strcomp
global __go_print_string
global __go_print_nl
;global __go_set_closure

global __load_idt
global __load_gdt
global __generic_isr
global __test_int
global __arbitrary_convert
global __call
global __remap_irq
global __reload_segments

;global __unsafe_get_addr

extern go.kernel.Kmain
extern go.video.Error
;extern go.runtime.StringComp
extern go.video.Print
extern go.video.NL

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

align 4
    dd MAGIC
    dd FLAGS
    dd CHECKSUM

STACKSIZE equ 0x4000  ; Define our stack size at 16k

loader:
    mov  esp, stack + STACKSIZE ; Setup stack pointer

    mov  [magic], eax
    mov  [mbd], ebx

    call go.kernel.Kmain   ; Jump to Go's kernel.Kmain

    ;cli
.hang:
    hlt
    jmp  .hang

; Go compatibility - noop'd
__go_runtime_error:
    jmp go.video.Error
;__go_register_gc_roots:
;    ret
__go_type_hash_identity:
    ret
__go_type_equal_identity:
    ret
;__go_strcmp:
;    jmp go.runtime.StringComp
__go_print_string:
    jmp go.video.Print
__go_print_nl:
    jmp go.video.NL
;__go_set_closure:
;    ret
    
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
	jmp 0x08:reload_cs
    reload_cs:
	mov ax, 0x10
	mov ds, ax
	mov es, ax
	mov fs, ax
	mov gs, ax
	mov ss, ax
	ret
    
__test_int:
    int 33
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
    
__arbitrary_convert: ret

__remap_irq:
    push eax
    mov al, 0x11
    out 0x20, al
    out 0xA0, al
    mov al, 0x20
    out 0x21, al
    mov al, 0x28
    out 0xA1, al
    mov al, 0x04
    out 0x21, al
    mov al, 0x02
    out 0xA1, al
    mov al, 0x01
    out 0x21, al
    out 0xA1, al
    mov al, 0x00
    out 0x21, al
    out 0xA1, al
    pop eax
    ret

extern go.video.PutChar
global __test_args
__test_args:
;    push ebp
;    mov ebp, esp
;    sub esp, 4
;    mov esp, 65
    ;push 0x41
    jmp go.video.PutChar
;    pop ebp
;    ret

section .bss

align 4
stack: resb STACKSIZE   ; Reserve 16k for stack
magic: resd 1
mbd:   resd 1
