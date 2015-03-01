package page

import (
	"ptr"
)

type PageEntryPacked uint64

type Page [512]PageEntryPacked

type DirPointerTable [4]uint64

type PageEntry struct{
	Address uint64
	Global, Size, Dirty, Accessed, CacheDisable, WriteThrough, User, ReadWrite, Present bool
}

const(
	PRESENT PageEntryPacked = 1 << iota
	READ_WRITE
	USER
	WRITE_THROUGH
	CACHE_DISABLED
	ACCESSED
	DIRTY
	SIZE
	GLOBAL
)

func PackEntry(entry PageEntry)PageEntryPacked{
	e := PageEntryPacked(entry.Address  & 0xFFFFFFFFFFFFF000)
	if entry.Global{
		e |= GLOBAL
	}
	if entry.Size{
		e |= SIZE
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

var(
	//pageAlignedEnd uintptr = (end & 0xFFFFF000) + 0x1000
	PageDir uintptr
	firstPage uintptr
	dirPtrTable uintptr
	mapl4 uintptr
)

func page(p uintptr)*Page{
	return (*Page)(ptr.GetAddr(p & 0xFFFFF000))
}

func ptrTable(p uintptr)*DirPointerTable{
	return (*DirPointerTable)(ptr.GetAddr(p & 0xFFFFFFE0))
}

//extern __kernel_end
func kernelEnd()uintptr

//extern __enable_paging
func enable(uintptr)

func Init(){

	mapl4 = (kernelEnd() & 0xFFFFF000) + 0x1000
	dirPtrTable = mapl4
	PageDir = mapl4 + 0x1000
	firstPage = mapl4 + 0x2000
	
	//page(mapl4)[0] = PackEntry(PageEntry{Address: uint64(dirPtrTable), ReadWrite: true, Present: true}) //superviser level, read/write, present
	ptrTable(dirPtrTable)[0] = uint64(PageDir) | 3

	for i:=0;i<512;i++{
		page(PageDir)[i] = READ_WRITE //superviser level, read/write, not present
		
		page(firstPage)[i] = PackEntry(PageEntry{Address: uint64(i<<12), ReadWrite: true, Present: true}) //superviser level, read/write, present
		
		//ptrTable(dirPtrTable)[i] = (uint64(i)<<12) | 3
	}
	
	page(PageDir)[0] |= PackEntry(PageEntry{Address: uint64(firstPage), ReadWrite: true, Present: true}) //superviser level, read/write, present
	
	//page(PageDir)[511] = 
	
	enable(dirPtrTable)
}