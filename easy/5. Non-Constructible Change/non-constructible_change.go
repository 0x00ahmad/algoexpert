package main

import "sort"

func NonConstructibleChange(coins []int) int {
	c_size := len(coins)
	if c_size == 0 {
		return 1
	}
	sort.Ints(coins)
	if coins[0] > 1 {
		return 1
	}
	curr_change := 0
	for i := 0; i < c_size; i++ {
		if coins[i] > curr_change+1 {
			return curr_change + 1
		}
		curr_change += coins[i]
	}
	return curr_change + 1
}
