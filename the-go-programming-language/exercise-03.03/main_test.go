package main

import (
	"math"
	"testing"
)

func TestIsFinite(t *testing.T) {
	tests := map[string]struct {
		input float64
		want  bool
	}{
		"Number":            {3.0, true},
		"Positive Infinity": {math.Inf(1), false},
		"Negative Infinity": {math.Inf(-1), false},
		"NaN":               {math.NaN(), false},
	}

	for name, test := range tests {
		got := isFinite(test.input)
		want := test.want

		if got != want {
			t.Errorf("%v: got %v want %v", name, got, want)
		}
	}
}

func TestMapRange(t *testing.T) {
	tests := map[string]struct {
		sourceMin   float64
		sourceMax   float64
		sourceValue float64
		targetMin   float64
		targetMax   float64
		targetValue float64
	}{
		"Simple test":        {100, 200, 150, 500, 700, 600},
		"Negative min test":  {-300, 0, -300, -200, 0, -200},
		"Negative half test": {-300, -100, -100, -200, 0, 0},
		"Negative max test":  {-300, 0, 0, -100, -50, -50},
		"Across 0":           {-300, 500, 100, -200, 400, 100},
		"RGB min":            {-0.1, 0.1, -0.1, 0, 255, 0},
		"RGB half":           {-0.1, 0.1, 0.0, 0, 255, 127.5},
		"RGB max":            {-0.1, 0.1, 0.1, 0, 255, 255},
	}

	for name, test := range tests {
		got := mapRange(test.sourceMin, test.sourceMax, test.sourceValue, test.targetMin, test.targetMax)
		want := test.targetValue

		if got != want {
			t.Errorf("%v: got %v want %v", name, got, want)
		}
	}
}
