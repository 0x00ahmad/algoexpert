package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func recurse(node *BinaryTree, depth int) int {
  if node == nil {
    return 0
  }
  return depth + recurse(node.Left, depth+1) + recurse(node.Right, depth+1)
}

func NodeDepths(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	level := 0
	return recurse(root, level)
}
