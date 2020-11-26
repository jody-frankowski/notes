package reverse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	input := []byte("Hello 世界")
	want := []byte("界世 olleH")

	fmt.Printf("%v\n", input)
	reverse(input)
	fmt.Printf("%v\n", input)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("have %v want %v", input, want)
	}
}
