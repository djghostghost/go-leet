package dp

import (
	"fmt"
	"math"
	"sort"
)

func longestStrChain(words []string) int {

	if words == nil || len(words) == 0 {
		return 0
	}

	sort.Sort(StringSlice(words))

	mark := make(map[string][]int)

	for i := 0; i < len(words); i++ {
		mark[words[i]] = make([]int, 26)
		for _, ch := range words[i] {
			mark[words[i]][ch-'a']++
		}
	}

	dp := make([]int, len(words))
	dp[0] = 1
	result := dp[0]
	for i := 1; i < len(words); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {

			if isChain(&mark, words[i], words[j]) {
				fmt.Printf("StringChain:%s------%s\n", words[i], words[j])
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
			}
		}

		result = int(math.Max(float64(result), float64(dp[i])))
	}

	return result
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return len(p[i]) < len(p[j])
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func isChain(mark *map[string][]int, a string, b string) bool {
	if len(a)-len(b) != 1 {
		return false
	}

	markA := (*mark)[a]
	markB := (*mark)[b]

	diff := 0
	for i := 0; i < 26; i++ {
		diff += int(math.Abs(float64(markA[i] - markB[i])))
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
