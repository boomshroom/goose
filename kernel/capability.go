package capability

import (
	"unsafe"
)

type Capability struct {
	OwnedProc uint64
	Ports     []uint16
	Pages     []uintptr
	Addrs     []uintptr
}

const (
	Print = iota
	Kbd
	num_capabilities
)

var printPorts = [2]uint16{*(*uint16)(unsafe.Pointer(uintptr(0x0463))), *(*uint16)(unsafe.Pointer(uintptr(0x0463))) + 1}
var printAddr = [1]uintptr{0x410}
var printPages = [2]uintptr{0xB0000, 0xB8000}

var kbdPort = [1]uint16{0x60}

func init() {

}

var Capabilities = [num_capabilities]Capability{
	Print: {Ports: printPorts[:], Pages: printPages[:], Addrs: printAddr[:]},
	Kbd:   {Ports: kbdPort[:]},
}
