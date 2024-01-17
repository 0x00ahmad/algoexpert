package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	if node.Right != nil {
		return findLeftMostLeaf(node.Right)
	}
	return findRightMostsParent(node)
}

func findLeftMostLeaf(node *BinaryTree) *BinaryTree {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func findRightMostsParent(node *BinaryTree) *BinaryTree {
	current := node
	for current.Parent != nil && current.Parent.Right == current {
		current = current.Parent
	}
	return current.Parent
}
