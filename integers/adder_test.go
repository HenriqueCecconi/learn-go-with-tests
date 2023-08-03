package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("expected '%d' but got '%d'", want, sum)
	}
}

func ExampleAdd() {
	sum := Add(4, 4)
	fmt.Println(sum)
	// Output: 8
}
