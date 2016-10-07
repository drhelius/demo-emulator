package memory

import (
	"fmt"

	"github.com/drhelius/demo-emulator/gb/input"
	"github.com/drhelius/demo-emulator/gb/timer"
	"github.com/drhelius/demo-emulator/gb/util"
	"github.com/drhelius/demo-emulator/gb/video"
)

var (
	memoryMap = make([]uint8, 0x10000)
	rom       []uint8
)

// SetupROM Receives the rom data
func SetupROM(r []uint8) {
	rom = r
}

// Read returns the 8 bit value at the 16 bit address of the memory
func Read(addr uint16) uint8 {
	if addr < 0x8000 {
		// ROM
		return rom[addr]
	}

	if addr >= 0xFF00 {
		switch addr {
		case 0xFF00:
			// P1
			return input.Read()
		case 0xFF07:
			// TAC
			return memoryMap[addr] | 0xF8
		case 0xFF0F:
			// IF
			return memoryMap[addr] | 0xE0
		case 0xFF41:
			// STAT
			return memoryMap[addr] | 0x80
		//case 0xFF44:
		//    return (m_pVideo->IsScreenEnabled() ? m_pMemory->Retrieve(0xFF44) : 0x00);
		case 0xFF4F:
			// VBK
			return memoryMap[addr] | 0xFE
		}
	}

	return memoryMap[addr]
}

// Write stores the 8 bit value at the 16 bit address of the memory
func Write(addr uint16, value uint8) {

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
			memoryMap[addr] = 0x00
			timer.DivCycles = 0
			break
		case 0xFF07:
			// TAC
			value &= 0x07
			currentTac := memoryMap[addr]
			if (currentTac & 0x03) != (value & 0x03) {
				memoryMap[0xFF05] = memoryMap[0xFF06]
				timer.TimaCycles = 0
			}
			memoryMap[addr] = value
		case 0xFF0F:
			// IF
			memoryMap[addr] = value & 0x1F
		case 0xFF40:
			// LCDC
			currentLcdc := memoryMap[addr]
			newLcdc := value
			memoryMap[addr] = newLcdc
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
			currentStat := memoryMap[addr] & 0x07
			newStat := (value & 0x78) | (currentStat & 0x07)
			memoryMap[addr] = newStat
		case 0xFF44:
			// LY
			currentLy := memoryMap[addr]
			if util.IsSetBit(currentLy, 7) && !util.IsSetBit(value, 7) {
				video.DisableScreen()
			}
		case 0xFF45:
			// LYC
			currentLyc := memoryMap[addr]
			if currentLyc != value {
				memoryMap[addr] = value
				lcdc := memoryMap[0xFF40]
				if util.IsSetBit(lcdc, 7) {
					video.CompareLYToLYC()
				}
			}
		case 0xFF46:
			// DMA
			memoryMap[addr] = value
			address := uint16(value) << 8
			if address >= 0x8000 && address < 0xE000 {
				for i := uint16(0x0000); i < 0x00A0; i++ {
					memoryMap[0xFE00+i] = memoryMap[address+i]
				}
			}
		case 0xFFFF:
			// IE
			memoryMap[addr] = value & 0x1F
			break
		default:
			memoryMap[addr] = value
		}
	}
}
