package main

var WHITE int = 0
var GREY int = 1
var BLACK int = 2

func CycleInGraph(edges [][]int) bool {
	colors := make([]int, len(edges))
	noOfNodes := len(edges)
	for node := 0; node < noOfNodes; node++ {
		if colors[node] != WHITE {
			continue
		}
		hasCycle := traverseAndColorize(node, edges, colors)
		if hasCycle {
			return true
		}
	}
	return false
}

func traverseAndColorize(node int, edges [][]int, colors []int) bool {
	colors[node] = GREY
	neighbours := edges[node]
	for _, neighbour := range neighbours {
		neighborColor := colors[neighbour]
		if neighborColor == GREY {
			return true
		}
		if neighborColor != WHITE {
			continue
		}
		hasCycle := traverseAndColorize(neighbour, edges, colors)
		if hasCycle {
			return true
		}
	}
	colors[node] = BLACK
	return false
}
