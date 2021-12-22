package functions

import (
	"fmt"
	"testing"
)

func Test_ForEach(t *testing.T) {
	numbers := testNumbers(199)

	tmp := 0

	ForEach(numbers, func(i int) { tmp++ })

	if tmp != len(numbers) {
		t.Errorf("Action should have been run %d times, was only run %d times", len(numbers), tmp)
	}

}

func Test_ForEachIndexed(t *testing.T) {
	numbers := testNumbers(199)

	tmp := 0

	ForEachIndexed(numbers, func(value int, index int) { tmp += (index + value) })

	if tmp != 39402 {
		t.Errorf("expected 39402, was only  %d  ", tmp)
	}

}

func Test_First(t *testing.T) {
	numbers := testNumbers(199)

	predicate := func(i int) bool { return i > 77 }
	first, ok := First(numbers, predicate)

	if !ok {
		t.Errorf("should have found a value matching the predicate")
	}

	if *first != 78 {
		t.Errorf("First number should have been 78, was %d", *first)
	}
}

func Test_Last(t *testing.T) {
	numbers := testNumbers(42)

	predicate := func(i int) bool { return i%13 == 0 }
	last, ok := Last(numbers, predicate)

	if !ok {
		t.Errorf("should have found a value matching the predicate")
	}

	if *last != 39 {
		t.Errorf("First number should have been 39, was %d", *last)
	}
}

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

	if Contains(numbers, 100) {
		t.Errorf("100 is not in slice")
	}

	if !Contains(numbers, 10) {
		t.Errorf("10 is in slice")
	}

}

func Test_ContainsAll(t *testing.T) {
	numbers := testNumbers(100)
	needles := []int{1, 2, 57}

	if !ContainsAll(numbers, needles) {
		t.Errorf("all needles should be present")
	}

	needlesNotUnion := []int{1, 2, 199}

	if ContainsAll(numbers, needlesNotUnion) {
		t.Errorf("199 should not be present")
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
