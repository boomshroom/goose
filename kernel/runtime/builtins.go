package runtime

import "unsafe"

type Array [1<<30]uint8

//extern __break
func breakPoint()

func MemCmp(str1, str2 *Array, count int)int32{
	breakPoint()
	for i := 0; i < count; i++{
		if str1[i] < str2[i]{
			return -1
		}else if str1[i] > str2[i]{
			return 1
		}
	}
	return 0
}

func MemCpy(dest, src *Array, count int)*Array{
	for i := 0; i<count; i++{
		dest[i] = src[i]
	}
	return dest
}

func MemMove(dest, src *Array, count int)*Array{
	if uintptr(unsafe.Pointer(dest)) < uintptr(unsafe.Pointer(src)){
		for i := 0; i<count; i++{
			dest[i] = src[i]
		}
	}else{
		for i := count-1; i>=0; i--{
			dest[i] = src[i]
		}
	}
	return dest
}