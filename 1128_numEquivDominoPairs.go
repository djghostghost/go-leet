package main

const MAX_LENGTH = 40001

func numEquivDominoPairs(dominoes [][]int) int {

	mark := make(map[int]int)

	min := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	count := 0
	for _, dominoe := range dominoes {
		value := min(dominoe[0], dominoe[1])*MAX_LENGTH + max(dominoe[0], dominoe[1])
		mark[value]++
	}
	for _, v := range mark {
		if v < 2 {
			continue
		}
		count += (v * (v - 1)) / 2
	}
	return count
}
