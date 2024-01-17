package main

import "sort"

func SmallestDifference(a1, a2 []int) []int {
	sort.Ints(a1)
	sort.Ints(a2)
	result := make([]int, 2)
	i, j := 0, 0
	min := 2147483648
	min_i, min_j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] == a2[j] {
			return []int{a1[i], a2[j]}
		}
		diff := a1[i] - a2[j]
		if diff < 0 {
			diff = -diff
		}
		if diff < min {
			min = diff
			min_i = i
			min_j = j
		}
		if a1[i] < a2[j] {
			i++
		} else {
			j++
		}
	}
	result[0] = a1[min_i]
	result[1] = a2[min_j]
	return result
}
