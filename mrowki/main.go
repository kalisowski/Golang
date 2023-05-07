package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	pointSize int32 = 4
)

func main() {
	var windowWidth, windowHeight int
	flag.IntVar(&windowWidth, "w", 800, "Window width")
	flag.IntVar(&windowHeight, "h", 600, "Window height")
	flag.Parse()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Fprintf(os.Stderr, "SDL2 initialization error: %v\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL2 Example", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Window creation error: %v\n", err)
		os.Exit(1)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Renderer creation error: %v\n", err)
		os.Exit(1)
	}
	defer renderer.Destroy()

	rand.Seed(time.Now().UnixNano())
	runtime.LockOSThread()

	quit := make(chan bool)
	x, y := int32(windowWidth/2), int32(windowHeight/2)

	for {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit <- true
			case *sdl.KeyboardEvent:
				keyEvent := event.(*sdl.KeyboardEvent)
				if keyEvent.State == sdl.PRESSED && keyEvent.Keysym.Sym == sdl.K_q {
					quit <- true
				} else if keyEvent.State == sdl.PRESSED && keyEvent.Keysym.Mod&sdl.KMOD_ALT != 0 && keyEvent.Keysym.Sym == sdl.K_F4 {
					quit <- true
				}
			}
		}

		x += int32(rand.Intn(3) - 1)
		y += int32(rand.Intn(3) - 1)

		if x <= 0 || y <= 0 || x >= int32(windowWidth) || y >= int32(windowHeight) {
			break
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.FillRect(&sdl.Rect{X: x, Y: y, W: pointSize, H: pointSize})
		renderer.Present()

		select {
		case <-quit:
			break
		default:
			sdl.Delay(1)
		}
	}

	window.Destroy()
	sdl.Quit()
}
