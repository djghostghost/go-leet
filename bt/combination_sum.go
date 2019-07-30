package bt

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {

	ress := make([][]int, 0)

	sort.Ints(candidates)
	var v func(int, int, []int)

	v = func(i int, t int, res []int) {

		if t == 0 && len(res) != 0 {
			dst := make([]int, len(res))
			copy(dst, res)
			ress = append(ress, dst)
			return
		}

		if t < 0 || i >= len(candidates) || (i < len(candidates) && t < candidates[i]) {
			return
		}

		num := candidates[i]
		v(i, t-num, append(res, num))
		v(i+1, t, res)

	}
	v(0, target, make([]int, 0))
	return ress

}
