package functions

import (
	"fmt"
	"testing"
)

func Test_Associate(t *testing.T) {
	numbers := testNumbers(100)

	mapd := Associate(numbers)

	if len(mapd) != len(numbers) {
		t.Errorf("resulting map and starting items should have same length")
	}

}

func Test_AssociateBy(t *testing.T) {
	numbers := testNumbers(100)

	mapd := AssociateBy(numbers, func(t int) string {
		return fmt.Sprintf("%d", t)
	})

	if len(mapd) != len(numbers) {
		t.Errorf("resulting map and starting items should have same length")
	}

}

func Test_Reduce(t *testing.T) {

	numbers := testNumbers(10)

	accumulator := func(s string, i int) string {
		return s + fmt.Sprintf("%d", i)
	}
	reduction := Reduce(numbers, accumulator, "hei")

	if reduction != "hei0123456789" {
		t.Errorf("Got unexpected string :%s", reduction)
	}
}

func Test_Filter(t *testing.T) {
	numbers := testNumbers(10)

	predicate := func(t int) bool {
		return t > 5
	}
	filtered := Filter(numbers, predicate)

	if len(filtered) != 4 {
		t.Errorf("Expected result of size 5, was %d", len(filtered))
	}

}

func Test_Map(t *testing.T) {

	numbers := testNumbers(10)

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

func Test_Any(t *testing.T) {
	numbers := testNumbers(100)

	if Any(numbers, func(t int) bool { return t > 100 }) {
		t.Errorf("100 is not in slice")
	}

	if !Any(numbers, func(t int) bool { return t == 10 }) {
		t.Errorf("10 is in slice")
	}

}

func Test_Contains(t *testing.T) {
	numbers := testNumbers(100)

	if Contains(numbers, func(t int) bool { return t > 100 }) {
		t.Errorf("100 is not in slice")
	}

	if !Any(numbers, func(t int) bool { return t == 10 }) {
		t.Errorf("10 is in slice")
	}

}

//testNumbers is an internal function for testing-purposes
func testNumbers(size int) []int {
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}
	return numbers
}
