package main

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Sort(sort.Reverse(sort.IntSlice(redShirtHeights)))
	sort.Sort(sort.Reverse(sort.IntSlice(blueShirtHeights)))
	var redShirtsInFirstRow bool

	if blueShirtHeights[0] > redShirtHeights[0] {
		redShirtsInFirstRow = true
	} else {
		redShirtsInFirstRow = false
	}

	for i := 0; i < len(redShirtHeights); i++ {
		if redShirtsInFirstRow {
			if redShirtHeights[i] >= blueShirtHeights[i] {
				return false
			}
		} else {
			if blueShirtHeights[i] >= redShirtHeights[i] {
				return false
			}
		}
	}

	return true
}
