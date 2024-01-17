package main

type doubleArray = [][]int

func MinimumPassesOfMatrix(matrix doubleArray) int {
	passes := convertNegatives(matrix)
	for _, row := range matrix {
		for _, elem := range row {
			if elem < 0 {
				return -1
			}
		}
	}
	return passes - 1
}

func convertNegatives(matrix doubleArray) int {
	passes := 0
	Q := getAllPositivePositions(matrix)
	for len(Q) > 0 {
		q_curr_len := len(Q)
		for q_curr_len > 0 {
			// pop the curr val from queue
			currPos := Q[0]
			Q = Q[1:]
			currRow, currCol := currPos[0], currPos[1]
			adjacentIndices := getAdjacentIndices(currRow, currCol, matrix)
			for _, adjacentPosition := range adjacentIndices {
				pRow, pCol := adjacentPosition[0], adjacentPosition[1]
				value := matrix[pRow][pCol]
				if value < 0 {
					matrix[pRow][pCol] *= -1
					Q = append(Q, adjacentPosition)
				}
			}
			q_curr_len--
		}
		passes++
	}
	return passes
}

func getAllPositivePositions(matrix doubleArray) doubleArray {
	poses := make(doubleArray, 0)
	for row := range matrix{
		for col := range matrix[row]{
			if matrix[row][col] > 0 {
				poses = append(poses, []int{row, col})
			}
		}
	}
	return poses
}

func getAdjacentIndices(r, c int, matrix doubleArray) doubleArray {
	adjacentPoses := make(doubleArray, 0)
	if r > 0 {
		adjacentPoses = append(adjacentPoses, []int{r - 1, c})
	}
	if r < len(matrix)-1 {
		adjacentPoses = append(adjacentPoses, []int{r + 1, c})
	}
	if c > 0 {
		adjacentPoses = append(adjacentPoses, []int{r, c - 1})
	}
	if c < len(matrix[0])-1 {
		adjacentPoses = append(adjacentPoses, []int{r, c + 1})
	}
	return adjacentPoses
}
