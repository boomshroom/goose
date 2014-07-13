;
; Adapted from osdev.org's Bare Bones tutorial http://wiki.osdev.org/Bare_Bones
;

global loader
global magic
global mbd

; Go compatibility
global __go_register_gc_roots
global __go_runtime_error

global __unsafe_get_addr

extern go.kernel.Kmain
extern go.video.Error

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

    cli
.hang:
    hlt
    jmp  .hang

__unsafe_get_addr:  ; Allows us to convert int -> ptr in go
    push ebp
    mov ebp,esp
    mov eax, [ebp+8]
    mov  esp,ebp
    pop  ebp
    ret

; Go compatibility - noop'd
__go_runtime_error:
    call go.video.Error
__go_register_gc_roots:
    ret

section .bss

align 4
stack: resb STACKSIZE   ; Reserve 16k for stack
magic: resd 1
mbd:   resd 1
