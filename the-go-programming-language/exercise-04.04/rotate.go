// Exercise 4.4: Write a version of rotate that operates in a single pass.
package rotate

func rotate(s []int, numberOfRotation int) {
	new := make([]int, len(s))

	copy(new[len(s)-numberOfRotation:], s[:numberOfRotation])
	copy(new[:], s[numberOfRotation:])

	copy(s, new)
}
