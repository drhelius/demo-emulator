package input

import (
	"fmt"

	"github.com/drhelius/demo-emulator/gb/util"
)

// Tick runs the input eumulation n cycles
func Tick(cycles uint32) {

}

// ButtonPressed tells the input system that a button has been pressed
func ButtonPressed(button util.GameboyButton) {
	fmt.Printf("button pressed %d\n", button)
}

// ButtonReleased tells the input system that a button has been released
func ButtonReleased(button util.GameboyButton) {
	fmt.Printf("button released %d\n", button)
}
