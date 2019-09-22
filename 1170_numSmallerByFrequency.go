package main

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {

	queryCounter := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		queryCounter[i] = f(queries[i])
	}

	wordCounter := make([]int, len(words))

	for i := 0; i < len(wordCounter); i++ {
		wordCounter[i] = f(words[i])
	}

	sort.Ints(wordCounter)

	result := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {

		counter := queryCounter[i]

		for j := 0; j < len(wordCounter); j++ {
			if counter < wordCounter[j] {
				result[i] = len(words) - j
				break
			}
		}

	}

	return result
}

func f(s string) int {

	counter := make([]int, 26)
	for _, ch := range s {
		counter[ch-'a']++
	}

	for _, count := range counter {
		if count > 0 {
			return count
		}
	}
	return 0
}
