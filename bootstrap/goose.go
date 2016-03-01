package main

import (
	"elf"
	"page"
)

//extern __get_kernel64
func kernel64() uintptr

//extern __get_kernel64_size
func kernel64Size() uintptr

//extern __enable_64bit
func enableLong(entry uintptr, pml4 *page.Page)

func main() {
	//page.Init()
	enableLong(elf.Parse(kernel64(), kernel64Size()), &page.Pages[0])
}
