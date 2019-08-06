package main

import (
	"bytes"
	"fmt"
)

func baseNeg2(N int) string {

	var result bytes.Buffer

	if N == 0 {
		return "0"
	}

	for N != 0 {

		result.WriteRune(rune('0' + N%(-2)))
		N = N / (-2)
		if N < 0 {
			N = ^N + 1
		}
		fmt.Println(N)
	}

	return result.String()

}
