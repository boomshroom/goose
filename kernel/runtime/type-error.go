package runtime

func TypeHashError(val, keySize uintptr){
	panicString("Hash of unhashable type")
}

func TypeEqualError(val, keySize uintptr){
	panicString("Comparing uncomparable types")
}