package main

import "fmt"

type Grams float64
type Pounds float64

func (grams Grams) String() string {
	return fmt.Sprintf("%g grams", grams)
}

func (pounds Pounds) String() string {
	return fmt.Sprintf("%g pounds", pounds)
}

func GramsToPounds(grams Grams) Pounds {
	return Pounds(grams / 453.59)
}

func PoundsToGrams(pounds Pounds) Grams {
	return Grams(pounds * 453.59)
}
