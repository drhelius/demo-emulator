package video

import "github.com/drhelius/demo-emulator/gb/util"

var (
	windowLine uint
)

// ResetWindowLine resetes the current line of the window
func ResetWindowLine() {
	wy := mem.GetMemoryMap()[0xFF4A]

	if (windowLine == 0) && (lyCounter < 144) && (lyCounter > wy) {
		windowLine = 144
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
