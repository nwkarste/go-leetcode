package utils

import ()

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	for k > 1 {
		var newList []*ListNode
		for i := k - 1; i > 0; i -= 2 {
			merge := merge2Lists(lists[i], lists[i-1])
			newList = append(newList, merge)
			if i == 2 {
				newList = append(newList, lists[0])
			}
		}
		k = len(newList)
		lists = newList
	}
	return lists[0]
}

func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	var prevNode *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		prevNode = list1
		list1 = list1.Next
	} else {
		prevNode = list2
		list2 = list2.Next
	}
	topNode := prevNode
	if prevNode.Next != nil {
		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				prevNode.Next = list1
				prevNode = list1
				list1 = list1.Next
			} else {
				prevNode.Next = list2
				prevNode = list2
				list2 = list2.Next
			}
		}
	}

	if list1 != nil {
		prevNode.Next = list1
	}
	if list2 != nil {
		prevNode.Next = list2
	}
	return topNode
}
