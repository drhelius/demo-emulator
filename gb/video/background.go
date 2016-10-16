package video

import "github.com/drhelius/demo-emulator/gb/util"

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
