package asm

//extern inportb
func inportB(uint16) uint8

//extern outportb
func outportB(uint16, uint8)

//extern enable_ints
func enableInts()

//extern io_wait
func ioWait()

func OutportB(port uint16, data uint8) {
	outportB(port, data)
}

func InportB(port uint16) uint8 {
	return inportB(port)
}

func EnableInts() {
	enableInts()
}

func IOWait(){
	ioWait()
}