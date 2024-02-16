package simengine

import (
	"runtime"
	"log"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Engine struct {
	Width int
	Height int
	Title string
	Background Color
	window *glfw.Window
}

func (e *Engine) Run() {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		log.Println(err)
		return
	}
	defer glfw.Terminate()
    
  glfw.WindowHint(glfw.ContextVersionMajor, 4);
  glfw.WindowHint(glfw.ContextVersionMinor, 6);
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile);
  glfw.WindowHint(glfw.Resizable, glfw.False);
    
  e.window, err = glfw.CreateWindow(e.Width, e.Height, e.Title, nil, nil);
	if err != nil {
		log.Println(err)
		return
	}
  e.window.MakeContextCurrent();

  err = gl.Init()
	if err != nil {
		log.Println(err)
		return
	}


	gl.ClearColor(e.Background.Red, e.Background.Green, e.Background.Blue, e.Background.Alpha)
  for !e.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT);
		glfw.PollEvents()
		e.window.SwapBuffers()
	}
}
