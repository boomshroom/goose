package elf

import "unsafe"

const programOffset = 0xFFFF800000000000 - 0x40000000

//type array [1<<30]uint8

type slice struct {
	ptr              uintptr
	length, capacity uintptr
}

func IsElf(program []uint8) bool {
	if program[0] != '\x7f' || program[1] != 'E' || program[2] != 'L' || program[3] != 'F' {
		return false
	}
	return program[4] == 2 && program[6] == 1 && (program[7] == 0 || program[7] == 0xF) && program[18] == 0x3E && program[20] == 1
}

func Parse(program uintptr, size uintptr) uintptr {
	s := slice{program, size, size}
	ident := *(*[]uint8)(unsafe.Pointer(&s))
	if !IsElf(ident) {
		return 0
	}
	//x64 := *(*uint8)(unsafe.Pointer(program+0x4)) == 2
	//bigEndian := *(*uint8)(unsafe.Pointer(program+0x5)) == 2
	//elfType := *(*uint16)(unsafe.Pointer(program+0x10))
	entryPoint := *(*uint64)(unsafe.Pointer(&ident[0x18]))
	pHdrOffset := *(*uint64)(unsafe.Pointer(&ident[0x20]))
	pEntSize := *(*uint16)(unsafe.Pointer(&ident[0x36]))
	pEntNum := *(*uint16)(unsafe.Pointer(&ident[0x38]))
	//progHeader := (*array)(unsafe.Pointer(uintptr(pHdrOffset) + program))

	for i := uintptr(pHdrOffset) + program; i < uintptr(pHdrOffset)+program+uintptr(pEntSize*pEntNum); i += uintptr(pEntSize) {
		progHeader := (*[7]uint64)(unsafe.Pointer(i))
		entType := (*[2]uint32)(unsafe.Pointer(i))[0]
		//entFlags := (*[2]uint32)(unsafe.Pointer(i))[1]
		if entType == 1 {
			dest := uintptr(progHeader[2] - programOffset) // Offset for kernel at 4G mark
			src := program + uintptr(progHeader[1])
			copy(src, dest, uintptr(progHeader[4]))
			zero(dest+uintptr(progHeader[4]), uintptr(progHeader[5]-progHeader[4]))
		}
	}
	return uintptr(entryPoint)
}

func copy(src, dest, size uintptr) {
	byteNum := uintptr(0)
	if size >= 8 {
		for ; byteNum <= uintptr(size-8); byteNum += 8 {
			*(*uint64)(unsafe.Pointer(dest + byteNum)) = *(*uint64)(unsafe.Pointer(src + byteNum))
		}
	}
	for ; byteNum < size; byteNum++ {
		*(*uint8)(unsafe.Pointer(dest + byteNum)) = *(*uint8)(unsafe.Pointer(src + byteNum))
	}
}

func zero(addr, size uintptr) {
	byteNum := uintptr(0)
	if size >= 8 {
		for ; byteNum <= uintptr(size-8); byteNum += 8 {
			*(*uint64)(unsafe.Pointer(addr + byteNum)) = 0
		}
	}
	for ; byteNum < size; byteNum++ {
		*(*uint8)(unsafe.Pointer(addr + byteNum)) = 0
	}
}
