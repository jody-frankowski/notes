// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

func eliminateAdjacentDuplicates(s *[]string) {
	if len(*s) == 0 {
		return
	}
	slice := *s
	previous := slice[0]
	j := 1
	for i := 1; i < len(slice); i++ {
		if slice[i] != previous {
			previous = slice[i]
			slice[j] = slice[i]
			j++
		}
	}
	*s = slice[:j]
}
