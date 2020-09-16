package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewLinkedList(val int, others ...int) *ListNode {
	ll := &ListNode{Val: val}
	current := ll
	for _, other := range others {
		current.Next = &ListNode{Val: other}
		current = current.Next
	}
	return ll
}

func NewLinkedList2(root int, children ...int) *ListNode {
	ll := &ListNode{root, nil}
	ll.Add(children...)
	return ll
}

func (ll *ListNode) toArray() []int {
	vals := []int{}
	current := ll
	for current != nil {
		vals = append(vals, current.Val)
		current = current.Next
	}
	return vals
}

func (ll *ListNode) Add(values ...int) {
	current := ll
	for current.Next != nil {
		current = current.Next
	}
	for value := range values {
		current.Next = &ListNode{value, nil}
		current = current.Next
	}
}

func (ll *ListNode) GetNth(n int) *ListNode {
	counter, current := 1, ll
	for counter < n {
		counter, current = counter+1, current.Next
	}
	return current
}

func TestLoop(t *testing.T) {
	ll := NewLinkedList2(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ll.GetNth(10).Next = ll.GetNth(5)
	output, expected := DetectCycle(ll), ll.GetNth(5)
	require.Equal(t, expected, output)
}

func TestReverse(t *testing.T) {
	input := NewLinkedList(0, 1, 2, 3, 4, 5)
	output := ReverseLinkedList(input)
	expected := []int{5, 4, 3, 2, 1, 0}
	require.NotNil(t, output)
	require.Equal(t, expected, output.toArray())
}

func TestMerge(t *testing.T) {
	list1 := NewLinkedList(2, 6, 7, 8)
	list2 := NewLinkedList(1, 3, 4, 5, 9, 10)
	output := MergeLinkedLists(list1, list2)
	expectedNodes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, output.toArray(), expectedNodes)
}
