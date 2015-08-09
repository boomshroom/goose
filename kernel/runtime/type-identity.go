package runtime

import "unsafe"

func TypeHashIdentity(key *Array, keySize int)(ret uint64){
	if keySize <= 8{
		ret := *(*uint64)(unsafe.Pointer(&key))
		return ret & (1<<uint(keySize << 3) - 1)
	}

	ret = 5381
	
	p := key[:keySize]
	for _, b := range p{
		ret = ret * 33 + uint64(b)
	}
	return
}

func TypeEqualIdentity(k1, k2 *Array, keySize int)bool{
	return MemCmp(k1, k2, keySize) == 0
}