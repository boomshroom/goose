
global loader

; Go compatibility
global __go_runtime_error
global __go_type_hash_identity
global __go_type_equal_identity
global __go_type_hash_error
global __go_type_equal_error
global __go_print_string
global __go_print_uint64
global __go_print_bool
global __go_print_nl

global __load_gdt
global __reload_segments
global __stack_ptr

extern go.kernel.Kmain
extern go.video.ErrCode
extern go.video.Print
extern go.video.PrintUint
extern go.video.PrintBool
extern go.video.NL
extern go.types.HashIdent
extern go.types.EqualIdent
extern go.types.HashError
extern go.types.EqualError

extern go.kernel64.Kmain

section .entry
loader: jmp __loader

section .text
	
STACKSIZE equ 0x4000
STACKPTR equ stack + STACKSIZE

__loader:
	;mov rsp, STACKPTR

	call go.kernel64.Kmain
	
	;mov ax, 0x1B
	;mov ds, ax
	;mov es, ax
	;mov fs, ax
	;mov gs, ax
	
	.kill:
		hlt
		jmp .kill
	
	;mov rax, rsp
	;push 0x1B
	;push rax
	;pushf
	;push 0x08
	;mov rax, __kill
	;push rax
	
	;iretd
	
	;.kill:
	;	hlt
	;	jmp .kill
		
__kill:
	hlt
	jmp __kill
		
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
		
__load_gdt:
	cli
	;add eax,  KERNEL_VIRTUAL_BASE
	lgdt [rdi]
	;mov ax, 0x20
	;.kill:
		;hlt
		;jmp .kill
	;ltr ax
	ret
    
__reload_segments:
	;jmp 0x00000008:reload_cs
	;reload_cs:
		mov ax, 0x10
		mov ds, ax
		mov es, ax
		mov fs, ax
		mov gs, ax
		mov ss, ax
	ret

__stack_ptr:
	mov rax, STACKPTR
	ret

section .bss
	
stack: resb STACKSIZE