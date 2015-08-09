package video

import (
	"color"
	"unsafe"
	//"asm"
	"runtime"
)

var x, y int
var termColor color.Color  = color.MakeColor(color.LIGHT_GRAY, color.BLACK)
var vidMem *[25][80]VidEntry = (*[25][80]VidEntry)(unsafe.Pointer(uintptr(0x000B8000)))
//var ioPort uint16
type VidEntry struct{
	Char byte
	Color color.Color
}

func init() {
	//ioPort = *(*uint16)(ptr.GetAddr(0x46C))
	Clear()
	runtime.ErrorPrint = errorMsg
}

func MakeEntry(char byte)VidEntry{
	return VidEntry{Char: char, Color: termColor}
}

func SetColor(c color.Color) {
	termColor = c
}

func Print(line string) {
	for _, ch := range line {
		PutChar(ch)
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
	vidMem[y][x] = MakeEntry(0)
	x = 0
	y++
}

func PutChar(c rune) {
	if c == '\n' {
		NL()
		updateCursor()
	} else if c == '\t' {
		vidMem[y][x].Color = termColor
		x += 4 - (x % 4)
		updateCursor()
	} else if c == '\b'{
		vidMem[y][x].Color = termColor
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
	vidMem[y][x] = MakeEntry(byte(c))
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
	vidMem[y][x].Color ^= color.MakeColor(color.WHITE,color.WHITE)
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
			vidMem[j][i] = VidEntry{Char: 0, Color: termColor}
		}
	}
	x = 0
	y = 0
	updateCursor()
}

func MoveCursor(dx, dy int){
	vidMem[y][x].Color = termColor
	x += dx
	y += dy
	updateCursor()
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

func errorMsg(err string){
	Println(err)
	Println("System Halted.")
	for {}
}

func Scroll() {
	for yVal := 1; yVal < 25; yVal++ {
		vidMem[yVal-1] = vidMem[yVal]
	}
	y = 24
}