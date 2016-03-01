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

func (color Color) Blink() Color {
	return color | BLINK
}

func (color Color) Bright() Color {
	return color | BRIGHT
}

func (color Color) Dark() Color {
	return color & (^BRIGHT & 255)
}
