package mbcs

import "fmt"

// RomOnly is the public type for Memory
type RomOnly struct {
	memoryMap []uint8
	rom       []uint8
}

// GetMemoryMap returns the memory array
func (m *RomOnly) GetMemoryMap() []uint8 {
	return m.memoryMap
}

// GetROM returns the rom array
func (m *RomOnly) GetROM() []uint8 {
	return m.rom
}

// Setup Receives the rom data and intializes memory
func (m *RomOnly) Setup(r []uint8) {
	m.rom = r
	m.memoryMap = make([]uint8, 0x10000)

	for i := 0; i < 0x100; i++ {
		m.memoryMap[0xFF00+i] = initialValuesForFFXX[i]
	}
}

// Read returns the 8 bit value at the 16 bit address of the memory
func (m *RomOnly) Read(addr uint16) uint8 {
	switch {
	case addr < 0x8000:
		// ROM
		return m.rom[addr]
	case addr >= 0xFF00:
		// IO Registers
		return ReadIO(addr, m.memoryMap)
	}
	return m.memoryMap[addr]
}

// Write stores the 8 bit value at the 16 bit address of the memory
func (m *RomOnly) Write(addr uint16, value uint8) {
	switch {
	case addr < 0x8000:
		// ROM
		fmt.Printf("** Attempting to write on ROM address %X %X\n", addr, value)
	case (addr >= 0xC000) && (addr < 0xFE00):
		WriteCommon(addr, value, m.memoryMap)
	case addr >= 0xFF00:
		// IO Registers
		WriteIO(addr, value, m.memoryMap, m)
	default:
		m.memoryMap[addr] = value
	}
}
