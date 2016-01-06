package syscall

import (
	"asm"
	"capability"
	"kbd"
	"proc"
	"unsafe"
	"video"
)

type inport struct {
	ptr  *byte
	port uint16
}

const (
	invalid = iota
	Print
	GetChar
	PrintInt
	InportB
	RequestCapability
	RegisterInterupt
)

var syscalls = [...]func(unsafe.Pointer){
	invalid: func(p unsafe.Pointer){
		i := *(*int)(unsafe.Pointer(&p))
		video.Error("Invalid syscall!", i, true)
	},
	Print: func(p unsafe.Pointer) {
		video.Print(*(*string)(p))
	},
	GetChar: func(p unsafe.Pointer) {
		kbd.Buffer = (*int)(p)
	},
	PrintInt: func(p unsafe.Pointer) {
		i := *(*uint64)(unsafe.Pointer(&p))
		if i < 128 {
			video.PutChar(rune(i))
		} else {
			video.PrintUint(i)
		}
	},
	InportB: func(p unsafe.Pointer) {
		arg := (*inport)(p)
		for _, port := range capability.Capabilities[proc.Current.Capability].Ports {
			if port == arg.port {
				*arg.ptr = asm.InportB(arg.port)
				continue
			}
		}
	},
	RequestCapability: func(p unsafe.Pointer) {
		capa := &capability.Capabilities[*(*uint64)(unsafe.Pointer(&p))]
		if capa.OwnedProc == 0 && proc.Current.Capability == 0 {
			capa.OwnedProc = proc.Current.Id
			proc.Current.Capability = *(*uint)(p)
		}
	},
	RegisterInterupt: func(p unsafe.Pointer){
		
	},
}

func Syscall() {
	l := proc.Current.SyscallLen
	if l > 40 {
		l = 40
	}
	for i := uint64(0); i < l; i++ {
		sys := proc.Current.Syscalls[i]
		if int(sys.Id) < len(syscalls) && syscalls[sys.Id] != nil {
			syscalls[sys.Id](sys.Args)
		}
	}
	proc.Current.SyscallLen = 0
}
