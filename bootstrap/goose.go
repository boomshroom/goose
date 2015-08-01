package main

import (
	"page"
	"elf"
)

//extern __get_kernel64
func kernel64()uintptr

//extern __enable_64bit
func enableLong(entry, pml4 uintptr)

func main() {
	//page.Init()
	enableLong(elf.Parse(kernel64()), page.Mapl4)
}
