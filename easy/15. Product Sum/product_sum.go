package main

type SpecialArray []interface{}

func _sum(array []interface{}, depth int) int {
	var sum int
	for _, v := range array {
		switch v.(type) {
		case SpecialArray:
			sum += _sum(v.(SpecialArray), depth+1)
		default:
			sum += v.(int)
		}
	}
	return sum * depth
}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return _sum(array, 1)
}

