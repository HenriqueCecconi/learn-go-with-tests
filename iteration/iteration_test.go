package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("test repetition with default amount", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaaaa"
		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
	t.Run("test user defined repetition amount", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		if got != want {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	concatenated := Repeat("a", 8)
	fmt.Println(concatenated)
	// Output: aaaaaaaa
}
