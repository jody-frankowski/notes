package main

import (
	"math"
	"testing"
)

func TestFeetToMeters(t *testing.T) {
	tests := map[string]struct {
		input Feet
		want  Meters
	}{
		"Simple Test": {input: 1, want: 0.3048},
		"10 Feet":     {input: 10, want: 3.048},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := FeetToMeters(test.input)
			want := test.want
			if math.Abs(float64(got)-float64(want)) > 0.001 {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestMetersToFeet(t *testing.T) {
	tests := map[string]struct {
		input Meters
		want  Feet
	}{
		"Simple Test": {input: 1, want: 3.2808},
		"10 Feet":     {input: 10, want: 32.808},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := MetersToFeet(test.input)
			want := test.want
			if math.Abs(float64(got)-float64(want)) > 0.001 {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
