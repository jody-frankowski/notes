package reverse

import "testing"

func TestReverse(t *testing.T) {
	input := [5]int{1, 2, 3, 4, 5}
	want := [5]int{5, 4, 3, 2, 1}

	reverse(&input)
	if input != want {
		t.Errorf("have %v want %v", input, want)
	}
}
