package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	romSize            = 0x8000
	entryPoint         = 0x0100
	cartInitAddress    = 0x0150
	defaultOutputPath  = "roms/minimal.gb"
)

func buildMinimalROM() []byte {
	rom := make([]byte, romSize) // 32KB

	copy(rom[entryPoint:], []byte{
		0x00,             // NOP
		0xC3, 0x50, 0x01, // JP 0x0150
	})

	copy(rom[cartInitAddress:], []byte{
		0x3E, 0x03, // LD A,0x03
		0xE0, 0x47, // LDH (0x47),A
		0x3E, 0x91, // LD A,0x91
		0xE0, 0x40, // LDH (0x40),A
		0x76, // HALT
	})

	return rom
}

func writeMinimalROM(outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, buildMinimalROM(), 0o644); err != nil {
		return fmt.Errorf("write ROM file: %w", err)
	}

	return nil
}

func main() {
	if err := writeMinimalROM(defaultOutputPath); err != nil {
		log.Fatal(err)
	}
}
