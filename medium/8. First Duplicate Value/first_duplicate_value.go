package main

import "math"

func FirstDuplicateValue(array []int) int {
	for i := 0; i < len(array); i++ {
		if array[int(math.Abs(float64(array[i])))-1] < 0 {
			return int(math.Abs(float64(array[i])))
		}
		array[int(math.Abs(float64(array[i])))-1] *= -1
	}
	return -1
}
