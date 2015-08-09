package main

import (
	_ "gdt"
	"video"
	"elf"
)

//extern __get_app
func getAppStart()uintptr

func main(){
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
	video.NL()
	video.Println("Reading App...")
	app := elf.Parse(getAppStart())
	video.Println("Loading App...")
	app.CopyToMem()
	video.Println("Launching App!")
	app.Func()()
}