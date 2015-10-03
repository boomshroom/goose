extern __go_init_main
extern main.main

;extern go.types.HashIdent
;extern go.types.EqualIdent

;global __go_type_equal_identity
;global __go_type_hash_identity

section .text

global _start
_start:
; Setup end of stack frame
	mov rsp, stack + 0xFFF
	;push rbp
	;push rbp
	mov rbp, rsp

	call __go_init_main
	call main.main

global __go_print_string
__go_print_string:

	mov [str_buf], rdi
	mov [str_buf+8], rsi

	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 1
	mov qword [rax+8], str_buf
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1

	ret

global __go_print_nl
__go_print_nl:

	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 1
	mov qword [rax+8], nl_str
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1

	ret

; Go compatibility
;__go_type_hash_identity:
;    jmp go.types.HashIdent
;__go_type_equal_identity:
;    jmp go.types.EqualIdent

%include "../proc.inc"

section data
	nl_start: db `\n`
	nl_end:
	nl_len equ nl_end - nl_start
	nl_str: dq nl_start, nl_len

section .bss
align 0x800
stack: resb 0x800
str_buf: resq 2 ; temp until malloc