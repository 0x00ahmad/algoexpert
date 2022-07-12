package main

import (
	"math"
)

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func BinaryTreeDepth(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	l := BinaryTreeDepth(tree.Left)
	r := BinaryTreeDepth(tree.Right)
	return int(math.Max(float64(l), float64(r)) + 1)
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	l := BinaryTreeDepth(tree.Left)
	r := BinaryTreeDepth(tree.Right)
	ld := BinaryTreeDiameter(tree.Left)
	rd := BinaryTreeDiameter(tree.Right)
	return int(math.Max(float64(l+r), math.Max(float64(ld), float64(rd))))
}
