package pit

import (
	"idt"
	"regs"
	"unsafe"
)

var ticks int

func handler(r *regs.Regs) {
	//TODO
}

func Init() {
	dummy := handler
	idt.AddIRQ(0, **(**uintptr)(unsafe.Pointer(&dummy)))
}
