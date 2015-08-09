
global loader

global __load_gdt
global __load_idt
global __reload_segments
global __stack_ptr
global __get_app
global __get_app_size
global __kernel_end
global __start_app

global __go_print_string
global __go_print_uint64
global __go_print_bool
global __go_print_nl

global __break

extern go.kernel.Kmain

extern go.video.Print
extern go.video.PrintUint
extern go.video.PrintBool
extern go.video.NL

extern main.main
extern __go_init_main
extern go.page.SetPageLoc
extern go.gdt.SetKernelStack

extern go.idt.IDT
;extern go.kernel.Kmain

extern _binary_hello_start
extern _binary_hello_size
extern kernel_end

section .entry
loader:
	mov rcx, __loader
	jmp rcx

section .text
	
STACKSIZE equ 0x4000
STACKPTR equ stack + STACKSIZE

__loader:
	;mov rsp, STACKPTR
	mov rdi, rax
	call go.page.SetPageLoc
	call __go_init_main
	call main.main
	;mov ax, 0x1B
	;mov ds, ax
	;mov es, ax
	;mov fs, ax
	;mov gs, ax
	
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

__break:
	xchg bx, bx
	ret

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

	mov ax, 0x28
	ltr ax
	ret

__load_idt:
	lidt [rdi]
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

__start_app:
	xchg bx,bx
	mov rbx, [rdi]

	mov ax, 0x23
	mov ds, ax
	mov es, ax
	mov fs, ax
	mov gs, ax

	mov rdi, rsp
	call go.gdt.SetKernelStack

	cli
	mov rax, rsp
	push 0x23
	push rax
	pushf
	push 0x1B
	push rbx

	xchg bx,bx

	iretq

__stack_ptr:
	mov rax, STACKPTR
	ret

__get_app:
	mov rax, _binary_hello_start
	ret 

__get_app_size:
	mov rax, _binary_hello_size
	ret 

__kernel_end:
	mov rax, kernel_end
	ret

section .bss
	
stack: resb STACKSIZE