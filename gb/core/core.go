package core

import (
	"fmt"
	"io/ioutil"

	"github.com/drhelius/demo-emulator/gb/cpu"
	"github.com/drhelius/demo-emulator/gb/input"
	"github.com/drhelius/demo-emulator/gb/mapper"
	"github.com/drhelius/demo-emulator/gb/mbcs"
	"github.com/drhelius/demo-emulator/gb/util"
	"github.com/drhelius/demo-emulator/gb/video"
)

var (
	ready   bool
	pallete = [12]uint8{
		0x87, 0x96, 0x03,
		0x4d, 0x6b, 0x03,
		0x2b, 0x55, 0x03,
		0x14, 0x44, 0x03}
)

// RunToVBlank runs a single frame of the emulator
// The emulator must run at 60fps
func RunToVBlank(colorFrameBuffer []uint8) {
	if ready {
		for vblank := false; !vblank; {
			var clockCycles = cpu.Tick()
			vblank = video.Tick(clockCycles)
			input.Tick(clockCycles)
		}

		for i, pixelCount := 0, util.GbWidth*util.GbHeight; i < pixelCount; i++ {
			colorFrameBuffer[i*4] = pallete[video.GbFrameBuffer[i]*3]         // red
			colorFrameBuffer[(i*4)+1] = pallete[(video.GbFrameBuffer[i]*3)+1] // green
			colorFrameBuffer[(i*4)+2] = pallete[(video.GbFrameBuffer[i]*3)+2] // blue
		}
	}
}

// LoadROM loads a new rom into the Emulator
// This fucntion must be called before running RunToVBlank
func LoadROM(filePath string) {

	fmt.Printf("loading rom \"%s\" ...\n", filePath)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fmt.Println("load rom ok")

	cartType := data[0x147]
	var m mapper.Mapper

	switch cartType {
	case 0x00:
		// NO MBC
		fallthrough
	case 0x08:
		// ROM
		// SRAM
		fallthrough
	case 0x09:
		// ROM
		// SRAM
		// BATT
		m = new(mbcs.RomOnly)
	case 0x01:
		// MBC1
		fallthrough
	case 0x02:
		// MBC1
		// SRAM
		fallthrough
	case 0x03:
		// MBC1
		// SRAM
		// BATT
		m = new(mbcs.MBC1)
	default:
		panic(fmt.Sprintf("Cartridge type not supported: %d", cartType))
	}

	m.Setup(data)
	cpu.SetMapper(m)
	video.SetMapper(m)

	ready = true
}

// ButtonPressed tells the emulator that a button has been pressed
func ButtonPressed(button util.GameboyButton) {
	input.ButtonPressed(button)
}

// ButtonReleased tells the emulator that a button has been released
func ButtonReleased(button util.GameboyButton) {
	input.ButtonReleased(button)
}
