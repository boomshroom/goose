package types

//import "video"

//extern __kill
func kill()

func HashIdent(key, keySize uintptr)uintptr{
	return 0
}

func EqualIdent(p1, p2, keySize uintptr)bool{
	return true
}
/*
func HashError(val, keySize uintptr)uintptr{
	kill()
	//video.Error(ErrorMsg[0], 0, true)
	return 0
}

func EqualError(v1, v2, keySize uintptr)uintptr{
	kill()
	//video.Error(ErrorMsg[1], 1, true)
	return 0
}
/*
//var ErrorMsg [2][40]byte

func Init(){
	//video.CopyStr(&ErrorMsg[0], "Unhashable Type Exception")
	//video.CopyStr(&ErrorMsg[1], "Incomparable Type Exception")
}*/