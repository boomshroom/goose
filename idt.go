package idt

import (
	"asm"
	"ptr"
	"regs"
	"unsafe"
	"video"
)

type IDTDescPacked uint64

type IDTDesc struct {
	Offset   uint32
	Selector uint16
	TypeAttr uint8
}

var IDT uint64

const size uint16 = 256

var table [size]IDTDescPacked

func Pack(desc IDTDesc) (i IDTDescPacked) {
	i = IDTDescPacked(desc.Offset & 0xFFFF)          //lower offset
	i |= IDTDescPacked(desc.Offset&0xFFFF0000) << 32 //higher offset
	i |= IDTDescPacked(desc.Selector) << 16
	i |= IDTDescPacked(desc.TypeAttr) << 40
	return
}

func SetupIDT() {
	IDT = uint64((size * 64) - 1)
	IDT |= (uint64(uintptr(unsafe.Pointer(&table))) & 0xFFFFFFFF) << 16
	loadTable()
	loadIDT()
	//genericInt()

	copyStr(&ErrorMsg[0], "Division By Zero Exception")
	copyStr(&ErrorMsg[1], "Debug Exception")
	copyStr(&ErrorMsg[2], "Non Maskable Interrupt Exception")
	copyStr(&ErrorMsg[3], "Breakpoint Exception")
	copyStr(&ErrorMsg[4], "Into Detected Overflow Exception")
	copyStr(&ErrorMsg[5], "Out of Bounds Exception")
	copyStr(&ErrorMsg[6], "Invalid Opcode Exception")
	copyStr(&ErrorMsg[7], "No Coprocessor Exception")
	copyStr(&ErrorMsg[8], "Double Fault Exception")
	copyStr(&ErrorMsg[9], "Coprocessor Segment Overrun Exception")
	copyStr(&ErrorMsg[10], "Bad TSS Exception")
	copyStr(&ErrorMsg[11], "Segment Not Present Exception")
	copyStr(&ErrorMsg[12], "Stack Fault Exception")
	copyStr(&ErrorMsg[13], "General Protection Fault Exception")
	copyStr(&ErrorMsg[14], "Page Fault Exception")
	copyStr(&ErrorMsg[15], "Unknown Interrupt Exception")
	copyStr(&ErrorMsg[16], "Coprocessor Fault Exception")
	copyStr(&ErrorMsg[17], "Alignment Check Exception (486+)")
	copyStr(&ErrorMsg[18], "Machine Check Exception (Pentium/586+)")
	copyStr(&ErrorMsg[19], "Reserved Exception")
}

func SetupIRQ() {
	remapIRQ()

	table[32] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq0)), Selector: 0x08, TypeAttr: 0x8E})
	table[33] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq1)), Selector: 0x08, TypeAttr: 0x8E})
	table[34] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq2)), Selector: 0x08, TypeAttr: 0x8E})
	table[35] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq3)), Selector: 0x08, TypeAttr: 0x8E})
	table[36] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq4)), Selector: 0x08, TypeAttr: 0x8E})
	table[37] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq5)), Selector: 0x08, TypeAttr: 0x8E})
	table[38] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq6)), Selector: 0x08, TypeAttr: 0x8E})
	table[39] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq7)), Selector: 0x08, TypeAttr: 0x8E})
	table[40] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq8)), Selector: 0x08, TypeAttr: 0x8E})
	table[41] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq9)), Selector: 0x08, TypeAttr: 0x8E})
	table[42] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq10)), Selector: 0x08, TypeAttr: 0x8E})
	table[43] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq11)), Selector: 0x08, TypeAttr: 0x8E})
	table[44] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq12)), Selector: 0x08, TypeAttr: 0x8E})
	table[45] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq13)), Selector: 0x08, TypeAttr: 0x8E})
	table[46] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq14)), Selector: 0x08, TypeAttr: 0x8E})
	table[47] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(irq15)), Selector: 0x08, TypeAttr: 0x8E})

	asm.EnableInts()
	//dummy := pitHandler
	//IrqRoutines[0] = **(**uintptr)(unsafe.Pointer(&dummy))
}

func Irq0() {
	//irq0()
}

//extern __load_idt
func loadIDT()

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
	
	asm.OutportB(0x21, master)
	asm.OutportB(0xA1, slave)
	//asm.OutportB(0x21, 0xFF)
	//asm.OutportB(0xA1, 0xFF)
}

var ErrorMsg [20][40]byte

func ISR(r *regs.Regs) {
	//video.PrintHex(uint64(r.IntNo), false, true, true, 8)

	if r.IntNo < 32 {
		
		if r.IntNo > 18 {
			video.Error(ErrorMsg[19], int(r.IntNo), true)
		} else {
			if r.ErrCode != 0{
				video.Print("Error Code: ")
				video.PrintHex(uint64(r.ErrCode), false, true,true, 8)
			}
			video.Error(ErrorMsg[r.IntNo], int(r.IntNo), true)
		}
	}
}

func copyStr(array *[40]byte, str string) {
	for i := 0; i < len(str); i++ {
		array[i] = str[i]
	}
}

var IrqRoutines [16]uintptr

func AddIRQ(index uint8, query uintptr) {
	IrqRoutines[index] = query
}

func RemoveIRQ(index uint8) {
	IrqRoutines[index] = 0
}

func IRQ(r *regs.Regs) {
	if r.IntNo == 7{
		asm.OutportB(0x20, 0x0B)
		irr := asm.InportB(0x20)
		if irr & 0x80 == 0 {
			return
		}
	}
	//video.PrintHex(uint64(r.IntNo), false, true, true, 4)
	var handler uintptr = IrqRoutines[r.IntNo-32]
	if handler != 0 {
		call(handler, r)
	}
	if r.IntNo >= 40 {
		asm.OutportB(0xA0, 0x20)
	}
	asm.OutportB(0x20, 0x20)
}

//extern __call
func call(ptr uintptr, r *regs.Regs)

//extern __arbitrary_convert
func PtrToFunc(ptr uintptr) func(r *regs.Regs) //Y U no allow (func())(unsafe.Pointer) ?

//extern go.pit.Handler
func pitHandler(r *regs.Regs)

//func Table()*[size]IDTDescPacked{
//	return (*[size]IDTDescPacked)(ptr.GetAddr(table))
//}

func loadTable() {
	table[0] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr0)), Selector: 0x08, TypeAttr: 0x8E})
	table[1] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr1)), Selector: 0x08, TypeAttr: 0x8E})
	table[2] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr2)), Selector: 0x08, TypeAttr: 0x8E})
	table[3] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr3)), Selector: 0x08, TypeAttr: 0x8E})
	table[4] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr4)), Selector: 0x08, TypeAttr: 0x8E})
	table[5] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr5)), Selector: 0x08, TypeAttr: 0x8E})
	table[6] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr6)), Selector: 0x08, TypeAttr: 0x8E})
	table[7] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr7)), Selector: 0x08, TypeAttr: 0x8E})
	table[8] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr8)), Selector: 0x08, TypeAttr: 0x8E})
	table[9] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr9)), Selector: 0x08, TypeAttr: 0x8E})
	table[10] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr10)), Selector: 0x08, TypeAttr: 0x8E})
	table[11] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr11)), Selector: 0x08, TypeAttr: 0x8E})
	table[12] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr12)), Selector: 0x08, TypeAttr: 0x8E})
	table[13] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr13)), Selector: 0x08, TypeAttr: 0x8E})
	table[14] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr14)), Selector: 0x08, TypeAttr: 0x8E})
	table[15] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr15)), Selector: 0x08, TypeAttr: 0x8E})
	table[16] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr16)), Selector: 0x08, TypeAttr: 0x8E})
	table[17] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr17)), Selector: 0x08, TypeAttr: 0x8E})
	table[18] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr18)), Selector: 0x08, TypeAttr: 0x8E})
	table[19] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr19)), Selector: 0x08, TypeAttr: 0x8E})
	table[20] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr20)), Selector: 0x08, TypeAttr: 0x8E})
	table[21] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr21)), Selector: 0x08, TypeAttr: 0x8E})
	table[22] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr22)), Selector: 0x08, TypeAttr: 0x8E})
	table[23] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr23)), Selector: 0x08, TypeAttr: 0x8E})
	table[24] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr24)), Selector: 0x08, TypeAttr: 0x8E})
	table[25] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr25)), Selector: 0x08, TypeAttr: 0x8E})
	table[26] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr26)), Selector: 0x08, TypeAttr: 0x8E})
	table[27] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr27)), Selector: 0x08, TypeAttr: 0x8E})
	table[28] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr28)), Selector: 0x08, TypeAttr: 0x8E})
	table[29] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr29)), Selector: 0x08, TypeAttr: 0x8E})
	table[30] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr30)), Selector: 0x08, TypeAttr: 0x8E})
	table[31] = Pack(IDTDesc{Offset: uint32(ptr.FuncToPtr(isr31)), Selector: 0x08, TypeAttr: 0x8E})
}

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
