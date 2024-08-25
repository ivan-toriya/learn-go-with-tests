package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := strings.Repeat("a", 5)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// BenchmarkBuiltInRepeat is a benchmark test for the built-in strings.Repeat function.
// It's approximately 3 times faster on my local env than the custom Repeat function.
func BenchmarkBuiltInRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
