package main

import "fmt"

type Pixel rune

const (
	Lit  Pixel = '#'
	Dark Pixel = ' '
)

type CRT struct {
	width  int
	height int
	lines  []Pixel

	cursorPos int
	spritePos int
}

func NewCRT(width, height int) *CRT {
	crt := CRT{
		width:     width,
		height:    height,
		lines:     make([]Pixel, width*height),
		cursorPos: 0,
	}

	// Initialize screen dark
	crt.Clear(Dark)
	return &crt
}

func (crt *CRT) Clear(pixel Pixel) {
	for i := 0; i < crt.height*crt.width; i++ {
		crt.lines[i] = pixel
	}
}

func (crt *CRT) Display() {
	for i := 0; i < crt.width*crt.height; i += crt.width {
		fmt.Println(string(crt.lines[i : i+crt.width]))
	}
}

func (crt *CRT) SetSpritePos(pos int) { crt.spritePos = pos }
func (crt *CRT) SpritePos() int       { return crt.spritePos }
func (crt *CRT) CursorPos() int       { return crt.cursorPos }

func (crt *CRT) Tick() {
	// Rest position after screen has been drawn completely
	if crt.cursorPos == crt.width*crt.height {
		crt.cursorPos = 0
	}

	d := crt.spritePos - (crt.cursorPos % crt.width)
	if d == -1 || d == 0 || d == 1 {
		crt.lines[crt.cursorPos] = Lit
	} else {
		crt.lines[crt.cursorPos] = Dark
	}
	crt.cursorPos += 1
}
