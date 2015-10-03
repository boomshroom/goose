package asm

//extern inportb
func inportB(uint16) uint8

//extern inport16
func inport16(uint16) uint16

//extern outportb
func outportB(uint16, uint8)

//extern enable_ints
func enableInts()

//extern io_wait
func ioWait()

//extern memcpy
func memcpy(dest, src, size uintptr)

func OutportB(port uint16, data uint8) {
	outportB(port, data)
}

func InportB(port uint16) uint8 {
	return inportB(port)
}

func Inport16(port uint16) uint16 {
	return inport16(port)
}


func EnableInts() {
	enableInts()
}

func IOWait(){
	ioWait()
}