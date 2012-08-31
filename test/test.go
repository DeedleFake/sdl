package main

import (
	"fmt"
	"github.com/DeedleFake/sdl"
	"github.com/DeedleFake/sdl/img"
	"time"
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

		nm, err := sdl.GetNumDisplayModes(i)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\tModes: %v\n", nm)
		for m := 0; m < nm; m++ {
			dm, err := sdl.GetDisplayMode(i, m)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t\t%v\n", dm)
		}
	}
	fmt.Println()

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

	ren, err := win.CreateRenderer(-1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer ren.Destroy()
	ren.SetDrawColor(100, 100, 255, sdl.ALPHA_OPAQUE)

	bmp, err := img.LoadTexture(ren, "test.bmp")
	if err != nil {
		panic(err)
	}
	defer bmp.Destroy()

	var rot float64
	var x, y int

	input := make(map[sdl.Keycode]bool)

	fps := time.Tick(time.Second / 60)
	for fps != nil {
		var ev sdl.Event
		for sdl.PollEvent(&ev) {
			switch ev := ev.(type) {
			case *sdl.KeyboardEvent:
				if ev.Type == sdl.KEYDOWN {
					input[ev.Keysym.Sym] = true
				} else {
					input[ev.Keysym.Sym] = false
				}
			case *sdl.QuitEvent:
				fps = nil
			}
		}

		if input[sdl.K_UP] {
			y--
		}
		if input[sdl.K_DOWN] {
			y++
		}
		if input[sdl.K_LEFT] {
			x--
		}
		if input[sdl.K_RIGHT] {
			x++
		}

		rot += .1

		ren.Clear()

		ren.CopyEx(
			bmp,
			nil,
			&sdl.Rect{int32(x), int32(y), 300, 300},
			rot,
			nil,
			sdl.FLIP_NONE,
		)

		ren.Present()

		if fps != nil {
			<-fps
		}
	}
}
