package main

import (
	"os"
)

func generateRom() {
	rom := make([]byte, 0x8000) // 32KB

	copy(rom[0x0100:], []byte{
		0x00,             // NOP
		0xC3, 0x50, 0x01, // JP 0x0150
	})

	copy(rom[0x0150:], []byte{
		0x3E, 0x03, // LD A,0x03
		0xE0, 0x47, // LDH (0x47),A
		0x3E, 0x91, // LD A,0x91
		0xE0, 0x40, // LDH (0x40),A
		0x76, // HALT
	})

	_ = os.MkdirAll("roms", 0o755)
	_ = os.WriteFile("roms/minimal.gb", rom, 0o644)
}

func main() {
	generateRom()
}
