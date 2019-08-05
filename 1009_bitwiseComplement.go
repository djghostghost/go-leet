package main

import (
	"fmt"
	"math"
)

func bitwiseComplement(N int) int {

	binStr := fmt.Sprintf("%b", N)

	result := float64(0)
	for i := len(binStr) - 1; i >= 0; i-- {
		result += float64('1'-binStr[i]) * math.Pow(float64(2), float64(len(binStr)-i-1))
	}

	return int(result)
}
