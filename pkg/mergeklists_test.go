package utils

import (
	"fmt"
	"testing"
)

func array2linkedlist(list []int) *ListNode {
	var start *ListNode
	var prevNode *ListNode
	if len(list) == 0 {
		return start
	}
	start = &ListNode{
		Val: list[0],
	}
	list = list[1:]
	prevNode = start
	for _, val := range list {
		node := &ListNode{
			Val: val,
		}
		prevNode.Next = node
		prevNode = node
	}
	return start
}

func linkedlist2array(list *ListNode) []int {
	var arrlist []int
	for list != nil {
		arrlist = append(arrlist, list.Val)
		list = list.Next
	}
	return arrlist
}

func TestMergeKLists(t *testing.T) {
	type testCase struct {
		desc     string
		lists    [][]int
		expected []int
	}
	testCases := []testCase{
		testCase{
			desc: "sample",
			lists: [][]int{
				{0},
				{1},
			},
			expected: []int{0, 1},
		},
		testCase{
			desc: "sample",
			lists: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		testCase{
			desc: "nil",
			lists: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
				{},
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		testCase{
			desc:     "nil",
			lists:    [][]int{},
			expected: []int{},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%s", i, tc.desc), func(t *testing.T) {
			lists := make([]*ListNode, len(tc.lists))
			for i, list := range tc.lists {
				lists[i] = array2linkedlist(list)
			}
			list := mergeKLists(lists)
			arrlist := linkedlist2array(list)
			if len(tc.expected) != len(arrlist) {
				t.Errorf("expected: %+v, got: %+v", tc.expected, arrlist)
			}
			for i, val := range tc.expected {
				if arrlist[i] != val {
					t.Errorf("expected: %+v, got: %+v", tc.expected, arrlist)
				}
			}
		})
	}

}
