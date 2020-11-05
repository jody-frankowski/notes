package main

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestBigFloatComplexAdd(t *testing.T) {
	tests := []struct {
		a    bigFloatComplex
		b    bigFloatComplex
		want bigFloatComplex
	}{
		{
			a:    bigFloatComplex{r: big.NewFloat(2), i: big.NewFloat(3)},
			b:    bigFloatComplex{r: big.NewFloat(5), i: big.NewFloat(7)},
			want: bigFloatComplex{r: big.NewFloat(7), i: big.NewFloat(10)},
		},
		{
			a:    bigFloatComplex{r: big.NewFloat(0), i: big.NewFloat(1)},
			b:    bigFloatComplex{r: big.NewFloat(1), i: big.NewFloat(1)},
			want: bigFloatComplex{r: big.NewFloat(1), i: big.NewFloat(2)},
		},
	}

	for _, test := range tests {
		have := bigFloatComplexAdd(test.a, test.b)

		if !bigFloatComplexIsEqual(have, test.want) {
			t.Errorf("have %v want %v", have, test.want)
		}
	}
}

func TestBigFloatComplexSub(t *testing.T) {
	tests := []struct {
		a    bigFloatComplex
		b    bigFloatComplex
		want bigFloatComplex
	}{
		{
			a:    bigFloatComplex{r: big.NewFloat(5), i: big.NewFloat(7)},
			b:    bigFloatComplex{r: big.NewFloat(2), i: big.NewFloat(7)},
			want: bigFloatComplex{r: big.NewFloat(3), i: big.NewFloat(0)},
		},
		{
			a:    bigFloatComplex{r: big.NewFloat(2), i: big.NewFloat(0)},
			b:    bigFloatComplex{r: big.NewFloat(3), i: big.NewFloat(0)},
			want: bigFloatComplex{r: big.NewFloat(-1), i: big.NewFloat(0)},
		},
	}

	for _, test := range tests {
		have := bigFloatComplexSub(test.a, test.b)

		if !bigFloatComplexIsEqual(have, test.want) {
			t.Errorf("have %v want %v", have, test.want)
		}
	}
}

func TestBigFloatComplexMul(t *testing.T) {
	tests := []struct {
		a    bigFloatComplex
		b    bigFloatComplex
		want bigFloatComplex
	}{
		{
			a:    bigFloatComplex{r: big.NewFloat(2), i: big.NewFloat(3)},
			b:    bigFloatComplex{r: big.NewFloat(5), i: big.NewFloat(7)},
			want: bigFloatComplex{r: big.NewFloat(-11), i: big.NewFloat(29)},
		},
		{
			a:    bigFloatComplex{r: big.NewFloat(4), i: big.NewFloat(0)},
			b:    bigFloatComplex{r: big.NewFloat(2), i: big.NewFloat(3)},
			want: bigFloatComplex{r: big.NewFloat(8), i: big.NewFloat(12)},
		},
	}

	for _, test := range tests {
		have := bigFloatComplexMul(test.a, test.b)

		if !bigFloatComplexIsEqual(have, test.want) {
			t.Errorf("have %v want %v", have, test.want)
		}
	}
}

func TestBigFloatComplexAbs(t *testing.T) {
	tests := []struct {
		input bigFloatComplex
		want  float64
	}{
		{
			input: bigFloatComplex{r: big.NewFloat(2), i: big.NewFloat(0)},
			want:  2,
		},
		{
			input: bigFloatComplex{r: big.NewFloat(0), i: big.NewFloat(2)},
			want:  2,
		},
		{
			input: bigFloatComplex{r: big.NewFloat(2.8284271247461903), i: big.NewFloat(2.8284271247461903)},
			want:  4,
		},
	}

	for _, test := range tests {
		have := bigFloatComplexAbs(test.input)

		if have != test.want {
			fmt.Println(math.Sqrt(2) + math.Sqrt(2))
			t.Errorf("have %v want %v", have, test.want)
		}
	}
}
