package main

func LongestPeak(array []int) int {
	longestPeakLength := 0
	i := 1

	for i < len(array)-1 {
		if !(array[i-1] < array[i] && array[i+1] < array[i]) {
			i++
			continue
		}
		left := i - 2
		for left >= 0 && array[left] < array[left+1] {
			left--
		}

		right := i + 2
		for right < len(array) && array[right] < array[right-1] {
			right++
		}
		currentPeakLength := right - left - 1
		if currentPeakLength > longestPeakLength {
			longestPeakLength = currentPeakLength
		}
		i = right
	}
	return longestPeakLength
}
