package main

import (
	"bytes"
	"testing"
)

func TestBuildMinimalROMShape(t *testing.T) {
	rom := buildMinimalROM()

	if got, want := len(rom), romSize; got != want {
		t.Fatalf("ROM size = %d, want %d", got, want)
	}

	entryBytes := []byte{0x00, 0xC3, 0x50, 0x01}
	if !bytes.Equal(rom[entryPoint:entryPoint+len(entryBytes)], entryBytes) {
		t.Fatalf("entry bytes mismatch: got % x, want % x", rom[entryPoint:entryPoint+len(entryBytes)], entryBytes)
	}

	initBytes := []byte{0x3E, 0x03, 0xE0, 0x47, 0x3E, 0x91, 0xE0, 0x40, 0x76}
	if !bytes.Equal(rom[cartInitAddress:cartInitAddress+len(initBytes)], initBytes) {
		t.Fatalf("init bytes mismatch: got % x, want % x", rom[cartInitAddress:cartInitAddress+len(initBytes)], initBytes)
	}
}
