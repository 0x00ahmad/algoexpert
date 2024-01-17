package main

func SpiralTraverse(array [][]int) []int {
	var result []int
	if len(array) == 0 {
		return result
	}
	row, row_end, col, col_end := 0, len(array)-1, 0, len(array[0])-1
	direction := 0
	for row <= row_end && col <= col_end {
		if direction == 0 {
			for i := col; i <= col_end; i++ {
				result = append(result, array[row][i])
			}
			row++
			direction = 1
		} else if direction == 1 {
			for i := row; i <= row_end; i++ {
				result = append(result, array[i][col_end])
			}
			col_end--
			direction = 2
		} else if direction == 2 {
			for i := col_end; i >= col; i-- {
				result = append(result, array[row_end][i])
			}
			row_end--
			direction = 3
		} else if direction == 3 {
			for i := row_end; i >= row; i-- {
				result = append(result, array[i][col])
			}
			col++
			direction = 0
		}
	}
	return result
}
