package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

// Exercise:

func TestCount(t *testing.T) {
	count := strings.Count("abcdabcd", "abc")
	expected := 2

	if count != expected {
		t.Errorf("go %q want %q", count, expected)
	}
}

func TestReplaceAll(t *testing.T) {
	replaced := strings.ReplaceAll("123 abc 123", "123", "456")
	expected := "456 abc 456"

	if replaced != expected {
		t.Errorf("go %q want %q", replaced, expected)
	}
}
