package main

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

// iterative implementation using a queue
func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		array = append(array, node.Name)
		for _, child := range node.Children {
			queue = append(queue, child)
		}
	}
	return array
}
