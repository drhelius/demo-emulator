package video

import (
	"github.com/drhelius/demo-emulator/gb/cpu"
	"github.com/drhelius/demo-emulator/gb/mapper"
	"github.com/drhelius/demo-emulator/gb/util"
)

var (
	// GbFrameBuffer is the internal Game Boy frame buffer
	GbFrameBuffer [util.GbWidth * util.GbHeight]uint8
	// ScreenEnabled keeps track of the screen state
	ScreenEnabled       bool
	statusMode          uint8
	statusModeCycles    uint
	subStatusModeCycles uint
	lyCounter           uint8
	vblankLine          uint8
	mem                 mapper.Mapper
	colorCacheBuffer    [util.GbWidth * util.GbHeight]uint8
)

func init() {
	statusMode = 1
	lyCounter = 144
	ScreenEnabled = true
}

// SetMapper injects the memory impl
func SetMapper(m mapper.Mapper) {
	mem = m
}

// Tick runs the video eumulation n cycles
// Then updates the frameBuffer and returns true if the simulation reached the vblank
func Tick(cycles uint) bool {
	vblank := false
	statusModeCycles += cycles

	if ScreenEnabled {
		switch statusMode {
		case 0:
			// During H-BLANK
			if statusModeCycles >= 204 {
				statusModeCycles -= 204
				statusMode = 2
				lyCounter++
				mem.GetMemoryMap()[0xFF44] = lyCounter
				CompareLYToLYC()

				if lyCounter == 144 {
					statusMode = 1
					vblankLine = 0
					subStatusModeCycles = statusModeCycles
					cpu.RequestInterrupt(cpu.InterruptVBlank)
					stat := mem.GetMemoryMap()[0xFF41]
					if util.IsSetBit(stat, 4) {
						cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
					}
					vblank = true
					windowLine = 0
				} else {
					stat := mem.GetMemoryMap()[0xFF41]
					if util.IsSetBit(stat, 5) {
						cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
					}
				}

				updateStatRegister()
			}
		case 1:
			// During V-BLANK
			subStatusModeCycles += cycles

			if subStatusModeCycles >= 456 {
				subStatusModeCycles -= 456
				vblankLine++

				if vblankLine <= 9 {
					lyCounter++
					mem.GetMemoryMap()[0xFF44] = lyCounter
					CompareLYToLYC()
				}
			}

			if (statusModeCycles >= 4104) && (subStatusModeCycles >= 4) && (lyCounter == 153) {
				lyCounter = 0
				mem.GetMemoryMap()[0xFF44] = lyCounter
				CompareLYToLYC()
			}

			if statusModeCycles >= 4560 {
				statusModeCycles -= 4560
				statusMode = 2
				updateStatRegister()
				stat := mem.GetMemoryMap()[0xFF41]
				if util.IsSetBit(stat, 5) {
					cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
				}
			}
		case 2:
			// During searching OAM RAM
			if statusModeCycles >= 80 {
				statusModeCycles -= 80
				statusMode = 3
				updateStatRegister()
			}
		case 3:
			// During transfering data to LCD driver
			if statusModeCycles >= 172 {
				statusModeCycles -= 172
				statusMode = 0
				scanLine(lyCounter)
				updateStatRegister()
				stat := mem.GetMemoryMap()[0xFF41]
				if util.IsSetBit(stat, 3) {
					cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
				}
			}
		}
	} else {
		if statusModeCycles >= 70224 {
			statusModeCycles -= 70224
			vblank = true
		}
	}

	return vblank
}

// EnableScreen enables the screen
func EnableScreen() {
	if !ScreenEnabled {
		ScreenEnabled = true
		statusMode = 0
		statusModeCycles = 0
		subStatusModeCycles = 0
		lyCounter = 0
		vblankLine = 0
		windowLine = 0

		mem.GetMemoryMap()[0xFF44] = lyCounter

		stat := mem.GetMemoryMap()[0xFF41]
		if util.IsSetBit(stat, 5) {
			cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
		}

		CompareLYToLYC()
	}
}

// DisableScreen disables the screen
func DisableScreen() {
	ScreenEnabled = false
	mem.GetMemoryMap()[0xFF44] = 0x00
	stat := mem.GetMemoryMap()[0xFF41]
	stat &= 0x7C
	mem.GetMemoryMap()[0xFF41] = stat
	statusMode = 0
	statusModeCycles = 0
	subStatusModeCycles = 0
	lyCounter = 0
}

// CompareLYToLYC compares LY counter with LYC register
func CompareLYToLYC() {
	if ScreenEnabled {
		lyc := mem.GetMemoryMap()[0xFF45]
		stat := mem.GetMemoryMap()[0xFF41]

		if lyc == lyCounter {
			stat = util.SetBit(stat, 2)
			if util.IsSetBit(stat, 6) {
				cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
			}
		} else {
			stat = util.UnsetBit(stat, 2)
		}

		mem.GetMemoryMap()[0xFF41] = stat
	}
}

func scanLine(line uint8) {
	lcdc := mem.GetMemoryMap()[0xFF40]

	if ScreenEnabled && util.IsSetBit(lcdc, 7) {
		renderBG(line)
		renderWindow(line)
		renderSprites(line)
	} else {
		var x uint8
		for ; x < util.GbWidth; x++ {
			GbFrameBuffer[(line*util.GbWidth)+x] = 0
		}
	}
}

func updateStatRegister() {
	// Updates the STAT register with current mode
	stat := mem.GetMemoryMap()[0xFF41]
	mem.GetMemoryMap()[0xFF41] = (stat & 0xFC) | (statusMode & 0x3)
}
