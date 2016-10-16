package video

import "github.com/drhelius/demo-emulator/gb/util"

var (
	spriteCacheBuffer [util.GbWidth * util.GbHeight]int
)

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
