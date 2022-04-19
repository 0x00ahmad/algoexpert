package main

func IsMonotonic(array []int) bool {
	if len(array) < 2 {
		return true
	}
	a_len := len(array)
	is_descreasing := true
	is_increasing := true

	for i := 1; i < a_len; i++ {
		if array[i] > array[i-1] {
			is_descreasing = false
		}
		if array[i] < array[i-1] {
			is_increasing = false
		}
	}
	return is_descreasing || is_increasing
}
