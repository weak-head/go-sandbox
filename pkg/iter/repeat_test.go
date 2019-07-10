package iter

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectRepeated := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepeated(t, expected, repeated)
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectRepeated(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("ab", 3)
	fmt.Println(repeated)
	// Output: ababab
}
