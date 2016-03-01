package runtime

import "unsafe"

type String struct {
	Ptr *Array
	Len int
}

var goStringTemp String

func GoString(ptr *uint8) string {
	if ptr == nil {
		return ""
	}
	l := 0
	first := (*Array)(unsafe.Pointer(ptr))
	for ; first[l] != 0; l++ {
	}
	goStringTemp = String{Ptr: first, Len: l}
	return *(*string)(unsafe.Pointer(&goStringTemp))
}

func StringsEqual(s1, s2 string) bool {
	return len(s1) == len(s2) && MemCmp((*String)(unsafe.Pointer(&s1)).Ptr, (*String)(unsafe.Pointer(&s2)).Ptr, len(s1)) == 0
}

func PtrStringsEqual(ps1, ps2 *string) bool {
	if ps1 == nil {
		return ps2 == nil
	} else if ps2 == nil {
		return false
	} else {
		return StringsEqual(*ps1, *ps2)
	}
}

// stringiter2 returns the rune that starts at s[k]
// and the index where the next rune starts.
func StringIter2(s string, k int) (int, rune) {
	if k >= len(s) {
		// 0 is end of iteration
		return 0, 0
	}

	c := s[k]
	//if c < runeself {
	return k + 1, rune(c)
	//}

	// multi-char rune
	//r, n := charntorune(s[k:])
	//return k + n, r
}
