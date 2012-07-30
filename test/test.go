package main

import (
	".."
	"time"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	win, err := sdl.CreateWindow(
		"Test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		640,
		480,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	time.Sleep(3 * time.Second)
}
