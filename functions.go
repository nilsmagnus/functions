package functions

func Associate[T comparable](items []T) map[T]T {
	result := make(map[T]T, len(items))
	for _, item := range items {
		result[item] = item
	}
	return result
}

func AssociateBy[T any, O comparable](items []T, key func(T) O) map[O]T {
	result := make(map[O]T, len(items))
	for _, item := range items {
		result[key(item)] = item
	}
	return result
}

func Any[T any](items []T, predicate func(T) bool) bool {
	for _, t := range items {
		if predicate(t) {
			return true
		}
	}
	return false
}

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

func Map[T any, O any](in []T, transform func(T) O) []O {
	out := make([]O, len(in))
	for i, t := range in {
		out[i] = transform(t)
	}
	return out
}

func Reduce[T any, O any](in []T, acc func(O, T) O, initial O) (result O) {

	result = initial

	for _, t := range in {
		result = acc(result, t)
	}

	return result
}
