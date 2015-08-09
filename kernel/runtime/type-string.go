package runtime

func TypeHashString(key *string, _ uint)(ret uint){
	ret = 5381
	for _, ch := range *key {
		ret = ret * 33 + uint(ch)
	}
	return
}

func TypeEqualString(k1, k2 *string, _ uint)bool{
	return PtrStringsEqual(k1, k2)
}