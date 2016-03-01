package main

//extern __scan_char
func scan(*int)

//extern __register_interupt
func register(*struct{i int64; f func()})

//extern __start_proc
func startApp(*string)

//extern __int_ret
func IntRet()

var syscall = struct{i int64; f func()}{i: 0, f: Int}

var i int

var proc = "proc2"

func main() {
	println("Hello from userspace!")
	//scan(&i)
	//register(&syscall)
	/*prev := 0
	shift := false
	for {
		if i == 0x2a || i == 0x36{
			i = 0
			shift = true
		}else if i == 0xaa || i == 0xb6 {
			i = 0
			shift = false
		}else if i != 0 && i != prev && i < 128 {
			char := kbdus[i&127]

			prev = i
			i = 0
			if shift && char >= 'a' && char <= 'z'{
				char -= 'a'-'A'
			}else if shift{
				char = shifted[char]
			}
			print(char)
		} else if i == 0 && prev != 0 {
			prev = 0
		}
	}*/
	startApp(&proc)
}

func Int(){
	println("keyboard message recieved!")
	IntRet()
}

var kbdus [128]uint8 = [128]uint8{
	0, 27, '1', '2', '3', '4', '5', '6', '7', '8', /* 9 */
	'9', '0', '-', '=', '\b', /* Backspace */
	'\t',               /* Tab */
	'q', 'w', 'e', 'r', /* 19 */
	't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\n', /* Enter key */
	0,                                                /* 29   - Control */
	'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', /* 39 */
	'\'', '`', 0, /* Left shift */
	'\\', 'z', 'x', 'c', 'v', 'b', 'n', /* 49 */
	'm', ',', '.', '/', 0, /* Right shift */
	'*',
	0,   /* Alt */
	' ', /* Space bar */
	0,   /* Caps lock */
	0,   /* 59 - F1 key ... > */
	0, 0, 0, 0, 0, 0, 0, 0,
	0, /* < ... F10 */
	0, /* 69 - Num lock*/
	0, /* Scroll Lock */
	0, /* Home key */
	0, /* Up Arrow */
	0, /* Page Up */
	'-',
	0, /* Left Arrow */
	0,
	0, /* Right Arrow */
	'+',
	0, /* 79 - End key*/
	0, /* Down Arrow */
	0, /* Page Down */
	0, /* Insert Key */
	0, /* Delete Key */
	0, 0, 0,
	0, /* F11 Key */
	0, /* F12 Key */
	0, /* All other keys are undefined */
}

var shifted [127]uint8 = [127]uint8{
	'1': '!',
	'2': '@',
	'3': '#',
	'4': '$',
	'5': '%',
	'6': '^',
	'7': '&',
	'8': '*',
	'9': '(',
	'0': ')',
	'-': '_',
	'=': '+',
	'`': '~',
	'[': '{',
	']': '}',
	'\\': '|',
	';': ':',
	'\'': '"',
	',': '<',
	'.': '>',
	'/': '?',
}