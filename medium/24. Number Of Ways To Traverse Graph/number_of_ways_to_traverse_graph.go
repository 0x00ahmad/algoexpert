package main

// the "optimal" accepted solution is to use tabulation to solve this, but we can do better
// use permutations to solve this
// time O(n+m) | space O(1)
func NumberOfWaysToTraverseGraph(width int, height int) int {
	maxRightOperations := width - 1
	maxDownOperations := height - 1

	numerator := factorial(maxRightOperations + maxDownOperations)
	denominator := factorial(maxRightOperations) * factorial(maxDownOperations)
	return numerator / denominator
}

// iterative
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

