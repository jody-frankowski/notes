// Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string
// concatenation.
package comma

import (
	"bytes"
)

func comma(s string) string {
	var buf bytes.Buffer

	start := len(s) % 3
	if start == 0 {
		// Don't start at 3 for the empty string
		if len(s) > 0 {
			start = 3
		}
	}
	buf.WriteString(s[0:start])
	for i := start; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
