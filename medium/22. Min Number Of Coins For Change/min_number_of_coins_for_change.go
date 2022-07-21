package main

import (
	"math"
)

// O(nd) time | O(n) space
func MinNumberOfCoinsForChange(n int, denoms []int) int {
	coins := make([]int, n+1)
	for i := range coins {
		coins[i] = math.MaxInt32
	}
	coins[0] = 0
	for _, denom := range denoms {
		for amount := range coins {
			if denom <= amount {
				coins[amount] = min(coins[amount], coins[amount-denom]+1)
			}
		}
	}
	if coins[n] == math.MaxInt32 {
		return -1
	}
	return coins[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
