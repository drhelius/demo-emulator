package cpu

import "github.com/drhelius/demo-emulator/gb/util"

var (
	divCycles  uint
	timaCycles uint
)

// ResetDivCycles sets divCycles to 0
func ResetDivCycles() {
	divCycles = 0
}

// ResetTimaCycles sets timaCycles to 0
func ResetTimaCycles() {
	timaCycles = 0
}

func updateTimers() {
	divCycles += clockCycles

	for divCycles >= 256 {
		divCycles -= 256
		div := mem.Read(0xFF04)
		div++
		mem.Write(0xFF04, div)
	}

	tac := mem.Read(0xFF07)

	// if tima is running
	if util.IsSetBit(tac, 2) {
		timaCycles += clockCycles

		var freq uint

		switch tac & 0x03 {
		case 0:
			freq = 1024
		case 1:
			freq = 16
		case 2:
			freq = 64
		case 3:
			freq = 256
		}

		for timaCycles >= freq {
			timaCycles -= freq
			tima := mem.Read(0xFF05)

			if tima == 0xFF {
				tima = mem.Read(0xFF06)
				RequestInterrupt(InterruptTimer)
			} else {
				tima++
			}

			mem.Write(0xFF05, tima)
		}
	}
}
