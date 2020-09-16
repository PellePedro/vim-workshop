package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Input: 0 -> 1 -> 2 -> 3 -> 4 -> 5
Output: 5 -> 4 -> 3 -> 2 -> 1
*/
// O(n) time | O(1) space where n is the number of nodes
func ReverseLinkedList(head *ListNode) *ListNode {
	var previousNode, currentNode *ListNode = nil, head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

/*
MergeLinkedList LeedCode 21
O(n + m) time | O(1) space - where
   n is the number of nodes in the first Linked List and
   m is the number of nodes in the second Linked List

1 -> 2 -> 4 -> 8 -> 9 -> 10 null
1 -> 3 -> 4
*/
func MergeLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := &ListNode{Val: -1}
	head := curr
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return head.Next
}

/*
142. Linked List Cycle II
https://leetcode.com/problems/linked-list-cycle-ii/
*/

// Approach1 : Hash Table

/* Approach 2 Floyd's Tortoise and are Algorithm

Floyd's Tortoise and Hare algorithm
Time Complexity: O(n)
Space Complexity: O(1)

Given a linked list, return the node where the cycle begins.
If there is no cycle, return null.
*/

func getIntersection(head *ListNode) *ListNode {
	turtoise := head
	hare := head
	for hare != nil && hare.Next != nil {
		turtoise, hare = turtoise.Next, hare.Next.Next
		if turtoise == hare {
			return turtoise
		}
	}
	return nil
}

func DetectCycle(head *ListNode) *ListNode {

	if head == nil { return nil }

	intersect := getIntersection(head)

	if intersect == nil { return nil }

	ptr1, ptr2 := head, intersect
	for ptr1 != ptr2 {
		ptr1, ptr2 = ptr1.Next, ptr2.Next
	}

	return ptr1
}


// 160. Intersection of Two Linked Lists
