package main

import "sort"

type Pair struct {
	val int
	idx int
}

// O(n log n) time | O(n) space
func TaskAssignment(k int, tasks []int) [][]int {
	if k == 1 {
		return [][]int{{0, 1}}
	}
	assignments := make([][]int, 0)
	sorted_tasks := make([]Pair, 0)
	for i, e := range tasks {
		sorted_tasks = append(sorted_tasks, Pair{e, i})
	}
	sort.Slice(sorted_tasks, func(a, b int) bool {
		return sorted_tasks[a].val < sorted_tasks[b].val
	})
	lower, upper := 0, len(tasks)-1
	for lower < upper {
		assignments = append(assignments, []int{sorted_tasks[lower].idx, sorted_tasks[upper].idx})
		upper--
		lower++
	}
	return assignments
}
