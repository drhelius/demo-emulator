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
	glfw.Setup()
	opengl.Setup(colorFrameBuffer[:])
	loadROM()
	loop()
	opengl.Teardown()
	glfw.Teardown()
}

func loadROM() {
	romPathPtr := flag.String("rom", "", "rom path")
	flag.Parse()

	core.LoadROM(*romPathPtr)
}

func loop() {
	fmt.Println("starting loop...")
	for !glfw.WindowClosed() {
		core.RunToVBlank(colorFrameBuffer[:])
		opengl.Render()
		glfw.Update()
	}
}
