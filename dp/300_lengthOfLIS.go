package dp

import "math"

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	result := 0

	dp := make([]int, len(nums))

	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {

			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		result = int(math.Max(float64(dp[i]), float64(result)))

	}
	return result
}
