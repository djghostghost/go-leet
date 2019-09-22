package dp

import "math"

func findNumberOfLIS(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	counts := make(map[int]int, len(nums))
	counts[0] = 1
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))

			}
		}
		//if dp[i] == 1 {
		//	counts[i] = 1
		//}

		max = int(math.Max(float64(max), float64(dp[i])))
	}

	for i := 1; i < len(nums); i++ {

		for j := i - 1; j >= 0; j-- {

			if nums[j] < nums[i] && dp[j]+1 == dp[i] {
				counts[i] += counts[j]
			}
		}

		if counts[i] == 0 {
			counts[i] = 1
		}

	}

	result := 0

	for i, val := range counts {
		if dp[i] == max {
			result += val
		}
	}

	return result
}
