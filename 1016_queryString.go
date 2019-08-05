package main

import (
	"fmt"
	"strings"
)

func queryString(S string, N int) bool {

	result := true
	for i := 1; i <= N; i++ {
		nStr := fmt.Sprintf("%b", i)
		result = result && strings.Contains(S, nStr)
		if !result {
			return false
		}
	}
	return result
}
