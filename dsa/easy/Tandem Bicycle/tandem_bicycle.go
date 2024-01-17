package main

import (
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	result := 0

	sort.Ints(redShirtSpeeds)
	if fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))
	} else {
		sort.Ints(blueShirtSpeeds)
	}

	for i := 0; i < len(redShirtSpeeds); i++ {
		rss := redShirtSpeeds[i]
		bss := blueShirtSpeeds[i]
		if rss > bss {
			result += rss
		} else {
			result += bss
		}
	}

	return result
}
