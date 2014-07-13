package video

//extern __unsafe_get_addr
func getAddr(addr uint32) *[25][80][2]byte

var x, y int
var Vidmem *[25][80][2]byte

func Init() {
	Vidmem = getAddr(0xB8000)
}

func Print(line string) {
	for i := 0; i < len(line); i++ {
		PutChar(rune(line[i]))
	}
}

func PutChar(c rune) {
	if c == '\n' {
		x = 0
		y++
	} else {
		Vidmem[y][x][0] = byte(c)
		Vidmem[y][x][1] = 0x07
		x++
		if x > 80 {
			x = 0
			y++
		}
	}
	// todo: no scrolling support
}

func Clear() {
	for i := 0; i < 80; i++ {
		for j := 0; j < 25; j++ {
			Vidmem[j][i][0] = 0
			Vidmem[j][i][1] = 0
		}
	}
	x = 0
	y = 0
}

func Error(i int){
	if i < 3 {
		Print("ERROR: index out of range")
		return
	}
	Print("ERROR")
	 //on a phone, can`t backspace
}
