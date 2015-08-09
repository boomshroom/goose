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
	mov rsp, stack + 128
	;push rbp
	;push rbp
	mov rbp, rsp

	call __go_init_main
	call main.main

global __go_print_string
__go_print_string:
	mov rdx, 0
	int 0x80
	ret

global __go_print_nl
__go_print_nl:
	mov rdx, 1
	int 0x80
	ret

; Go compatibility
;__go_type_hash_identity:
;    jmp go.types.HashIdent
;__go_type_equal_identity:
;    jmp go.types.EqualIdent

;section .bss
stack: resb 128