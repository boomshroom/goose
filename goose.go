package kernel

import (
	"idt"
	"video"
	"gdt"
	"pit"
	"kbd"
	"page"
	"types"
	//"ptr"
)

//extern __kernel_end
func kernelEnd()uintptr

//extern __kernel_start
func kernelStart()uintptr

func Kmain() {
	types.Init()
	gdt.SetupGDT()
	idt.SetupIDT()
	idt.SetupIRQ()
	video.Init()
	video.Clear()
	page.Init()
	pit.Init()
	kbd.Init()
	
	video.Print("  ______  _____       _____    _          \n")
	video.Print(" / _____)/ ___ \\     / ___ \\  | |         \n")
	video.Print("| /  ___| |   | |___| |   | |  \\ \\   ____ \n")
	video.Print("| | (___) |   | (___) |   | |   \\ \\ / _  )\n")
	video.Print("| \\____/| |___| |   | |___| |____) | (/ / \n")
	video.Print(" \\_____/ \\_____/     \\_____(______/ \\____)\n")
	video.Print("                                GO-OSe\n")
	video.Print("Proof of concept Golang <golang.org> x86 kernel\n")
	video.Print("by Tom Gascoigne <tom.gascoigne.me>\n")
	video.Print("and Angelo B\n")
	
	//page.Init()*/
	//video.PrintHex(*(*uint64)(ptr.GetAddr(0xA0000000)), true, true, true, 8)
	//println(^uint64(0x20-1))
	//video.PrintHex(uint64(kernelStart()), true, true, true, 8)
	//video.PrintHex(uint64(kernelEnd()-kernelStart()), true, true, true, 8)
	//video.PrintHex((uint64(kernelEnd())& 0xFFFFF000) + 0x2000, true, true, true, 8)
	//video.PrintHex(uint64(page.PageDir), true, true, true, 8)
	//for i:=0; i<23; i++{
	//	video.PrintHex(uint64((*[1024]uint32)(ptr.GetAddr(page.PageDir))[i]), true, true, true, 8)
	//}
}
