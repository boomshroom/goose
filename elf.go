package elf

import "unsafe"

func IsElf(program uintptr)bool{
	magic :=*(*uint32)(unsafe.Pointer(program)) == 0x464c457f // 7f + ELF
	version := *(*uint8)(unsafe.Pointer(program+0x6)) == 1
	abi := *(*uint8)(unsafe.Pointer(program+0x7))
	abiCheck := abi == 0 || abi == 0xF //0xF unused abi
	arch := *(*uint16)(unsafe.Pointer(program+0x12)) == 0x3 //x86
	version2 := *(*uint32)(unsafe.Pointer(program+0x14)) == 1
	return magic && version && abiCheck && arch && version2
}

func ParseElf(program uintptr){
	if !IsElf(program){
		return
	}
	x64 := *(*uint8)(unsafe.Pointer(program+0x4)) == 2
	bigEndian := *(*uint8)(unsafe.Pointer(program+0x5)) == 2
	elfType := *(*uint16)(unsafe.Pointer(program+0x10))
	entryPoint := *(*uint32)(unsafe.Pointer(program+0x18))
}