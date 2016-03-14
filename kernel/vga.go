package vga

import (
	"video"
	"color"
	"unsafe"
)

type framebuffer struct{
	x, y int
	termColor color.VGAColor
	vidMem *[25][80]VidEntry
}

type VidEntry struct {
	Char  byte
	Color color.VGAColor
}

var fb = framebuffer{
	termColor: color.MakeColor(color.LIGHT_GRAY, color.BLACK),
	vidMem: (*[25][80]VidEntry)(unsafe.Pointer(uintptr(0x000B8000))),
}

func init(){
	video.SetPrinter(&fb)
}

func SetFrameBuffer(){
	fb.Clear()
	video.SetPrinter(&fb)
}

func (f *framebuffer) SetColor(c color.RGBA32) {
	if int(c.B) + int(c.G) + int(c.R) >= 384 {
		f.termColor = color.VGAColor(((c.B>>7) & 1) | ((c.G>>6) & 2) | ((c.R>>5) & 4)).Bright()
	}else{
		f.termColor = color.VGAColor(((c.B>>7) & 1) | ((c.G>>6) & 2) | ((c.R>>5) & 4)).Bright()
	}
}

func (f *framebuffer) PutChar(c rune) {
	if c == '\n' {
		f.vidMem[f.y][f.x] = f.MakeEntry(0)
		f.x = 0
		f.y++
	} else if c == '\t' {
		f.vidMem[f.y][f.x].Color = f.termColor
		f.x += 4 - (f.x % 4)
	} else if c == '\b' {
		f.vidMem[f.y][f.x].Color = f.termColor
		f.x--
	} else {
		if f.y > 24 {
			f.Scroll()
		}
		f.vidMem[f.y][f.x] = f.MakeEntry(byte(c))
		f.x++
		if f.x > 80 {
			f.x = 0
			f.y++
		}
	}
	f.updateCursor()
}

func (f *framebuffer) Clear(){
	*f.vidMem = [25][80]VidEntry{}
}

func (f *framebuffer) updateCursor() {
	if f.y > 24 {
		f.Scroll()
	}

	f.vidMem[f.y][f.x].Color ^= color.MakeColor(color.WHITE, color.WHITE)
}

func (f *framebuffer)MakeEntry(char byte) VidEntry {
	return VidEntry{Char: char, Color: f.termColor}
}

func (f *framebuffer) Scroll() {
	for yVal := 1; yVal < 25; yVal++ {
		f.vidMem[yVal-1] = f.vidMem[yVal]
	}
	f.y = 24
}