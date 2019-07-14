package bt

func combinationSum(candidates []int, target int) [][]int {

	ress := make([][]int, 0)
	combinationSumHelper(candidates, 0, target, make([]int, 0), &ress)
	return ress

}

func combinationSumHelper(candiates []int, i int, target int, res []int, ress *[][]int) {

	if target < 0 || i >= len(candiates) {
		return
	}

	if target == 0 && len(res) != 0 {
		copRes := make([]int, len(res))
		copy(copRes,res)
		*ress = append(*ress, copRes)
		return
	}

	num := candiates[i]
	combinationSumHelper(candiates, i+1, target-num, append(res, num), ress)
	combinationSumHelper(candiates, i+1, target, res, ress)

}
