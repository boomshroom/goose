global __go_runtime_error
global __go_type_hash_identity
global __go_type_equal_identity
global __go_type_hash_error
global __go_type_equal_error
global __go_type_hash_string
global __go_type_equal_string
global __go_memcmp
global __go_strcmp
global __go_copy
global __go_tdn_unsafe.Pointer
global _get_ptr_desc

global runtime.stringiter2
global __go_register_gc_roots

extern go.runtime.RuntimeError

extern go.runtime.TypeHashIdentity
extern go.runtime.TypeEqualIdentity
extern go.runtime.TypeHashError
extern go.runtime.TypeEqualError
extern go.runtime.TypeHashString
extern go.runtime.TypeEqualString
extern go.runtime.MemCmp
extern go.runtime.StrCmp
extern go.runtime.Copy
extern go.runtime.StringIter2
extern go.runtime.UnsafePointerDesc


; Printing builtin functions are in loader.s and direct straight to video.go

__go_runtime_error:
    jmp go.runtime.RuntimeError
__go_type_hash_identity:
    jmp go.runtime.TypeHashIdentity
__go_type_equal_identity:
    jmp go.runtime.TypeEqualIdentity
__go_type_hash_error:
    jmp go.runtime.TypeHashError
__go_type_equal_error:
    jmp go.runtime.TypeEqualError
__go_type_hash_string:
    jmp go.runtime.TypeHashString
__go_type_equal_string:
    jmp go.runtime.TypeEqualString
__go_memcmp:
	jmp go.runtime.MemCmp
__go_strcmp:
	jmp go.runtime.StrCmp
__go_copy:
    jmp go.runtime.Copy
runtime.stringiter2:
	jmp go.runtime.StringIter2

__go_register_gc_roots:
	ret

_get_ptr_desc:
	mov rax, __go_tdn_unsafe.Pointer
	ret

section .bss
__go_tdn_unsafe.Pointer: resb 80