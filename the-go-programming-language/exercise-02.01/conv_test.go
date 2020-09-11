package tempconv

import (
	"math"
	"testing"
)

func TestKToC(t *testing.T) {
	testKToC := func(kelvinInput Kelvin, celsiusExpected Celsius) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			got := KToC(kelvinInput)
			want := Celsius(celsiusExpected)

			if math.Abs(float64(got)-float64(want)) > 0.1 {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}

	t.Run("Absolute Zero", testKToC(0, AbsoluteZeroC))
	t.Run("Freezing", testKToC(273.15, FreezingC))
	t.Run("Boiling", testKToC(373.15, BoilingC))
}

func TestCToK(t *testing.T) {
	testCToK := func(celsiusInput Celsius, kelvinExpected Kelvin) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			got := CToK(celsiusInput)
			want := kelvinExpected

			if math.Abs(float64(got)-float64(want)) > 0.1 {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}

	t.Run("Absolute Zero", testCToK(AbsoluteZeroC, 0))
	t.Run("Freezing", testCToK(FreezingC, 273.15))
	t.Run("Boiling", testCToK(BoilingC, 373.15))
}

func TestCToF(t *testing.T) {
	testCToF := func(celsiusInput Celsius, fahrenheitExpected Fahrenheit) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			got := CToF(celsiusInput)
			want := fahrenheitExpected

			if math.Abs(float64(got)-float64(want)) > 0.1 {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}

	t.Run("Absolute Zero", testCToF(AbsoluteZeroC, -459.67))
	t.Run("Freezing", testCToF(FreezingC, 32))
	t.Run("Boiling", testCToF(BoilingC, 212))
}
