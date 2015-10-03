package tables

import (
	"video"
	"unsafe"
	"runtime"
)

var MultibootTable *MBTable

func SetTable(t *MBTable){
	if t == nil{
		video.Error("Multiboot not valid", -1, true)
	}
	MultibootTable = t
}

type MBTable struct{
	Flags Flags
	MemLower, MemUpper uint32
	BootDevice BootDevice
	cmd uint32
	ModsCount, modsAddr uint32
	syms [4]uint32
	mmapLength, mmapAddr uint32
	drivesLength, drivesAddr uint32
	configTable uint32
	bootloaderName uint32
	apmTable uint32

}

func (t *MBTable)Command()string{
	if t.Flags & CmdLine == 0{
		return ""
	}
	return runtime.GoString((*uint8)(unsafe.Pointer(uintptr(t.cmd))))
}

func (t *MBTable)Mods()[]Mod{
	if t.Flags & Mods == 0{
		return nil
	}
	array := (*[1<<30]Mod)(unsafe.Pointer(uintptr(t.modsAddr)))
	return array[:t.ModsCount:t.ModsCount]
}

func (t *MBTable)MMap()MemoryMap{
	if t.Flags & MMap == 0{
		return nil
	}
	array := (*[1<<30]MemorySegment)(unsafe.Pointer(uintptr(t.mmapAddr)))
	l := t.mmapLength/uint32(unsafe.Sizeof(MemorySegment))
	return MemoryMap(array[:l:l])
}

func (t *MBTable)APMTable()*APM{
	if t.Flags & APMTable == 0{
		return nil
	}
	return (*APM)(unsafe.Pointer(uintptr(t.apmTable)))
}

type Flags uint32

const (
	Mem Flags = 1 << iota
	BootDev
	CmdLine
	Mods
	Aout
	ELF
	MMap
	Drives
	CfgTable
	BootLoaderName
	APMTable
	GraphicsTable
)

type BootDevice struct{
	Drive uint8
	Parts [3]uint8
}

type Mod struct{
	start, end uint32
	name uint32
	_ uint32
}

func (m *Mod)Bytes()[]uint8{
	array := (*runtime.Array)(unsafe.Pointer(uintptr(m.start)))
	l := m.end - m.start
	return array[:l:l]
}

func (m *Mod)Name()string{
	return runtime.GoString((*uint8)(unsafe.Pointer(uintptr(m.name))))
}
/*
type AoutSyms struct{
	TabSize
}*/

type MemoryMap []MemorySegment

type MemorySegment struct{
	size uint32
	baseAddrLow, baseAddrHigh uint32
	lengthLow, lengthHigh uint32
	memType uint32
}

func (m *MemorySegment)Base()uintptr{
	return uintptr(m.baseAddrLow) | uintptr(m.baseAddrHigh)<<32
}

func (m *MemorySegment)Length()uintptr{
	return uintptr(m.lengthLow) | uintptr(m.lengthHigh)<<32
}

func (m *MemorySegment)End()uintptr{
	return m.Base() + m.Length()
}

func (m *MemorySegment)Accessable()bool{
	return m.memType == 1
}

func (m *MemorySegment)Block()[]uint8{
	l:=m.Length()
	return (*runtime.Array)(unsafe.Pointer(m.Base()))[:l:l]
}

func (m *MemorySegment)MemBlock()MemBlock{
	return MemBlock{Start: m.Base(), End: m.End()}
}

type MemBlock struct{
	Start, End uintptr
}

type APM struct{
	Version uint16
	cseg uint16
	offset uint32
	cseg16, dseg uint16
	flags uint16
	csegLen, cseg16Len, dsegLen uint16
}