package main

import (
	"crypto/sha256"
	"testing"
)

func TestCountSHA256BitsDifferent(t *testing.T) {
	test := func(a [32]uint8, b [32]uint8, want int) func(*testing.T) {
		return func(t *testing.T) {
			have := countSHA256BitsDifferent(a, b)
			if have != want {
				t.Errorf("have %v want %v", have, want)
			}
		}
	}

	t.Run("0 differences", test([32]uint8{1, 2, 3}, [32]uint8{1, 2, 3}, 0))
	t.Run("1 difference", test([32]uint8{1, 2, 3}, [32]uint8{1, 2, 3, 2}, 1))
	t.Run("2 differences", test([32]uint8{1, 2, 3}, [32]uint8{3, 2, 3, 2}, 2))
	t.Run("3 differences", test([32]uint8{1, 2, 3, 4}, [32]uint8{3, 2, 3, 2}, 3))
	t.Run("125 differences", test(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")), 125))
}
