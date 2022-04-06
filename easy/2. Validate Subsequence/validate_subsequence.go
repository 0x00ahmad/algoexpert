package main

func IsValidSubsequence(array []int, sequence []int) bool {
	i := 0
	for _, v := range array {
		if i < len(sequence) && v == sequence[i] {
			i++
		}
	}
	return i == len(sequence)
}

