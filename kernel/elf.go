package elf

import (
	"page"
	"proc"
	"runtime"
	"unsafe"
	"video"
)

var KernelElf *Program

type Program struct {
	Header
}

func (p *Program) Func() func() {
	dummy := &p.Entry
	return *(*func())(unsafe.Pointer(&dummy))
}

func (p *Program) CopyToMem() {
	if proc.NumProcs >= 4 {
		video.Error("Too many loaded processes", int(proc.NumProcs), true)
	}
	proc.NumProcs++
	newID := proc.NumProcs

	proc.Procs[newID] = proc.Proc{
		Id:          newID,
		RIP:         uint64(p.Entry),
		RSP:         0x7FFFFFFFEFF0,
		StackPage:   page.NewPage(0x7FFFFFFFE000, page.K, page.PRESENT|page.USER|page.READ_WRITE),
		SyscallPage: page.NewPage(0x7FFFFFFFF000, page.K, page.PRESENT|page.USER|page.READ_WRITE),
		ElfHeader:   unsafe.Pointer(p),
	}
	
	headers := (*[1 << 30]ProgramHeader)(unsafe.Pointer(p.Phoff + uintptr(unsafe.Pointer(p))))
	for i := uint16(0); i < p.Phnum; i++ {
		headers[i].CopyToMem(p, newID)
	}
}

type Header struct {
	Ident
	Type                                                 uint16
	Machine                                              uint16
	Version                                              uint32
	Entry                                                uintptr
	Phoff                                                uintptr
	Shoff                                                uintptr
	Flags                                                uint32
	Ehsize, Phentsize, Phnum, Shentsize, Shnum, Shstrndx uint16
}

func (h *Header) ProgHeaders() []ProgramHeader {
	return (*[1 << 30]ProgramHeader)(unsafe.Pointer(h.Phoff + uintptr(unsafe.Pointer(h))))[:h.Phnum:h.Phnum]
}

func (h *Header) SectHeaders() []SectionHeader {
	return (*[1 << 30]SectionHeader)(unsafe.Pointer(h.Shoff + uintptr(unsafe.Pointer(h))))[:h.Shnum:h.Shnum]
}

type Ident struct {
	Magic      uint32
	Class      uint8
	Data       uint8
	Version    uint8
	ABI        uint8
	AVIVersion uint8
	_          uint32
}

type ProgramHeader struct {
	Type   uint32
	Flags  uint32
	Offset uintptr
	Vaddr  *runtime.Array
	Paddr  uintptr
	Filesz uintptr
	Memsz  uintptr
	Align  uint
}

func (p *ProgramHeader) Prog(parent *Program) []uint8 {
	return (*[1 << 30]uint8)(unsafe.Pointer(p.Offset + uintptr(unsafe.Pointer(parent))))[:p.Filesz:p.Filesz]
}

func (p *ProgramHeader) CopyToMem(parent *Program, procID uint64) {
	if p.Type == 1 {
		pageProps := page.PRESENT | page.USER

		if p.Flags&2 != 0 {
			pageProps |= page.READ_WRITE
		}
		process := &proc.Procs[procID]
		for i := uintptr(0); i < (p.Memsz+0x1000)&^0xFFF; i += 0x1000 {
			vAddr := uintptr(unsafe.Pointer(p.Vaddr))+i
			process.Pages[process.NumPages] = proc.PageMapping{
				Physical: page.NewPage(vAddr, page.K, pageProps),
				Virtual: vAddr,
			}
			process.NumPages++
			if process.NumPages > 10 {
				video.Error("Too many pages alocated by process", process.NumPages, true)
			}
		}

		loc := p.Vaddr[:p.Memsz:p.Memsz]
		print("") // Fails to copy 2nd process if removed
		copy(loc[:p.Filesz], p.Prog(parent))
		zero(loc[p.Filesz:], p.Memsz-p.Filesz)
	}
}

type SectionType uint32

type SectionHeader struct {
	Name uint32
	Type SectionType
	Flags uint64
	Addr uint64
	Offset uintptr
	Size uintptr
	Link uint32
	Info uint32
	AddrAlign uint64
	EntSize uint64
}

const(
	Null SectionType = iota
	ProgBits
	SymTable
	StrTable
	// others not nessisary at the moment
)

func PrintAddress(addr uintptr){
	video.Print(KernelElf.LookupSymbol(addr))
	video.PutChar(' ')
	println(addr)
}

func (h *Header) SymSect()[]Symbol{
	for _, sect := range h.SectHeaders(){
		if sect.Type == SymTable {
			return (*[1 << 30]Symbol)(unsafe.Pointer(sect.Offset + uintptr(unsafe.Pointer(h))))[:uint64(sect.Size)/sect.EntSize:uint64(sect.Size)/sect.EntSize]
		}
	}
	return nil
}

func (h *Header) LookupSymbol(addr uintptr)string{
	var s *Symbol
	l := len(h.SymSect())
	for i:=0; i< l; i++{
		sym := &h.SymSect()[i]
		if sym.Value > s.Value && sym.Value < addr {
			s = sym
		}
	}
	if s == nil {
		return ""
	}

	for i, sect := range h.SectHeaders(){
		if sect.Type == StrTable && i != int(h.Shstrndx) {
			strTable := unsafe.Pointer(sect.Offset + uintptr(unsafe.Pointer(h)))
			strLen := 0
			for ; (*runtime.Array)(strTable)[strLen+int(s.Name)]!=0; strLen++{}
			ret := runtime.String{Ptr: (*runtime.Array)(unsafe.Pointer(uintptr(strTable)+uintptr(s.Name))), Len: strLen}
			return *(*string)(unsafe.Pointer(&ret))
		}
	}
	return ""
}

type Symbol struct{
	Name uint32
	Info uint8
	Other uint8
	Index uint16
	Value uintptr
	Size uint64
}

func (program *Program) IsElf() bool {
	if program.Ident.Magic != 0x464C457f { // "\x7fELF"
		return false
	}
	return program.Class == 2 && program.Data == 1 && program.Ident.Version == 1 && (program.ABI == 0 || program.ABI == 0xF) && program.Machine == 0x3E && program.Version == 1
}

func Parse(program *uint8) (prog *Program) {
	prog = (*Program)(unsafe.Pointer(program))
	if !prog.IsElf() {
		return nil
	}
	return
}

func zero(addr []uint8, size uintptr) {
	if len(addr)!=int(size) {
		video.Println("Provided size different than slice len")
		println(len(addr))
		println(size)
	}
	byteNum := uintptr(0)
	if size >= 8 {
		for ; byteNum <= uintptr(size-8); byteNum += 8 {
			if byteNum >= uintptr(len(addr)){
				println(byteNum)
				println(len(addr))
			}
			*(*uint64)(unsafe.Pointer(&addr[byteNum])) = 0
		}
	}
	for ; byteNum < size; byteNum++ {
		addr[byteNum] = 0
	}
}
