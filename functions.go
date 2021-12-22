package functions

func Filter[T any](items []T, keep func(T) bool) []T {
	n := 0
	for _, x := range items {
		if keep(x) {
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
