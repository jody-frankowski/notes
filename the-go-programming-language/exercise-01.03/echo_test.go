// Benchmark the different echo implementations
// Run with `go test -bench .`
package echo

import (
	"os"
	"testing"
)

var args = []string{"aaaa", "bbbb", "cccc", "1111", "2222", "3333"}

func BenchmarkEcho1(t *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < t.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(t *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < t.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(t *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	for i := 0; i < t.N; i++ {
		echo3(args)
	}
}
