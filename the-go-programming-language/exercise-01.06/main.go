// Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding more
// values to palette and then displaying them by changing the third argument of SetColorIndex in
// some interesting way.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var palette = []color.Color{
	color.White,
	color.RGBA{0xd1, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x66, 0x22, 0xff},
	color.RGBA{0xff, 0xda, 0x21, 0xff},
	color.RGBA{0x33, 0xdd, 0x00, 0xff},
	color.RGBA{0x11, 0x33, 0xcc, 0xff},
	color.RGBA{0x22, 0x00, 0x66, 0xff},
	color.RGBA{0x33, 0x00, 0x44, 0xff},
}

const (
	whiteIndex  = 0 // first color in palette
	RedIndex    = 1 // first color in palette
	OrangeIndex = 2 // first color in palette
	YellowIndex = 3 // first color in palette
	GreenIndex  = 4 // first color in palette
	BlueIndex   = 5 // first color in palette
	IndigoIndex = 6 // first color in palette
	VioletIndex = 7 // first color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
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
		var color uint8 = RedIndex
		pixel := 0
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), color%8)
			if pixel%20 == 0 {
				color++
			}
			pixel++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
