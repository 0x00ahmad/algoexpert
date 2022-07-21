package main

import (
	"math"
)

// time O(n) | space O(1)
func KadanesAlgorithm(array []int) int {
	maxSum := array[0]
	currentSum := array[0]
	for i := 1; i < len(array); i++ {
		currentSum = int(math.Max(float64(array[i]), float64(currentSum+array[i])))
		maxSum = int(math.Max(float64(maxSum), float64(currentSum)))
	}
	return maxSum
}
