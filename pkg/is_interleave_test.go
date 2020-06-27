package utils

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	type testCase struct {
		desc     string
		s1       string
		s2       string
		s3       string
		expected bool
	}
	testCases := []testCase{
		testCase{
			desc:     "failure2",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
		},
		testCase{
			desc:     "failure",
			s1:       "bc",
			s2:       "bad",
			s3:       "bcbad",
			expected: true,
		},
		testCase{
			desc:     "long",
			s1:       "bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa",
			s2:       "babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab",
			s3:       "babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab",
			expected: false,
		},
		testCase{
			desc:     "timeout",
			s1:       "bc",
			s2:       "ad",
			s3:       "cbad",
			expected: false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			out := isInterleave(tc.s1, tc.s2, tc.s3)
			if out != tc.expected {
				t.Errorf("expected: %+v, got: %+v", tc.expected, out)
			}
		})
	}

}
