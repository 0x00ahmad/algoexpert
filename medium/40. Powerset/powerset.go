package main

import "math"

func Powerset(array []int) [][]int {
	powerset := [][]int{}
	for _, num := range array {
		l := len(powerset)
		for i := 0; i < l; i++ {
			subset := powerset[i]
			new_subset := append([]int{}, subset...)
			new_subset = append(new_subset, num)
			powerset = append(powerset, new_subset)
		}
	}
	return powerset
}

func PowersetBitWise(array []int) [][]int {
	powerset := [][]int{}
	upper_bound := int(math.Pow(2, float64(len(array))))
	for x := 0; x < upper_bound; x++ {
		subset := []int{}
		for idx := 0; idx < len(array); idx++ {
			if x&int(math.Pow(2, float64(idx))) > 0 {
				subset = append(subset, array[idx])
			}
		}
		powerset = append(powerset, subset)
	}
	return powerset
}
