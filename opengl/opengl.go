package opengl

import (
	"fmt"
	"runtime"

	"github.com/drhelius/demo-emulator/gb/util"
	"github.com/go-gl/gl/v2.1/gl"
)

var (
	texture          uint32
	viewportWidth    = util.GbWidth
	viewportHeight   = util.GbHeight
	colorFrameBuffer []uint8
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Setup initializes OpenGL
func Setup(fb []uint8) {

	colorFrameBuffer = fb

	if err := gl.Init(); err != nil {
		panic(err)
	}

	fmt.Println("gl init ok")

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, util.GbWidth, util.GbHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&colorFrameBuffer[0]))

	fmt.Println("gl setup ok")
}

// Render draws the current frame
func Render() {
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
}

// Teardown shutd downs OpenGL
func Teardown() {
	gl.DeleteTextures(1, &texture)
}
