package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	if top.Name == descendantOne.Name || top.Name == descendantTwo.Name {
		return top
	}

	d1 := getNodeDepth(top, descendantOne)
	d2 := getNodeDepth(top, descendantTwo)

	if d1 != d2 {
		for d1 > d2 {
			d1--
			descendantOne = descendantOne.Ancestor
		}
		for d2 > d1 {
			d2--
			descendantTwo = descendantTwo.Ancestor
		}
	}

	for descendantOne.Name != descendantTwo.Name {
		descendantOne = descendantOne.Ancestor
		descendantTwo = descendantTwo.Ancestor
	}

	return descendantOne
}

func getNodeDepth(top, descendantOne *AncestralTree) int {
	depth := 0
	is_ancestor_equal_to_top := false
	curr := descendantOne
	for !is_ancestor_equal_to_top {
		if curr.Name == top.Name {
			is_ancestor_equal_to_top = true
			break
		}
		curr = curr.Ancestor
		depth += 1
	}
	return depth
}
