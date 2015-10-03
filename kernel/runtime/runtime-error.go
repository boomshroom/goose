package runtime

var errorMsgs [11]string = [...]string{
	"Slice Index out of Bounds Error",
	"Array Index out of Bounds Error",
	"String Index out of Bounds Error",
	"Slice Slice out of Bounds Error",
	"Array Slice out of Bounds Error",
	"String Slice out of Bounds Error",
	"Nil Pointer Error",
	"Make slice out of bounds",
	"Make map out of bounds",
	"Make chan out of bounds",
	"Division By Zero Error",
}

func RuntimeError(code int32){
	if code >= int32(len(errorMsgs)){
		panicString("Unknown Runtime Error")
	}else{
		panicString(errorMsgs[code])
	}
}