package utils

import (
	"fmt"
	"testing"
)

func TestKnightProbability(t *testing.T) {
	type testCase struct {
		desc     string
		N        int
		K        int
		r        int
		c        int
		expected float64
	}
	testCases := []testCase{
		testCase{
			desc:     "sample1",
			N:        3,
			K:        2,
			r:        0,
			c:        0,
			expected: 0.0625,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			prob := knightProbability(tc.N, tc.K, tc.r, tc.c)
			if prob != tc.expected {
				t.Errorf("expected: %+v, got: %+v", tc.expected, prob)
			}
		})
	}

}
