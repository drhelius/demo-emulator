package memoryimpl

import (
	"fmt"

	"github.com/drhelius/demo-emulator/gb/cpu"
	"github.com/drhelius/demo-emulator/gb/input"
	"github.com/drhelius/demo-emulator/gb/util"
	"github.com/drhelius/demo-emulator/gb/video"
)

// Memory is the public type for Memory
type Memory struct {
	memoryMap []uint8 //= make([]uint8, 0x10000)
	rom       []uint8
}

// Setup Receives the rom data and intializes memory
func (m *Memory) Setup(r []uint8) {
	m.rom = r
	m.memoryMap = make([]uint8, 0x10000)
}

// Read returns the 8 bit value at the 16 bit address of the memory
func (m *Memory) Read(addr uint16) uint8 {
	if addr < 0x8000 {
		// ROM
		return m.rom[addr]
	}

	if addr >= 0xFF00 {
		switch addr {
		case 0xFF00:
			// P1
			return input.Read()
		case 0xFF07:
			// TAC
			return m.memoryMap[addr] | 0xF8
		case 0xFF0F:
			// IF
			return m.memoryMap[addr] | 0xE0
		case 0xFF41:
			// STAT
			return m.memoryMap[addr] | 0x80
		case 0xFF44:
			if video.ScreenEnabled {
				return m.memoryMap[0xFF44]
			}
			return 0x00
		case 0xFF4F:
			// VBK
			return m.memoryMap[addr] | 0xFE
		}
	}

	return m.memoryMap[addr]
}

// Write stores the 8 bit value at the 16 bit address of the memory
func (m *Memory) Write(addr uint16, value uint8) {

	if addr < 0x8000 {
		// ROM
		fmt.Printf("** Attempting to write on ROM address %X %X\n", addr, value)
		return
	}

	if addr >= 0xFF00 {
		switch addr {
		case 0xFF00:
			// P1
			input.Write(value)
		case 0xFF04:
			// DIV
			m.memoryMap[addr] = 0x00
			cpu.ResetDivCycles()
			break
		case 0xFF07:
			// TAC
			value &= 0x07
			currentTac := m.memoryMap[addr]
			if (currentTac & 0x03) != (value & 0x03) {
				m.memoryMap[0xFF05] = m.memoryMap[0xFF06]
				cpu.ResetTimaCycles()
			}
			m.memoryMap[addr] = value
		case 0xFF0F:
			// IF
			m.memoryMap[addr] = value & 0x1F
		case 0xFF40:
			// LCDC
			currentLcdc := m.memoryMap[addr]
			newLcdc := value
			m.memoryMap[addr] = newLcdc
			if !util.IsSetBit(currentLcdc, 5) && util.IsSetBit(newLcdc, 5) {
				video.ResetWindowLine()
			}
			if util.IsSetBit(newLcdc, 7) {
				video.EnableScreen()
			} else {
				video.DisableScreen()
			}
		case 0xFF41:
			// STAT
			currentStat := m.memoryMap[addr] & 0x07
			newStat := (value & 0x78) | (currentStat & 0x07)
			m.memoryMap[addr] = newStat
		case 0xFF44:
			// LY
			currentLy := m.memoryMap[addr]
			if util.IsSetBit(currentLy, 7) && !util.IsSetBit(value, 7) {
				video.DisableScreen()
			}
		case 0xFF45:
			// LYC
			currentLyc := m.memoryMap[addr]
			if currentLyc != value {
				m.memoryMap[addr] = value
				lcdc := m.memoryMap[0xFF40]
				if util.IsSetBit(lcdc, 7) {
					video.CompareLYToLYC()
				}
			}
		case 0xFF46:
			// DMA
			m.memoryMap[addr] = value
			address := uint16(value) << 8
			if address >= 0x8000 && address < 0xE000 {
				for i := uint16(0x0000); i < 0x00A0; i++ {
					m.memoryMap[0xFE00+i] = m.memoryMap[address+i]
				}
			}
		case 0xFFFF:
			// IE
			m.memoryMap[addr] = value & 0x1F
			break
		default:
			m.memoryMap[addr] = value
		}
	}
}
