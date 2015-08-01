package elf

import (
	"unsafe"
	"types"
)

func IsElf(program *types.Array)bool{
	if program[0] != 0x7f || program[1] != 'E' || program[2] != 'L' || program[3] != 'F' {
		return false
	}
	return program[4] == 2 && program[6] == 1 && (program[7]  == 0 || program[7] == 0xF) && program[18] == 0x3E && program[20] == 1
}

func Parse(program uintptr)uintptr{
	ident := (*types.Array)(unsafe.Pointer(program))
	if !IsElf(ident){
		return 0
	}
	//x64 := *(*uint8)(unsafe.Pointer(program+0x4)) == 2
	//bigEndian := *(*uint8)(unsafe.Pointer(program+0x5)) == 2
	//elfType := *(*uint16)(unsafe.Pointer(program+0x10))
	entryPoint := *(*uintptr)(unsafe.Pointer(&ident[0x18]))
	pHdrOffset := *(*uintptr)(unsafe.Pointer(&ident[0x20]))
	pEntSize := *(*uint16)(unsafe.Pointer(&ident[0x36]))
	pEntNum := *(*uint16)(unsafe.Pointer(&ident[0x38]))
	//progHeader := (*types.Array)(unsafe.Pointer(uintptr(pHdrOffset) + program))
	
	for i:=pHdrOffset + program; i < pHdrOffset + program + uintptr(pEntSize*pEntNum); i+=uintptr(pEntSize){
		progHeader := (*[7]uintptr)(unsafe.Pointer(i))
		entType := (*[2]uint32)(unsafe.Pointer(i))[0]
		//entFlags := (*[2]uint32)(unsafe.Pointer(i))[1]
		if entType == 1 {
			dest := progHeader[2]
			src := program + progHeader[1]
			copy(src, dest, progHeader[4])
			zero(dest+progHeader[4], progHeader[5]-progHeader[4])
		}
	}
	return entryPoint
}

func copy(src, dest, size uintptr) {
	byteNum := uintptr(0)
	if size >= 8{
		for ; byteNum <= uintptr(size-8); byteNum += 8 {
			*(*uint64)(unsafe.Pointer(dest+byteNum)) = *(*uint64)(unsafe.Pointer(src+byteNum))
		}
	}
	for ; byteNum < size; byteNum++{
		*(*uint8)(unsafe.Pointer(dest+byteNum)) = *(*uint8)(unsafe.Pointer(src+byteNum))
	}
}

func zero(addr, size uintptr){
	byteNum := uintptr(0)
	if size >= 8{
		for ; byteNum <= uintptr(size-8); byteNum += 8 {
			*(*uint64)(unsafe.Pointer(addr+byteNum)) = 0
		}
	}
	for ; byteNum < size; byteNum++{
		*(*uint8)(unsafe.Pointer(addr+byteNum)) = 0
	}
}