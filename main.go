package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/drhelius/demo-emulator/gb/core"
	"github.com/drhelius/demo-emulator/gb/util"
)

const zoom = 5

var (
	texture          uint32
	viewportWidth    = util.GbWidth
	viewportHeight   = util.GbHeight
	colorFrameBuffer [util.GbWidth * util.GbHeight * 4]uint8
	window           *glfw.Window
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	fmt.Println("glfw init ok")

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	var err error
	window, err = glfw.CreateWindow(160*zoom, 144*zoom, "Codemotion GB", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetKeyCallback(onKey)
	window.MakeContextCurrent()

	fmt.Println("window ok")

	setup()

	fmt.Println("starting loop...")

	loop()

	destroy()
}

func setup() {

	romPathPtr := flag.String("rom", "", "rom path")
	flag.Parse()

	core.LoadROM(*romPathPtr)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	fmt.Println("gl init ok")

	glfw.SwapInterval(1)

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, util.GbWidth, util.GbHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&colorFrameBuffer[0]))

	fmt.Println("gl setup ok")
}

func loop() {
	for !window.ShouldClose() {
		core.RunToVBlank(colorFrameBuffer[:])
		render()
		glfw.PollEvents()
	}
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {

	switch glfw.Key(k) {
	case glfw.KeyEnter:
		if action == glfw.Press {
			core.ButtonPressed(util.StartButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.StartButton)
		}
	case glfw.KeySpace:
		if action == glfw.Press {
			core.ButtonPressed(util.SelectButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.SelectButton)
		}
	case glfw.KeyS:
		if action == glfw.Press {
			core.ButtonPressed(util.AButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.AButton)
		}
	case glfw.KeyA:
		if action == glfw.Press {
			core.ButtonPressed(util.BButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.BButton)
		}
	case glfw.KeyUp:
		if action == glfw.Press {
			core.ButtonPressed(util.UpButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.UpButton)
		}
	case glfw.KeyDown:
		if action == glfw.Press {
			core.ButtonPressed(util.DownButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.DownButton)
		}
	case glfw.KeyLeft:
		if action == glfw.Press {
			core.ButtonPressed(util.LeftButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.LeftButton)
		}
	case glfw.KeyRight:
		if action == glfw.Press {
			core.ButtonPressed(util.RightButton)
		} else if action == glfw.Release {
			core.ButtonReleased(util.RightButton)
		}
	default:
		return
	}
}

func render() {
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, util.GbWidth, util.GbHeight, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&colorFrameBuffer[0]))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, float64(viewportWidth), 0.0, float64(viewportHeight), -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)

	gl.Begin(gl.QUADS)
	gl.TexCoord2d(0.0, 1.0)
	gl.Vertex2d(0.0, 0.0)
	gl.TexCoord2d(1.0, 1.0)
	gl.Vertex2d(float64(viewportWidth), 0.0)
	gl.TexCoord2d(1.0, 0.0)
	gl.Vertex2d(float64(viewportWidth), float64(viewportHeight))
	gl.TexCoord2d(0.0, 0.0)
	gl.Vertex2d(0.0, float64(viewportHeight))
	gl.End()

	window.SwapBuffers()
}

func destroy() {
	gl.DeleteTextures(1, &texture)
	fmt.Println("destroyed")
}
