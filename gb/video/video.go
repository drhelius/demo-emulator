package video

import (
	"github.com/drhelius/demo-emulator/gb/util"
)

var (
	// GbFrameBuffer is the internal Game Boy frame buffer
	GbFrameBuffer [util.GbWidth * util.GbHeight]uint8
	// ScreenEnabled keeps track of the screen state
	ScreenEnabled bool
)

// Tick runs the video eumulation n cycles
// Then updates the frameBuffer and returns true if the simulation reached the vblank
func Tick(cycles uint32) bool {
	return true
}

// EnableScreen enables the screen
func EnableScreen() {

}

// DisableScreen disables the screen
func DisableScreen() {

}

// ResetWindowLine resets the line counter
func ResetWindowLine() {

}

// CompareLYToLYC
func CompareLYToLYC() {

}
