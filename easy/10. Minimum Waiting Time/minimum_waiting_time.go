package main

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	minWaitingTime := 0
	q_len := len(queries)
	for i := 0; i < q_len; i++ {
		minWaitingTime += (q_len - (i + 1)) * queries[i]
	}
	return minWaitingTime
}

