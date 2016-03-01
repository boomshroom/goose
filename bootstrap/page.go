package page

import (
	"ptr"
	"unsafe"
)

type PageEntryPacked uint64

type Page [512]PageEntryPacked

type PageEntry struct {
	Address                                                                              uint64
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
	e := PageEntryPacked(entry.Address & 0xFFFFFFFFFFFFF000)
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
	return PageEntry{Address: uint64(entry) & 0xFFFFFFFFFFFFF000,
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
	return uintptr(entry) & 0xFFFFF000
}

var (
	Pages *[5]Page = getPages()
)

func page(p uintptr) *Page {
	return (*Page)(ptr.GetAddr(p & 0xFFFFF000))
}

//extern __get_pages
func getPages() *[5]Page

//extern __enable_paging
func enable(*Page)

func init() {
	mapl4 := &Pages[0]
	dirPtrTable := &Pages[1]
	bootstrapPage := &Pages[2]
	kernelPage := &Pages[3]
	kernelPtrTableHigh := &Pages[4]

	mapl4[0] = PageEntry{Address: uint64(uintptr(unsafe.Pointer(dirPtrTable))), Present: true}.Pack()
	mapl4[256] = PageEntry{Address: uint64(uintptr(unsafe.Pointer(kernelPtrTableHigh))), Present: true}.Pack()
	dirPtrTable[0] = PageEntry{Address: uint64(uintptr(unsafe.Pointer(bootstrapPage))), Present: true}.Pack()
	dirPtrTable[1] = PageEntry{Address: uint64(uintptr(unsafe.Pointer(kernelPage))), Present: true}.Pack()
	kernelPage[0] = PageEntry{Address: 0x200000, Large: true, ReadWrite: true, Present: true}.Pack()
	kernelPtrTableHigh[0] = PageEntry{Address: uint64(uintptr(unsafe.Pointer(kernelPage))), Present: true, Global: true}.Pack()
	bootstrapPage[0] = PageEntry{Address: 0, Large: true, Present: true}.Pack()

	for i := 1; i < 512; i++ {
		if i != 256 {
			mapl4[i] = 0
		}
		if i != 1 {
			dirPtrTable[i] = 0
		}
		kernelPtrTableHigh[i] = 0
		kernelPage[i] = 0
		bootstrapPage[i] = 0
	}

	enable(dirPtrTable)
}
