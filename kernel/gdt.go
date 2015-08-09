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

var TSS segment.TSSPacked = segment.TSSPacked{} //segment.TSS{RSP: [3]uint64{0: stack()}}.Pack()

var gdtr segment.TablePtrPacked = segment.TablePtr{Size: unsafe.Sizeof(Table), Ptr: uintptr(unsafe.Pointer(&Table))}.Pack()

func init(){

	Table[5], Table[6] = segment.SystemDesc{Base:uint64(uintptr(unsafe.Pointer(&TSS[0]))), Limit:uint32(unsafe.Sizeof(TSS))-1, Type: segment.TSSAvail}.Pack().Decompose()
	loadGDT(&gdtr)
	//reloadSegments()
}

func SetKernelStack(stack uintptr){
	TSS.SetKernelStack(stack)
}

//extern __stack_ptr
func stack()uint64

//extern __load_gdt
func loadGDT(*segment.TablePtrPacked)

//extern __reload_segments
func reloadSegments()