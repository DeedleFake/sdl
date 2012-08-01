package main

import (
	".."
	"fmt"
	"io"
	"os"
	"time"
)

func LoadTexture(ren *sdl.Renderer, r io.ReadSeeker) (*sdl.Texture, error) {
	rw := sdl.RWFromReadSeeker(r)
	bmp, err := sdl.LoadBMP_RW(rw, true)
	if err != nil {
		return nil, err
	}
	defer bmp.Free()

	bmpT, err := ren.CreateTextureFromSurface(bmp)
	if err != nil {
		return nil, err
	}

	return bmpT, nil
}

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

	file, err := os.Open("test.bmp")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bmp, err := LoadTexture(ren, file)
	if err != nil {
		panic(err)
	}
	defer bmp.Destroy()

	ren.Clear()
	ren.Copy(bmp, nil, nil)
	ren.Present()

	time.Sleep(3 * time.Second)
}
