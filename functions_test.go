package functions

import (
	"fmt"
	"testing"
)

func Test_Filter(t *testing.T) {
	numbers := make([]int, 10)
	for i, _ := range numbers {
		numbers[i] = i
	}

	filtered := Filter(numbers, func(t int) bool {
		return t > 5
	})

	if len(filtered) != 4 {
		t.Errorf("Expected result of size 5, was %d", len(filtered))
	}

}
func Test_Map(t *testing.T) {

	numbers := make([]int, 10)
	for i, _ := range numbers {
		numbers[i] = i*10 + 14
	}

	transform := func(t int) string {
		return fmt.Sprintf("%d", t)
	}
	strings := Map(numbers, transform)

	if len(strings) != len(numbers) {
		t.Errorf("input and output should be of equal lengths")
	}

	for i, transformed := range strings {
		if transform(numbers[i]) != transformed {
			t.Errorf("Number wash not mapped correctly, was %s, expected %s", transformed, transform(numbers[i]))
		}
	}

}
