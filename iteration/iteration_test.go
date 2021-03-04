package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("should pass on custom count", func(t *testing.T) {
		result := Repeat("a", 7)
		expected := "aaaaaaa"

		if expected != result {
			t.Errorf("expected %q got %q", expected, result)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	result := Repeat("na", 2)

	fmt.Println("ba" + result)

	// Output: banana
}
