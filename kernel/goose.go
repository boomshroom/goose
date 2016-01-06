package main

import (
	"elf"
	_ "gdt"
	_ "idt"
	_ "kbd"
	_ "syscall"
	"tables"
	"video"
	//"unsafe"
)

//extern __start_app
func startApp(func())

//extern __break
func breakPoint()

var mods []tables.Mod

func main() {
	mods = tables.MultibootTable.Mods()
	//gdt.SetupGDT()
	//pit.Init()
	//kbd.Init()
	//video.Init()

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

	if tables.MultibootTable.Flags&tables.Mods == 0 {
		video.Println("Mods dissabled")
	} else {
		if len(mods) == 0 {
			video.Print("No ")
		}
		video.Println("Modules loaded")
		for i := 0; i < len(mods); i++ {
			video.Println(mods[i].Name())

			video.Println("Reading App...")
			app := elf.Parse(&mods[i].Bytes()[0])
			video.Println("Loading App...")
			app.CopyToMem()
			video.Println("Launching App!")

			println(app.Entry)

			startApp(app.Func())
		}
	}
}
