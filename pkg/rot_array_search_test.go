package utils

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	type testCase struct {
		desc     string
		nums     []int
		target   int
		expected int
	}
	testCases := []testCase{
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   6,
			expected: 2,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: 0,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   8,
			expected: -1,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: 5,
		},
		testCase{
			desc:     "sample",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   2,
			expected: 6,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.expected {
				t.Errorf("expected: %+v, got: %+v", tc.expected, got)
			}
		})
	}

}
