package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type testCase struct {
		desc     string
		nums     []int
		expected []int
	}
	testCases := []testCase{
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: []int{1},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			root := &TreeNode{Val: 1}
			got := inorderTraversal(root)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected: %+v, got: %+v", tc.expected, got)
			}
		})
	}

}
