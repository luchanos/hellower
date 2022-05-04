package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	maxVal := math.Inf(-1)
	for _, val := range numbers {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func main() {
	fmt.Println(maximum(1, 2, 3, 4, 5, 6, 7))
}
