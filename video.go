package video

import (
	"color"
	"ptr"
)

var x, y int
var termColor color.Color
var vidMem uintptr

func vidPtr() *[25][80][2]byte {
	return (*[25][80][2]byte)(ptr.GetAddr(vidMem))
}

func Init() {
	vidMem = 0xB8000
	termColor = color.MakeColor(color.LIGHT_GRAY, color.BLACK)
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
	vidPtr()[y][x][0] = 0
	vidPtr()[y][x][1] = 0
	x = 0
	y++
}

func PutChar(c rune) {
	if c == '\n' {
		NL()
		updateCursor()
	} else if c == '\t' {
		x += 4 - (x % 4)
		updateCursor()
	} else if c == '\b'{
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
	vidPtr()[y][x][0] = byte(c)
	vidPtr()[y][x][1] = byte(termColor)
	x++
	if x > 80 {
		x = 0
		y++
	}
	updateCursor()
}

func updateCursor(){
	vidPtr()[y][x][0] = byte('_')
	vidPtr()[y][x][1] = byte(termColor | 0x80)
}

func Clear() {
	for i := 0; i < 80; i++ {
		for j := 0; j < 25; j++ {
			vidPtr()[j][i][0] = 0
			vidPtr()[j][i][1] = 0
		}
	}
	x = 0
	y = 0
	updateCursor()
}

func MoveCursor(dx, dy int){
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

func Scroll() {
	for yVal := 1; yVal < 25; yVal++ {
		vidPtr()[yVal-1] = vidPtr()[yVal]
	}
	y = 24
}