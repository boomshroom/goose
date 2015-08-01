package gdt

import (
	"unsafe"
	//"video"
	"segment"
)

const size uint16 = 8
var Table [size]segment.Seg64

var tss segment.TSSPacked

var gdtr segment.TablePtrPacked
func loadGDTR(){
	gdtr[0] = uint16(unsafe.Sizeof(Table))
	gdtr[1] = uint16(uintptr(unsafe.Pointer(&Table)))
	gdtr[2] = uint16(uintptr(unsafe.Pointer(&Table)) >> 16)
	gdtr[3] = uint16(uintptr(unsafe.Pointer(&Table)) >> 32)
	gdtr[4] = uint16(uintptr(unsafe.Pointer(&Table)) >> 48)
}

var err [40]byte

func init(){
	for i := 0; i < len("GDT entry too large"); i++ {
		err[i] = "GDT entry too large"[i]
	}

	loadTable()
	tss = segment.TSSPacked{} //segment.TSS{RSP: [3]uint64{0: stack()}}.Pack()
	gdtr = segment.TablePtr{Size: unsafe.Sizeof(Table), Ptr: uintptr(unsafe.Pointer(&Table))}.Pack()
	loadGDT(&gdtr)
	//reloadSegments()
}

//extern __stack_ptr
func stack()uint64

//extern __load_gdt
func loadGDT(*segment.TablePtrPacked)

//extern __reload_segments
func reloadSegments()

func loadTable(){
	Table[0] = segment.Seg64(0)
	Table[1] = segment.CodeDataDesc{Code: true}.Pack()
	Table[2] = segment.CodeDataDesc{Code: false}.Pack()
	Table[4] = segment.CodeDataDesc{Code: true, User: true}.Pack()
	Table[5] = segment.CodeDataDesc{Code: false, User: true}.Pack()
	Table[6], Table[7] = segment.SystemDesc{Base:uint64(uintptr(unsafe.Pointer(&tss[0]))), Limit:uint32(unsafe.Sizeof(tss))-1, Type: segment.TSSAvail}.Pack().Decompose()
}