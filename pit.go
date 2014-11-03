package pit

import (
	"idt"
	"regs"
	"unsafe"
	"video"
)

var ticks int

func handler(r *regs.Regs) {
	/*ticks++
	if ticks%36 == 0 {
		//video.Println("Tick")
	} else if ticks%36 == 18 {
		//video.Println("Tock")
	}*/
}

func Init() {
	dummy := handler
	idt.AddIRQ(0, **(**uintptr)(unsafe.Pointer(&dummy)))
	video.Println("PIT Initialized")
}
