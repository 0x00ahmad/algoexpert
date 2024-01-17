package main

// start the root for the tree from the middle of the vector
func MinHeightBST(array []int) *BST {
	if len(array) == 0 {
		return nil
	}
	mid := len(array) / 2
	tree := &BST{Value: array[mid]}

	arrayLeft := array[:mid]
	arrayRight := array[mid+1:]

	tree.Left = MinHeightBST(arrayLeft)
	tree.Right = MinHeightBST(arrayRight)

	return tree
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
