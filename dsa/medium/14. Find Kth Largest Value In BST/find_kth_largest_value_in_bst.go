package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type TraversalInfo struct {
	visitedNodeCount int
	latestNodeVal    int
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	info := TraversalInfo{0, -1}
	reverseInOrderFindValue(tree, k, &info)
	return info.latestNodeVal
}

func reverseInOrderFindValue(tree *BST, k int, info *TraversalInfo) {
	if tree == nil || info.visitedNodeCount >= k {
		return
	}
	reverseInOrderFindValue(tree.Right, k, info)
	if info.visitedNodeCount < k {
		info.visitedNodeCount = info.visitedNodeCount + 1
		info.latestNodeVal = tree.Value
		reverseInOrderFindValue(tree.Left, k, info)
	}
}
