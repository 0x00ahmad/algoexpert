package main

func MoveElementToEnd(array []int, toMove int) []int {
	left := 0
	right := len(array) - 1

	for left < right {
		for array[left] != toMove && left < right {
			left++
		}

		for array[right] == toMove && left < right {
			right--
		}

		if left < right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}
	return array
}

