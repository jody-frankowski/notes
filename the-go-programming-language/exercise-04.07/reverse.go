// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?
package reverse

func reverseBytes(slice []byte, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		slice[start], slice[end] = slice[end], slice[start]
	}
}

func reverse(slice []byte) {
	reverseBytes(slice, 0, len(slice)-1)

	var sequenceStart int
	var sequenceStarted bool = false
	for i, byte := range slice {
		if byte>>7 == 0b0 {
			continue
		} else if !sequenceStarted {
			sequenceStart = i
			sequenceStarted = true
		}

		if byte>>6 != 0b10 {
			reverseBytes(slice, sequenceStart, i)
			sequenceStarted = false
		}
	}
}
