package color

type Color uint8

const (
	BLACK Color = iota
	BLUE
	GREEN
	CYAN
	RED
	MAGENTA
	BROWN
	LIGHT_GRAY
	DARK_GRAY
	LIGHT_BLUE
	LIGHT_GREEN
	LIGHT_CYAN
	LIGHT_RED
	LIGHT_MAGENTA
	YELLOW
	WHITE
	BLINK  = 128
	BRIGHT = 8
)

func MakeColor(foreground, background Color) Color {
	return (background << 4) | (foreground & 15)
}

func Blink(color Color) Color {
	return color | BLINK
}

func Bright(color Color) Color {
	return color | BRIGHT
}

func Dark(color Color) Color {
	return color & (^BRIGHT & 255)
}
