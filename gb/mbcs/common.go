package mbcs

// WriteCommon emulates the internal RAM 0xC000-0xD000
// and the mirror of the internal RAM 0xE000-0xFE00
func WriteCommon(addr uint16, value uint8, mem []uint8) {
	switch {
	case (addr >= 0xC000) && (addr < 0xDE00):
		mem[addr] = value
		mem[addr+0x2000] = value
	case (addr >= 0xE000) && (addr < 0xFE00):
		mem[addr] = value
		mem[addr-0x2000] = value
	default:
		mem[addr] = value
	}
}
