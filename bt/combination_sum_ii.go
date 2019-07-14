package bt

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {

	ress := make([][]int, 0)

	counter := make(map[int]int, 0)

	for _, ele := range candidates {
		counter[ele]++
	}

	eles := make([]int, 0)

	for k := range counter {
		eles = append(eles, k)
	}
	sort.Ints(eles)

	var v func(i int, target int, res []int)
	v = func(i int, t int, res []int) {

		if t == 0 && len(res) != 0 {

			dst := make([]int, len(res))
			copy(dst, res)
			ress = append(ress, dst)
			return
		}

		if i >= len(eles) || t < 0 || eles[i]>t {
			return
		}

		ele := eles[i]
		sum := 0
		eleSlice := make([]int, 0)
		for j := 0; j < counter[ele]; j++ {
			sum += ele
			if sum > target {
				break
			}
			eleSlice = append(eleSlice, ele)
			v(i+1, t-sum, append(res, eleSlice...))
		}
		v(i+1, t, res)
	}
	v(0, target, make([]int, 0))
	return ress
}

func combinationSum2Helper(counter *map[int]int, eles []int, i int, target int, res []int, ress *[][]int) {

	if target == 0 && len(res) != 0 {

		dst := make([]int, len(res))
		copy(dst, res)
		*ress = append(*ress, dst)
		return
	}

	if i >= len(eles) || target < 0 {
		return
	}

	ele := eles[i]

	sum := 0
	eleSlice := make([]int, 0)
	for j := 0; j < (*counter)[ele]; j++ {
		sum += ele
		eleSlice = append(eleSlice, ele)

		combinationSum2Helper(counter, eles, i+1, target-sum, append(res, eleSlice...), ress)
	}
	combinationSum2Helper(counter, eles, i+1, target, res, ress)
}
