package main

import (
	"math"
)

func pathInZigZagTree(label int) []int {

	depth := int(math.Ceil(math.Log2(float64(label))))

	if label&(label-1) == 0 {
		depth += 1
	}

	path := make([]int, 1)
	path[0] = label

	for i := depth; i >= 2; i-- {
		if i%2 == 0 {
			minmax := int(3*math.Pow(float64(2), float64(i-1)) - 1)
			origin := minmax - path[len(path)-1]
			if origin%2 == 0 {
				path = append(path, origin/2)
			} else {
				path = append(path, (origin-1)/2)
			}
		} else {

			minmax := int(3*math.Pow(float64(2), float64(i-2)) - 1)

			if path[len(path)-1]%2 == 0 {
				path = append(path, minmax-(path[len(path)-1]/2))
			} else {
				path = append(path, minmax-(path[len(path)-1]-1)/2)
			}
		}
	}
	////反转数组
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
	return path

}
