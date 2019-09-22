package dp

import (
	"math"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {

	if envelopes == nil || len(envelopes) == 0 {
		return 0
	}
	sort.Sort(Envelops(envelopes))

	dp := make([]int, len(envelopes))

	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {

			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
			}
		}

		max = int(math.Max(float64(max), float64(dp[i])))
	}

	return max
}

type Envelops [][]int

func (p Envelops) Len() int {
	return len(p)
}

func (p Envelops) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Envelops) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}
