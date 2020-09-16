package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBinarySearchTree(value int) *BinarySearchTree {
	return &BinarySearchTree{Value: value}
}

func TestCase1(t *testing.T) {
	root := NewBinarySearchTree(10)
	root.Left = NewBinarySearchTree(5)
	root.Left.Left = NewBinarySearchTree(2)
	root.Left.Left.Left = NewBinarySearchTree(1)
	root.Left.Right = NewBinarySearchTree(5)
	root.Right = NewBinarySearchTree(15)
	root.Right.Right = NewBinarySearchTree(22)

	inOrder := []int{1, 2, 5, 5, 10, 15, 22}
	//preOrder := []int{10, 5, 2, 1, 5, 15, 22}
	//postOrder := []int{1, 2, 5, 5, 22, 15, 10}

	require.Equal(t, inOrder, root.InOrderTraverse([]int{}))
	//	require.Equal(t, preOrder, root.PreOrderTraverse([]int{}))
	//	require.Equal(t, postOrder, root.PostOrderTraverse([]int{}))
}
