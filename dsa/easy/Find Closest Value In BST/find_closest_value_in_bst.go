package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	var current_closest int = tree.Value
	current_node := tree

  for current_node != nil {
    if math.Abs(float64(target - current_closest)) > math.Abs(float64(target - current_node.Value)) {
      current_closest = current_node.Value
    }
    if target < current_node.Value {
      current_node = current_node.Left
    } else {
      current_node = current_node.Right
    }
  }

	return current_closest
}
