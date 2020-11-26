// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see
// unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import "unicode"

func squashAdjacentUnicodeSpaces(s *[]byte) {
	if len(*s) == 0 {
		return
	}
	runes := []rune(string(*s))
	j := 1
	for i := 1; i < len(runes); i++ {
		if !unicode.IsSpace(runes[i-1]) && unicode.IsSpace(runes[i]) {
			runes[j] = []rune(" ")[0]
			j++
		} else if !unicode.IsSpace(runes[i]) {
			runes[j] = runes[i]
			j++
		}
	}
	runes = runes[:j]
	*s = []byte(string(runes))
}
