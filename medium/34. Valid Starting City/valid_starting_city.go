package main

func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	currGas := 0

	targetValidCity := 0
	minGas := 0

	for idx := 1; idx < len(distances); idx++ {
		currGas += (fuel[idx-1] * mpg) - distances[idx-1]
		if currGas < minGas {
			minGas = currGas
			targetValidCity = idx
		}
	}

	return targetValidCity
}
