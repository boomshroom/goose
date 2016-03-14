BITS 64
global loader

global __load_gdt
global __load_idt
global __reload_segments
global __stack_ptr
global __kernel_end
global __start_app
global __unwind_stack
global __enter_int

global __go_print_string
global __go_print_uint64
global __go_print_int64
global __go_print_pointer
global __go_print_bool
global __go_print_nl

global __break

extern go.kernel.Kmain

extern go.video.Print
extern go.video.PrintUint
extern go.video.PrintBool
extern go.elf.PrintAddress
extern go.video.NL

extern main.main
extern __go_init_main
extern go.page.SetPageLoc
extern go.tables.SetTable
extern go.gdt.SetKernelStack
extern go.elf.KernelElf
extern go.syscall.Syscall
extern go.proc.Procs
extern go.proc.CurrentID
extern go.proc.NumProcs

extern go.idt.IDT
;extern go.kernel.Kmain

extern kernel_end

section .entry
loader:
	mov rdx, qword __loader
	jmp rdx

section .text
	
STACKSIZE equ 0x4000
STACKPTR equ stack + STACKSIZE

__loader:
	;mov rsp, STACKPTR
	mov rdi, rax
	mov rax, qword go.elf.KernelElf
	mov qword [rax], rcx
	call go.page.SetPageLoc
	
	call __go_init_main

	mov rdi, rbx
	call go.tables.SetTable

	xchg bx,bx

	call main.main
		
__kill:
	hlt
	jmp __kill

__break:
	xchg bx,bx
	ret

__go_print_string:
    jmp go.video.Print
__go_print_uint64:
__go_print_int64:
__go_print_pointer:
    jmp go.video.PrintUint
__go_print_bool:
    jmp go.video.PrintBool
__go_print_nl:
    jmp go.video.NL
		
__load_gdt:
	cli
	;add eax,  KERNEL_VIRTUAL_BASE
	lgdt [rdi]

	mov ax, 0x28
	ltr ax

	mov rdi, STACKPTR
	call go.gdt.SetKernelStack

	ret

__load_idt:
	lidt [rdi]

	; enable syscall/sysret
	mov rax, 1
	cpuid
	and rdx, 1<<5
	test rdx, rdx
	jnz .msr
		mov rdi, qword msr_err
		mov rsi, qword msr_err_len
		call go.video.Print
		jmp __kill
	.msr:

	mov rax, 0x80000001
	cpuid
	and rdx, 1<<11
	test rdx, rdx
	jnz .sys
		mov rdi, qword syscall_err
		mov rsi, qword syscall_err_len
		call go.video.Print
		jmp __kill
	.sys:

	mov ecx, 0xC0000080
	rdmsr
	or eax, 1
	wrmsr

	xor eax, eax
	mov edx, 0x8 + (0xb << 16)
	or ecx, 1
	wrmsr

	mov rdx, qword __syscall
	mov eax, edx
	shr rdx, 32
	add ecx, 1
	wrmsr

	mov ecx, 0xC0000084
	rdmsr
	or eax, 1<<9
	wrmsr

	ret
    
__reload_segments:
	mov ax, 0x10
	;mov ds, ax
	;mov es, ax
	mov fs, ax
	mov gs, ax
	;mov ss, ax

	mov rax, rsp
	push 0
	push rax
	pushf
	push 0x8
	mov rax, qword .reloadss
	push rax
	iretq

.reloadss:
	ret

__start_app:
	mov ecx, 0xC0000102
	mov rax, go.proc.CurrentID
	mov qword [rax], 1
	mov rdx, go.proc.Procs
	lea rdx, [rdx + proc.size]
	mov eax, edx
	shr rdx, 32
	wrmsr

	mov rcx, [rdi]

	pushf
	pop r11
	mov rsp, 0x7FFFFFFFEFF0
	mov rbp, rsp
	; NASM/YASM seem to think that REX.W is a symbol

	db 0x48
	sysret

__enter_int:

	push rbp
	mov [rel temp_rsp], rsp
	mov rsp, 0x7FFFFFFFEFF0

	mov ecx, 0xC0000082
	mov rdx, qword .return
	mov eax, edx
	shr rdx, 32
	wrmsr

	mov rcx, [rdi]
	pushf
	pop r11

	; NASM/YASM seem to think that REX.W is a symbol
	db 0x48
	sysret
	.return:

	push r11
	popf

	mov rsp, qword [rel temp_rsp]
	pop rbp

	mov ecx, 0xC0000082
	mov rdx, qword __syscall
	mov eax, edx
	shr rdx, 32
	wrmsr

	ret

__syscall: ; flush syscalls and halt
	swapgs
	mov [gs:proc.rsp], rsp
	mov rsp, qword STACKPTR
	sti
	mov [gs:proc.rip], rcx ; rip
	mov [gs:proc.rax], rax
	mov [gs:proc.rbx], rbx
	mov [gs:proc.flags], r11 ; flags
	
	;call go.syscall.Syscall

	;mov rcx, [gs:proc.rip] ; rip
	;mov rax, [gs:proc.rax]
	;mov rbx, [gs:proc.rbx]
	;mov r11, [gs:proc.flags] ; flags
	;cli
	;mov rsp, [gs:proc.rsp]
	;swapgs

	;sti
	;xchg bx,bx
	call go.syscall.Syscall
	hlt
	;jmp .kill

	;iretq
	db 0x48
	sysret


__unwind_stack:
    pop rbx
    mov r12, rsp
    mov rsp, rbp
    pop rbp ; as though returning from Video.errPrint()

    pop rdi
    call go.video.PrintUint
    call go.video.NL

    ;pop rbx ; instruction that called Video.errPrint()
    mov r13, rsp
    mov rsp, rbp
    pop rbp ; as though returning from runtimeError()

    pop rdi
    call go.video.PrintUint
    call go.video.NL

	mov r13, rsp
    mov rsp, rbp
    pop rbp ; as though returning from runtimeError()

    pop rdi
    call go.elf.PrintAddress
    ;call go.video.NL

    mov r13, rsp
    mov rsp, rbp
    pop rbp ; as though returning from whatever failed()

    pop rdi
    call go.elf.PrintAddress
    ;call go.video.NL

    mov rsp, r12
    push rbx
    ret

__stack_ptr:
	mov rax, qword STACKPTR
	ret

__kernel_end:
	mov rax, qword kernel_end
	ret

%include "../proc.inc"

section .data
msr_err: db "Model specific registers not available"
msr_err_end:
msr_err_len equ msr_err_end - msr_err

syscall_err: db "Syscall/Sysret instructions not available"
syscall_err_end:
syscall_err_len equ syscall_err_end - syscall_err

section .bss
	
stack: resb STACKSIZE
temp_rsp: resq 1