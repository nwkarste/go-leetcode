package utils

import (
	"fmt"
	"testing"
)

func TestNthPersonGetsNthSeat(t *testing.T) {
	type testCase struct {
		desc     string
		n        int
		expected float64
	}
	testCases := []testCase{
		testCase{
			desc:     "1",
			n:        1,
			expected: 1,
		},
		testCase{
			desc:     "2",
			n:        2,
			expected: .5,
		},
		testCase{
			desc:     "3",
			n:        3,
			expected: .5,
		},
		testCase{
			desc:     "200",
			n:        200,
			expected: .5,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			prob := nthPersonGetsNthSeat(tc.n)
			if prob != tc.expected {
				t.Errorf("expected: %+v, got: %+v", tc.expected, prob)
			}
		})
	}

}
