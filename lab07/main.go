package main

import (
	"Libka/structures"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Moje Okno", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(255, 255, 255, 255)

	for {
		renderer.Clear()

		renderer.Present()
		window.UpdateSurface()

		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			return
		}
	}
	sdl.Quit()
}

func test() {
	xinit := structures.Point[float64]{
		X: 1.0,
		Y: 2.0,
	}
	yinit := structures.Point[float64]{
		X: 3.0,
		Y: 5.0,
	}
	fmt.Println(xinit)
	x := xinit.Move(yinit)
	fmt.Println(x)
	x = x.Add(xinit)
	fmt.Println(x)
	fmt.Println(yinit)
	fmt.Println(x.Distance(yinit))

	rect := structures.Rect[float64]{
		Min: xinit,
		Max: yinit,
	}
	fmt.Println(rect.Corners())
}
