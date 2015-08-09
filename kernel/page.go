package page

import (
	"unsafe"
)

/*
// Helper functions for finding locations of virtual memory
// Copy to playground to determine values
func indecies(addr uint64) [5]uint64 {
	return [...]uint64{
		4: addr >> 39 & 0x1FF,
		3: (addr >> 30) & 0x1FF,
		2: (addr >> 21) & 0x1FF,
		1: (addr >> 12) & 0x1FF,
		0: addr & 0x1000,
	}
}

func invIndecies(PT, PD, PDP, PML4 uint64) uint64 {
	return (PT << 12) | (PD << 21) | (PDP << 30) | (PML4 << 39)
}
*/

func NewPage(address uintptr, size PageSize, props PageEntryPacked){
	MapAddress(address, 0x400000, size, props)
}

func MapAddress(logical, physical uintptr, size PageSize, props PageEntryPacked){
	ml4Entry := &pml4[(logical >> 39) & 0x1FF]
	if *ml4Entry & PRESENT == 0{
		*ml4Entry = PageEntry{Address: nextPhysAddr(), Present: true}.Pack()
	}

	dptEntry := &ml4Entry.NextLevel()[(logical >> 30) & 0x1FF]
	if size == G {
		*dptEntry = PageEntry{Address: (*Page)(unsafe.Pointer(physical)), Present: true, Large: true}.Pack()
		dptEntry.SetProp(props, true)
		return
	}else if *dptEntry & (PRESENT|LARGE) != PRESENT{
		*dptEntry = PageEntry{Address: nextPhysAddr(), Present: true}.Pack()
	}

	pdtEntry := &dptEntry.NextLevel()[(logical >> 21) & 0x1FF]
	if size == M {
		*pdtEntry = PageEntry{Address: (*Page)(unsafe.Pointer(physical)), Present: true, Large: true}.Pack()
		pdtEntry.SetProp(props, true)
		return
	}else if *pdtEntry & (PRESENT|LARGE) != PRESENT{
		*pdtEntry = PageEntry{Address: nextPhysAddr(), Present: true}.Pack()
	}

	pageEntry := &pdtEntry.NextLevel()[(logical >> 12) & 0x1FF]
	*pageEntry = PageEntry{Address: (*Page)(unsafe.Pointer(physical)), Present: true}.Pack()
	pageEntry.SetProp(props, true)
}

func nextPhysAddr()*Page{
	l:=len(stack)
	stack = stack[:l+1]
	return (*Page)(unsafe.Pointer(pml4.PhysAddr(uintptr(unsafe.Pointer(&stack[l])))))
	//return (*Page)(unsafe.Pointer(&stack[l]))
}

type PageEntryPacked uintptr

type Page [512]PageEntryPacked

func (p *Page)GetEntry(addr uintptr, level Level)*PageEntryPacked{

	if level < PDT {
		return nil
	}
	entry := &p[addr >> 39]
	if !entry.HasProp(PRESENT){
		return nil
	}else if entry.HasProp(LARGE){
		return entry
	}else{
		return entry.NextLevel().GetEntry(addr << 9, level-1)
	}
}

func (p *Page)PhysAddr(logical uintptr)uintptr{
	ml4Entry := &pml4[(logical >> 39) & 0x1FF]
	if *ml4Entry & PRESENT == 0{
		return 0
	}

	dptEntry := &ml4Entry.NextLevel()[(logical >> 30) & 0x1FF]
	if *dptEntry & PRESENT == 0{
		return 0
	}else if *dptEntry & LARGE != 0{
		return dptEntry.Address() | (logical & (1<<30 - 1))
	}

	pdtEntry := &dptEntry.NextLevel()[(logical >> 21) & 0x1FF]
	if *pdtEntry & PRESENT == 0{
		return 0
	}else if *pdtEntry & LARGE != 0{
		return pdtEntry.Address() | (logical & (1<<21 - 1))
	}

	pageEntry := &pdtEntry.NextLevel()[(logical >> 12) & 0x1FF]
	if *pageEntry & PRESENT == 0{
		return 0
	}
	return pdtEntry.Address() | (logical & (1<<12 - 1))
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
	Address *Page
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
	e := PageEntryPacked(unsafe.Pointer(entry.Address)) &^ 0xFFF
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
	return PageEntry{Address: (*Page)(unsafe.Pointer(entry &^ 0xFFF)), 
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
	return uintptr(entry) &^ 0xFFF
}

func(entry PageEntryPacked)HasProp(prop PageEntryPacked)bool{
	return (entry & prop)!=0
}

func(entry *PageEntryPacked)SetProp(prop PageEntryPacked, set bool){
	if set{
		*entry = *entry | (prop & 0xFFF)
	}else{
		*entry = *entry & ^(prop & 0xFFF)
	}
}

func(entry PageEntryPacked)NextLevel()*Page{
	if entry & (PRESENT|LARGE) != PRESENT{
		return nil
	}
	e := entry &^ 0xFFF
	if e >= 0x200000{
		//breakPoint()
		return (*Page)(unsafe.Pointer(e + (0xFFFF800000000000 - 0x200000)))
	}
	return (*Page)(unsafe.Pointer(e))
}

func SetPageLoc(page *Page){
	pml4 = page
}

var(
	stack []Page
	pml4 *Page
)

//extern __kernel_end
func kernelEnd()uintptr

func init(){
	stackBegin := (kernelEnd() &^ 0xFFF) + 0x1000
	stack = (*[1<<30]Page)(unsafe.Pointer(stackBegin))[:0:(0xFFFF800000200000 - stackBegin)>>12]
	//video.Println("Page initialized")
	/*Mapl4 = (kernelEnd() & 0xFFFFF000) + 0x1000
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
	
	enable(dirPtrTable)*/
}
