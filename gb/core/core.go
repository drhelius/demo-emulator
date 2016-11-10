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

var pallete = [12]uint8{
	// R     G     B
	0x87, 0x96, 0x03, // color 0
	0x4d, 0x6b, 0x03, // color 1
	0x2b, 0x55, 0x03, // color 2
	0x14, 0x44, 0x03} // color 3

// RunToVBlank runs a single frame of the emulator
// The emulator must run at 60fps
func RunToVBlank(colorFrameBuffer []uint8) {

	// keep updating each system
	// until the vblank is reached
	for vblank := false; !vblank; {
		clockCycles := cpu.Tick()
		vblank = video.Tick(clockCycles)
		input.Tick(clockCycles)
	}

	// the frame buffer of the Game Boy encodes color
	// as a four shades of gray (or green)
	// we have to transform these 4 colors to RGB by using
	// a predefined pallete
	for i, pixelCount := 0, util.GbWidth*util.GbHeight; i < pixelCount; i++ {
		colorFrameBuffer[i*4] = pallete[video.GbFrameBuffer[i]*3]         // red
		colorFrameBuffer[(i*4)+1] = pallete[(video.GbFrameBuffer[i]*3)+1] // green
		colorFrameBuffer[(i*4)+2] = pallete[(video.GbFrameBuffer[i]*3)+2] // blue
	}
}

// LoadROM loads a new rom into the emulator
// this fucntion must be called before running RunToVBlank
func LoadROM(filePath string) {

	fmt.Printf("loading ROM \"%s\"...\n", filePath)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	cartType := data[0x147]
	var m mapper.Mapper

	// check the memory mapper of the ROM
	switch cartType {
	case 0x00:
		fmt.Println("found ROM")
		m = new(mbcs.RomOnly)
	case 0x08:
		fmt.Println("found ROM + SRAM")
		m = new(mbcs.RomOnly)
	case 0x09:
		fmt.Println("found ROM + SRAM + BATT")
		m = new(mbcs.RomOnly)
	case 0x01:
		fmt.Println("found MBC1")
		m = new(mbcs.MBC1)
	case 0x02:
		fmt.Println("found MBC1 + SRAM")
		m = new(mbcs.MBC1)
	case 0x03:
		fmt.Println("found MBC1 + SRAM + BATT")
		m = new(mbcs.MBC1)
	default:
		// MBC1 and ROM are the only cartridges supported
		panic(fmt.Sprintf("cartridge type not supported: %d", cartType))
	}

	// once we know which is the correct memory mapper
	// we inject it into the systems that use the memory
	m.Setup(data)
	cpu.SetMapper(m)
	video.SetMapper(m)
}

// ButtonPressed tells the emulator that a button has been pressed
func ButtonPressed(button util.GameboyButton) {
	input.ButtonPressed(button)
}

// ButtonReleased tells the emulator that a button has been released
func ButtonReleased(button util.GameboyButton) {
	input.ButtonReleased(button)
}
