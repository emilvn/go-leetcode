package main

// # MERGE TWO SORTED LISTS
//
// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Example 1:
//
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:
//
// Input: list1 = [], list2 = []
// Output: []
// Example 3:
//
// Input: list1 = [], list2 = [0]
// Output: [0]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	d := &ListNode{}
	t := d

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			t.Next = list1
			list1 = list1.Next
		} else {
			t.Next = list2
			list2 = list2.Next
		}
		t = t.Next
	}

	if list1 != nil {
		t.Next = list1
	} else {
		t.Next = list2
	}

	return d.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
