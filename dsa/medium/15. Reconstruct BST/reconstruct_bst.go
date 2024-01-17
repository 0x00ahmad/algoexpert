package main

import (
	"math"
)

// This is an input class. Do not edit. ~ Ok :|
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type TreeInfo struct {
	rootIdx int
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	TreeInfo := &TreeInfo{}
	return constructBSTFromRange(math.MinInt32, math.MaxInt32, preOrderTraversalValues, TreeInfo)
}

func constructBSTFromRange(lower, upper int, preOrderTraversalValues []int, subTreeInfo *TreeInfo) *BST {
	if subTreeInfo.rootIdx == len(preOrderTraversalValues) {
		return nil
	}

	rootValue := preOrderTraversalValues[subTreeInfo.rootIdx]
	// the < is not <= because we want to include duplicates
	if rootValue < lower || rootValue >= upper {
		return nil
	}

	root := &BST{Value: rootValue}
	subTreeInfo.rootIdx++
	root.Left = constructBSTFromRange(lower, rootValue, preOrderTraversalValues, subTreeInfo)
	root.Right = constructBSTFromRange(rootValue, upper, preOrderTraversalValues, subTreeInfo)

	return root
}
