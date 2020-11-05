// Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same
// fractal using four different representations of numbers: complex64, complex128, big.Float, and
// big.Rat. (The latter two types are found in the math/big package. Float uses arbitrary but
// bounded-precision floating-point; Rat uses unbounded-precision rational numbers.) How do they
// compare in performance and memory usage? At what zoom levels do rendering artifacts become
// visible?
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleMandelbrot)
	http.ListenAndServe(":8000", nil)
}

func handleMandelbrot(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		x = 0
	}
	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		y = 0
	}
	zoom, err := strconv.ParseFloat(r.URL.Query().Get("zoom"), 64)
	if err != nil {
		zoom = 1
	}
	writeMandelbrot(w, x, y, zoom)
}

func writeMandelbrot(w io.Writer, x, y, zoom float64) {
	var xmin, ymin, xmax, ymax = x - (2 / zoom), y - (2 / zoom), x + (2 / zoom), y + (2 / zoom)
	const (
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := bigFloatComplex{r: big.NewFloat(x), i: big.NewFloat(y)}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotBigFloatComplex(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrotBigFloatComplex(z bigFloatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	v := bigFloatComplex{r: big.NewFloat(0), i: big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		v = bigFloatComplexAdd(bigFloatComplexMul(v, v), z)
		if z.r.Cmp(big.NewFloat(-1)) == 0 && z.i.Cmp(big.NewFloat(1)) == 0 {
			fmt.Println(v)
		}
		if bigFloatComplexAbs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

type bigFloatComplex struct {
	r *big.Float
	i *big.Float
}

func (b bigFloatComplex) String() string {
	return fmt.Sprintf("r: %v, i: %v", b.r, b.i)
}

func bigFloatComplexIsEqual(a, b bigFloatComplex) bool {
	return a.r.Cmp(b.r) == 0 && a.i.Cmp(b.i) == 0
}

func bigFloatComplexAdd(a, b bigFloatComplex) bigFloatComplex {
	r := new(big.Float)
	i := new(big.Float)

	return bigFloatComplex{r: r.Add(a.r, b.r), i: i.Add(a.i, b.i)}
}

func bigFloatComplexSub(a, b bigFloatComplex) bigFloatComplex {
	r := new(big.Float)
	i := new(big.Float)

	return bigFloatComplex{r: r.Sub(a.r, b.r), i: i.Sub(a.i, b.i)}
}

func bigFloatComplexMul(a, b bigFloatComplex) bigFloatComplex {
	r := new(big.Float)
	i := new(big.Float)
	tmp := new(big.Float)
	tmp2 := new(big.Float)

	r.Sub(tmp.Mul(a.r, b.r), tmp2.Mul(a.i, b.i))
	i.Add(tmp.Mul(a.r, b.i), tmp2.Mul(a.i, b.r))
	return bigFloatComplex{r: r, i: i}
}

func bigFloatComplexAbs(a bigFloatComplex) float64 {
	tmp := new(big.Float)
	tmp2 := new(big.Float)

	tmp.Mul(a.r, a.r)
	tmp2.Mul(a.i, a.i)
	tmp.Add(tmp, tmp2)
	tmp.Sqrt(tmp)

	value, _ := tmp.Float64()
	return value
}
