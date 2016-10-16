package cpu

import "github.com/drhelius/demo-emulator/gb/util"

var (
	serialBit    int
	serialCycles uint
)

func updateSerial() {
	sc := mem.Read(0xFF02)

	if util.IsSetBit(sc, 7) && util.IsSetBit(sc, 0) {
		serialCycles += clockCycles

		if serialBit < 0 {
			serialBit = 0
			serialCycles = 0
			return
		}

		if serialCycles >= 512 {
			if serialBit > 7 {
				mem.Write(0xFF02, sc&0x7F)
				RequestInterrupt(InterruptSerial)
				serialBit = -1
				return
			}

			sb := mem.Read(0xFF01)
			sb <<= 1
			sb |= 0x01
			mem.Write(0xFF01, sb)

			serialCycles -= 512
			serialBit++
		}
	}
}
