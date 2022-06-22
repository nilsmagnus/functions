package functions

// Associate creates a map of comparables where key == value
func Associate[T comparable](items []T) map[T]T {
	result := make(map[T]T, len(items))
	for _, item := range items {
		result[item] = item
	}
	return result
}

// AssociateBy returns a map of key O with value T. If for any items A, B key(A)== key(B), only B is in the resulting map
func AssociateBy[T any, O comparable](items []T, key func(T) O) map[O]T {
	result := make(map[O]T, len(items))
	for _, item := range items {
		result[key(item)] = item
	}
	return result
}

// Any checks if predicate is true for any item in items.
func Any[T any](items []T, predicate func(T) bool) bool {
	for _, t := range items {
		if predicate(t) {
			return true
		}
	}
	return false
}

// First returns the first element of items matching predicate p. Returns item and a bool, the bool indicates if anything was found
func First[T any](items []T, predicate func(T) bool) (*T, bool) {
	for _, i := range items {
		if predicate(i) {
			return &i, true
		}
	}
	return nil, false
}

// Last returns the last element of items matching predicate p. Returns item and a bool, the bool indicates if anything was found
func Last[T any](items []T, predicate func(T) bool) (*T, bool) {
	// search backwards in items
	for i := len(items) - 1; i > 0; i-- {
		if predicate(items[i]) {
			return &items[i], true
		}
	}
	return nil, false
}

// ForEach runs the action on every item of items
func ForEach[T any](items []T, action func(a T)) {
	for _, item := range items {
		action(item)
	}
}

// ForEachIndexed runs the action on every item of items
func ForEachIndexed[T any](items []T, action func(a T, index int)) {
	for i, item := range items {
		action(item, i)
	}
}

func Distinct[T comparable](items []T) []T {
	vals := map[T]T{}
	for _, item := range items {
		vals[item] = item
	}

	res := make([]T, 0, len(vals))

	for _, t := range vals {
		res = append(res, t)
	}
	return res

}

//Contains checks needle is in items.
func Contains[T comparable](items []T, needle T) bool {
	for _, t := range items {
		if needle == t {
			return true
		}
	}
	return false
}

//ContainsAll checks if all needles exists in items
func ContainsAll[T comparable](items []T, needles []T) bool {
	for _, needle := range needles {
		if !Contains(items, needle) {
			return false
		}
	}
	return true
}

//Filter returns a slice of items matching the predicate
func Filter[T any](items []T, predicate func(T) bool) []T {
	n := 0
	for _, x := range items {
		if predicate(x) {
			items[n] = x
			n++
		}
	}
	return items[:n]
}

// Map converts a slice of T to a slice of O using the transform-function
func Map[T any, O any](in []T, transform func(T) O) []O {
	out := make([]O, len(in))
	for i, t := range in {
		out[i] = transform(t)
	}
	return out
}
// MapIndexed converts a slice of T to a slice of O using the transform-function
func MapIndexed[T any, O any](in []T, transform func(T,int) O) []O {
	out := make([]O, len(in))
	for i, t := range in {
		out[i] = transform(t, i)
	}
	return out
}

// Reduce reduces the input slice, in, to a resulting O using the accumulator-function
func Reduce[T any, O any](in []T, acc func(O, T) O, initial O) (result O) {

	result = initial

	for _, t := range in {
		result = acc(result, t)
	}

	return result
}

// Reverse a slice, return new slice with reversed values
func Reverse[T any](s []T)[]T{
	out := make([]T, len(s))
	
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = s[j], s[i]
	}
	return out
}
