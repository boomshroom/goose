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

print_string:

	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 1
	mov qword [rax+8], rdi
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1
	ret

global __go_print_string
__go_print_string:

	mov [str_buf], rdi
	mov [str_buf+8], rsi

	mov rdi, str_buf
	jmp print_string

global __go_print_nl
__go_print_nl:

	mov rdi, nl_str
	jmp print_string

global __go_print_int64
global __go_print_uint64
__go_print_int64:
__go_print_uint64:

	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 3
	mov qword [rax+8], rdi
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1

	ret

global __scan_char
__scan_char:
	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 2
	mov qword [rax+8], rdi
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1

	ret
	
global __request
__request:
	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 4
	mov qword [rax+8], rdi
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1

global __register_interupt
	mov rax, current_proc + proc.syscall_len
	mov rax, [rax]
	shl rax, 4
	mov rdx, current_proc + proc.syscalls
	add rax, rdx
	mov qword [rax], 4
	mov qword [rax+8], rdi
	mov rax, current_proc + proc.syscall_len
	add qword [rax], 1

global __go_runtime_error
__go_runtime_error:
	mov rdi, err_str
	call print_string
	.kill:
	jmp .kill

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

	err_start: db "Runtime Error", `\n`
	err_end:
	err_len equ err_end - err_start
	err_str: dq err_start, err_len

section .bss
align 0x800
stack: resb 0x800
str_buf: resq 2 ; temp until malloc