package runtime

var ErrorPrint func(string) = func(string){
	for{}
}

func panicString(s string){
	ErrorPrint(s)
}