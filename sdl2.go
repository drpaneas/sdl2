package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b, a byte
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4
	if index <= len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
		pixels[index+3] = c.a
	}
}

func main() {
	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_BORDERLESS)
	if err != nil {
		errMsg := fmt.Errorf("\n[ERROR]:\tWindow Creation failed.\n[E_LOG]:\t%s", err)
		fmt.Println(errMsg)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		errMsg := fmt.Errorf("\n[ERROR]:\tRenderer Creation failed.\n[E_LOG]:\t%s", err)
		fmt.Println(errMsg)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		errMsg := fmt.Errorf("\n[ERROR]:\tTexture Creation failed.\n[E_LOG]:\t%s", err)
		fmt.Println(errMsg)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{255, 100, 90, 0}, pixels)
		}
	}

	err = tex.Update(nil, pixels, winWidth*4)
	if err != nil {
		errMsg := fmt.Errorf("\n[ERROR]:\tTexture Update failed.\n[E_LOG]:\t%s", err)
		fmt.Println(errMsg)
		return
	}

	err = renderer.Copy(tex, nil, nil)
	if err != nil {
		errMsg := fmt.Errorf("\n[ERROR]:\tRenderer failed to copy the texture.\n[E_LOG]:\t%s", err)
		fmt.Println(errMsg)
		return
	}

	renderer.Present()

	sdl.Delay(2000) // Περίμενε για 2 δευτερόλεπτα
}
