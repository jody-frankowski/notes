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
