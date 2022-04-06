package main

func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]int)
	n_len := len(array)

	for i := 0; i < n_len; i++ {
		num := array[i]
		diff := target - num
		if _, ok := m[diff]; ok {
			return []int{diff, num}
		}
		m[num] = i
	}

	return []int{}
}
