	.file	"hello.go"
	.text
.Ltext0:
	.section	.go_export,"",@progbits
	.byte	0x76,0x31,0x3b,0x0a
	.byte	0x70,0x61,0x63,0x6b,0x61,0x67,0x65,0x20
	.byte	0x6d,0x61,0x69,0x6e
	.byte	0x3b,0x0a
	.byte	0x70,0x6b,0x67,0x70,0x61,0x74,0x68,0x20
	.byte	0x6d,0x61,0x69,0x6e
	.byte	0x3b,0x0a
	.byte	0x70,0x72,0x69,0x6f,0x72,0x69,0x74,0x79,0x20,0x31,0x3b,0x0a
	.byte	0x66,0x75,0x6e,0x63,0x20
	.byte	0x49,0x6e,0x74
	.byte	0x20,0x28
	.byte	0x29
	.byte	0x3b,0x0a
	.byte	0x66,0x75,0x6e,0x63,0x20
	.byte	0x49,0x6e,0x74,0x52,0x65,0x74
	.byte	0x20,0x28
	.byte	0x29
	.byte	0x3b,0x0a
	.byte	0x63,0x68,0x65,0x63,0x6b,0x73,0x75,0x6d,0x20,0x41,0x45,0x34,0x39
	.byte	0x33,0x31,0x39,0x32,0x45,0x43,0x46,0x35,0x34,0x43,0x34,0x31,0x35
	.byte	0x36,0x42,0x30,0x30,0x46,0x34,0x31,0x38,0x38,0x42,0x44,0x36,0x33
	.byte	0x34,0x41,0x39,0x35,0x39,0x32,0x34,0x44,0x41,0x43,0x3b,0x0a
	.globl	main.IntRet$descriptor
	.section	.rodata.main.IntRet$descriptor,"a",@progbits
	.align 8
	.type	main.IntRet$descriptor, @object
	.size	main.IntRet$descriptor, 8
main.IntRet$descriptor:
	.quad	__int_ret
	.globl	main.Int$descriptor
	.section	.rodata.main.Int$descriptor,"a",@progbits
	.align 8
	.type	main.Int$descriptor, @object
	.size	main.Int$descriptor, 8
main.Int$descriptor:
	.quad	main.Int
	.section	.rodata
.LC0:
	.byte	0x48,0x65,0x6c,0x6c,0x6f,0x20,0x66,0x72,0x6f,0x6d,0x20,0x75,0x73
	.byte	0x65,0x72,0x73,0x70,0x61,0x63,0x65,0x21
	.zero	1
	.text
	.globl	main.main
	.type	main.main, @function
main.main:
.LFB0:
	.file 1 "hello.go"
	.loc 1 14 0
	.cfi_startproc
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	pushq	%rbx
	subq	$8, %rsp
	.cfi_offset 3, -24
	.loc 1 15 0
	movl	$.LC0, %eax
	movl	$21, %edx
	movq	%rax, %rcx
	movq	%rdx, %rbx
	movq	%rdx, %rax
	movq	%rcx, %rdi
	movq	%rax, %rsi
	call	__go_print_string
	call	__go_print_nl
	.loc 1 17 0
	movl	$main.Int$descriptor, %edi
	call	__register_interupt
	.loc 1 14 0
	addq	$8, %rsp
	popq	%rbx
	popq	%rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE0:
	.size	main.main, .-main.main
	.section	.rodata
.LC1:
	.byte	0x6b,0x65,0x79,0x62,0x6f,0x61,0x72,0x64,0x20,0x6d,0x65,0x73,0x73
	.byte	0x61,0x67,0x65,0x20,0x72,0x65,0x63,0x69,0x65,0x76,0x65,0x64,0x21
	.zero	1
	.text
	.globl	main.Int
	.type	main.Int, @function
main.Int:
.LFB1:
	.loc 1 44 0
	.cfi_startproc
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	pushq	%rbx
	subq	$8, %rsp
	.cfi_offset 3, -24
	.loc 1 45 0
	movl	$.LC1, %eax
	movl	$26, %edx
	movq	%rax, %rcx
	movq	%rdx, %rbx
	movq	%rdx, %rax
	movq	%rcx, %rdi
	movq	%rax, %rsi
	call	__go_print_string
	call	__go_print_nl
	.loc 1 46 0
	call	__int_ret
	.loc 1 44 0
	addq	$8, %rsp
	popq	%rbx
	popq	%rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE1:
	.size	main.Int, .-main.Int
	.globl	__go_init_main
	.type	__go_init_main, @function
__go_init_main:
.LFB2:
	.loc 1 1 0
	.cfi_startproc
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	.loc 1 1 0
	popq	%rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE2:
	.size	__go_init_main, .-__go_init_main
	.local	main.i
	.comm	main.i,8,8
	.data
	.align 32
	.type	main.kbdus, @object
	.size	main.kbdus, 128
main.kbdus:
	.byte	0
	.byte	27
	.byte	49
	.byte	50
	.byte	51
	.byte	52
	.byte	53
	.byte	54
	.byte	55
	.byte	56
	.byte	57
	.byte	48
	.byte	45
	.byte	61
	.byte	8
	.byte	9
	.byte	113
	.byte	119
	.byte	101
	.byte	114
	.byte	116
	.byte	121
	.byte	117
	.byte	105
	.byte	111
	.byte	112
	.byte	91
	.byte	93
	.byte	10
	.byte	0
	.byte	97
	.byte	115
	.byte	100
	.byte	102
	.byte	103
	.byte	104
	.byte	106
	.byte	107
	.byte	108
	.byte	59
	.byte	39
	.byte	96
	.byte	0
	.byte	92
	.byte	122
	.byte	120
	.byte	99
	.byte	118
	.byte	98
	.byte	110
	.byte	109
	.byte	44
	.byte	46
	.byte	47
	.byte	0
	.byte	42
	.byte	0
	.byte	32
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	45
	.byte	0
	.byte	0
	.byte	0
	.byte	43
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.byte	0
	.zero	38
	.align 32
	.type	main.shifted, @object
	.size	main.shifted, 127
main.shifted:
	.zero	39
	.byte	34
	.zero	4
	.byte	60
	.byte	95
	.byte	62
	.byte	63
	.byte	41
	.byte	33
	.byte	64
	.byte	35
	.byte	36
	.byte	37
	.byte	94
	.byte	38
	.byte	42
	.byte	40
	.zero	1
	.byte	58
	.zero	1
	.byte	43
	.zero	29
	.byte	123
	.byte	124
	.byte	125
	.zero	2
	.byte	126
	.zero	30
	.text
.Letext0:
	.section	.debug_info,"",@progbits
.Ldebug_info0:
	.long	0x79
	.value	0x4
	.long	.Ldebug_abbrev0
	.byte	0x8
	.uleb128 0x1
	.long	.LASF2
	.byte	0x16
	.long	.LASF3
	.long	.LASF4
	.quad	.Ltext0
	.quad	.Letext0-.Ltext0
	.long	.Ldebug_line0
	.uleb128 0x2
	.long	.LASF0
	.byte	0x1
	.byte	0xe
	.quad	.LFB0
	.quad	.LFE0-.LFB0
	.uleb128 0x1
	.byte	0x9c
	.uleb128 0x2
	.long	.LASF1
	.byte	0x1
	.byte	0x2c
	.quad	.LFB1
	.quad	.LFE1-.LFB1
	.uleb128 0x1
	.byte	0x9c
	.uleb128 0x3
	.long	.LASF5
	.byte	0x1
	.byte	0x1
	.long	.LASF6
	.quad	.LFB2
	.quad	.LFE2-.LFB2
	.uleb128 0x1
	.byte	0x9c
	.byte	0
	.section	.debug_abbrev,"",@progbits
.Ldebug_abbrev0:
	.uleb128 0x1
	.uleb128 0x11
	.byte	0x1
	.uleb128 0x25
	.uleb128 0xe
	.uleb128 0x13
	.uleb128 0xb
	.uleb128 0x3
	.uleb128 0xe
	.uleb128 0x1b
	.uleb128 0xe
	.uleb128 0x11
	.uleb128 0x1
	.uleb128 0x12
	.uleb128 0x7
	.uleb128 0x10
	.uleb128 0x17
	.byte	0
	.byte	0
	.uleb128 0x2
	.uleb128 0x2e
	.byte	0
	.uleb128 0x3f
	.uleb128 0x19
	.uleb128 0x3
	.uleb128 0xe
	.uleb128 0x3a
	.uleb128 0xb
	.uleb128 0x3b
	.uleb128 0xb
	.uleb128 0x11
	.uleb128 0x1
	.uleb128 0x12
	.uleb128 0x7
	.uleb128 0x40
	.uleb128 0x18
	.uleb128 0x2116
	.uleb128 0x19
	.byte	0
	.byte	0
	.uleb128 0x3
	.uleb128 0x2e
	.byte	0
	.uleb128 0x3f
	.uleb128 0x19
	.uleb128 0x3
	.uleb128 0xe
	.uleb128 0x3a
	.uleb128 0xb
	.uleb128 0x3b
	.uleb128 0xb
	.uleb128 0x6e
	.uleb128 0xe
	.uleb128 0x11
	.uleb128 0x1
	.uleb128 0x12
	.uleb128 0x7
	.uleb128 0x40
	.uleb128 0x18
	.uleb128 0x2117
	.uleb128 0x19
	.byte	0
	.byte	0
	.byte	0
	.section	.debug_aranges,"",@progbits
	.long	0x2c
	.value	0x2
	.long	.Ldebug_info0
	.byte	0x8
	.byte	0
	.value	0
	.value	0
	.quad	.Ltext0
	.quad	.Letext0-.Ltext0
	.quad	0
	.quad	0
	.section	.debug_line,"",@progbits
.Ldebug_line0:
	.section	.debug_str,"MS",@progbits,1
.LASF5:
	.string	"main.__go_init_main"
.LASF3:
	.string	"hello.go"
.LASF6:
	.string	"__go_init_main"
.LASF4:
	.string	"/home/angelo-arch/src/goose/testapp"
.LASF0:
	.string	"main.main"
.LASF2:
	.string	"GNU Go 5.2.0 -mtune=generic -march=x86-64"
.LASF1:
	.string	"main.Int"
	.ident	"GCC: (GNU) 5.2.0"
