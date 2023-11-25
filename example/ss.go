package main

import (
	"github.com/justhx0r/screenshot"
	"image/png"
	"os"
)

//garble:controlflow flatten_passes=1 junk_jumps=201 block_splits=max flatten_hardening=xor
func main() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./ss.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}
