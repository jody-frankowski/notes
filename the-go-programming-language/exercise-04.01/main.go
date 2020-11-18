// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256
// hashes. (See PopCount from Section 2.6.2.)
package main

func countSHA256BitsDifferent(a [32]uint8, b [32]uint8) int {
	differentCount := 0

	for i := range a {
		for j := 0; j < 8; j++ {
			if (a[i] >> j & 1) != (b[i] >> j & 1) {
				differentCount++
			}
		}
	}

	return differentCount
}
