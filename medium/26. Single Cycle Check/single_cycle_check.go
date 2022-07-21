package main

// time O(n) | space O(1)
func HasSingleCycle(array []int) bool {
	array_length := len(array)
	STARTING_INDEX := 0
	visited_indices := 0
	curr_idx := STARTING_INDEX
	for visited_indices < array_length {
		if visited_indices > 0 && curr_idx == STARTING_INDEX {
			return false
		}
		visited_indices++
		curr_idx = get_next_index(curr_idx, array[curr_idx], array_length)
	}
	return curr_idx == STARTING_INDEX
}

func get_next_index(curr, jump_by, array_len int) int {
	next_idx := (curr + jump_by) % array_len
	if next_idx < 0 {
		return next_idx + array_len
	}
	return next_idx
}
