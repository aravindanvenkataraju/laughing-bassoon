// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

//!-main
// Packages not needed by version in book.

//!+main

var palette = []color.Color{color.Black, color.RGBA{0x20, 0xff, 0x20, 0xff}, color.RGBA{0xff, 0x20, 0x20, 0xff}, color.RGBA{0x20, 0x20, 0xff, 0xff}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	redIndex   = 2
	blueIndex  = 3
)

func lissajous(out io.Writer, paramDelay int, paramCycles int) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(paramCycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.3), size+int(y*size+0.3),
				greenIndex)
			img.SetColorIndex(size+int(x*size+0.6), size+int(y*size+0.6),
				redIndex)
			img.SetColorIndex(size+int(x*size+1), size+int(y*size+1),
				blueIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, paramDelay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
