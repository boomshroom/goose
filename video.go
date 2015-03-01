package video

import (
	"color"
	"ptr"
	//"asm"
)

var x, y int
var termColor color.Color
var vidMem uintptr
//var ioPort uint16
type VidEntry struct{
	Char byte
	Color color.Color
}

func vidPtr() *[25][80]VidEntry {
	return (*[25][80]VidEntry)(ptr.GetAddr(vidMem))
}

func Init() {
	vidMem = 0xB8000
	termColor = color.MakeColor(color.LIGHT_GRAY, color.BLACK)
	initErrs()
	//ioPort = *(*uint16)(ptr.GetAddr(0x46C))
}

func MakeEntry(char byte)VidEntry{
	return VidEntry{Char: char, Color: termColor}
}

func SetColor(c color.Color) {
	termColor = c
}

func Print(line string) {
	for i := 0; i < len(line); i++ {
		PutChar(rune(line[i]))
	}
}

func Println(line string) {
	Print(line)
	NL()
}

func PrintHex(num uint64, caps, prefix, newline bool, digits int8) {
	if prefix {
		if caps {
			Print("0X")
		} else {
			Print("0x")
		}
	}
	nonzero := false
	for i := int8(16); i > -1; i-- {
		digit := uint8(num>>uint(i*4)) & 0xF
		if digit != 0 || nonzero || i < digits {
			nonzero = true
			PutChar(Int4ToHex(digit, caps))
		}
	}
	if newline {
		NL()
	}
}

func PrintUint(num uint64){
	PrintHex(num, false, true, false, 8)
}

func PrintBool(b bool){
	if b{
		Print("true")
	}else{
		Print("false")
	}
}

func Int4ToHex(digit uint8, caps bool) rune {
	if digit < 10 {
		return rune(digit + '0')
	} else if caps {
		return rune(digit - 0xA + 'A')
	} else {
		return rune(digit - 0xA + 'a')
	}
}

func NL() {
	vidPtr()[y][x] = MakeEntry(0)
	x = 0
	y++
}

func PutChar(c rune) {
	if c == '\n' {
		NL()
		updateCursor()
	} else if c == '\t' {
		vidPtr()[y][x].Color = termColor
		x += 4 - (x % 4)
		updateCursor()
	} else if c == '\b'{
		vidPtr()[y][x].Color = termColor
		x--
		updateCursor()
	} else{
		PutCharRaw(c)
	}
}
func PutCharRaw(c rune) {
	if y > 24 {
		Scroll()
	}
	vidPtr()[y][x] = MakeEntry(byte(c))
	x++
	if x > 80 {
		x = 0
		y++
	}
	updateCursor()
}
var check = true

func updateCursor(){
	//vidPtr()[y][x] = MakeEntry(' ')
	vidPtr()[y][x].Color ^= color.MakeColor(color.WHITE,color.WHITE)
/*
	pos:= uint16(y)*80 + uint16(x)
	asm.OutportB(ioPort, 0x0F)
	asm.OutportB(ioPort+1, uint8(pos))
	asm.OutportB(ioPort, 0x0F)
	asm.OutportB(ioPort+1, uint8(pos>>8))*/
}

func Clear() {
	for i := 0; i < 80; i++ {
		for j := 0; j < 25; j++ {
			vidPtr()[j][i] = VidEntry{Char: 0, Color: termColor}
		}
	}
	x = 0
	y = 0
	updateCursor()
}

func MoveCursor(dx, dy int){
	vidPtr()[y][x].Color = termColor
	x += dx
	y += dy
	updateCursor()
}

func CopyStr(array *[40]byte, str string) {
	if len(str) > 40{
		Error(ErrorMsg[0], len(str), true)
	}
	for i := 0; i < len(str); i++ {
		array[i] = str[i]
	}
}

var ErrorMsg [13][40]byte

func initErrs(){
	CopyStr(&ErrorMsg[0], "Error message too long")
	CopyStr(&ErrorMsg[1], "Slice Index out of Bounds Exception")
	CopyStr(&ErrorMsg[2], "Array Index out of Bounds Exception")
	CopyStr(&ErrorMsg[3], "String Index out of Bounds Exception")
	CopyStr(&ErrorMsg[4], "Slice Slice out of Bounds Exception")
	CopyStr(&ErrorMsg[5], "Array Slice out of Bounds Exception")
	CopyStr(&ErrorMsg[6], "String Slice out of Bounds Exception")
	CopyStr(&ErrorMsg[7], "Nil Pointer Exception")
	CopyStr(&ErrorMsg[8], "Make slice out of bounds")
	CopyStr(&ErrorMsg[9], "Make map out of bounds")
	CopyStr(&ErrorMsg[10], "Make chan out of bounds")
	CopyStr(&ErrorMsg[11], "Division By Zero Exception")
	CopyStr(&ErrorMsg[12], "Unknown Exception")
	
}

func ErrCode(code int32){
	if code > 10{
		Error(ErrorMsg[12], int(code), true)
	}
	Error(ErrorMsg[code+1], int(code), true)
}

func Error(errorMsg [40]byte, errorCode int, halt bool) {
	Print("ERROR: ")
	if errorCode != -1 {
		PrintHex(uint64(errorCode), false, true, false, 2)
		PutChar(' ')
	}
	for i := 0; i < 40; i++ {
		PutChar(rune(errorMsg[i]))
	}
	NL()
	if halt {
		Println("System Halted.")
		for {
		}
	}
}

func Scroll() {
	for yVal := 1; yVal < 25; yVal++ {
		vidPtr()[yVal-1] = vidPtr()[yVal]
	}
	y = 24
}