package runtime

import "unsafe"

func StrCmp(s1, s2 string)int{
	var l int
	if len(s1) < len(s2){
		l = len(s1)
	}else{
		l = len(s2)
	}

	if i:= MemCmp((*String)(unsafe.Pointer(&s1)).Ptr, (*String)(unsafe.Pointer(&s2)).Ptr, l); i!=0{
		return int(i)
	}else if len(s1) < len(s2){
		return -1
	}else if len(s1) > len(s2) {
		return 1
	}else{
		return 0
	}
}