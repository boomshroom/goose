package proc

import "unsafe"

var Current = (*Proc)(unsafe.Pointer(uintptr(0x7FFFFFFFF000)))

type Proc struct{
	Id uint64
	PhysicalAddr uintptr
	RSP, RIP, RAX, RBX uint64
	Flags uint64
	SyscallLen uint64
	Syscalls [40]Syscall
}

type Syscall struct{
	Id uint64
	Args unsafe.Pointer
}