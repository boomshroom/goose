package video

import (
	//"asm"
	"proc"
	"runtime"
	"color"
)

func SetPrinter(p Printer) {
	printer = p
}

func PrintCurrent() {
	Print("Attempting to kill proc ")
	PrintHex(proc.CurrentID, false, true, true, 0)
}

var printer Printer

type Printer interface {
	PutChar(rune)
	SetColor(color.RGBA32)
}

func init() {
	runtime.ErrorPrint = errorMsg
}

func Print(line string) {
	for _, ch := range line {
		PutChar(ch)
	}
}

func PutChar(ch rune){
	printer.PutChar(ch)
}

func Println(line string) {
	Print(line)
	PutChar('\n')
}

func NL(){
	PutChar('\n')
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
		PutChar('\n')
	}
}

func PrintUint(num uint64) {
	PrintHex(num, false, true, false, 8)
}

func PrintBool(b bool) {
	if b {
		Print("true")
	} else {
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

func Error(errorMsg string, errorCode int, halt bool) {
	Print("ERROR: ")
	if errorCode != -1 {
		PrintHex(uint64(errorCode), false, true, false, 2)
		PutChar(' ')
	}
	Println(errorMsg)
	if halt {
		Println("System Halted.")
		for {
		}
	}
}

//extern __unwind_stack
func unwindStack()

func errorMsg(err string) {
	Println(err)
	unwindStack()
	Println("System Halted.")
	for {
	}
}
