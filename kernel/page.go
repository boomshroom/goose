package page

import (
	"unsafe"
	//"tables"
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

var nextPage uintptr = 0x400000

func NewPage(address uintptr, size PageSize, props PageEntryPacked) (physical PageEntryPacked) {
	if len(freeStack) != 0 && size == K {
		physical = MapAddress(address, freeStack[len(freeStack)-1].Address(), K, props)
		freeStack = freeStack[:len(freeStack)-1]
		return
	}
	physical = MapAddress(address, nextPage, size, props)
	switch size {
	case G:
		nextPage = (nextPage&^0x40000000 - 1) + 0x40000000
	case M:
		nextPage = (nextPage&^0x200000 - 1) + 0x200000
	default:
		nextPage = nextPage&^0xFFF + 0x1000
	}
	return
}

func (p *PageEntryPacked) Enable(logical uintptr, size PageSize) {
	*p = MapAddress(logical, p.Address(), size, (*p)&0xFFF)
}

func MapAddress(logical, physical uintptr, size PageSize, props PageEntryPacked) PageEntryPacked {

	ml4Entry := &pml4[(logical>>39)&0x1FF]
	if *ml4Entry&PRESENT == 0 {
		*ml4Entry = PageEntry{Address: nextPhysAddr(), Present: true}.Pack()
		ml4Entry.SetProp(props, true)
	}

	dptEntry := &ml4Entry.NextLevel()[(logical>>30)&0x1FF]
	if size == G {
		*dptEntry = PageEntry{Address: (*Page)(unsafe.Pointer(physical)), Present: true, Large: true}.Pack()
		dptEntry.SetProp(props, true)
		return *dptEntry
	} else if *dptEntry&(PRESENT|LARGE) != PRESENT {
		*dptEntry = PageEntry{Address: nextPhysAddr(), Present: true}.Pack()
		for i := range dptEntry.NextLevel() {
			dptEntry.NextLevel()[i] = 0
		}
	}

	pdtEntry := &dptEntry.NextLevel()[(logical>>21)&0x1FF]
	if size == M {
		*pdtEntry = PageEntry{Address: (*Page)(unsafe.Pointer(physical)), Present: true, Large: true}.Pack()
		pdtEntry.SetProp(props, true)
		return *pdtEntry
	} else if *pdtEntry&(PRESENT|LARGE) != PRESENT {
		addr := nextPhysAddr()
		*pdtEntry = PageEntry{Address: addr, Present: true}.Pack()
	}

	if props&USER != 0 {
		ml4Entry.SetProp(USER, true)
		dptEntry.SetProp(USER, true)
		pdtEntry.SetProp(USER, true)
		if props&READ_WRITE != 0 {
			ml4Entry.SetProp(READ_WRITE, true)
			dptEntry.SetProp(READ_WRITE, true)
			pdtEntry.SetProp(READ_WRITE, true)
		}
	}

	pageEntry := &pdtEntry.NextLevel()[(logical>>12)&0x1FF]
	*pageEntry = PageEntry{Address: (*Page)(unsafe.Pointer(physical)), Present: true}.Pack()
	pageEntry.SetProp(props, true)
	return *pageEntry
}

func nextPhysAddr() *Page {
	l := len(stack)
	stack = stack[:l+1]
	stack[l] = Page{}
	return (*Page)(unsafe.Pointer(pml4.PhysAddr(uintptr(unsafe.Pointer(&stack[l])))))
}

type PageEntryPacked uintptr

type Page [512]PageEntryPacked

func (p *Page) GetEntry(addr uintptr, level Level) *PageEntryPacked {

	if level < PDT {
		return nil
	}
	entry := &p[addr>>39]
	if !entry.HasProp(PRESENT) {
		return nil
	} else if entry.HasProp(LARGE) {
		return entry
	} else {
		return entry.NextLevel().GetEntry(addr<<9, level-1)
	}
}

func (p *Page) PhysAddr(logical uintptr) uintptr {
	ml4Entry := &pml4[(logical>>39)&0x1FF]
	if *ml4Entry&PRESENT == 0 {
		return 0
	}

	dptEntry := &ml4Entry.NextLevel()[(logical>>30)&0x1FF]
	if *dptEntry&PRESENT == 0 {
		return 0
	} else if *dptEntry&LARGE != 0 {
		return dptEntry.Address() | (logical & (1<<30 - 1))
	}

	pdtEntry := &dptEntry.NextLevel()[(logical>>21)&0x1FF]
	if *pdtEntry&PRESENT == 0 {
		return 0
	} else if *pdtEntry&LARGE != 0 {
		return pdtEntry.Address() | (logical & (1<<21 - 1))
	}

	pageEntry := &pdtEntry.NextLevel()[(logical>>12)&0x1FF]
	if *pageEntry&PRESENT == 0 {
		return 0
	}
	return pdtEntry.Address() | (logical & (1<<12 - 1))
}

type PageSize uint8

const (
	K PageSize = iota // 4 KiB pages
	M                 // 2 MiB pages
	G                 // 1 GiB pages
)

type Level int8

const (
	PML4T Level = 4
	PDPT  Level = 3
	PDT   Level = 2
	PT    Level = 1
)

type PageEntry struct {
	Address                                                                              *Page
	Global, Large, Dirty, Accessed, CacheDisable, WriteThrough, User, ReadWrite, Present bool
}

const (
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

func (entry PageEntry) Pack() PageEntryPacked {
	e := PageEntryPacked(unsafe.Pointer(entry.Address)) &^ 0xFFF
	if entry.Global {
		e |= GLOBAL
	}
	if entry.Large {
		e |= LARGE
	}
	if entry.Dirty {
		e |= DIRTY
	}
	if entry.Accessed {
		e |= ACCESSED
	}
	if entry.CacheDisable {
		e |= CACHE_DISABLED
	}
	if entry.WriteThrough {
		e |= WRITE_THROUGH
	}
	if entry.User {
		e |= USER
	}
	if entry.ReadWrite {
		e |= READ_WRITE
	}
	if entry.Present {
		e |= PRESENT
	}
	return e
}

func (entry PageEntryPacked) Unpack() PageEntry {
	return PageEntry{Address: (*Page)(unsafe.Pointer(entry &^ 0xFFF)),
		Global:       entry&GLOBAL != 0,
		Large:        entry&LARGE != 0,
		Dirty:        entry&DIRTY != 0,
		Accessed:     entry&ACCESSED != 0,
		CacheDisable: entry&CACHE_DISABLED != 0,
		WriteThrough: entry&WRITE_THROUGH != 0,
		User:         entry&USER != 0,
		ReadWrite:    entry&READ_WRITE != 0,
		Present:      entry&PRESENT != 0}
}

func (entry PageEntryPacked) Address() uintptr {
	return uintptr(entry) &^ 0xFFF
}

func (entry PageEntryPacked) HasProp(prop PageEntryPacked) bool {
	return (entry & prop) != 0
}

func (entry *PageEntryPacked) SetProp(prop PageEntryPacked, set bool) {
	if set {
		*entry = *entry | (prop & 0xFFF)
	} else {
		*entry = *entry & ^(prop & 0xFFF)
	}
}

func (entry PageEntryPacked) NextLevel() *Page {
	if entry&(PRESENT|LARGE) != PRESENT {
		return nil
	}
	e := entry &^ 0xFFF
	if e >= 0x200000 {
		return (*Page)(unsafe.Pointer(e + (0xFFFF800000000000 - 0x200000)))
	}
	return (*Page)(unsafe.Pointer(e))
}

func (entry *PageEntryPacked) Free() {
	entry.SetProp(PRESENT, false)
	freeStack = freeStack[:len(freeStack)+1]
	freeStack[len(freeStack)-1] = *entry
}

func SetPageLoc(page *Page) {
	pml4 = page
}

var (
	stack     []Page
	freeStack []PageEntryPacked
	pml4      *Page
)

//extern __kernel_end
func kernelEnd() uintptr

func init() {
	stackBegin := (kernelEnd() &^ 0xFFF) + 0x1000
	stack = (*[1 << 30]Page)(unsafe.Pointer(stackBegin))[:0 : ((0xFFFF800000200000-stackBegin)>>12)-1]
	freeStack = (*[512]PageEntryPacked)(unsafe.Pointer(uintptr(0xFFFF8000001FF000)))[:0]

	pml4.GetEntry(0x40000000, PDPT).SetProp(PRESENT, false)
	//for i := range stack{
	//	stack[i] = Page{}
	//}
	//tables.MultibootTable = (*tables.MBTable)(unsafe.Pointer(uintptr(0xFFFFFFFFFFFFF000)))
}
