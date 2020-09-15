package main

import (
	"math"
	"testing"
)

func TestGramsToPounds(t *testing.T) {
	tests := map[string]struct {
		input Grams
		want  Pounds
	}{
		"Simple Test": {input: 1000, want: 2.204},
		"3000 Grams":  {input: 3000, want: 6.613},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := GramsToPounds(test.input)
			want := test.want
			if math.Abs(float64(got)-float64(want)) > 0.001 {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestPoundsToGrams(t *testing.T) {
	tests := map[string]struct {
		input Pounds
		want  Grams
	}{
		"Simple Test": {input: 1, want: 453.59},
		"10 Pounds":   {input: 10, want: 4535.9},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := PoundsToGrams(test.input)
			want := test.want
			if math.Abs(float64(got)-float64(want)) > 0.001 {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
