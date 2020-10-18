package main

import (
	"image/color"
	"reflect"
	"testing"
)

func TestAvgColor(t *testing.T) {
	tests := []struct {
		input []color.Color
		want  color.Color
	}{
		{
			input: []color.Color{
				color.RGBA{0, 0, 0, 0},
				color.RGBA{0, 0, 0, 0},
				color.RGBA{0, 0, 0, 0},
				color.RGBA{0, 0, 0, 0},
			},
			want: color.RGBA{0, 0, 0, 0},
		},
		{
			input: []color.Color{
				color.RGBA{255, 255, 255, 255},
				color.RGBA{255, 255, 255, 255},
				color.RGBA{255, 255, 255, 255},
				color.RGBA{255, 255, 255, 255},
			},
			want: color.RGBA{255, 255, 255, 255},
		},
		{
			input: []color.Color{
				color.RGBA{90, 120, 150, 70},
				color.RGBA{10, 40, 30, 233},
				color.RGBA{20, 60, 40, 1},
				color.RGBA{30, 80, 255, 101},
			},
			want: color.RGBA{37, 75, 118, 101},
		},
	}

	for _, test := range tests {
		got := avgColor(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
