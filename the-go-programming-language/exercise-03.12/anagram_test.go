package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	test := func(a string, b string, want bool) func(*testing.T) {
		return func(t *testing.T) {
			have := isAnagram(a, b)
			if have != want {
				t.Errorf("have %v want %v", have, want)
			}
		}
	}

	t.Run("Simple True", test("abc", "cba", true))
	t.Run("Simple False", test("abc", "ab", false))
	t.Run("Repeating letters", test("aaabbbccc", "abcabcbac", true))
	t.Run("Repeating letters but not enough", test("aabbbccc", "abcabcbac", false))
}
