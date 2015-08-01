package page

import (
	"ptr"
)

type PageEntryPacked uint64

type Page [512]PageEntryPacked

type PageEntry struct{
	Address uint64
	Global, Large, Dirty, Accessed, CacheDisable, WriteThrough, User, ReadWrite, Present bool
}

const(
	PRESENT PageEntryPacked = 1 << iota
	READ_WRITE
	USER
	WRITE_THROUGH
	CACHE_DISABLED
	ACCESSED
	DIRTY
	LARGE
	GLOBAL
)

func (entry PageEntry)Pack()PageEntryPacked{
	e := PageEntryPacked(entry.Address & 0xFFFFFFFFFFFFF000)
	if entry.Global{
		e |= GLOBAL
	}
	if entry.Large{
		e |= LARGE
	}
	if entry.Dirty{
		e |= DIRTY
	}
	if entry.Accessed{
		e |= ACCESSED
	}
	if entry.CacheDisable{
		e |= CACHE_DISABLED
	}
	if entry.WriteThrough{
		e |= WRITE_THROUGH
	}
	if entry.User{
		e |= USER
	}
	if entry.ReadWrite{
		e |= READ_WRITE
	}
	if entry.Present{
		e |= PRESENT
	}
	return e
}

func (entry PageEntryPacked)Unpack()PageEntry{
	return PageEntry{Address: uint64(entry) & 0xFFFFFFFFFFFFF000, 
		Global: entry & GLOBAL !=0, 
		Large: entry & LARGE !=0, 
		Dirty: entry & DIRTY !=0, 
		Accessed: entry & ACCESSED !=0, 
		CacheDisable: entry & CACHE_DISABLED !=0, 
		WriteThrough: entry & WRITE_THROUGH !=0, 
		User: entry & USER !=0, 
		ReadWrite: entry & READ_WRITE !=0, 
		Present: entry & PRESENT !=0}
}

func(entry PageEntryPacked)Address()uintptr{
	return uintptr(entry) & 0xFFFFF000
}

var(
	kernelPage uintptr
	dirPtrTable uintptr
	Mapl4 uintptr
	bootstrapPage uintptr
)

func page(p uintptr)*Page{
	return (*Page)(ptr.GetAddr(p & 0xFFFFF000))
}

//extern __kernel_end
func kernelEnd()uintptr

//extern __enable_paging
func enable(uintptr)

func init(){
	Mapl4 = (kernelEnd() & 0xFFFFF000) + 0x1000
	dirPtrTable = Mapl4 + 0x1000
	bootstrapPage = Mapl4 + 0x2000
	kernelPage = Mapl4 + 0x3000
	
	page(Mapl4)[0] = PageEntry{Address: uint64(dirPtrTable), ReadWrite: true, Present:true}.Pack()
	page(dirPtrTable)[0] = PageEntry{Address: uint64(bootstrapPage), Present:true}.Pack()
	page(kernelPage)[0] = PageEntry{Address: 0x200000, Large: true, ReadWrite: true, Present:true}.Pack()
	page(dirPtrTable)[3] = PageEntry{Address: uint64(kernelPage), Present:true}.Pack()
	page(dirPtrTable)[4] = PageEntry{Address: uint64(kernelPage), Present:true}.Pack()
	page(bootstrapPage)[0] = PageEntry{Address: 0, Large: true, Present:true, ReadWrite:true}.Pack()
	
	for i:=1; i<512; i++{
		page(Mapl4)[i] = 0
		if i!=4 && i !=3{
			page(dirPtrTable)[i] = 0
		}
		page(kernelPage)[i]  = 0
		page(bootstrapPage)[i] = 0
	}
	
	enable(dirPtrTable)
}
