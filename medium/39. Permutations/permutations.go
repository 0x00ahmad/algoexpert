package main

func GetPermutations(array []int) [][]int {
	permutations := [][]int{}
	generatePermutationsInPlace(0, array, &permutations)
	return permutations
}

// INFO: this one is more efficient than the copy method in terms of time - O(n n!)
func generatePermutationsInPlace(currIdx int, array []int, permutations *[][]int) {
	if currIdx == len(array)-1 {
		newPerm := make([]int, len(array))
		copy(newPerm, array)
		*permutations = append(*permutations, newPerm)
		return
	}

	for idx := currIdx; idx < len(array); idx++ {
		swap(array, currIdx, idx)
		generatePermutationsInPlace(currIdx+1, array, permutations)
		swap(array, currIdx, idx)
	}
}

func swap(array []int, a, b int) {
	array[a], array[b] = array[b], array[a]
}

// INFO: this one is time-wise more costly with upper bound of time - O(n^2 n!)
func generatePermutationsCopyMethod(array []int, currentP []int, permutations *[][]int) {
	if len(array) == 0 && len(currentP) != 0 {
		*permutations = append(*permutations, currentP)
	}
	for i := range array {
		newArr := make([]int, i)
		copy(newArr, array[:i])
		newArr = append(newArr, array[i+1:]...)
		newPerm := make([]int, len(currentP))
		copy(newPerm, currentP)
		newPerm = append(newPerm, array[i])
		generatePermutations(newArr, newPerm, permutations)
	}
}
