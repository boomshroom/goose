package idt

import (
	"asm"
	"unsafe"
	"video"
	"segment"
)

var IDT = segment.TablePtr{Size: unsafe.Sizeof(table), Ptr: uintptr(unsafe.Pointer(&table))}.Pack()

const size uint16 = 256

var table = [size]segment.Seg128{
	0x0: segment.GateDesc{Offset: funcToPtr(isr0), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1: segment.GateDesc{Offset: funcToPtr(isr1), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2: segment.GateDesc{Offset: funcToPtr(isr2), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x3: segment.GateDesc{Offset: funcToPtr(isr3), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x4: segment.GateDesc{Offset: funcToPtr(isr4), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x5: segment.GateDesc{Offset: funcToPtr(isr5), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x6: segment.GateDesc{Offset: funcToPtr(isr6), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x7: segment.GateDesc{Offset: funcToPtr(isr7), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x8: segment.GateDesc{Offset: funcToPtr(isr8), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x9: segment.GateDesc{Offset: funcToPtr(isr9), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0xa: segment.GateDesc{Offset: funcToPtr(isr10), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0xb: segment.GateDesc{Offset: funcToPtr(isr11), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0xc: segment.GateDesc{Offset: funcToPtr(isr12), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0xd: segment.GateDesc{Offset: funcToPtr(isr13), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0xe: segment.GateDesc{Offset: funcToPtr(isr14), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0xf: segment.GateDesc{Offset: funcToPtr(isr15), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x10: segment.GateDesc{Offset: funcToPtr(isr16), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x11: segment.GateDesc{Offset: funcToPtr(isr17), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x12: segment.GateDesc{Offset: funcToPtr(isr18), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x13: segment.GateDesc{Offset: funcToPtr(isr19), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x14: segment.GateDesc{Offset: funcToPtr(isr20), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x15: segment.GateDesc{Offset: funcToPtr(isr21), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x16: segment.GateDesc{Offset: funcToPtr(isr22), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x17: segment.GateDesc{Offset: funcToPtr(isr23), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x18: segment.GateDesc{Offset: funcToPtr(isr24), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x19: segment.GateDesc{Offset: funcToPtr(isr25), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1a: segment.GateDesc{Offset: funcToPtr(isr26), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1b: segment.GateDesc{Offset: funcToPtr(isr27), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1c: segment.GateDesc{Offset: funcToPtr(isr28), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1d: segment.GateDesc{Offset: funcToPtr(isr29), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1e: segment.GateDesc{Offset: funcToPtr(isr30), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x1f: segment.GateDesc{Offset: funcToPtr(isr31), Selector: 0x08, Type: segment.Interupt}.Pack(),

	//0x20: segment.GateDesc{Offset: funcToPtr(irq0), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x20: segment.Seg128{},
	0x21: segment.GateDesc{Offset: funcToPtr(irq1), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x22: segment.GateDesc{Offset: funcToPtr(irq2), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x23: segment.GateDesc{Offset: funcToPtr(irq3), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x24: segment.GateDesc{Offset: funcToPtr(irq4), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x25: segment.GateDesc{Offset: funcToPtr(irq5), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x26: segment.GateDesc{Offset: funcToPtr(irq6), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x27: segment.GateDesc{Offset: funcToPtr(irq7), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x28: segment.GateDesc{Offset: funcToPtr(irq8), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x29: segment.GateDesc{Offset: funcToPtr(irq9), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2a: segment.GateDesc{Offset: funcToPtr(irq10), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2b: segment.GateDesc{Offset: funcToPtr(irq11), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2c: segment.GateDesc{Offset: funcToPtr(irq12), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2d: segment.GateDesc{Offset: funcToPtr(irq13), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2e: segment.GateDesc{Offset: funcToPtr(irq14), Selector: 0x08, Type: segment.Interupt}.Pack(),
	0x2f: segment.GateDesc{Offset: funcToPtr(irq15), Selector: 0x08, Type: segment.Interupt}.Pack(),

	0x80: segment.GateDesc{Offset: funcToPtr(syscall), Selector: 0x08, Type: segment.Interupt, User: true}.Pack(),
}

// Hack to get the address of a function
// Borowed from reflect.MakeFunc()
func funcToPtr(f func()) uintptr {
	dummy := f
	return **(**uintptr)(unsafe.Pointer(&dummy))
}

func init() {

	loadIDT(&IDT)
	remapIRQ()

	asm.EnableInts()
}

//extern __load_idt
func loadIDT(*segment.TablePtrPacked)

func remapIRQ() {
	master := asm.InportB(0x21)
	slave := asm.InportB(0xA1)
	
	asm.OutportB(0x20, 0x11)
	asm.IOWait()
	asm.OutportB(0xA0, 0x11)
	asm.IOWait()
	asm.OutportB(0x21, 0x20)
	asm.IOWait()
	asm.OutportB(0xA1, 0x28)
	asm.IOWait()
	asm.OutportB(0x21, 0x04)
	asm.IOWait()
	asm.OutportB(0xA1, 0x02)
	asm.IOWait()
	
	asm.OutportB(0x21, 0x01)
	asm.IOWait()
	asm.OutportB(0xA1, 0x01)
	asm.IOWait()

	//asm.OutportB(0xA1, 0xff)
	//asm.OutportB(0x21, 0xff)
	
	asm.OutportB(0x21, master)
	asm.OutportB(0xA1, slave)
}

type intsStruct struct{
	errMsgs [20]string
	errHandlers [20]func(uint32)
}

var Interrupts = intsStruct{
	errMsgs: [20]string{
		"Division By Zero Exception",
		"Debug Exception",
		"Non Maskable Interrupt Exception",
		"Breakpoint Exception",
		"Into Detected Overflow Exception",
		"Out of Bounds Exception",
		"Invalid Opcode Exception",
		"No Coprocessor Exception",
		"Double Fault Exception",
		"Coprocessor Segment Overrun Exception",
		"Bad TSS Exception",
		"Segment Not Present Exception",
		"Stack Fault Exception",
		"General Protection Fault Exception",
		"Page Fault Exception",
		"Unknown Interrupt Exception",
		"Floating-Point Math Exception",
		"Alignment Check Exception (486+)",
		"Machine Check Exception (Pentium/586+)",
		"Reserved Exception",
	},
	errHandlers: [20]func(uint32) {
		0xD: func(errCode uint32){
			index := (errCode>>3)&(1<<13 -1)
			switch (errCode>>1)&3{
			case 0:
				video.Print("GDT ")
				switch index{
				case 0:
					video.Println("Null Descriptor")
				case 1:
					video.Println("Kernel Code Descriptor")
				case 2:
					video.Println("Kernel Data Descriptor")
				case 3:
					video.Println("User Code Descriptor")
				case 4:
					video.Println("User Data Descriptor")
				default:
					video.Print("Descriptor: ")
					video.PrintUint(uint64(index))
					video.NL()
				}
			case 1, 3:
				video.Print("IDT Descriptor: ")
				video.PrintUint(uint64(index))
				video.NL()
			case 2:
				video.Println("LDT Descriptor: ")
				video.PrintUint(uint64(index))
				video.NL()
			}
			if errCode & 1 != 0 {
				video.Println("External Source")
			}
		},
		0xE: func(errCode uint32){
			if errCode&1 == 0{
				video.Println("Page not present")
			}else{
				video.Println("Page protection violation")
			}
			if errCode&2 != 0{
				video.Println("Atempted page write")
			}
			if errCode&4 != 0{
				video.Println("Userspace active")
			}
			if errCode&8 != 0{
				video.Println("Read 1 in reserved field")
			}
			if errCode&16 != 0{
				video.Println("Atempted instruction fetch")
			}
		},
	},
}


func ISR(intNo, errCode, rip uint64) {

	if intNo < 32 {
		
		if intNo > 18 {
			video.Error(Interrupts.errMsgs[19], int(intNo), false)
		} else {
			video.Error(Interrupts.errMsgs[intNo], int(intNo), false)
		}
	}

	video.Print("Interrupt occured at ")
	video.PrintUint(rip)
	video.NL()
	if Interrupts.errHandlers[intNo] != nil{
		Interrupts.errHandlers[intNo](uint32(errCode))
	}
	if errCode != 0 {
		video.Print("Error code: ")
		video.PrintUint(errCode)
	}
	for{}
}

func Syscall(str string, id uint){
	switch id{
	case 0:
		video.Print(str)
	case 1:
		video.NL()
	}
	
}

var IrqRoutines [16]uintptr

func AddIRQ(index uint8, query uintptr) {
	IrqRoutines[index] = query
}

func RemoveIRQ(index uint8) {
	IrqRoutines[index] = 0
}

func IRQ(intNo, errCode, rip uint) {
	video.Print("Interrupt Request Query: ")
	video.PrintUint(uint64(intNo))
	video.NL()
	video.Print("At: ")
	video.PrintUint(uint64(rip))
	video.NL()
	if intNo == 7{
		asm.OutportB(0x20, 0x0B)
		irr := asm.InportB(0x20)
		if irr & 0x80 == 0 {
			return
		}
	}
	handler := &IrqRoutines[intNo-32]
	if *handler != 0 {
		(*(*func())(unsafe.Pointer(handler)))()
	}
	if intNo >= 40 {
		asm.OutportB(0xA0, 0x20)
	}
	asm.OutportB(0x20, 0x20)
}

//extern go.pit.Handler
func pitHandler(intNo, errCode uint)

//extern __isr0
func isr0()

//extern __isr1
func isr1()

//extern __isr2
func isr2()

//extern __isr3
func isr3()

//extern __isr4
func isr4()

//extern __isr5
func isr5()

//extern __isr6
func isr6()

//extern __isr7
func isr7()

//extern __isr8
func isr8()

//extern __isr9
func isr9()

//extern __isr10
func isr10()

//extern __isr11
func isr11()

//extern __isr12
func isr12()

//extern __isr13
func isr13()

//extern __isr14
func isr14()

//extern __isr15
func isr15()

//extern __isr16
func isr16()

//extern __isr17
func isr17()

//extern __isr18
func isr18()

//extern __isr19
func isr19()

//extern __isr20
func isr20()

//extern __isr21
func isr21()

//extern __isr22
func isr22()

//extern __isr23
func isr23()

//extern __isr24
func isr24()

//extern __isr25
func isr25()

//extern __isr26
func isr26()

//extern __isr27
func isr27()

//extern __isr28
func isr28()

//extern __isr29
func isr29()

//extern __isr30
func isr30()

//extern __isr31
func isr31()

//extern __irq0
func irq0()

//extern __irq1
func irq1()

//extern __irq2
func irq2()

//extern __irq3
func irq3()

//extern __irq4
func irq4()

//extern __irq5
func irq5()

//extern __irq6
func irq6()

//extern __irq7
func irq7()

//extern __irq8
func irq8()

//extern __irq9
func irq9()

//extern __irq10
func irq10()

//extern __irq11
func irq11()

//extern __irq12
func irq12()

//extern __irq13
func irq13()

//extern __irq14
func irq14()

//extern __irq15
func irq15()

//extern __syscall
func syscall()