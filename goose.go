package kernel

import (
	//	"color"
	"idt"
	"video"
	//"ptr"
	"gdt"
	"pit"
	//"unsafe"
	"kbd"
)

//extern __test_int
func testInt()

//extern __test_args
func testArgs(c rune)

//func Kmain() {
func Kmain(mdb uintptr, magic uint16) {
	video.Init()
	video.Clear()
	gdt.SetupGDT()
	idt.SetupIDT()
	idt.SetupIRQ()
	pit.Init()
	kbd.Init()
	//println("Hello")
	video.Print("  ______  _____       _____    _          \n")
	video.Print(" / _____)/ ___ \\     / ___ \\  | |         \n")
	video.Print("| /  ___| |   | |___| |   | |  \\ \\   ____ \n")
	video.Print("| | (___) |   | (___) |   | |   \\ \\ / _  )\n")
	video.Print("| \\____/| |___| |   | |___| |____) | (/ / \n")
	video.Print(" \\_____/ \\_____/     \\_____(______/ \\____)\n")
	video.Print("                                GO-OSe\n")
	video.Print("Proof of concept Golang <golang.org> x86 kernel\n")
	video.Print("by Tom Gascoigne <tom.gascoigne.me>\n")
	video.Print("and Angelo B <mbulfone@gmail.com>\n")
	//video.PrintHex(uint64(idt.Pack(idt.IDTDesc{Offset:uint32(ptr.FuncToPtr(genericInt)), Selector:0x08, TypeAttr:0x8E})), false, true, true, 16)
	//video.PrintHex(uint64(ptr.FuncToPtr(genericInt)), false, true, true, 8)
	//video.PrintHex(uint64(idt.IrqRoutines[0]), false, true, true, 8)
	//video.PrintHex(*(*uint64)(unsafe.Pointer(idt.IrqRoutines[0])), false, true, true, 8)
	//video.PrintHex(gdt.GDT, false, true, true, 8)
	//video.Println("0x08  0x8e")
	//testInt()
	//idt.Irq0()
	//video.Print("Divided by zero?")
	//zero := 0
	//_ = 1/zero
	//_ = 2/zero
	//testArgs('a')
	//video.PrintHex(uint64(video))
	//idt.PtrToFunc(idt.IrqRoutines[0])(nil)
}
