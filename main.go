package main

import (
	"flag"
	"fmt"

	"github.com/drhelius/demo-emulator/gb/core"
	"github.com/drhelius/demo-emulator/gb/util"
	"github.com/drhelius/demo-emulator/glfw"
	"github.com/drhelius/demo-emulator/opengl"
)

var colorFrameBuffer [util.GbWidth * util.GbHeight * 4]uint8

func main() {
	// init windowing system
	glfw.Setup()

	// init rendering system
	opengl.Setup(colorFrameBuffer[:])

	loadROM()

	loop()

	opengl.Teardown()
	glfw.Teardown()
}

func loadROM() {
	// get the rom path from a command line argument
	romPathPtr := flag.String("rom", "", "rom path")
	flag.Parse()

	// tell the emulator to load the rom from a path
	core.LoadROM(*romPathPtr)
}

func loop() {
	fmt.Println("starting loop...")

	// while the user doesn't close the window
	for !glfw.WindowClosed() {

		// run the emulation one frame
		// it should be called 60 times in a second
		core.RunToVBlank(colorFrameBuffer[:])

		// render de results
		opengl.Render()

		// present the rendering in a window
		// and update input
		glfw.Update()
	}
}
