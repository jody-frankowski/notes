package main

import (
	"reflect"
	"testing"
)

func TestSquashAdjacentUnicodeSpaces(t *testing.T) {
	test := func(inputString string, wantString string) func(*testing.T) {
		return func(t *testing.T) {
			input := []byte(inputString)
			want := []byte(wantString)
			squashAdjacentUnicodeSpaces(&input)
			if !reflect.DeepEqual(input, want) {
				t.Errorf("have %v want %v", input, want)
			}
		}
	}

	t.Run("Simple", test("Hello \t\n\v\f\r 世界", "Hello 世界"))
	t.Run("Empty string", test("", ""))
	t.Run("1 space", test(" ", " "))
	t.Run("2 space", test("  ", " "))
	t.Run("Only whitespace", test(" \n \t \v \r \f ", " "))
	t.Run("Whitespace at the beginning", test("  a", " a"))
	t.Run("Whitespace at the end", test("a  ", "a "))
	t.Run("Whitespace at both end", test("  a  ", " a "))
}
