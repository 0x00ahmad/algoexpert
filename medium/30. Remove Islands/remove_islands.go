package main

// O(wh) | O(wh) space
func RemoveIslands(matrix [][]int) [][]int {
	mr_len := len(matrix)
	mc_len := len(matrix[0])
	for i := 0; i < mr_len; i++ {
		for j := 0; j < mc_len; j++ {
			isRowBorder := i == 0 || i == mr_len-1
			isColBorder := j == 0 || j == mc_len-1
			isBorder := isRowBorder || isColBorder
			if !isBorder || matrix[i][j] != 1 {
				continue
			}
			switchBorderOnesToTwo(matrix, i, j)
		}
	}
	for i := 0; i < mr_len; i++ {
		for j := 0; j < mc_len; j++ {
			switch matrix[i][j] {
			case 1:
				matrix[i][j] = 0
			case 2:
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}

func switchBorderOnesToTwo(matrix [][]int, sRow, sCol int) {
	stack := [][]int{{sRow, sCol}}

	var currentPosition []int
	for len(stack) > 0 {
		currentPosition = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cRow, cCol := currentPosition[0], currentPosition[1]
		matrix[cRow][cCol] = 2
		neighbours := getNeighbours(matrix, cRow, cCol)
		for _, neighbour := range neighbours {
			row, col := neighbour[0], neighbour[1]
			if matrix[row][col] != 1 {
				continue
			}
			stack = append(stack, neighbour)
		}
	}
}

func getNeighbours(matrix [][]int, row, col int) [][]int {
	neighbours := make([][]int, 0)
	numRows := len(matrix)
	numCols := len(matrix[0])
	if row-1 >= 0 { // above
		neighbours = append(neighbours, []int{row - 1, col})
	}
	if row+1 < numRows { // below
		neighbours = append(neighbours, []int{row + 1, col})
	}
	if col-1 >= 0 { // left
		neighbours = append(neighbours, []int{row, col - 1})
	}
	if col+1 < numCols { // right
		neighbours = append(neighbours, []int{row, col + 1})
	}
	return neighbours
}
