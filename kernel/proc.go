package proc

import (
	"page"
	"unsafe"
)

var Current = (*Syscalls)(unsafe.Pointer(uintptr(0x7FFFFFFFF000)))

type Syscalls struct {
	Len      uint64
	Syscalls [40]Syscall
}

var Procs [20]Proc
var CurrentID uint64
var NumProcs uint64

type Proc struct {
	Id uint64
	//PhysicalAddr       uintptr
	RSP, RIP, RAX, RBX uint64
	Flags              uint64
	Capability         uint
	StackPage          page.PageEntryPacked
	SyscallPage        page.PageEntryPacked
	ElfHeader          unsafe.Pointer // avoid circular dependency, may replace with custom debug setup
	NumPages           int
	Pages              [10]PageMapping
}

type PageMapping struct {
	Physical page.PageEntryPacked
	Virtual  uintptr
}

func (p *Proc) Enable() {
	p.StackPage.Enable(0x7FFFFFFFE000, page.K)
	p.SyscallPage.Enable(0x7FFFFFFFF000, page.K)
	for i := 0; i < p.NumPages; i++ {
		pg := &p.Pages[i]
		pg.Physical.Enable(pg.Virtual, page.K)
	}
}

func KillProc() {
	p := &Procs[CurrentID]
	for pg := 0; pg < p.NumPages; pg++ {
		p.Pages[pg].Physical.Free()
	}
	p.StackPage.Free()
	p.SyscallPage.Free()
	for i := CurrentID; i < NumProcs; i++ {
		p = &Procs[i]
		*p = Procs[i+1]
		p.Id = i
	}
	NumProcs--
}

type Syscall struct {
	Id   uint64
	Args unsafe.Pointer
}
