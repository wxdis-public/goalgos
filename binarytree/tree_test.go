package tree_test

import (
	"testing"

	"github.com/wxdis-public/goalgos/tree"
)

func TestFind(t *testing.T) {
	tt := map[string]struct {
		values []int
		value  int
		expect bool
	}{
		"Value exists": {
			values: []int{2, 3, 4, 5, 6},
			value:  5,
			expect: true,
		},
		"Single value exists": {
			values: []int{1},
			value:  1,
			expect: true,
		},
		"Lower value doesn't exist": {
			values: []int{4, 5, 6, 7},
			value:  3,
			expect: false,
		},
		"Higher value doesn't exist": {
			values: []int{4, 5, 6, 7},
			value:  8,
			expect: false,
		},
		"Value doesn't exist in empty tree": {
			values: []int{},
			value:  8,
			expect: false,
		},
	}
	for tn, tc := range tt {
		testtree := tree.New(tc.values...)
		res := testtree.Find(tc.value)
		if tc.expect && res == nil {
			t.Errorf("%v: Expected to find value '%v' in the tree but didn't", tn, tc.value)
		} else if tc.expect && res != nil {
			if tc.value != res.Value {
				t.Errorf("%v: Expected to find value '%v' in the tree but instead found '%v'", tn, tc.value, res.Value)
			}
		} else if !tc.expect && res != nil {
			t.Errorf("%v: Didn't expect to find value '%v' in the tree but got '%v'", tn, tc.value, res.Value)
		}
	}
}
