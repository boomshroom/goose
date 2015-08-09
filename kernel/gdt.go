package gdt

import (
	"unsafe"
	//"video"
	"segment"
)

const size uint16 = 7
var Table [size]segment.Seg64 = [size]segment.Seg64{
	0: 0,
	1: segment.CodeDataDesc{Code: true}.Pack(),
	2: segment.CodeDataDesc{Code: false}.Pack(),
	3: segment.CodeDataDesc{Code: true, User: true}.Pack(),
	4: segment.CodeDataDesc{Code: false, User: true}.Pack(),
}

var tss segment.TSSPacked = segment.TSSPacked{} //segment.TSS{RSP: [3]uint64{0: stack()}}.Pack()

var gdtr segment.TablePtrPacked = segment.TablePtr{Size: unsafe.Sizeof(Table), Ptr: uintptr(unsafe.Pointer(&Table))}.Pack()

func loadGDTR(){
	gdtr[0] = uint16(unsafe.Sizeof(Table))
	gdtr[1] = uint16(uintptr(unsafe.Pointer(&Table)))
	gdtr[2] = uint16(uintptr(unsafe.Pointer(&Table)) >> 16)
	gdtr[3] = uint16(uintptr(unsafe.Pointer(&Table)) >> 32)
	gdtr[4] = uint16(uintptr(unsafe.Pointer(&Table)) >> 48)
}

var err = "GDT entry too large"

func init(){

	Table[5], Table[6] = segment.SystemDesc{Base:uint64(uintptr(unsafe.Pointer(&tss[0]))), Limit:uint32(unsafe.Sizeof(tss))-1, Type: segment.TSSAvail}.Pack().Decompose()
	loadGDT(&gdtr)
	//reloadSegments()
}

//extern __stack_ptr
func stack()uint64

//extern __load_gdt
func loadGDT(*segment.TablePtrPacked)

//extern __reload_segments
func reloadSegments()