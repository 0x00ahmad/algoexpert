package main

// O(wh) time & space
func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			sizes = traverseNode(i, j, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, sizes []int) []int {
	curr_size := 0
	// for bfs it's a queue, for dfs it's a stack
	nodes_to_visit := [][]int{{i, j}}
	for len(nodes_to_visit) > 0 {
		curr_node := nodes_to_visit[0]
		nodes_to_visit = nodes_to_visit[1:]
		i := curr_node[0]
		j := curr_node[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}
		curr_size++
		// find and explore neighbours
		unvisited_neighbors := getUnvisitedNeighbors(i, j, matrix, visited)
		nodes_to_visit = append(nodes_to_visit, unvisited_neighbors...)
	}
	if curr_size > 0 {
		sizes = append(sizes, curr_size)
	}
	return sizes
}

func getUnvisitedNeighbors(i, j int, matrix [][]int, visited [][]bool) [][]int {
	unvisited := make([][]int, 0)

	if i > 0 && !visited[i-1][j] {
		unvisited = append(unvisited, []int{i - 1, j})
	}

	if i < len(matrix)-1 && !visited[i+1][j] {
		unvisited = append(unvisited, []int{i + 1, j})

	}

	if j > 0 && !visited[i][j-1] {
		unvisited = append(unvisited, []int{i, j - 1})

	}

	if j < len(matrix[0])-1 && !visited[i][j+1] {
		unvisited = append(unvisited, []int{i, j + 1})
	}

	return unvisited
}
