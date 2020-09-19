package popcount

import "testing"

func TestPopCount(t *testing.T) {
	compareCount := func(input uint64, want int) func(*testing.T) {
		return func(t *testing.T) {
			got := PopCount(input)
			want := want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}

	t.Run("11", compareCount(11, 3))
	t.Run("0", compareCount(0, 0))
	t.Run("uint64 max", compareCount(0xFFFFFFFFFFFFFFFF, 64))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFF)
	}
}

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOriginal(0xFFFFFFFF)
	}
}
