package randomGenerator

import (
	"fmt"
	"testing"
)

func TestGeradorAula(t *testing.T) {
	got := ExemploAula(10)
	want := [10]int{7, 5, 6, 1, 8, 0, 4, 2, 3, 7}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("want %d, get %d at position %d of the vector", want[i], v, i)
		}
		fmt.Printf("want %d, get %d at position %d of the vector", want[i], v, i)
	}

}
