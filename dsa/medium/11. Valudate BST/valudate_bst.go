package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	return _validateBST(tree, math.MinInt32, math.MaxInt32)
}

func _validateBST(tree *BST, min_val, max_val int) bool {
	if tree.Value < min_val || tree.Value >= max_val {
		return false
	}

	if tree.Left != nil && !_validateBST(tree.Left, min_val, tree.Value) {
		return false
	}

	if tree.Right != nil && !_validateBST(tree.Right, tree.Value, max_val) {
		return false
	}

	return true
}
