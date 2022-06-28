package binarysearch

import "golang.org/x/exp/constraints"

// BinarySearch : Check if a value 'target' is present in the supplied list 'vals'.
// Returns the index of the value if found, and '-1' if the value is not found in
// the list.
func BinarySearch[T constraints.Ordered](vals []T, target T) int {
	l, h := 0, len(vals)-1
	for l <= h {
		p := (l + h) >> 1
		if vals[p] == target {
			return p
		} else if vals[p] < target {
			l = p + 1
		} else {
			h = p - 1
		}
	}
	return -1
}
