package mmap

import (
	"runtime"
	"unsafe"
)
var MMap MemoryMap

type MemoryMap []MemorySegment

type MemorySegment struct {
	size                      uint32
	baseAddrLow, baseAddrHigh uint32
	lengthLow, lengthHigh     uint32
	memType                   uint32
}

func (m *MemorySegment) Base() uintptr {
	return uintptr(m.baseAddrLow) | uintptr(m.baseAddrHigh)<<32
}

func (m *MemorySegment) Length() uintptr {
	return uintptr(m.lengthLow) | uintptr(m.lengthHigh)<<32
}

func (m *MemorySegment) End() uintptr {
	return m.Base() + m.Length()
}

func (m *MemorySegment) Accessable() bool {
	return m.memType == 1
}

func (m *MemorySegment) Block() []uint8 {
	l := m.Length()
	return (*runtime.Array)(unsafe.Pointer(m.Base()))[:l:l]
}

func (m *MemorySegment) MemBlock() MemBlock {
	return MemBlock{Start: m.Base(), End: m.End()}
}

type MemBlock struct {
	Start, End uintptr
}