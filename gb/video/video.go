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
	windowLine          uint
	mem                 mapper.Mapper
	spriteCacheBuffer   [util.GbWidth * util.GbHeight]int
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

	if ScreenEnabled {
		statusModeCycles += cycles

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

// ResetWindowLine resetes the current line of the window
func ResetWindowLine() {
	wy := mem.GetMemoryMap()[0xFF4A]

	if (windowLine == 0) && (lyCounter < 144) && (lyCounter > wy) {
		windowLine = 144
	}
}

func updateStatRegister() {
	// Updates the STAT register with current mode
	stat := mem.GetMemoryMap()[0xFF41]
	mem.GetMemoryMap()[0xFF41] = (stat & 0xFC) | (statusMode & 0x3)
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

func renderBG(line uint8) {
	lcdc := mem.GetMemoryMap()[0xFF40]
	lineWidth := uint(line) * uint(util.GbWidth)

	if util.IsSetBit(lcdc, 0) {
		var tiles uint = 0x8800
		if util.IsSetBit(lcdc, 4) {
			tiles = 0x8000
		}
		var maploc uint = 0x9800
		if util.IsSetBit(lcdc, 3) {
			maploc = 0x9C00
		}

		scx := mem.GetMemoryMap()[0xFF43]
		scy := mem.GetMemoryMap()[0xFF42]
		lineAdjusted := line + scy
		y32 := (uint(lineAdjusted) / 8) * 32
		pixely := uint(lineAdjusted) % 8
		pixely2 := pixely * 2

		var x uint
		for ; x < 32; x++ {
			var tile uint8

			if tiles == 0x8800 {
				tile = uint8(int(int8(mem.GetMemoryMap()[maploc+y32+x])) + 128)
			} else {
				tile = mem.Read(uint16(maploc + y32 + x))
			}

			mapOffsetX := x * 8
			tile16 := uint(tile) * 16
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

				position := lineWidth + uint(bufferX)

				colorCacheBuffer[position] = pixel & 0x03

				palette := mem.GetMemoryMap()[0xFF47]
				color := (palette >> (pixel * 2)) & 0x03
				GbFrameBuffer[position] = color
			}
		}
	} else {
		var x uint
		for ; x < util.GbWidth; x++ {
			position := lineWidth + x
			GbFrameBuffer[position] = 0
			colorCacheBuffer[position] = 0
		}
	}
}

func renderWindow(line uint8) {
	if windowLine > 143 {
		return
	}

	lcdc := mem.GetMemoryMap()[0xFF40]
	if !util.IsSetBit(lcdc, 5) {
		return
	}

	wx := int(mem.GetMemoryMap()[0xFF4B]) - 7
	if wx > 159 {
		return
	}

	wy := mem.GetMemoryMap()[0xFF4A]
	if (wy > 143) || (wy > line) {
		return
	}

	var tilesAddr uint = 0x8800
	if util.IsSetBit(lcdc, 4) {
		tilesAddr = 0x8000
	}
	var mapAddr uint = 0x9800
	if util.IsSetBit(lcdc, 6) {
		mapAddr = 0x9C00
	}

	lineAdjusted := windowLine
	y32 := (lineAdjusted / 8) * 32
	pixely := lineAdjusted % 8
	pixely2 := pixely * 2
	lineWidth := uint(line) * util.GbWidth

	var x uint
	for ; x < 32; x++ {
		var tile int

		if tilesAddr == 0x8800 {
			tile = int(int8(mem.GetMemoryMap()[mapAddr+y32+x]))
			tile += 128
		} else {
			tile = int(mem.GetMemoryMap()[mapAddr+y32+x])
		}

		mapOffsetX := x * 8
		tile16 := uint(tile) * 16
		tileAddress := tilesAddr + tile16 + pixely2

		byte1 := mem.GetMemoryMap()[tileAddress]
		byte2 := mem.GetMemoryMap()[tileAddress+1]

		var pixelx uint8
		for ; pixelx < 8; pixelx++ {
			bufferX := (int(mapOffsetX) + int(pixelx) + wx)

			if (bufferX < 0) || (bufferX >= util.GbWidth) {
				continue
			}

			var pixel uint8
			if (byte1 & (0x1 << (7 - pixelx))) != 0 {
				pixel = 1
			}
			if (byte2 & (0x1 << (7 - pixelx))) != 0 {
				pixel |= 2
			}

			position := lineWidth + uint(bufferX)
			colorCacheBuffer[position] = pixel & 0x03

			palette := mem.GetMemoryMap()[0xFF47]
			color := (palette >> (pixel * 2)) & 0x03
			GbFrameBuffer[position] = color

		}
	}
	windowLine++
}

func renderSprites(line uint8) {

	lcdc := mem.GetMemoryMap()[0xFF40]

	if !util.IsSetBit(lcdc, 1) {
		return
	}

	spriteHeight := 8
	var spriteHeightMask uint8 = 0xFF
	if util.IsSetBit(lcdc, 2) {
		spriteHeight = 16
		spriteHeightMask = 0xFE
	}

	lineWidth := int(line) * util.GbWidth

	sprite := 39
	for ; sprite >= 0; sprite-- {
		sprite4 := sprite * 4
		spriteY := int(mem.GetMemoryMap()[0xFE00+sprite4]) - 16

		if (spriteY > int(line)) || ((spriteY + spriteHeight) <= int(line)) {
			continue
		}

		spriteX := int(mem.GetMemoryMap()[0xFE00+sprite4+1]) - 8

		if (spriteX < -7) || (spriteX >= util.GbWidth) {
			continue
		}

		spriteTile16 := int(mem.GetMemoryMap()[0xFE00+sprite4+2]&spriteHeightMask) * 16
		spriteFlags := mem.GetMemoryMap()[0xFE00+sprite4+3]

		spritePallete := util.IsSetBit(spriteFlags, 4)
		xflip := util.IsSetBit(spriteFlags, 5)
		yflip := util.IsSetBit(spriteFlags, 6)
		aboveBG := !util.IsSetBit(spriteFlags, 7)

		tiles := 0x8000
		pixelY := int(line) - spriteY
		if yflip {
			pixelY = (spriteHeight - 1) - (int(line) - spriteY)
		}

		pixelY2 := 0
		offset := 0

		if (spriteHeight == 16) && (pixelY >= 8) {
			pixelY2 = (pixelY - 8) * 2
			offset = 16
		} else {
			pixelY2 = pixelY * 2
		}

		tileAddress := tiles + spriteTile16 + pixelY2 + offset

		byte1 := mem.GetMemoryMap()[tileAddress]
		byte2 := mem.GetMemoryMap()[tileAddress+1]

		var pixelx int
		for ; pixelx < 8; pixelx++ {

			var pixelxFlipped = 7 - uint(pixelx)
			if xflip {
				pixelxFlipped = uint(pixelx)
			}

			var pixel uint8

			if (byte1 & (0x01 << pixelxFlipped)) != 0 {
				pixel = 0x01
			}
			if (byte2 & (0x01 << pixelxFlipped)) != 0 {
				pixel |= 0x02
			}

			if pixel == 0 {
				continue
			}

			bufferX := spriteX + pixelx

			if (bufferX < 0) || (bufferX >= util.GbWidth) {
				continue
			}

			position := lineWidth + bufferX

			colorCache := colorCacheBuffer[position]

			spriteCache := spriteCacheBuffer[position]
			if util.IsSetBit(colorCache, 3) && (spriteCache < spriteX) {
				continue
			}

			if !aboveBG && ((colorCache & 0x03) != 0) {
				continue
			}

			colorCacheBuffer[position] = util.SetBit(colorCache, 3)
			spriteCacheBuffer[position] = spriteX

			var paletteAddr uint16 = 0xFF48
			if spritePallete {
				paletteAddr = 0xFF49
			}

			palette := mem.GetMemoryMap()[paletteAddr]
			color := (palette >> (pixel * 2)) & 0x03
			GbFrameBuffer[position] = color
		}
	}
}
