package main

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		array = append(array, node.Name)
		queue = append(queue, node.Children...)
	}
	return array
}
