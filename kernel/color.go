package color

type VGAColor uint8

const (
	BLACK VGAColor = iota
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

func MakeColor(foreground, background VGAColor) VGAColor {
	return (background << 4) | (foreground & 15)
}

func (color VGAColor) Blink() VGAColor {
	return color | BLINK
}

func (color VGAColor) Bright() VGAColor {
	return color | BRIGHT
}

func (color VGAColor) Dark() VGAColor {
	return color & (^BRIGHT & 255)
}

func (c VGAColor)BGRA32()BGRA32{
	if c&BRIGHT == 0 {
		return BGRA32{R: uint8((c&RED)<<5), G: uint8((c&GREEN)<<6), B: uint8((c&BLUE)<<7)}
	}else{
		return BGRA32{R: uint8(((c&RED)<<6)-1), G: uint8(((c&GREEN)<<7)-1), B: uint8(((c&BLUE)<<8)-1)}
	}
}

type RGBA32 struct{
	R, G, B, A uint8
}

func (c RGBA32)VGAColor()VGAColor{
	if int(c.B) + int(c.G) + int(c.R) >= 384 {
		return VGAColor(((c.B>>7) & 1) | ((c.G>>6) & 2) | ((c.R>>5) & 4)).Bright()
	}else{
		return VGAColor(((c.B>>7) & 1) | ((c.G>>6) & 2) | ((c.R>>5) & 4)).Bright()
	}
}

func (c RGBA32)BGRA32()BGRA32{
	return BGRA32{B: c.B, G: c.G, R: c.R, A: c.A}
}

type BGRA32 struct{
	B, G, R, A uint8
}

func (c BGRA32)VGAColor()VGAColor{
	if int(c.B) + int(c.G) + int(c.R) >= 384 {
		return VGAColor(((c.B>>7) & 1) | ((c.G>>6) & 2) | ((c.R>>5) & 4)).Bright()
	}else{
		return VGAColor(((c.B>>7) & 1) | ((c.G>>6) & 2) | ((c.R>>5) & 4)).Bright()
	}
}