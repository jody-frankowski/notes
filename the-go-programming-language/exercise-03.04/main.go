// Exercise 3.4: Following the approach of the Lissajous example in Section 1.7, construct a
// webserver that computes surfaces and writes SVG data to the client. The server must set the
// Content-Type header like this:
// w.Header().Set("Content-Type", "image/svg+xml")
// (This step was not required in the Lissajous example because the server uses standard heuristics
// to recognize common formats like PNG from the first 512 bytes of the response, and generates the
// proper header.) Allow the client to specify values like height, width, and color as HTTP request
// parameters.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		color := "grey"
		width := 600
		height := 320

		keys, ok := r.URL.Query()["color"]
		if ok {
			color = keys[len(keys)-1]
		}
		keys, ok = r.URL.Query()["width"]
		if ok {
			request_width, err := strconv.Atoi(keys[len(keys)-1])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "400: %v", err)
				return
			} else {
				width = request_width
			}
		}
		keys, ok = r.URL.Query()["height"]
		if ok {
			request_height, err := strconv.Atoi(keys[len(keys)-1])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "400: %v", err)
				return
			} else {
				height = request_height
			}
		}

		w.Header().Set("Content-Type", "image/svg+xml")

		generateSVG(w, width, height, color)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generateSVG(out io.Writer, width, height int, color string) {
	cells := 100                            // number of grid cells
	xyrange := 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit
	angle := math.Pi / 6                    // angle of x, y axes (=30°)

	var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

	f := func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return math.Sin(r) / r
	}

	corner := func(i, j int) (float64, float64) {
		// Find point (x,y) at corner of cell (i,j).
		x := xyrange * (float64(i)/float64(cells) - 0.5)
		y := xyrange * (float64(j)/float64(cells) - 0.5)

		// Compute surface height z.
		z := f(x, y)

		// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
		sx := float64(width)/2 + (x-y)*cos30*xyscale
		sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
		return sx, sy
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}
