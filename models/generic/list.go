package generic

// RemoveDuplicatesFromList remove duplicate items from a generic list.
// Order of items is not preserved.
func RemoveDuplicatesFromList[T comparable](in []T) (out []T) {
	out = []T{}
	m := make(map[T]bool)

	var itemExists bool
	for _, x := range in {
		itemExists, _ = m[x]
		if !itemExists {
			m[x] = true
			out = append(out, x)
		}
	}

	return out
}

//TODO: Add a non-destructive version of this function to preserve items order.
