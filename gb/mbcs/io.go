package mbcs

import (
	"github.com/drhelius/demo-emulator/gb/cpu"
	"github.com/drhelius/demo-emulator/gb/input"
	"github.com/drhelius/demo-emulator/gb/mapper"
	"github.com/drhelius/demo-emulator/gb/util"
	"github.com/drhelius/demo-emulator/gb/video"
)

// ReadIO returns the 8 bit value at the 16 bit address of the memory
func ReadIO(addr uint16, mem []uint8) uint8 {
	switch addr {
	case 0xFF00:
		// P1
		return input.Read()
	case 0xFF07:
		// TAC
		return mem[addr] | 0xF8
	case 0xFF0F:
		// IF
		return mem[addr] | 0xE0
	case 0xFF41:
		// STAT
		return mem[addr] | 0x80
	case 0xFF44:
		if video.ScreenEnabled {
			return mem[0xFF44]
		}
		return 0x00
	case 0xFF4F:
		// VBK
		return mem[addr] | 0xFE
	}

	return mem[addr]
}

// WriteIO stores the 8 bit value at the 16 bit address of the memory
func WriteIO(addr uint16, value uint8, mem []uint8, m mapper.Mapper) {
	switch addr {
	case 0xFF00:
		// P1
		input.Write(value)
	case 0xFF04:
		// DIV
		cpu.ResetDivCycles()
		mem[addr] = 0x00
	case 0xFF07:
		// TAC
		value &= 0x07
		currentTac := mem[addr]
		if (currentTac & 0x03) != (value & 0x03) {
			cpu.ResetTimaCycles()
			mem[0xFF05] = mem[0xFF06]
		}
		mem[addr] = value
	case 0xFF0F:
		// IF
		mem[addr] = value & 0x1F
	case 0xFF40:
		// LCDC
		mem[addr] = value
		if util.IsSetBit(value, 7) {
			video.EnableScreen()
		} else {
			video.DisableScreen()
		}
	case 0xFF41:
		// STAT
		currentStat := mem[addr] & 0x07
		newStat := (value & 0x78) | (currentStat & 0x07)
		mem[addr] = newStat
		lcdc := mem[0xFF40]
		if util.IsSetBit(lcdc, 7) {
			video.CompareLYToLYC()
		}
		mem[addr] = value
	case 0xFF44:
		// LY
		currentLy := mem[addr]
		if util.IsSetBit(currentLy, 7) && !util.IsSetBit(value, 7) {
			video.DisableScreen()
		}
		mem[addr] = value
	case 0xFF45:
		// LYC
		currentLyc := mem[addr]
		if currentLyc != value {
			mem[addr] = value
			lcdc := mem[0xFF40]
			if util.IsSetBit(lcdc, 7) {
				video.CompareLYToLYC()
			}
		}
	case 0xFF46:
		// DMA
		mem[addr] = value
		address := uint16(value) << 8
		if address >= 0x8000 && address < 0xE000 {
			for i := uint16(0x0000); i < 0x00A0; i++ {
				m.Write(0xFE00+i, m.Read(address+i))
			}
		}
	case 0xFFFF:
		// IE
		mem[addr] = value & 0x1F
	default:
		mem[addr] = value
	}
}
