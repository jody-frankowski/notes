package main

import "fmt"

type Feet float64
type Meters float64

func (feet Feet) String() string {
	return fmt.Sprintf("%gft", feet)
}

func (meters Meters) String() string {
	return fmt.Sprintf("%gm", meters)
}

func FeetToMeters(feet Feet) Meters {
	return Meters(feet / 3.2808)
}

func MetersToFeet(meters Meters) Feet {
	return Feet(meters * 3.2808)
}
