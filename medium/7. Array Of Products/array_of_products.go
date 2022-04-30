package main

func ArrayOfProducts(array []int) []int {
	out := make([]int, len(array))
	for i := range array {
		product := 1
		for j := range array {
			if i != j {
				product *= array[j]
			}
		}
		out[i] = product
	}
	return out
}
