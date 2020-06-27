package utils

import (
	"testing"
)

func TestFindOrder(t *testing.T) {
	type testCase struct {
		reqs      [][]int
		courseNum int
		expected  []int
	}
	testCases := []testCase{
		testCase{
			reqs: [][]int{
				{1, 0},
				{3, 1},
				{3, 2},
				{0, 2},
			},
			courseNum: 4,
			expected:  []int{2, 0, 1, 3},
		},
		testCase{
			reqs: [][]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{3, 2},
			},
			courseNum: 4,
			expected:  []int{0, 1, 2, 3},
		},
	}
	for _, tc := range testCases {
		order := findOrder(tc.courseNum, tc.reqs)
		for i := range order {
			if order[i] != tc.expected[i] {
				t.Errorf("expected: %+v, got: %+v", tc.expected, order)
			}
		}
	}

}
