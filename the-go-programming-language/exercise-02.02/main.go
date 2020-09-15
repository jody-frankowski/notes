// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads numbers
// from its command-line arguments or from the standard input if there are no arguments, and
// converts each number into units like temperature in Celsius and Fahrenheit, length in feet and
// meters, weight in pounds and kilograms, and the like.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", Feet(t), FeetToMeters(Feet(t)), Meters(t), MetersToFeet(Meters(t)))
		fmt.Printf("%s = %s, %s = %s\n", Grams(t), GramsToPounds(Grams(t)), Pounds(t), PoundsToGrams(Pounds(t)))
	}
}
