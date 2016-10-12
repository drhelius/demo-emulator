package mbcs

import "fmt"

// MBC1 is the public type for Memory
type MBC1 struct {
	memoryMap       []uint8
	rom             []uint8
	ram             []uint8
	mode            uint8
	romBank         uint16
	ramBank         uint16
	ramEnabled      bool
	romBankHighBits uint8
	higherROMBank   uint16
	higherRAMBank   uint16
	ramSize         uint8
}

// GetMemoryMap returns the memory array
func (m *MBC1) GetMemoryMap() []uint8 {
	return m.memoryMap
}

// GetROM returns the rom array
func (m *MBC1) GetROM() []uint8 {
	return m.rom
}

// Setup Receives the rom data and intializes memory
func (m *MBC1) Setup(r []uint8) {
	m.rom = r
	m.memoryMap = make([]uint8, 0x10000)
	m.ram = make([]uint8, 0x8000)

	for i := 0; i < 0x100; i++ {
		m.memoryMap[0xFF00+i] = initialValuesForFFXX[i]
	}

	m.romBank = 1
	m.ramSize = m.rom[0x149]
	fmt.Printf("the size of ther RAM is %X\n", m.ramSize)
	switch m.ramSize {
	case 0x00:
		fallthrough
	case 0x01:
		fallthrough
	case 0x02:
		m.higherRAMBank = 0x00
	case 0x04:
		m.higherRAMBank = 0x0F
	default:
		m.higherRAMBank = 0x03
		break
	}
	fmt.Printf("the higher RAM bank is %X\n", m.higherRAMBank)
	m.higherROMBank = uint16(max(pow2Ceil(len(m.rom)/0x4000), 2) - 1)
	fmt.Printf("the higher ROM bank is %X\n", m.higherROMBank)
}

// Read returns the 8 bit value at the 16 bit address of the memory
func (m *MBC1) Read(addr uint16) uint8 {
	switch {
	case (addr >= 0x0000) && (addr < 0x4000):
		// ROM bank 0
		return m.rom[addr]
	case (addr >= 0x4000) && (addr < 0x8000):
		// ROM bank X
		return m.rom[(addr-0x4000)+(m.romBank*0x4000)]
	case (addr >= 0xA000) && (addr < 0xC000):
		// RAM bank
		if m.ramEnabled {
			if m.mode == 0 {
				return m.ram[addr-0xA000]
			}
			return m.ram[(addr-0xA000)+(m.ramBank*0x2000)]
		}
		fmt.Printf("** Attempting to read from disabled RAM %X\n", addr)
		return 0xFF
	case addr >= 0xFF00:
		// IO Registers
		return ReadIO(addr, m.memoryMap)
	}
	return m.memoryMap[addr]
}

// Write stores the 8 bit value at the 16 bit address of the memory
func (m *MBC1) Write(addr uint16, value uint8) {
	switch {
	case (addr >= 0x0000) && (addr < 0x2000):
		//fmt.Printf("** write %X %X\n", addr, value)
		if m.ramSize > 0 {
			m.ramEnabled = ((value & 0x0F) == 0x0A)
			//fmt.Println("** ramEnabled ", m.ramEnabled)
		}
	case (addr >= 0x2000) && (addr < 0x4000):
		//fmt.Printf("** write %X %X\n", addr, value)
		if m.mode == 0 {
			m.romBank = uint16((value & 0x1F) | (m.romBankHighBits << 5))
		} else {
			m.romBank = uint16(value & 0x1F)
		}

		if m.romBank == 0x00 || m.romBank == 0x20 || m.romBank == 0x40 || m.romBank == 0x60 {
			m.romBank++
		}

		m.romBank &= m.higherROMBank
		//fmt.Println("** romBank ", m.romBank)
	case (addr >= 0x4000) && (addr < 0x6000):
		//fmt.Printf("** write %X %X\n", addr, value)
		if m.mode == 1 {
			m.ramBank = uint16(value & 0x03)
			m.ramBank &= m.higherRAMBank
			//fmt.Println("** ramBank ", m.ramBank)
		} else {
			m.romBankHighBits = value & 0x03
			m.romBank = (m.romBank & 0x1F) | (uint16(m.romBankHighBits) << 5)

			if m.romBank == 0x00 || m.romBank == 0x20 || m.romBank == 0x40 || m.romBank == 0x60 {
				m.romBank++
			}

			m.romBank &= m.higherROMBank
			//fmt.Println("** romBank ", m.romBank)
		}
	case (addr >= 0x6000) && (addr < 0x8000):
		//fmt.Printf("** write %X %X\n", addr, value)
		if (m.ramSize != 3) && ((value & 0x01) != 0) {
			fmt.Printf("** Attempting to change MBC1 to mode 1 with incorrect RAM banks %X %X\n", addr, value)
		} else {
			m.mode = value & 0x01
		}
		//fmt.Println("** mode ", m.mode)
	case (addr >= 0xA000) && (addr < 0xC000):
		//fmt.Printf("** write %X %X\n", addr, value)
		if m.ramEnabled {
			if m.mode == 0 {
				m.ram[addr-0xA000] = value
			} else {
				m.ram[(addr-0xA000)+(m.ramBank*0x2000)] = value
			}
		} else {
			fmt.Printf("** Attempting to write to disabled RAM %X\n", addr)
		}
	case (addr >= 0xC000) && (addr < 0xFE00):
		WriteCommon(addr, value, m.memoryMap)
	case addr >= 0xFF00:
		// IO Registers
		WriteIO(addr, value, m.memoryMap)
	default:
		m.memoryMap[addr] = value
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func pow2Ceil(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n++
	return n
}
