package binarysearch_test

import (
	"testing"

	"github.com/wxdis-public/goalgos/binarysearch"
)

func TestBinarySearch(t *testing.T) {
	tt := map[string]struct {
		values   []int
		target   int
		expected int
	}{
		"Target present in short array": {
			values:   []int{1, 2, 3, 4, 5},
			target:   4,
			expected: 3,
		},
		"Target not present in short array": {
			values:   []int{1, 2, 3, 5},
			target:   4,
			expected: -1,
		},
		"Target above upper bound of array": {
			values:   []int{1, 2, 3, 5},
			target:   7,
			expected: -1,
		},
		"Target below lower bound of array": {
			values:   []int{3, 4, 5, 6},
			target:   1,
			expected: -1,
		},
		"Target is first character": {
			values:   []int{1, 2, 3, 5},
			target:   1,
			expected: 0,
		},
		"Target is last character": {
			values:   []int{1, 2, 3, 5},
			target:   5,
			expected: 3,
		},
		"Nil array": {
			values:   nil,
			target:   0,
			expected: -1,
		},
	}

	for tn, tc := range tt {
		res := binarysearch.BinarySearch(tc.values, tc.target)
		if res != tc.expected {
			t.Errorf("%s failed: expected '%v' but got '%v'", tn, tc.expected, res)
		}

	}
}
