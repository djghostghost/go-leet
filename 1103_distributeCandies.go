package main

import "math"

func distributeCandies(candies int, num_people int) []int {

	result := make([]int, num_people)

	i := 0
	left := candies
	add := 1
	for left > 0 {

		candie := int(math.Min(float64(add), float64(left)))
		result[i] += candie
		left -= candie

		add++
		i = (i + 1) % num_people
	}
	return result
}
