package utils

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	type testCase struct {
		desc     string
		nums     []int
		expected [][]int
	}
	testCases := []testCase{
		testCase{
			desc: "1",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			perms := permute(tc.nums)
			t.Errorf("expected: %+v, got: %+v", tc.expected, perms)
		})
	}

}
