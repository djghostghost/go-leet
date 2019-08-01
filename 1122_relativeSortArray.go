package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {

	mark := make(map[int]int)

	for _, ele := range arr1 {
		mark[ele]++
	}

	result := make([]int, 0)

	for _, ele := range arr2 {

		count := mark[ele]
		for i := 0; i < count; i++ {
			result = append(result, ele)
			mark[ele]--
		}
	}

	temp := make([]int, 0)

	for key, val := range mark {

		for i := 0; i < val; i++ {
			temp = append(temp, key)
		}

	}

	sort.Ints(temp)

	result = append(result, temp...)
	return result
}
