package main

import (
	"fmt"
	"os"
)

var opcodes = [5]string{"0x00", "0xC3", "0x3E", "0xE0", "0x76"}

type CPU struct {
	A      byte   // accumulator register - used for arithmetic and logical operations
	PC     uint16 // program counter - points to the next instruction to execute
	Halted bool   // indicates if the CPU is halted - if true, the CPU will not execute any instructions
}

type MMU struct {
	memory [0x10000]byte // 64KB of memory - 0x10000 is 65536 bytes in hexadecimal - in 'go' this `memory [0x10000]byte` denotes an array of length 65536
}

func (m *MMU) Read(address uint16) byte {
	return m.memory[address] // return the value at the address
}

func (m *MMU) Write(address uint16, value byte) {
	if address < 0x8000 {
		return // ROM is read-only - ignore writes
	}
	// When we use the complete MMU struct written above, we should have many more 'if' statements to control where the value is written
	m.memory[address] = value
}

func (m *MMU) LoadROM(path string) error {
	rom, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read ROM %q: %w", path, err)
	}

	if len(rom) > 0x8000 {
		rom = rom[:0x8000]
	}

	copy(m.memory[:0x8000], rom)
	return nil
}

type GameBoy struct {
	CPU CPU
	MMU MMU
}

func (gb *GameBoy) Step() int {
	// FETCH
	opcode := gb.MMU.Read(gb.CPU.PC)
	gb.CPU.PC++

	// DECODE
	// inside each case: EXECUTE
	switch opcode {
	case 0x00: // NOP
		return 4

	case 0x3E: // LD A, n
		// program counter points to a memory address. take the value on that address (n) and put it in the accumulator register (A)
		gb.CPU.A = gb.MMU.Read(gb.CPU.PC)
		// increment program counter so it points to the next memory address
		gb.CPU.PC++
		return 8

	case 0x76: // HALT - don't do anything for 4 cpu cycles
		gb.CPU.Halted = true
		return 4

	case 0xC3: // JP nn - jump to address nn. the address that the cpu should jump to is stored in the next two addresses.
		lo := gb.MMU.Read(gb.CPU.PC)           // e.g 0x50
		hi := gb.MMU.Read(gb.CPU.PC + 1)       // e.g. 0x01
		gb.CPU.PC = uint16(hi)<<8 | uint16(lo) // this returns 0x0150. analysis:
		/*
			|-| hi = 0x01
			|-| lo = 0x50
			|-| uint16(hi)<<8 -> this turns 0x01 to 0x0100 (left shift by 8 bits) ->  this will become the left part of the bitwise OR operation between it and the value at the next address, which in this example is 0x50.
			|-| uint16(lo) -> this is 0x50
			|-| | -> bitwise OR operation
			|-| gb.CPU.PC = 0x0150 (0x01 comes first because it is first in memory: this is the little-endian architecture)
			**/

		return 16

	case 0xE0: // LDH (n), A
		offset := gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		gb.MMU.Write(0xFF00+uint16(offset), gb.CPU.A)
		return 12

	default:
		panic(fmt.Sprintf("Unknown opcode: 0x%02X", opcode))
	}
}

func NewGameBoy() GameBoy {
	return GameBoy{
		CPU: CPU{
			PC: 0x0100,
		},
	}
}
