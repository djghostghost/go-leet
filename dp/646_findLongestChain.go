package dp

import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {

	sort.Sort(pair(pairs))
	dp := make([]int, len(pairs))
	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(pairs); i++ {

		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
			}
		}
		max = int(math.Max(float64(dp[i]), float64(max)))
	}

	return max
}

type pair [][]int

func (p pair) Len() int {
	return len(p)
}

func (p pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pair) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}
