package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("expected %d but got %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(42, 1)
	fmt.Println(sum)
	// Output: 43
}
