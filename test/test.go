package main

import (
	".."
	"fmt"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	nd := sdl.GetNumVideoDrivers()
	fmt.Printf("Drivers: %v\n", nd)
	for i := 0; i < nd; i++ {
		fmt.Printf("\t%v\n", sdl.GetVideoDriver(i))
	}
	fmt.Printf("Using: %v\n", sdl.GetCurrentVideoDriver())
	fmt.Println()

	nd = sdl.GetNumVideoDisplays()
	fmt.Printf("Displays: %v\n", nd)
	for i := 0; i < nd; i++ {
		rect, err := sdl.GetDisplayBounds(i)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\t%v\n", rect)
	}

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
}
