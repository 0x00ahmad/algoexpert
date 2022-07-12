package main

import (
	"math"
)

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	if len(array) == 2 {
		return int(math.Max(float64(array[0]), float64(array[1])))
	}
	return int(math.Max(float64(array[0]+MaxSubsetSumNoAdjacent(array[2:])), float64(MaxSubsetSumNoAdjacent(array[1:]))))
}
