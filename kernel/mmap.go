package mmap

var MMap MemoryMap

type MemoryMap interface{
	Length()int
	Get(int)MemorySegment
}

type MemorySegment interface{
	Base() uintptr
	Length() uint
	Pages() uint
	End() uintptr
	Available() bool
}