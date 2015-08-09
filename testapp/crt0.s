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
	mov rbp, 0x1000
	push rbp
	push rbp
	mov rbp, rsp

	call __go_init_main
	call main.main

; Go compatibility
;__go_type_hash_identity:
;    jmp go.types.HashIdent
;__go_type_equal_identity:
;    jmp go.types.EqualIdent