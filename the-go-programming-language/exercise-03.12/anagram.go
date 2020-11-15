// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other, that
// is, they contain the same letters in a different order.
package anagram

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isAnagram(a string, b string) bool {
	return sortString(a) == sortString(b)
}
