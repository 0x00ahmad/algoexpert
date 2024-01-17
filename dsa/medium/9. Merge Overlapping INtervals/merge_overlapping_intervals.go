package main

// import the sort, fmt and the math package
import (
	"math"
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sortedIntervals := make([][]int, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i][0] < sortedIntervals[j][0]
	})

	var mergedIntervals [][]int
	mergedIntervals = append(mergedIntervals, sortedIntervals[0])
	currentInterval := sortedIntervals[0]

	for i := 1; i < len(sortedIntervals); i++ {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart := sortedIntervals[i][0]
		nextIntervalEnd := sortedIntervals[i][1]

		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = int(math.Max(float64(currentIntervalEnd), float64(nextIntervalEnd)))
		} else {
			mergedIntervals = append(mergedIntervals, sortedIntervals[i])
			currentInterval = sortedIntervals[i]
		}
	}
	return mergedIntervals
}
