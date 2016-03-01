struc syscalls
	;.id:			resq	1
	;.phys:			resq	1
	;.rsp:			resq	1
	;.rip:			resq	1
	;.rax:			resq	1
	;.rbx:			resq	1
	;.flags:		resq	1
	;.capabilities	resq	1
	.syscall_len:	resq	1
	.syscalls:		resq	80
endstruc

struc proc
	.id:			resq	1
	;.phys:			resq	1
	.rsp:			resq	1
	.rip:			resq	1
	.rax:			resq	1
	.rbx:			resq	1
	.flags:			resq	1
	.capabilities:	resq	1
	.stack_page:	resq	1
	.syscall_page:	resq	1
	.elf_header:	resq	1
	.num_pages:		resq	1
	.pages:			resq	20
	.size:
endstruc

current_syscalls equ 0x7FFFFFFFF000