package elf

import (
	"unsafe"
	//"runtime"
	"page"
)

type Array [1<<30]uint8

type Program struct{
	Header
}

func (p *Program)Func()func(){
	dummy := &p.Entry
	return *(*func())(unsafe.Pointer(&dummy))
}

func (p *Program)CopyToMem(){
	headers := (*[1<<30]ProgramHeader)(unsafe.Pointer(p.Phoff + uintptr(unsafe.Pointer(p))))
	for i:=uint16(0); i<p.Phnum; i++ {
	//for _, progHeader := range p.ProgHeaders(){
		headers[i].CopyToMem(p)
	}
}

type Header struct{
	Ident
	Type uint16
	Machine uint16
	Version uint32
	Entry uintptr
	Phoff uintptr
	Shoff uintptr
	Flags uint32
	Ehsize, Phentsize, Phnum, Shentsize, Shnum, Shstrndx uint16
}

func (h *Header)ProgHeaders()[]ProgramHeader{
	return (*[1<<30]ProgramHeader)(unsafe.Pointer(h.Phoff + uintptr(unsafe.Pointer(h))))[:h.Phnum:h.Phnum]
}

type Ident struct{
	Magic uint32
	Class uint8
	Data uint8
	Version uint8
	ABI uint8
	AVIVersion uint8
	_ uint32
}

type ProgramHeader struct{
	Type uint32
	Flags uint32
	Offset uintptr
	Vaddr *Array
	Paddr uintptr
	Filesz uintptr
	Memsz uintptr
	Align uint
}

func (p *ProgramHeader)Prog(parent *Program)[]uint8{
	return (*[1<<30]uint8)(unsafe.Pointer(p.Offset + uintptr(unsafe.Pointer(parent))))[:p.Filesz:p.Filesz]
}

//extern __break
func breakPoint()

func (p *ProgramHeader)CopyToMem(parent *Program){
	//breakPoint()
	if p.Type == 1 {
			pageProps := page.PRESENT | page.USER

			//if p.Flags & 2 != 0{
				pageProps |= page.READ_WRITE
			//}
			for i:=uintptr(0); i < (p.Memsz + 0x1000) &^ 0xFFF; i+=0x1000{
				page.NewPage(uintptr(unsafe.Pointer(p.Vaddr))+i, page.K, pageProps)
			}
			loc := p.Vaddr[:p.Memsz:p.Memsz]
			//breakPoint()
			copy(loc[:p.Filesz], p.Prog(parent))
			//breakPoint()
			zero(loc[p.Filesz:], p.Memsz-p.Filesz)
		}
}

func (program *Program)IsElf()bool{
	if program.Ident.Magic != 0x464C457f {// "\x7fELF"
		return false
	}
	return program.Class == 2 && program.Data == 1 && program.Ident.Version == 1 && (program.ABI  == 0 || program.ABI == 0xF) && program.Machine == 0x3E && program.Version == 1
}

func Parse(program uintptr)(prog *Program){
	prog = (*Program)(unsafe.Pointer(program))
	if !prog.IsElf(){
		return nil
	}
	return
}

func Copy(src, dest *Array, count int)*Array{
	for i := 0; i<count; i++{
		dest[i] = src[i]
	}
	return dest
}

func zero(addr []uint8, size uintptr){
	byteNum := uintptr(0)
	if size >= 8{
		for ; byteNum <= uintptr(size-8); byteNum += 8 {
			*(*uint64)(unsafe.Pointer(&addr[byteNum])) = 0
		}
	}
	for ; byteNum < size; byteNum++{
		addr[byteNum] = 0
	}
}