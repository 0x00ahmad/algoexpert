package main

import "sort"

// The first typical three sum optimal solution
func ThreeNumberSum(array []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(array)
	a_len := len(array)

	for i := 0; i < a_len; i++ {
		if i > 0 && array[i] == array[i-1] {
			continue
		}
		j := i + 1
		k := a_len - 1
		for j < k {
			sum := array[i] + array[j] + array[k]

			if sum == target {
				result = append(result, []int{array[i], array[j], array[k]})
				j++
				k--
				for j < k && array[j] == array[j-1] {
					j++
				}
				for j < k && array[k] == array[k+1] {
					k--
				}
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return result

}
