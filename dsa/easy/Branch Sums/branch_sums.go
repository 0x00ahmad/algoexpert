package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func recurse(node *BinaryTree, sum int, result *[]int) {
  if node == nil {
    return
  }

  sum += node.Value
  if node.Left == nil && node.Right == nil {
    *result = append(*result, sum)
    return
  }
  if node.Left != nil {
    recurse(node.Left, sum, result)
  }
  if node.Right != nil {
    recurse(node.Right, sum, result)
  }
}


func BranchSums(root *BinaryTree) []int {
	result := []int{}
  recurse(root, 0, &result)
  return result
}
