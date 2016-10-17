package input

import (
	"github.com/drhelius/demo-emulator/gb/cpu"
	"github.com/drhelius/demo-emulator/gb/util"
)

var (
	joypadState uint8
	p1          uint8
	inputCycles uint
)

func init() {
	joypadState = 0xFF
	p1 = 0xFF
}

// Tick runs the input eumulation n cycles
func Tick(cycles uint) {
	inputCycles += cycles

	// Joypad Poll Speed (64 Hz)
	if inputCycles >= 65536 {
		inputCycles -= 65536
		update()
	}
}

// Read returns the P1 register
func Read() uint8 {
	return p1
}

// Write stores the P1 register
func Write(value uint8) {
	p1 = (p1 & 0xCF) | (value & 0x30)
	update()
}

// ButtonPressed tells the input system that a button has been pressed
func ButtonPressed(button util.GameboyButton) {
	joypadState = util.UnsetBit(joypadState, uint8(button))
}

// ButtonReleased tells the input system that a button has been released
func ButtonReleased(button util.GameboyButton) {
	joypadState = util.SetBit(joypadState, uint8(button))
}

func update() {
	current := p1 & 0xF0

	switch current & 0x30 {
	case 0x10:
		topJoypad := (joypadState >> 4) & 0x0F
		current |= topJoypad
	case 0x20:
		bottomJoypad := joypadState & 0x0F
		current |= bottomJoypad
	case 0x30:
		current |= 0x0F
	}

	if (p1 & ^current & 0x0F) != 0 {
		cpu.RequestInterrupt(cpu.InterruptJoypad)
	}

	p1 = current
}
