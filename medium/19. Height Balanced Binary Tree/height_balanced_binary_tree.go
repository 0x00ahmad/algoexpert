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

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	if tree == nil {
		return true
	}
	leftHeight := HeightBalancedBinaryTree(tree.Left)
	rightHeight := HeightBalancedBinaryTree(tree.Right)
	if leftHeight && rightHeight && math.Abs(Height(tree.Left)-Height(tree.Right)) <= 1 {
		return true
	}
	return false
}

func Height(tree *BinaryTree) float64 {
	if tree == nil {
		return 0
	}
	return math.Max(Height(tree.Left), Height(tree.Right)) + 1
}
