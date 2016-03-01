package syscall

import (
	"asm"
	"capability"
	"elf"
	"idt"
	"kbd"
	"page"
	"proc"
	"tables"
	"unsafe"
	"video"
)

type inport struct {
	ptr  *byte
	port uint16
}

type interupt struct {
	intNum uint64
	entry  func()
}

const (
	invalid = iota
	Print
	GetChar
	PrintInt
	InportB
	RequestCapability
	RegisterInterupt
	StartProc
)

//extern __start_app
func startApp(func())

var syscalls = [...]func(unsafe.Pointer){
	invalid: func(p unsafe.Pointer) {
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
		for _, port := range capability.Capabilities[proc.Procs[proc.CurrentID].Capability].Ports {
			if port == arg.port {
				*arg.ptr = asm.InportB(arg.port)
				continue
			}
		}
	},
	RequestCapability: func(p unsafe.Pointer) {
		capa := &capability.Capabilities[*(*uint64)(unsafe.Pointer(&p))]
		if capa.OwnedProc == 0 && proc.Procs[proc.CurrentID].Capability == 0 {
			capa.OwnedProc = proc.CurrentID
			proc.Procs[proc.CurrentID].Capability = *(*uint)(p)
		}
	},
	RegisterInterupt: func(p unsafe.Pointer) {
		intReq := *(*interupt)(p)
		if intReq.intNum < 15 && idt.IrqRoutines[intReq.intNum+1] == nil {
			page.NewPage(0x7FFFFFFFE000, page.K, page.PRESENT|page.USER|page.READ_WRITE)
			idt.AddIRQ(uint8(intReq.intNum+1), intReq.entry)
		}
	},
	StartProc: func(p unsafe.Pointer) {
		cmd := *(*string)(p)
		for i := 0; i < len(tables.Modules); i++ {
			if tables.Modules[i].Name() == cmd {
				app := elf.Parse(&tables.Modules[i].Bytes()[0])
				app.CopyToMem()
				return
			}
		}
	},
}

//extern __break
func breakPoint()

func Syscall(kill bool) {
	l := proc.Current.Len
	if l > 40 {
		l = 40
	}
	for i := uint64(0); i < l; i++ {
		sys := proc.Current.Syscalls[i]
		if int(sys.Id) < len(syscalls) && syscalls[sys.Id] != nil {
			syscalls[sys.Id](sys.Args)
		}
	}
	proc.Current.Len = 0

	if proc.NumProcs == 0 {
		video.Error("No more processes running!", 0, true)
	}

	if proc.CurrentID == 0 {
		proc.CurrentID = 1
	}

	if kill {
		//video.Print("Killing process ")
		//video.PrintHex(proc.CurrentID, false, true, true, 0)
		proc.KillProc()
	}

	proc.CurrentID++
	if proc.CurrentID > proc.NumProcs {
		proc.CurrentID = 1
	}
	p := &proc.Procs[proc.CurrentID]
	p.Enable()
}
