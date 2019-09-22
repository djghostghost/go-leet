package dp

import "math"

func increasingTriplet(nums []int) bool {

	if nums == nil || len(nums) < 3 {
		return false
	}

	dp := make([]int, len(nums))

	dp[0] = 1

	for i := 1; i < len(nums); i++ {

		dp[i] = 1
		for j := 0; j < i; j++ {

			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
				if dp[i] >= 3 {
					return true
				}
			}

		}
	}
	return false

}
