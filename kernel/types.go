package types

import "video"

type Array [1<<30]uint8 // Max size of array aperantly

func HashIdent(key, keySize uintptr)uintptr{
	return 0
}

func EqualIdent(p1, p2, keySize uintptr)bool{
	return true
}

func HashError(val, keySize uintptr)uintptr{
	video.Error(ErrorMsg[0], 0, true)
	return 0
}

func EqualError(v1, v2, keySize uintptr)uintptr{
	video.Error(ErrorMsg[1], 1, true)
	return 0
}

func MemCmp(v1, v2 *Array, size uint)int{
	for i:=uint(0); i<size; i++{
		if v1[i] < v2[i]{
			return -1
		}else if v1[i] > v2[i] {
			return 1
		}
	}
	return 0
}

var ErrorMsg [2][40]byte

func Init(){
	video.CopyStr(&ErrorMsg[0], "Unhashable Type Exception")
	video.CopyStr(&ErrorMsg[1], "Incomparable Type Exception")
}