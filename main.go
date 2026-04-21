package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var WIDTH float32 = 640
var HEIGHT float32 = 576

func add(a, b int) int {
	return a + b
}

type Game struct {
	gb GameBoy
}

func (g *Game) Update() error {
	if g.gb.CPU.Halted {
		return nil
	}

	// Approximate DMG CPU budget per frame (4194304 / 59.7 ~= 70224 T-cycles).
	const targetCyclesPerFrame = 70224
	cyclesThisFrame := 0

	for cyclesThisFrame < targetCyclesPerFrame && !g.gb.CPU.Halted {
		cyclesThisFrame += g.gb.Step()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	lcdc := g.gb.MMU.Read(0xFF40)
	lcdEnabled := (lcdc & 0x80) != 0 // this is the 7th bit in the LCDC register. bit 7 controls the LCD on/off state. if it is 0, the LCD is off.
	if !lcdEnabled {
		// LCD off behavior (which on DMG is displayed as a white "whiter" than color #0. here it is the same, for simplicity)
		screen.Fill(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
		return
	}

	// LCD on: render using palette/BGP
	bgp := g.gb.MMU.Read(0xFF47)
	background := colorFromBGPColor0(bgp)
	screen.Fill(background)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(WIDTH), int(HEIGHT)
}

func colorFromBGPColor0(bgp byte) color.RGBA {
	color0 := bgp & 0x03

	switch color0 {
	case 0:
		return color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	case 1:
		return color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF}
	case 2:
		return color.RGBA{R: 0x55, G: 0x55, B: 0x55, A: 0xFF}
	default:
		return color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
	}
}

func main() {
	ebiten.SetWindowSize(int(WIDTH), int(HEIGHT))
	ebiten.SetWindowTitle("GameBoy Emulator")

	gb := NewGameBoy()
	if err := gb.MMU.LoadROM("roms/minimal.gb"); err != nil {
		log.Fatal(err)
	}

	game := &Game{gb: gb}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
