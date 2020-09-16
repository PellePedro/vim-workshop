package main

type BinarySearchTree struct {
	Value int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func (tree *BinarySearchTree) InOrderTraverse(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array)
	}
	return array
}
