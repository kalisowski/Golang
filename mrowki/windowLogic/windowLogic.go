package windowLogic

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Width  int32
	Height int32
}

func CreateWindow(title string, width, height int32) (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, nil, fmt.Errorf("SDL2 initialization error: %v", err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, fmt.Errorf("Window creation error: %v", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, fmt.Errorf("Renderer creation error: %v", err)
	}

	return window, renderer, nil
}

func DestroyWindow(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

func HandleEvents(window *sdl.Window, quit chan bool) {
	for {
		event := sdl.PollEvent()
		if event == nil {
			continue
		}

		switch event.(type) {
		case *sdl.QuitEvent:
			quit <- true
			return
		case *sdl.KeyboardEvent:
			keyEvent := event.(*sdl.KeyboardEvent)
			if keyEvent.State == sdl.PRESSED && keyEvent.Keysym.Sym == sdl.K_q {
				quit <- true
				return
			} else if keyEvent.State == sdl.PRESSED && keyEvent.Keysym.Mod&sdl.KMOD_ALT != 0 && keyEvent.Keysym.Sym == sdl.K_F4 {
				quit <- true
				return
			}
		}
	}
}
