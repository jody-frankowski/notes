package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5}
	want := []int{2, 3, 4, 5, 0, 1}

	rotate(input, 2)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("have %v want %v", input, want)
	}
}
