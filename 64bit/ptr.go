package ptr

import "unsafe"

// Hack to get the address of a function
// Borowed from reflect.MakeFunc()
func FuncToPtr(f func()) uintptr {
	dummy := f
	return **(**uintptr)(unsafe.Pointer(&dummy))
}

func GetAddr(addr uintptr) unsafe.Pointer {
	return (unsafe.Pointer(addr))
}
