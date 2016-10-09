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
	statusModeCycles    uint32
	subStatusModeCycles uint32
	lyCounter           uint8
	vblankLine          uint8
	mem                 mapper.Mapper
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
func Tick(cycles uint32) bool {
	vblank := false
	statusModeCycles += cycles

	switch statusMode {
	case 0:
		// During H-BLANK
		if statusModeCycles >= 204 {
			statusModeCycles -= 204
			statusMode = 2
			//fmt.Printf("statusMode: %d\n", statusMode)
			lyCounter++
			mem.Write(0xFF44, lyCounter)
			CompareLYToLYC()

			if lyCounter == 144 {
				statusMode = 1
				vblankLine = 0
				subStatusModeCycles = statusModeCycles
				cpu.RequestInterrupt(cpu.InterruptVBlank)
				vblank = true
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
				mem.Write(0xFF44, lyCounter)
				CompareLYToLYC()
			}
		}

		if (statusModeCycles >= 4104) && (subStatusModeCycles >= 4) && (lyCounter == 153) {
			lyCounter = 0
			mem.Write(0xFF44, lyCounter)
			CompareLYToLYC()
		}

		if statusModeCycles >= 4560 {
			statusModeCycles -= 4560
			statusMode = 2
			//fmt.Printf("statusMode: %d\n", statusMode)
			updateStatRegister()
		}
	case 2:
		// During searching OAM RAM
		if statusModeCycles >= 80 {
			statusModeCycles -= 80
			statusMode = 3
			//fmt.Printf("statusMode: %d\n", statusMode)
			updateStatRegister()
		}
	case 3:
		// During transfering data to LCD driver
		if statusModeCycles >= 172 {
			statusModeCycles -= 172
			statusMode = 0
			//fmt.Printf("statusMode: %d\n", statusMode)
			scanLine(lyCounter)
			updateStatRegister()
		}
	}

	return vblank
}

// EnableScreen enables the screen
func EnableScreen() {
	/*
		ScreenEnabled = true
	*/
}

// DisableScreen disables the screen
func DisableScreen() {
	/*
		ScreenEnabled = false
		mem.Write(0xFF44, 0x00)
		stat := mem.Read(0xFF41)
		stat &= 0x7C
		mem.Write(0xFF41, stat)
		statusMode = 0
		statusModeCycles = 0
		subStatusModeCycles = 0
		lyCounter = 0
	*/
}

// CompareLYToLYC compares LY counter with LYC register
func CompareLYToLYC() {
	if ScreenEnabled {
		lyc := mem.Read(0xFF45)
		stat := mem.Read(0xFF41)

		if lyc == lyCounter {
			stat = util.SetBit(stat, 2)
			if util.IsSetBit(stat, 6) {
				cpu.RequestInterrupt(cpu.InterruptLCDSTAT)
			}
		} else {
			stat = util.UnsetBit(stat, 2)
		}

		mem.Write(0xFF41, stat)
	}
}

func updateStatRegister() {
	// Updates the STAT register with current mode
	stat := mem.Read(0xFF41)
	mem.Write(0xFF41, (stat&0xFC)|(statusMode&0x3))
}

func scanLine(line uint8) {
	//lcdc := mem.Read(0xFF40)

	//if ScreenEnabled && util.IsSetBit(lcdc, 7) {
	renderBG(line)
	//renderWindow(line);
	//renderSprites(line);
	/*} else {
		var x uint8
		for ; x < util.GbWidth; x++ {
			GbFrameBuffer[(line*util.GbWidth)+x] = 0
		}
	}*/
}

func renderBG(line uint8) {
	//fmt.Printf("renderBG: %d\n", line)
	lcdc := mem.Read(0xFF40)
	lineWidth := uint32(line) * uint32(util.GbWidth)

	//if util.IsSetBit(lcdc, 0) {
	var tiles uint32 = 0x8800
	if util.IsSetBit(lcdc, 4) {
		tiles = 0x8000
	}
	var maploc uint32 = 0x9800
	if util.IsSetBit(lcdc, 3) {
		maploc = 0x9C00
	}

	scx := mem.Read(0xFF43)
	scy := mem.Read(0xFF42)
	lineAdjusted := line + scy
	y32 := (uint32(lineAdjusted) / 8) * 32
	pixely := uint32(lineAdjusted) % 8
	pixely2 := pixely * 2

	var x uint32
	for ; x < 32; x++ {
		var tile uint8

		if tiles == 0x8800 {
			tile = uint8(int32(int8(mem.Read(uint16(maploc+y32+x)))) + 128)
		} else {
			tile = mem.Read(uint16(maploc + y32 + x))
		}

		mapOffsetX := x * 8
		tile16 := uint32(tile) * 16
		tileAddress := tiles + tile16 + pixely2

		byte1 := mem.Read(uint16(tileAddress))
		byte2 := mem.Read(uint16(tileAddress) + 1)

		var pixelx uint8
		for ; pixelx < 8; pixelx++ {
			bufferX := uint8(mapOffsetX) + pixelx - scx

			if bufferX >= util.GbWidth {
				continue
			}
			var pixel uint8
			if (byte1 & (0x1 << (7 - pixelx))) != 0 {
				pixel = 1
			}
			if (byte2 & (0x1 << (7 - pixelx))) != 0 {
				pixel |= 2
			}

			position := lineWidth + uint32(bufferX)

			palette := mem.Read(0xFF47)
			color := (palette >> (pixel * 2)) & 0x03
			GbFrameBuffer[position] = color

		}
	}
	/*} else {
		var x uint32
		for ; x < util.GbWidth; x++ {
			position := lineWidth + x
			GbFrameBuffer[position] = 0
		}
	}*/
}
