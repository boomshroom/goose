package page

import (
	"ptr"
)

/*
// Helper function for finding locations of virtual memory
// Copy to playground to determine values
func indecies(addr uint64)[5]uint64{
	return [...]uint64{
		4: addr>>39,
		3: (addr>>30) & 0x1FF,
		2: (addr>>21) & 0x1FF,
		1: (addr>>12) & 0x1FF,
		0: addr & 0x1000,
	}
}
*/

func MapAddress(logical, physical uintptr, size PageSize, props PageEntryPacked){
	ml4Entry := &page(Mapl4)[logical >> 39]

	if !ml4Entry.HasProp(PRESENT){

	}

	dptIndex := logical >> 30
	pdtIndex := logical >> 21
	tableIndex := logical >> 12


}

type PageEntryPacked uintptr

type Page [512]PageEntryPacked

func (p *Page)GetEntry(addr uintptr, level Level)*PageEntryPacked{

	if level < PDT {
		return 0
	}
	entry := p[addr >> 39]
	if !entry.HasProp(PRESENT){
		return 0
	}else if entry.HasProp(LARGE){
		return entry
	}else{
		return entry.NextLevel(addr << 9, level-1)
	}
}

type PageSize uint8

const(
	K PageSize = iota // 4 KiB pages
	M // 2 MiB pages
	G // 1 GiB pages
)

type Level int8

const (
	PML4T Level = 4
	PDPT Level = 3
	PDT Level = 2
	PT Level = 1
)

type PageEntry struct{
	Address uintptr
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
	return PageEntry{Address: uintptr(entry) & 0xFFFFFFFFFFFFF000, 
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
	return uintptr(entry) & 0xFFFFFFFFFFFFF000
}

func(entry PageEntryPacked)HasProp(prop PageEntryPacked)bool{
	return (entry & prop)!=0
}

func(entry PageEntryPacked)NextLevel()*Page{
	if !entry.HasProp(PRESENT) || entry.HasProp(LARGE){
		return nil
	}
	return page(uintptr(entry))
}

var(
	stackBase uintptr
	stackLength
)

func page(p uintptr)*Page{
	return (*Page)(ptr.GetAddr(p & 0xFFFFFFFFFFFFF000))
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
	page(dirPtrTable)[0] = PageEntry{Address: uint64(bootstrapPage), ReadWrite: true, Present:true}.Pack()
	page(dirPtrTable)[3] = PageEntry{Address: uint64(kernelPage), ReadWrite: true, Present:true}.Pack()
	page(kernelPage)[0] = PageEntry{Address: 0x200000, Large: true, ReadWrite: true, Present:true}.Pack()
	page(bootstrapPage)[0] = PageEntry{Address: 0, Large: true, Present:true, ReadWrite:true}.Pack()
	
	for i:=1; i<512; i++{
		page(Mapl4)[i] = 0
		if i!=3{
			page(dirPtrTable)[i]  = 0
		}
		page(kernelPage)[i]  = 0
		page(bootstrapPage)[i] = 0
	}
	
	enable(dirPtrTable)
}
