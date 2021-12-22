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

//Contains checks if predicate is true for any item in items.
func Contains[T any](items []T, needle T) bool {
	for _, t := range items {
		if needle == t {
			return true
		}
	}
	return false
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

// Reduce reduces the input slice, in, to a resulting O using the accumulator-function
func Reduce[T any, O any](in []T, acc func(O, T) O, initial O) (result O) {

	result = initial

	for _, t := range in {
		result = acc(result, t)
	}

	return result
}
